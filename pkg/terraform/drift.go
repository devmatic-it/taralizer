// Package terraform drift engine (Phase A3): structural drift detection
// between a taralizer model and a directory of Terraform files.
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
	"strings"

	"github.com/devmatic-it/taralizer/pkg/taralizer"
)

// DriftItem is one structural discrepancy between a taralizer model and a
// Terraform directory. Model-side findings (MissingAssets, BrokenRefs) carry
// the model entry id and its terraform_ref; Terraform-side findings
// (Unmodeled, StaleMarkers) carry a file:line location.
type DriftItem struct {
	ID       string // model entry id, or the resource address for Unmodeled
	Ref      string // terraform_ref (model-side findings); empty otherwise
	Location string // "file:line" for Terraform-side findings; empty otherwise
}

// DriftReport categorizes the structural drift between a model and a
// Terraform directory:
//
//   - MissingAssets:  model entries whose terraform_ref has no matching
//     resource block in the Terraform directory (deleted or renamed).
//   - BrokenRefs:     model entries whose terraform_ref is not a valid
//     type.name HCL address (corrupt model).
//   - Unmodeled:      resource blocks in the Terraform directory that no
//     model entry references (added outside the model).
//   - StaleMarkers:   taralizer export markers (# taralizer:asset=<id> /
//     # taralizer:boundary=<id>) whose id no longer exists in the model.
type DriftReport struct {
	ModelPath string
	TFDir     string

	MissingAssets []DriftItem
	BrokenRefs    []DriftItem
	Unmodeled     []DriftItem
	StaleMarkers  []DriftItem

	Warnings []string
}

// Count returns the total number of drift findings across all categories.
func (r DriftReport) Count() int {
	return len(r.MissingAssets) + len(r.BrokenRefs) + len(r.Unmodeled) + len(r.StaleMarkers)
}

// markerRef is one taralizer export marker found in a .tf file.
type markerRef struct {
	kind string // "asset" or "boundary"
	id   string
	file string
	line int // 1-based line of the marker comment
}

// DetectDrift compares the taralizer model at modelPath with the Terraform
// directory tfDir and returns a categorized drift report. Model entries
// without a terraform_ref have no Terraform linkage and are not checked.
// Files with syntax errors abort detection with an error, consistent with
// ImportFromDir.
func DetectDrift(modelPath, tfDir string) (DriftReport, error) {
	report := DriftReport{ModelPath: modelPath, TFDir: tfDir}

	model, err := taralizer.Load(modelPath)
	if err != nil {
		return report, fmt.Errorf("load model %s: %w", modelPath, err)
	}

	files, err := listTerraformFiles(tfDir)
	if err != nil {
		return report, fmt.Errorf("list terraform files: %w", err)
	}

	tfAddrs := map[string]string{} // HCL address -> file (first occurrence)
	type fileResources struct {
		file      string
		resources []resourceRef
	}
	var parsedFiles []fileResources
	var markers []markerRef

	for _, file := range files {
		resources, err := parseResourceBlocks(file)
		if err != nil {
			return report, fmt.Errorf("parse %s: %w", file, err)
		}
		parsedFiles = append(parsedFiles, fileResources{file: file, resources: resources})
		for _, res := range resources {
			addr := res.typ + "." + res.name
			if first, dup := tfAddrs[addr]; dup {
				report.Warnings = append(report.Warnings,
					fmt.Sprintf("duplicate resource address %s in %s and %s: using the first occurrence", addr, first, file))
				continue
			}
			tfAddrs[addr] = file
		}

		found, err := scanMarkers(file)
		if err != nil {
			return report, fmt.Errorf("scan markers in %s: %w", file, err)
		}
		markers = append(markers, found...)
	}

	// Model side.
	modelIDs := map[string]bool{}
	refOwners := map[string]int{} // HCL address -> number of model entries referencing it

	addEntry := func(id, ref string) {
		modelIDs[id] = true
		if ref == "" {
			return // no Terraform linkage; not drift-checkable
		}
		if !validTerraformRef(ref) {
			report.BrokenRefs = append(report.BrokenRefs, DriftItem{ID: id, Ref: ref})
			return
		}
		refOwners[ref]++
		if _, ok := tfAddrs[ref]; !ok {
			report.MissingAssets = append(report.MissingAssets, DriftItem{ID: id, Ref: ref})
		}
	}
	for _, asset := range model.TechnicalAssets {
		addEntry(asset.Id, asset.TerraformRef)
	}
	for _, boundary := range model.TrustBoundaries {
		addEntry(boundary.Id, boundary.TerraformRef)
	}

	// Terraform side: resources no model entry references.
	for _, pf := range parsedFiles {
		for _, res := range pf.resources {
			addr := res.typ + "." + res.name
			if refOwners[addr] == 0 {
				report.Unmodeled = append(report.Unmodeled, DriftItem{
					ID:       addr,
					Location: fmt.Sprintf("%s:%d", pf.file, res.line),
				})
			}
		}
	}

	// Stale markers: marker id absent from the model.
	for _, m := range markers {
		if !modelIDs[m.id] {
			report.StaleMarkers = append(report.StaleMarkers, DriftItem{
				ID:       m.id,
				Location: fmt.Sprintf("%s:%d (%s marker)", m.file, m.line, m.kind),
			})
		}
	}

	return report, nil
}

// scanMarkers returns the taralizer export markers in one .tf file, in line
// order. A marker is a comment line (any indentation) starting with
// AssetMarkerPrefix or BoundaryMarkerPrefix.
func scanMarkers(path string) ([]markerRef, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}

	var markers []markerRef
	for i, line := range strings.Split(string(content), "\n") {
		trimmed := strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(trimmed, AssetMarkerPrefix):
			markers = append(markers, markerRef{kind: "asset", id: strings.TrimPrefix(trimmed, AssetMarkerPrefix), file: path, line: i + 1})
		case strings.HasPrefix(trimmed, BoundaryMarkerPrefix):
			markers = append(markers, markerRef{kind: "boundary", id: strings.TrimPrefix(trimmed, BoundaryMarkerPrefix), file: path, line: i + 1})
		}
	}
	return markers, nil
}

// validTerraformRef reports whether ref is a well-formed HCL resource
// address: exactly one dot separating non-empty type and name parts.
func validTerraformRef(ref string) bool {
	parts := strings.Split(ref, ".")
	return len(parts) == 2 && parts[0] != "" && parts[1] != ""
}

// Markdown renders the report as a markdown document. The header always
// shows the model path, Terraform directory and total finding count; each
// category with at least one finding gets a section with a table, in the
// order Missing assets, Broken refs, Unmodeled resources, Stale markers.
// Warnings (e.g. duplicate resource addresses) are listed last when present.
// A report with zero findings renders the header followed by
// "No drift detected."
func (r DriftReport) Markdown() string {
	var b strings.Builder

	b.WriteString("# Taralizer Drift Report\n\n")
	fmt.Fprintf(&b, "- Model: `%s`\n", r.ModelPath)
	fmt.Fprintf(&b, "- Terraform dir: `%s`\n", r.TFDir)
	fmt.Fprintf(&b, "- Findings: %d\n\n", r.Count())

	if r.Count() == 0 {
		b.WriteString("No drift detected.\n")
		return b.String()
	}

	// refTable renders model-side findings (ID + terraform_ref).
	refTable := func(title string, items []DriftItem) {
		if len(items) == 0 {
			return
		}
		fmt.Fprintf(&b, "## %s (%d)\n\n", title, len(items))
		b.WriteString("| ID | terraform_ref |\n| --- | --- |\n")
		for _, it := range items {
			fmt.Fprintf(&b, "| `%s` | `%s` |\n", it.ID, it.Ref)
		}
		b.WriteString("\n")
	}

	// locTable renders Terraform-side findings (identifier + file:line).
	locTable := func(title, firstCol string, items []DriftItem) {
		if len(items) == 0 {
			return
		}
		fmt.Fprintf(&b, "## %s (%d)\n\n", title, len(items))
		fmt.Fprintf(&b, "| %s | Location |\n| --- | --- |\n", firstCol)
		for _, it := range items {
			fmt.Fprintf(&b, "| `%s` | `%s` |\n", it.ID, it.Location)
		}
		b.WriteString("\n")
	}

	refTable("Missing assets", r.MissingAssets)
	refTable("Broken refs", r.BrokenRefs)
	locTable("Unmodeled resources", "Resource", r.Unmodeled)
	locTable("Stale markers", "Marker ID", r.StaleMarkers)

	if len(r.Warnings) > 0 {
		fmt.Fprintf(&b, "## Warnings (%d)\n\n", len(r.Warnings))
		for _, w := range r.Warnings {
			fmt.Fprintf(&b, "- %s\n", w)
		}
		b.WriteString("\n")
	}

	return b.String()
}
