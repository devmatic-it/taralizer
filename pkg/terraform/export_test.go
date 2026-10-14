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
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/devmatic-it/taralizer/pkg/taralizer"
)

const bankOfAnthosModel = "../../examples/gcp/bank_of_anthos.yaml"

// loadBankOfAnthos loads the Bank of Anthos example model used by the export
// acceptance tests.
func loadBankOfAnthos(t *testing.T) taralizer.Report {
	t.Helper()
	report, err := taralizer.Load(bankOfAnthosModel)
	if err != nil {
		t.Fatalf("load %s: %v", bankOfAnthosModel, err)
	}
	return report
}

// resourceFiles returns the exported per-resource .tf files in outDir, i.e.
// every file except provider.tf and variables.tf.
func resourceFiles(t *testing.T, outDir string) []string {
	t.Helper()
	entries, err := os.ReadDir(outDir)
	if err != nil {
		t.Fatalf("read %s: %v", outDir, err)
	}
	var files []string
	for _, entry := range entries {
		if entry.IsDir() || entry.Name() == "provider.tf" || entry.Name() == "variables.tf" {
			continue
		}
		files = append(files, entry.Name())
	}
	return files
}

// assertMarkerPrecedesResource checks that the single resource block in content
// is directly preceded by a taralizer marker comment.
func assertMarkerPrecedesResource(t *testing.T, file, content string) {
	t.Helper()
	lines := strings.Split(content, "\n")
	resourceIdx := -1
	for i, line := range lines {
		if strings.HasPrefix(line, "resource ") {
			resourceIdx = i
			break
		}
	}
	if resourceIdx == -1 {
		t.Fatalf("%s: no resource block found", file)
	}
	prev := ""
	for i := resourceIdx - 1; i >= 0; i-- {
		if strings.TrimSpace(lines[i]) != "" {
			prev = lines[i]
			break
		}
	}
	if !strings.HasPrefix(prev, AssetMarkerPrefix) && !strings.HasPrefix(prev, BoundaryMarkerPrefix) {
		t.Errorf("%s: resource block not preceded by a taralizer marker, got %q", file, prev)
	}
}

func TestExportToTerraformBankOfAnthos(t *testing.T) {
	report := loadBankOfAnthos(t)
	outDir := t.TempDir()

	refs, warnings, err := ExportToTerraform(report, googleProfilePath, outDir, false)
	if err != nil {
		t.Fatalf("ExportToTerraform: %v", err)
	}

	for _, base := range []string{"provider.tf", "variables.tf"} {
		if _, err := os.Stat(filepath.Join(outDir, base)); err != nil {
			t.Errorf("expected %s to be written: %v", base, err)
		}
	}

	files := resourceFiles(t, outDir)
	if len(files) < 8 {
		t.Fatalf("expected at least 8 resource files, got %d: %v", len(files), files)
	}

	for _, file := range files {
		content, err := os.ReadFile(filepath.Join(outDir, file))
		if err != nil {
			t.Fatalf("read %s: %v", file, err)
		}
		assertMarkerPrecedesResource(t, file, string(content))
		if strings.Count(string(content), "resource ") != 1 {
			t.Errorf("%s: expected exactly one resource block", file)
		}
	}

	if len(refs) != len(files) {
		t.Fatalf("expected %d refs, got %d", len(files), len(refs))
	}
	for id, address := range refs {
		parts := strings.SplitN(address, ".", 2)
		if len(parts) != 2 || parts[1] == "" {
			t.Errorf("ref %q: address %q is not of the form type.name", id, address)
			continue
		}
		if _, err := os.Stat(filepath.Join(outDir, parts[1]+".tf")); err != nil {
			t.Errorf("ref %q: expected file %s.tf for address %q", id, parts[1], address)
		}
	}

	if len(warnings) == 0 {
		t.Error("expected warnings for unmapped technologies (kubernetes-network-policies, cloud-services)")
	}
	for _, w := range warnings {
		t.Logf("warning: %s", w)
	}
}

func TestExportRefusesNonEmptyDir(t *testing.T) {
	report := loadBankOfAnthos(t)
	outDir := t.TempDir()

	existing := filepath.Join(outDir, "existing.tf")
	if err := os.WriteFile(existing, []byte("# user file\n"), 0644); err != nil {
		t.Fatalf("seed existing.tf: %v", err)
	}

	if _, _, err := ExportToTerraform(report, googleProfilePath, outDir, false); err == nil {
		t.Fatal("expected error for non-empty target dir without force")
	} else if !strings.Contains(err.Error(), "not empty") {
		t.Fatalf("expected 'not empty' error, got: %v", err)
	}
	if _, err := os.Stat(filepath.Join(outDir, "provider.tf")); !os.IsNotExist(err) {
		t.Error("expected no provider.tf to be written when the dir is refused")
	}

	if _, _, err := ExportToTerraform(report, googleProfilePath, outDir, true); err != nil {
		t.Fatalf("ExportToTerraform with force: %v", err)
	}
	if _, err := os.Stat(filepath.Join(outDir, "provider.tf")); err != nil {
		t.Errorf("expected provider.tf after forced export: %v", err)
	}
	content, err := os.ReadFile(existing)
	if err != nil || string(content) != "# user file\n" {
		t.Errorf("existing.tf modified by forced export: %v / %q", err, content)
	}
}

func TestExportUnmappedTechnologyWarnings(t *testing.T) {
	report := taralizer.Report{
		TechnicalAssets: []taralizer.TechnicalAsset{
			{Id: "mapped-asset", Name: "Mapped", Technology: "vpc"},
			{Id: "unmapped-asset", Name: "Unmapped", Technology: "kubernetes-pod"},
			{Id: "unknown-asset", Name: "Unknown", Technology: "no-such-tech"},
			{Id: "bare-asset", Name: "Bare"},
		},
	}
	outDir := t.TempDir()

	refs, warnings, err := ExportToTerraform(report, googleProfilePath, outDir, false)
	if err != nil {
		t.Fatalf("ExportToTerraform: %v", err)
	}

	if len(refs) != 1 || refs["mapped-asset"] == "" {
		t.Fatalf("expected exactly one ref for mapped-asset, got %v", refs)
	}

	wantSubstrings := []string{
		`unmapped-asset`, `kubernetes-pod`, "no terraform mapping",
		`unknown-asset`, `no-such-tech`, "not found in profile",
		`bare-asset`, "no technology set",
	}
	joined := strings.Join(warnings, "\n")
	for _, want := range wantSubstrings {
		if !strings.Contains(joined, want) {
			t.Errorf("warnings missing %q; got: %s", want, joined)
		}
	}

	files := resourceFiles(t, outDir)
	if len(files) != 1 || files[0] != "mapped_asset.tf" {
		t.Fatalf("expected only mapped_asset.tf, got %v", files)
	}
}

func TestExportDuplicateResourceNames(t *testing.T) {
	report := taralizer.Report{
		TechnicalAssets: []taralizer.TechnicalAsset{
			{Id: "dup", Name: "First", Technology: "vpc"},
			{Id: "dup", Name: "Second", Technology: "subnet"},
		},
	}
	outDir := t.TempDir()

	if _, _, err := ExportToTerraform(report, googleProfilePath, outDir, false); err == nil {
		t.Fatal("expected error for duplicate HCL resource names")
	} else if !strings.Contains(err.Error(), "duplicate HCL resource name") {
		t.Fatalf("expected duplicate-name error, got: %v", err)
	}

	entries, err := os.ReadDir(outDir)
	if err != nil {
		t.Fatalf("read outDir: %v", err)
	}
	if len(entries) != 0 {
		t.Errorf("expected nothing to be written on duplicate names, got %d files", len(entries))
	}
}

func TestExportScaffoldSubstitution(t *testing.T) {
	report := taralizer.Report{
		TechnicalAssets: []taralizer.TechnicalAsset{
			{Id: "my-vpc", Name: "My VPC", Technology: "vpc"},
		},
	}
	outDir := t.TempDir()

	refs, _, err := ExportToTerraform(report, googleProfilePath, outDir, false)
	if err != nil {
		t.Fatalf("ExportToTerraform: %v", err)
	}

	if want := "google_compute_network.my_vpc"; refs["my-vpc"] != want {
		t.Fatalf("expected ref %q, got %q", want, refs["my-vpc"])
	}

	content, err := os.ReadFile(filepath.Join(outDir, "my_vpc.tf"))
	if err != nil {
		t.Fatalf("read my_vpc.tf: %v", err)
	}
	text := string(content)
	for _, want := range []string{
		"# taralizer:asset=my-vpc",
		`resource "google_compute_network" "my_vpc" {`,
		// The HCL label keeps the underscored name, but the GCP resource
		// name (scaffold {name}) must be hyphenated: GCP forbids underscores.
		`name                    = "my-vpc"`,
	} {
		if !strings.Contains(text, want) {
			t.Errorf("my_vpc.tf missing %q; got:\n%s", want, text)
		}
	}
	if strings.Contains(text, "{name}") {
		t.Errorf("my_vpc.tf still contains unsubstituted placeholder:\n%s", text)
	}
}

// writeBackFixture is a minimal model exercising both target sections, an
// inline comment that must survive the edit, and a data_assets section whose
// items must never be touched.
const writeBackFixture = `title: Test Model
technical_assets:
  - id: vpc
    name: VPC
    technology: vpc # values from google.yaml

  - id: unmapped-thing
    name: Unmapped
    technology: kubernetes-pod

trust_boundaries:
  - id: cloud
    name: Cloud boundary

data_assets:
  - id: should-not-touch
    name: Data
`

func writeModelFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "model.yaml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write model: %v", err)
	}
	return path
}

// TestGcpResourceName is table-driven: GCP name rules (no underscores, leading
// letter) differ from HCL label rules, so the two sanitizers must diverge.
func TestGcpResourceName(t *testing.T) {
	cases := []struct {
		id   string
		want string
	}{
		{"virtual_network", "virtual-network"},
		{"my-vpc", "my-vpc"},
		{"My_VPC-01", "my-vpc-01"},
		{"ops_services", "ops-services"},
		{"a__b--c", "a-b-c"},
		{"-leading-trailing-", "leading-trailing"},
		{"123start", "123start"}, // GCP rules vary per resource type; we only guarantee shape
		{"", "resource"},
		{"!!!", "resource"},
	}
	for _, tc := range cases {
		if got := gcpResourceName(tc.id); got != tc.want {
			t.Errorf("gcpResourceName(%q) = %q, want %q", tc.id, got, tc.want)
		}
	}
}

func TestWriteBackTerraformRefsInserts(t *testing.T) {
	path := writeModelFile(t, writeBackFixture)

	refs := map[string]string{
		"vpc":   "google_compute_network.vpc",
		"cloud": "google_project.cloud",
	}
	written, err := WriteBackTerraformRefs(path, refs)
	if err != nil {
		t.Fatalf("WriteBackTerraformRefs: %v", err)
	}
	if written != 2 {
		t.Fatalf("expected 2 lines written, got %d", written)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read model: %v", err)
	}
	text := string(content)

	wantLines := []string{
		"  - id: vpc\n    terraform_ref: google_compute_network.vpc\n    name: VPC",
		"  - id: cloud\n    terraform_ref: google_project.cloud\n    name: Cloud boundary",
	}
	for _, want := range wantLines {
		if !strings.Contains(text, want) {
			t.Errorf("model missing %q; got:\n%s", want, text)
		}
	}

	// The ref must be inserted directly below the "- id:" line.
	lines := strings.Split(text, "\n")
	for i, l := range lines {
		if strings.TrimSpace(l) == "- id: vpc" && i+1 < len(lines) {
			if lines[i+1] != "    terraform_ref: google_compute_network.vpc" {
				t.Errorf("terraform_ref not directly below id line, got %q", lines[i+1])
			}
		}
	}

	// Untouched content must survive verbatim: inline comment, blank lines,
	// the unmapped item, and the data_assets section.
	for _, want := range []string{
		"technology: vpc # values from google.yaml",
		"  - id: unmapped-thing\n    name: Unmapped",
		"data_assets:\n  - id: should-not-touch\n    name: Data",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("model lost %q; got:\n%s", want, text)
		}
	}

	// Exactly one terraform_ref per exported object, none elsewhere.
	if got := strings.Count(text, "terraform_ref:"); got != 2 {
		t.Errorf("expected exactly 2 terraform_ref lines, got %d:\n%s", got, text)
	}

	// The edited document must still load as a valid model with refs set.
	report, err := taralizer.Load(path)
	if err != nil {
		t.Fatalf("Load edited model: %v", err)
	}
	if report.TechnicalAssets[0].TerraformRef != "google_compute_network.vpc" {
		t.Errorf("asset vpc TerraformRef = %q", report.TechnicalAssets[0].TerraformRef)
	}
	if report.TrustBoundaries[0].TerraformRef != "google_project.cloud" {
		t.Errorf("boundary cloud TerraformRef = %q", report.TrustBoundaries[0].TerraformRef)
	}
}

func TestWriteBackTerraformRefsUpdatesExisting(t *testing.T) {
	fixture := strings.Replace(writeBackFixture,
		"  - id: vpc\n    name: VPC",
		"  - id: vpc\n    terraform_ref: google_compute_network.stale\n    name: VPC", 1)
	path := writeModelFile(t, fixture)

	// Re-running with the same ref must be idempotent (byte-identical file).
	refs := map[string]string{"vpc": "google_compute_network.stale"}
	if _, err := WriteBackTerraformRefs(path, refs); err != nil {
		t.Fatalf("first WriteBackTerraformRefs: %v", err)
	}
	afterFirst, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read model: %v", err)
	}

	if _, err := WriteBackTerraformRefs(path, refs); err != nil {
		t.Fatalf("second WriteBackTerraformRefs: %v", err)
	}
	afterSecond, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read model: %v", err)
	}
	if string(afterFirst) != string(afterSecond) {
		t.Error("write-back is not idempotent; second run changed the file")
	}

	// A changed ref must update in place, not add a second line.
	refs["vpc"] = "google_compute_network.fresh"
	if _, err := WriteBackTerraformRefs(path, refs); err != nil {
		t.Fatalf("third WriteBackTerraformRefs: %v", err)
	}
	text, _ := os.ReadFile(path)
	if strings.Count(string(text), "terraform_ref:") != 1 {
		t.Errorf("expected exactly 1 terraform_ref line, got:\n%s", text)
	}
	if !strings.Contains(string(text), "    terraform_ref: google_compute_network.fresh\n") {
		t.Errorf("ref not updated in place; got:\n%s", text)
	}
}

func TestWriteBackTerraformRefsMissingID(t *testing.T) {
	path := writeModelFile(t, writeBackFixture)
	before, _ := os.ReadFile(path)

	refs := map[string]string{
		"vpc":      "google_compute_network.vpc",
		"ghost-id": "google_project.ghost",
	}
	if _, err := WriteBackTerraformRefs(path, refs); err == nil {
		t.Fatal("expected error for ref id missing from model")
	} else if !strings.Contains(err.Error(), "ghost-id") {
		t.Fatalf("error should name the missing id, got: %v", err)
	}

	after, _ := os.ReadFile(path)
	if string(before) != string(after) {
		t.Error("model file modified despite missing ref id")
	}
}

func TestWriteBackTerraformRefsEmpty(t *testing.T) {
	path := writeModelFile(t, writeBackFixture)
	before, _ := os.ReadFile(path)

	written, err := WriteBackTerraformRefs(path, nil)
	if err != nil || written != 0 {
		t.Fatalf("expected (0, nil) for empty refs, got (%d, %v)", written, err)
	}
	after, _ := os.ReadFile(path)
	if string(before) != string(after) {
		t.Error("model file modified for empty refs")
	}
}

func TestWriteBackTerraformRefsQuotedAndCommentedIDs(t *testing.T) {
	fixture := `technical_assets:
  - id: "quoted-id" # the asset
    name: Quoted
`
	path := writeModelFile(t, fixture)

	refs := map[string]string{"quoted-id": "google_project.quoted_id"}
	if _, err := WriteBackTerraformRefs(path, refs); err != nil {
		t.Fatalf("WriteBackTerraformRefs: %v", err)
	}

	text, _ := os.ReadFile(path)
	if !strings.Contains(string(text), "  - id: \"quoted-id\" # the asset\n    terraform_ref: google_project.quoted_id\n") {
		t.Errorf("quoted/commented id not handled; got:\n%s", text)
	}
}

func TestExportAndWriteBackEndToEnd(t *testing.T) {
	// Copy the example model so the repo fixture is never modified.
	src, err := os.ReadFile(bankOfAnthosModel)
	if err != nil {
		t.Fatalf("read example model: %v", err)
	}
	path := writeModelFile(t, string(src))

	report := loadBankOfAnthos(t)
	outDir := filepath.Join(t.TempDir(), "tf")
	refs, _, err := ExportToTerraform(report, googleProfilePath, outDir, false)
	if err != nil {
		t.Fatalf("ExportToTerraform: %v", err)
	}
	if len(refs) == 0 {
		t.Fatal("expected at least one exported ref")
	}

	written, err := WriteBackTerraformRefs(path, refs)
	if err != nil {
		t.Fatalf("WriteBackTerraformRefs: %v", err)
	}
	if written != len(refs) {
		t.Fatalf("expected %d lines written, got %d", len(refs), written)
	}

	updated, err := taralizer.Load(path)
	if err != nil {
		t.Fatalf("Load edited model: %v", err)
	}

	check := func(kind, id, got string) {
		t.Helper()
		want, ok := refs[id]
		if ok && got != want {
			t.Errorf("%s %q: TerraformRef = %q, want %q", kind, id, got, want)
		}
		if !ok && got != "" {
			t.Errorf("%s %q: unexpected TerraformRef %q", kind, id, got)
		}
	}
	for _, asset := range updated.TechnicalAssets {
		check("asset", asset.Id, asset.TerraformRef)
	}
	for _, boundary := range updated.TrustBoundaries {
		check("boundary", boundary.Id, boundary.TerraformRef)
	}

	// Inline comments from the original model must survive the surgical edit.
	text, _ := os.ReadFile(path)
	if !strings.Contains(string(text), "# values: public, internal") {
		t.Error("inline comments lost during write-back")
	}
}
