// Drift detection tests (Phase A3.1).
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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/devmatic-it/taralizer/pkg/taralizer"
)

// writeDriftFixture writes the model and tf files into a fresh temp dir and
// returns (modelPath, tfDir).
func writeDriftFixture(t *testing.T, model taralizer.Report, tfFiles map[string]string) (string, string) {
	t.Helper()

	dir := t.TempDir()
	modelPath := filepath.Join(dir, "model.yaml")
	require.NoError(t, taralizer.Save(model, modelPath))

	tfDir := filepath.Join(dir, "tf")
	require.NoError(t, os.MkdirAll(tfDir, 0755))
	for name, content := range tfFiles {
		require.NoError(t, os.WriteFile(filepath.Join(tfDir, name), []byte(content), 0644))
	}
	return modelPath, tfDir
}

func TestDetectDrift(t *testing.T) {
	const vmTF = `# taralizer:asset=vm
resource "google_compute_instance" "vm" {
  name         = "vm"
  machine_type = "e2-medium"
}
`
	const vpcTF = `# taralizer:boundary=vpc
resource "google_compute_network" "vpc" {
  name = "vpc"
}
`

	cases := []struct {
		name          string
		model         taralizer.Report
		tfFiles       map[string]string
		wantMissing   []DriftItem
		wantBroken    []DriftItem
		wantUnmodeled []DriftItem
		wantStale     []DriftItem
		wantWarnings  int
	}{
		{
			name: "in_sync_model_and_tf_produce_no_findings",
			model: taralizer.Report{
				TechnicalAssets: []taralizer.TechnicalAsset{{Id: "vm", Technology: "compute-instance", TerraformRef: "google_compute_instance.vm"}},
				TrustBoundaries: []taralizer.TrustBoundary{{Id: "vpc", Technology: "network", TerraformRef: "google_compute_network.vpc"}},
			},
			tfFiles: map[string]string{"vm.tf": vmTF, "vpc.tf": vpcTF},
		},
		{
			name: "model_ref_without_tf_resource_is_missing",
			model: taralizer.Report{
				TechnicalAssets: []taralizer.TechnicalAsset{{Id: "vm", Technology: "compute-instance", TerraformRef: "google_compute_instance.vm"}},
			},
			tfFiles: map[string]string{},
			wantMissing: []DriftItem{
				{ID: "vm", Ref: "google_compute_instance.vm"},
			},
		},
		{
			name: "malformed_refs_are_broken_not_missing",
			model: taralizer.Report{
				TechnicalAssets: []taralizer.TechnicalAsset{
					{Id: "no-dot", Technology: "compute-instance", TerraformRef: "google_compute_instance"},
					{Id: "empty-type", Technology: "compute-instance", TerraformRef: ".vm"},
					{Id: "empty-name", Technology: "compute-instance", TerraformRef: "google_compute_instance."},
					{Id: "two-dots", Technology: "compute-instance", TerraformRef: "a.b.c"},
				},
			},
			tfFiles: map[string]string{},
			wantBroken: []DriftItem{
				{ID: "no-dot", Ref: "google_compute_instance"},
				{ID: "empty-type", Ref: ".vm"},
				{ID: "empty-name", Ref: "google_compute_instance."},
				{ID: "two-dots", Ref: "a.b.c"},
			},
		},
		{
			name: "tf_resource_without_model_ref_is_unmodeled",
			model: taralizer.Report{
				TrustBoundaries: []taralizer.TrustBoundary{{Id: "vpc", Technology: "network", TerraformRef: "google_compute_network.vpc"}},
			},
			// vm.tf has no marker here so the case isolates Unmodeled; a
			// marker for an id absent from the model would also be stale.
			tfFiles: map[string]string{
				"vm.tf":  "resource \"google_compute_instance\" \"vm\" {\n  name = \"vm\"\n}\n",
				"vpc.tf": vpcTF,
			},
			wantUnmodeled: []DriftItem{
				{ID: "google_compute_instance.vm", Location: filepath.Join("tf", "vm.tf") + ":1"},
			},
		},
		{
			name: "marker_id_absent_from_model_is_stale",
			model: taralizer.Report{
				TechnicalAssets: []taralizer.TechnicalAsset{{Id: "vm", Technology: "compute-instance", TerraformRef: "google_compute_instance.vm"}},
			},
			tfFiles: map[string]string{
				"vm.tf":  vmTF,
				"old.tf": "# taralizer:asset=old-vm\nresource \"google_compute_instance\" \"old_vm\" {\n  name = \"old\"\n}\n",
			},
			wantStale: []DriftItem{
				{ID: "old-vm", Location: filepath.Join("tf", "old.tf") + ":1 (asset marker)"},
			},
			// old_vm is also unmodeled; vm.tf's marker matches the model.
			wantUnmodeled: []DriftItem{
				{ID: "google_compute_instance.old_vm", Location: filepath.Join("tf", "old.tf") + ":2"},
			},
		},
		{
			name: "entries_without_terraform_ref_are_not_checked",
			model: taralizer.Report{
				TechnicalAssets: []taralizer.TechnicalAsset{{Id: "manual", Technology: "compute-instance"}},
			},
			tfFiles: map[string]string{},
		},
		{
			name:  "duplicate_tf_address_warns_and_reports_each_occurrence_when_unmodeled",
			model: taralizer.Report{},
			tfFiles: map[string]string{
				"a.tf": "resource \"google_compute_disk\" \"d\" {\n  name = \"d\"\n}\n",
				"b.tf": "resource \"google_compute_disk\" \"d\" {\n  name = \"d\"\n}\n",
			},
			wantUnmodeled: []DriftItem{
				{ID: "google_compute_disk.d", Location: filepath.Join("tf", "a.tf") + ":1"},
				{ID: "google_compute_disk.d", Location: filepath.Join("tf", "b.tf") + ":1"},
			},
			wantWarnings: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			modelPath, tfDir := writeDriftFixture(t, tc.model, tc.tfFiles)

			// Expected locations are written relative to the fixture root;
			// DetectDrift reports absolute paths.
			root := filepath.Dir(tfDir)
			withRoot := func(items []DriftItem) []DriftItem {
				if items == nil {
					return nil // keep nil vs empty distinct for assert.Equal
				}
				out := make([]DriftItem, len(items))
				for i, it := range items {
					out[i] = it
					if it.Location != "" {
						out[i].Location = filepath.Join(root, it.Location)
					}
				}
				return out
			}

			report, err := DetectDrift(modelPath, tfDir, nil)
			require.NoError(t, err)

			assert.Equal(t, tc.wantMissing, report.MissingAssets)
			assert.Equal(t, tc.wantBroken, report.BrokenRefs)
			assert.Equal(t, withRoot(tc.wantUnmodeled), report.Unmodeled)
			assert.Equal(t, withRoot(tc.wantStale), report.StaleMarkers)
			assert.Len(t, report.Warnings, tc.wantWarnings)
		})
	}
}

func TestDetectDriftErrors(t *testing.T) {
	t.Run("missing_model_file", func(t *testing.T) {
		dir := t.TempDir()
		tfDir := filepath.Join(dir, "tf")
		require.NoError(t, os.MkdirAll(tfDir, 0755))

		_, err := DetectDrift(filepath.Join(dir, "nope.yaml"), tfDir, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "load model")
	})

	t.Run("missing_tf_dir", func(t *testing.T) {
		dir := t.TempDir()
		modelPath := filepath.Join(dir, "model.yaml")
		require.NoError(t, taralizer.Save(taralizer.Report{}, modelPath))

		_, err := DetectDrift(modelPath, filepath.Join(dir, "nope"), nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "list terraform files")
	})

	t.Run("tf_path_is_a_file", func(t *testing.T) {
		dir := t.TempDir()
		modelPath := filepath.Join(dir, "model.yaml")
		require.NoError(t, taralizer.Save(taralizer.Report{}, modelPath))

		tfFile := filepath.Join(dir, "main.tf")
		require.NoError(t, os.WriteFile(tfFile, []byte("resource \"a\" \"b\" {}\n"), 0644))

		_, err := DetectDrift(modelPath, tfFile, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "not a directory")
	})

	t.Run("syntax_error_aborts", func(t *testing.T) {
		dir := t.TempDir()
		modelPath := filepath.Join(dir, "model.yaml")
		require.NoError(t, taralizer.Save(taralizer.Report{}, modelPath))

		tfDir := filepath.Join(dir, "tf")
		require.NoError(t, os.MkdirAll(tfDir, 0755))
		require.NoError(t, os.WriteFile(filepath.Join(tfDir, "bad.tf"), []byte("resource \"a\" \"b\" { oops\n"), 0644))

		_, err := DetectDrift(modelPath, tfDir, nil)
		require.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "syntax error"), "got: %v", err)
	})
}

func TestDriftReportMarkdown(t *testing.T) {
	t.Run("clean_report", func(t *testing.T) {
		r := DriftReport{ModelPath: "model.yaml", TFDir: "./tf"}
		want := `# Taralizer Drift Report

- Model: ` + "`model.yaml`" + `
- Terraform dir: ` + "`./tf`" + `
- Findings: 0

No drift detected.
`
		assert.Equal(t, want, r.Markdown())
	})

	t.Run("all_categories_rendered_with_tables", func(t *testing.T) {
		r := DriftReport{
			ModelPath: "model.yaml",
			TFDir:     "./tf",
			MissingAssets: []DriftItem{
				{ID: "vm", Ref: "google_compute_instance.vm"},
				{ID: "vpc", Ref: "google_compute_network.vpc"},
			},
			BrokenRefs:   []DriftItem{{ID: "x", Ref: "bad-ref"}},
			Unmodeled:    []DriftItem{{ID: "google_compute_disk.d", Location: "tf/a.tf:1"}},
			StaleMarkers: []DriftItem{{ID: "old-vm", Location: "tf/old.tf:1 (asset marker)"}},
			Warnings:     []string{"duplicate resource address google_compute_disk.d in tf/a.tf and tf/b.tf: using the first occurrence"},
		}

		got := r.Markdown()
		for _, want := range []string{
			"# Taralizer Drift Report",
			"- Findings: 5",
			"## Missing assets (2)",
			"| ID | terraform_ref |",
			"| `vm` | `google_compute_instance.vm` |",
			"| `vpc` | `google_compute_network.vpc` |",
			"## Broken refs (1)",
			"| `x` | `bad-ref` |",
			"## Unmodeled resources (1)",
			"| Resource | Location |",
			"| `google_compute_disk.d` | `tf/a.tf:1` |",
			"## Stale markers (1)",
			"| Marker ID | Location |",
			"| `old-vm` | `tf/old.tf:1 (asset marker)` |",
			"## Warnings (1)",
			"- duplicate resource address google_compute_disk.d in tf/a.tf and tf/b.tf: using the first occurrence",
		} {
			assert.Contains(t, got, want)
		}
	})

	t.Run("empty_categories_omitted", func(t *testing.T) {
		r := DriftReport{
			ModelPath:     "m.yaml",
			TFDir:         "./tf",
			MissingAssets: []DriftItem{{ID: "vm", Ref: "a.b"}},
		}

		got := r.Markdown()
		assert.Contains(t, got, "## Missing assets (1)")
		for _, absent := range []string{"Broken refs", "Unmodeled resources", "Stale markers", "Warnings"} {
			assert.NotContains(t, got, absent)
		}
	})
}

// TestDetectDriftAttribute covers attribute-level drift (Phase A4): model
// attributes declared for a technology are compared against the static
// Terraform literal at each profile-declared path. It uses the real google
// profile so attribute relevance comes from profiles/google.yaml.
func TestDetectDriftAttribute(t *testing.T) {
	profile, err := LoadGoogleProfile(googleProfilePath)
	require.NoError(t, err)

	const vmTF = "resource \"google_compute_instance\" \"vm\" {\n  name                = \"vm\"\n  can_ip_forward      = true\n  allow_stopped_state = false\n}\n"

	cases := []struct {
		name          string
		model         taralizer.Report
		tfFiles       map[string]string
		wantAttrDrift []DriftItem
		wantWarnings  int
		wantLoc       string // filename expected in the drift item's Location (Contains)
	}{
		{
			name: "attribute_match_no_drift",
			model: taralizer.Report{
				TechnicalAssets: []taralizer.TechnicalAsset{{
					Id: "vm", Technology: "web-server", TerraformRef: "google_compute_instance.vm",
					Attributes: map[string]string{"can_ip_forward": "true", "allow_stopped_state": "false"},
				}},
			},
			tfFiles: map[string]string{"vm.tf": vmTF},
		},
		{
			name: "attribute_mismatch_is_drift",
			model: taralizer.Report{
				TechnicalAssets: []taralizer.TechnicalAsset{{
					Id: "vm", Technology: "web-server", TerraformRef: "google_compute_instance.vm",
					Attributes: map[string]string{"can_ip_forward": "false"}, // TF has true
				}},
			},
			tfFiles: map[string]string{"vm.tf": vmTF},
			wantAttrDrift: []DriftItem{{
				ID: "vm", Ref: "google_compute_instance.vm", Attribute: "can_ip_forward",
				ModelValue: "false", TFValue: "true",
			}},
			wantLoc: "vm.tf",
		},
		{
			name: "attribute_non_literal_warns",
			model: taralizer.Report{
				TechnicalAssets: []taralizer.TechnicalAsset{{
					Id: "vm", Technology: "web-server", TerraformRef: "google_compute_instance.vm",
					Attributes: map[string]string{"can_ip_forward": "true"},
				}},
			},
			tfFiles: map[string]string{
				"vm.tf": "resource \"google_compute_instance\" \"vm\" {\n  can_ip_forward = var.enable\n}\n",
			},
			wantWarnings: 1, // non-literal value: warning, not a drift finding
		},
		{
			name: "undeclared_attribute_skipped",
			model: taralizer.Report{
				TechnicalAssets: []taralizer.TechnicalAsset{{
					Id: "vm", Technology: "web-server", TerraformRef: "google_compute_instance.vm",
				}}, // no Attributes declared
			},
			tfFiles: map[string]string{"vm.tf": vmTF},
		},
		{
			name: "nested_attribute_mismatch",
			model: taralizer.Report{
				TechnicalAssets: []taralizer.TechnicalAsset{{
					Id: "db", Technology: "database-sql", TerraformRef: "google_sql_database_instance.db",
					Attributes: map[string]string{"require_ssl": "false"}, // TF has true
				}},
			},
			tfFiles: map[string]string{
				"db.tf": "resource \"google_sql_database_instance\" \"db\" {\n  name = \"db\"\n  settings {\n    require_ssl = true\n  }\n}\n",
			},
			wantAttrDrift: []DriftItem{{
				ID: "db", Ref: "google_sql_database_instance.db", Attribute: "require_ssl",
				ModelValue: "false", TFValue: "true",
			}},
			wantLoc: "db.tf",
		},
		{
			name: "technology_without_attributes_skipped",
			model: taralizer.Report{
				TechnicalAssets: []taralizer.TechnicalAsset{{
					Id: "ds", Technology: "database-nosql", TerraformRef: "google_firestore_database.ds",
				}},
			},
			tfFiles: map[string]string{
				"ds.tf": "resource \"google_firestore_database\" \"ds\" {\n  name = \"ds\"\n}\n",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			modelPath, tfDir := writeDriftFixture(t, tc.model, tc.tfFiles)

			report, err := DetectDrift(modelPath, tfDir, profile)
			require.NoError(t, err)

			require.Equal(t, len(tc.wantAttrDrift), len(report.AttributeDrift))
			for i, want := range tc.wantAttrDrift {
				got := report.AttributeDrift[i]
				assert.Equal(t, want.ID, got.ID)
				assert.Equal(t, want.Ref, got.Ref)
				assert.Equal(t, want.Attribute, got.Attribute)
				assert.Equal(t, want.ModelValue, got.ModelValue)
				assert.Equal(t, want.TFValue, got.TFValue)
				if tc.wantLoc != "" {
					assert.Contains(t, got.Location, tc.wantLoc)
				}
			}
			if tc.wantWarnings > 0 {
				assert.GreaterOrEqual(t, len(report.Warnings), tc.wantWarnings)
			}
		})
	}
}

func TestValidTerraformRef(t *testing.T) {
	cases := []struct {
		ref  string
		want bool
	}{
		{"google_compute_instance.vm", true},
		{"a.b", true},
		{"no-dot", false},
		{".vm", false},
		{"type.", false},
		{"a.b.c", false},
		{"", false},
	}
	for _, tc := range cases {
		t.Run(tc.ref, func(t *testing.T) {
			assert.Equal(t, tc.want, validTerraformRef(tc.ref))
		})
	}
}
