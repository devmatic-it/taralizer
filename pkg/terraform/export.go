// Package terraform export engine (Phase A1): renders taralizer model objects
// into Terraform scaffolds driven by a mapping profile.
// Copyright 2021 taralizer authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package terraform

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/devmatic-it/taralizer/pkg/taralizer"
	"gopkg.in/yaml.v3"
)

// Marker comment prefixes written directly above each exported resource block.
// Drift detection (Phase A3) parses these to map resources back to model objects;
// the asset/boundary distinction keeps the mapping unambiguous even when an
// asset id and a boundary id collide.
const (
	AssetMarkerPrefix    = "# taralizer:asset="
	BoundaryMarkerPrefix = "# taralizer:boundary="
)

// hclInvalidChars matches characters that cannot appear in an HCL resource name.
// Hyphens are rejected too: although HCL identifiers accept them, normalizing
// to underscores keeps names shell-friendly and makes "my-vpc"/"my_vpc"
// collisions detectable by the duplicate-name check.
var hclInvalidChars = regexp.MustCompile(`[^a-zA-Z0-9_]`)

// exportPlan is the fully computed description of one resource to write.
type exportPlan struct {
	id       string // model object id (asset or boundary)
	marker   string // full marker comment line, e.g. "# taralizer:asset=vpc"
	techType string // terraform resource type, e.g. "google_compute_network"
	hclName  string // sanitized HCL resource name derived from the model id
	gcpName  string // GCP-safe resource name (lowercase, hyphens) for scaffold {name}
	file     string // file name inside outDir, e.g. "vpc.tf"
	body     string // rendered resource body (scaffold with {name} substituted)
}

// ExportToTerraform renders the exportable technical assets and trust boundaries
// of report into outDir as Terraform scaffolds, using the mapping profile at
// profilePath.
//
// For every model object whose technology maps to a non-empty terraform resource
// type with a scaffold, one file <name>.tf is written containing the marker
// comment (# taralizer:asset=<id> or # taralizer:boundary=<id>) followed by the
// resource block. provider.tf and variables.tf are always written. Objects whose
// technology is missing, unknown to the profile, or unmapped produce a warning —
// they are never dropped silently.
//
// outDir must not exist or be empty unless force is true. Nothing is written if
// validation fails (unknown profile, non-empty dir without force, duplicate HCL
// resource names).
//
// It returns the map of model object id to HCL address (e.g.
// "google_compute_network.vpc") for terraform_ref write-back, the list of
// warnings, and an error if the export could not be completed.
func ExportToTerraform(report taralizer.Report, profilePath string, outDir string, force bool) (map[string]string, []string, error) {
	profile, err := LoadGoogleProfile(profilePath)
	if err != nil {
		return nil, nil, fmt.Errorf("load profile %s: %w", profilePath, err)
	}

	techByID := make(map[string]Technology, len(profile.Technologies))
	for _, tech := range profile.Technologies {
		techByID[tech.ID] = tech
	}

	var (
		plans    []exportPlan
		warnings []string
	)

	addObject := func(kind, id, technology string) {
		if technology == "" {
			warnings = append(warnings, fmt.Sprintf("%s %q: no technology set; skipped", kind, id))
			return
		}
		tech, ok := techByID[technology]
		if !ok {
			warnings = append(warnings, fmt.Sprintf("%s %q: technology %q not found in profile; skipped", kind, id, technology))
			return
		}
		if tech.Terraform == "" {
			warnings = append(warnings, fmt.Sprintf("%s %q: technology %q has no terraform mapping; skipped", kind, id, technology))
			return
		}

		hclName := sanitizeHCLName(id)
		marker := AssetMarkerPrefix + id
		if kind == "boundary" {
			marker = BoundaryMarkerPrefix + id
		}
		plans = append(plans, exportPlan{
			id:       id,
			marker:   marker,
			techType: tech.Terraform,
			hclName:  hclName,
			gcpName:  gcpResourceName(id),
			file:     hclName + ".tf",
			body:     renderScaffold(tech.Scaffold, gcpResourceName(id)),
		})
	}

	for _, asset := range report.TechnicalAssets {
		addObject("asset", asset.Id, asset.Technology)
	}
	for _, boundary := range report.TrustBoundaries {
		addObject("boundary", boundary.Id, boundary.Technology)
	}

	// Detect duplicate HCL resource names before writing anything: two model
	// objects sanitizing to the same name would collide in one file and produce
	// an ambiguous terraform_ref.
	seen := make(map[string]string, len(plans)) // hclName -> first model id
	for _, plan := range plans {
		if first, dup := seen[plan.hclName]; dup {
			return nil, warnings, fmt.Errorf("duplicate HCL resource name %q (model ids %q and %q)", plan.hclName, first, plan.id)
		}
		seen[plan.hclName] = plan.id
	}

	if err := ensureOutDir(outDir, force); err != nil {
		return nil, warnings, err
	}

	if err := os.WriteFile(filepath.Join(outDir, "provider.tf"), []byte(providerTF), 0644); err != nil {
		return nil, warnings, fmt.Errorf("write provider.tf: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "variables.tf"), []byte(variablesTF), 0644); err != nil {
		return nil, warnings, fmt.Errorf("write variables.tf: %w", err)
	}

	refs := make(map[string]string, len(plans))
	for _, plan := range plans {
		content := plan.marker + "\n" +
			fmt.Sprintf("resource %q %q {\n", plan.techType, plan.hclName) +
			plan.body +
			"}\n"
		if err := os.WriteFile(filepath.Join(outDir, plan.file), []byte(content), 0644); err != nil {
			return nil, warnings, fmt.Errorf("write %s: %w", plan.file, err)
		}
		refs[plan.id] = plan.techType + "." + plan.hclName
	}

	return refs, warnings, nil
}

// yamlPlainScalar matches values that are safe to write as unquoted YAML
// scalars. HCL addresses (type.name) always match; anything else is quoted.
var yamlPlainScalar = regexp.MustCompile(`^[A-Za-z0-9_.-]+$`)

// WriteBackTerraformRefs persists the refs returned by ExportToTerraform into
// the model file at modelPath: for every exported object id, its technical
// asset or trust boundary entry gains a terraform_ref field set to the HCL
// address (e.g. google_compute_network.vpc).
//
// The edit is surgical — it inserts or updates one line per object inside the
// existing technical_assets / trust_boundaries list items — so comments,
// ordering, and formatting elsewhere in the file are preserved. A full YAML
// round-trip would drop inline comments and reformat the whole document.
//
// Behavior:
//   - an existing terraform_ref line inside the item is updated in place
//     (re-running export after a scaffold change refreshes refs idempotently);
//   - otherwise the line is inserted directly below the item's "- id:" line;
//   - ids that appear more than once in a section are all updated; the marker
//     comment (which carries the raw model id) remains the drift anchor;
//   - if any ref id is not found in the file, nothing is written and an error
//     lists the missing ids (the model path likely does not match the report);
//   - before writing, the edited document is parsed as a taralizer.Report to
//     guarantee the file cannot be corrupted.
//
// It returns the number of lines inserted or updated.
func WriteBackTerraformRefs(modelPath string, refs map[string]string) (int, error) {
	if len(refs) == 0 {
		return 0, nil
	}

	raw, err := os.ReadFile(modelPath)
	if err != nil {
		return 0, fmt.Errorf("read model %s: %w", modelPath, err)
	}

	lines := strings.Split(string(raw), "\n")
	items, err := findRefItems(lines)
	if err != nil {
		return 0, fmt.Errorf("scan model %s: %w", modelPath, err)
	}

	found := make(map[string]bool, len(refs))
	written := 0
	// Process in reverse so earlier insertions do not shift later indices.
	for i := len(items) - 1; i >= 0; i-- {
		item := items[i]
		ref, ok := refs[item.id]
		if !ok {
			continue
		}
		found[item.id] = true
		lines = upsertTerraformRef(lines, item, ref)
		written++
	}

	var missing []string
	for id := range refs {
		if !found[id] {
			missing = append(missing, id)
		}
	}
	if len(missing) > 0 {
		sort.Strings(missing)
		return 0, fmt.Errorf("model %s does not contain terraform_ref targets: %s", modelPath, strings.Join(missing, ", "))
	}

	content := strings.Join(lines, "\n")
	var probe taralizer.Report
	if err := yaml.Unmarshal([]byte(content), &probe); err != nil {
		return 0, fmt.Errorf("refuse to write %s: edited document no longer parses as a taralizer model: %w", modelPath, err)
	}

	if err := os.WriteFile(modelPath, []byte(content), 0644); err != nil {
		return 0, fmt.Errorf("write model %s: %w", modelPath, err)
	}
	return written, nil
}

// refItem is one "- id: <id>" list entry inside technical_assets or
// trust_boundaries, with enough geometry to edit it.
type refItem struct {
	idLine  int    // index of the "- id:" line
	dashCol int    // column of the sequence dash
	id      string // unquoted model id
}

// findRefItems scans lines for "- id:" entries belonging to the top-level
// technical_assets and trust_boundaries sections. Nested lists (e.g.
// communication_links) are ignored because their items do not start with an
// inline "id:" key at the section's item indentation.
func findRefItems(lines []string) ([]refItem, error) {
	const (
		noSection = iota
		assetsSection
		boundariesSection
	)
	section := noSection
	var items []refItem

	for i, line := range lines {
		trimmed := strings.TrimLeft(line, " ")
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		isItem := strings.HasPrefix(trimmed, "- ") || trimmed == "-"
		if !isItem && !strings.HasPrefix(line, " ") { // column 0: a top-level key
			switch {
			case strings.HasPrefix(trimmed, "technical_assets:"):
				section = assetsSection
			case strings.HasPrefix(trimmed, "trust_boundaries:"):
				section = boundariesSection
			default:
				section = noSection
			}
			continue
		}
		if section == noSection || !isItem {
			continue
		}
		dashCol := len(line) - len(trimmed)
		rest := strings.TrimSpace(strings.TrimPrefix(trimmed, "-"))
		if !strings.HasPrefix(rest, "id:") {
			continue
		}
		value := strings.TrimSpace(strings.TrimPrefix(rest, "id:"))
		if idx := strings.Index(value, " #"); idx >= 0 { // trailing comment
			value = strings.TrimSpace(value[:idx])
		}
		id := stripYAMLQuotes(value)
		if id == "" {
			return nil, fmt.Errorf("line %d: empty id in technical_assets/trust_boundaries item", i+1)
		}
		items = append(items, refItem{idLine: i, dashCol: dashCol, id: id})
	}

	return items, nil
}

// upsertTerraformRef inserts or updates the terraform_ref line for one item
// and returns the new lines (exactly one line is inserted or updated).
func upsertTerraformRef(lines []string, item refItem, ref string) []string {
	fieldCol := item.dashCol + 2 // column of the "id" key in "- id:"
	newLine := strings.Repeat(" ", fieldCol) + "terraform_ref: " + yamlSafeScalar(ref)

	end := len(lines)
	for k := item.idLine + 1; k < len(lines); k++ {
		l := lines[k]
		if strings.TrimSpace(l) == "" {
			continue
		}
		if len(l)-len(strings.TrimLeft(l, " ")) <= item.dashCol {
			end = k
			break
		}
	}
	for k := item.idLine + 1; k < end; k++ {
		t := strings.TrimLeft(lines[k], " ")
		if t == "terraform_ref:" || strings.HasPrefix(t, "terraform_ref:") {
			lines[k] = newLine
			return lines
		}
	}

	out := make([]string, 0, len(lines)+1)
	out = append(out, lines[:item.idLine+1]...)
	out = append(out, newLine)
	out = append(out, lines[item.idLine+1:]...)
	return out
}

// stripYAMLQuotes removes one pair of matching surrounding quotes from a YAML
// scalar.
func stripYAMLQuotes(s string) string {
	if len(s) >= 2 &&
		((s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'')) {
		return s[1 : len(s)-1]
	}
	return s
}

// yamlSafeScalar renders ref as a YAML scalar, quoting it unless it is plain-safe.
func yamlSafeScalar(s string) string {
	if yamlPlainScalar.MatchString(s) {
		return s
	}
	return `"` + strings.ReplaceAll(s, `"`, `\"`) + `"`
}

// ensureOutDir creates outDir if needed and refuses to proceed when it already
// exists with content, unless force is set.
func ensureOutDir(outDir string, force bool) error {
	entries, err := os.ReadDir(outDir)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(outDir, 0755)
		}
		return fmt.Errorf("inspect target directory %s: %w", outDir, err)
	}
	if len(entries) > 0 && !force {
		return fmt.Errorf("target directory %s is not empty (use force to proceed)", outDir)
	}
	return nil
}

// sanitizeHCLName converts a model id into a valid HCL resource name: invalid
// characters (including hyphens) become underscores, empty results fall back to
// "resource", and a leading digit is prefixed so the name stays an identifier.
func sanitizeHCLName(id string) string {
	name := hclInvalidChars.ReplaceAllString(id, "_")
	if name == "" {
		name = "resource"
	}
	if name[0] >= '0' && name[0] <= '9' {
		name = "t_" + name
	}
	return name
}

// gcpResourceName converts a model id into a GCP-safe resource name: lowercase,
// every run of non-alphanumeric characters collapsed to one hyphen, leading and
// trailing hyphens trimmed. GCP name rules (networks, subnets, projects, KMS
// key rings, ...) forbid underscores and require a leading letter, so this is
// deliberately different from sanitizeHCLName. Empty results fall back to
// "resource".
func gcpResourceName(id string) string {
	name := strings.ToLower(id)
	name = gcpInvalidChars.ReplaceAllString(name, "-")
	name = strings.Trim(name, "-")
	if name == "" {
		name = "resource"
	}
	return name
}

// gcpInvalidChars matches runs of characters that cannot appear in a GCP
// resource name (anything not already lowercase alphanumeric).
var gcpInvalidChars = regexp.MustCompile(`[^a-z0-9]+`)

// renderScaffold substitutes the {name} placeholder with gcpName and indents
// every non-empty line by two spaces so the body nests inside a resource block.
func renderScaffold(scaffold, gcpName string) string {
	body := strings.ReplaceAll(scaffold, "{name}", gcpName)
	lines := strings.Split(strings.TrimRight(body, "\n"), "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) != "" {
			lines[i] = "  " + line
		}
	}
	return strings.Join(lines, "\n") + "\n"
}

// providerTF is the fixed provider/required_providers preamble written to every
// export. The google provider reads project and region from variables so the
// scaffold validates without hardcoded account data.
const providerTF = `terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }
}

provider "google" {
  project = var.project
  region  = var.region
}
`

// variablesTF declares the project and region variables referenced by provider.tf.
const variablesTF = `variable "project" {
  description = "GCP project ID for exported resources"
  type        = string
}

variable "region" {
  description = "GCP region for exported resources"
  type        = string
}
`
