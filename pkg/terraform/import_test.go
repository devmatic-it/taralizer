// Copyright 2021 taralizer authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
)

const importTestProfile = `name: Test Cloud
description: test profile
terraform_provider: google
technologies:
  - id: "vpc"
    name: "VPC network"
    type: "trust_boundary"
    terraform: "google_compute_network"
  - id: "web-server"
    name: "Web server"
    type: "technical_asset"
    terraform: "google_compute_instance"
  - id: "database-sql"
    name: "SQL database"
    type: "technical_asset"
    terraform: "google_sql_database_instance"
`

// writeImportFixture writes the given .tf files (name -> content) into a temp
// dir and returns the dir path plus the profile path.
func writeImportFixture(t *testing.T, files map[string]string) (string, string) {
	t.Helper()
	dir := t.TempDir()
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0644); err != nil {
			t.Fatalf("write fixture %s: %v", name, err)
		}
	}
	profilePath := filepath.Join(t.TempDir(), "profile.yaml")
	if err := os.WriteFile(profilePath, []byte(importTestProfile), 0644); err != nil {
		t.Fatalf("write profile: %v", err)
	}
	return dir, profilePath
}

func TestImportFromDir(t *testing.T) {
	cases := []struct {
		name        string
		files       map[string]string
		wantAssets  int
		wantBounds  int
		wantWarnings []string
		wantErr     string
	}{
		{
			name: "mapped_resources_become_assets_and_boundaries",
			files: map[string]string{
				"network.tf": `resource "google_compute_network" "vpc" {
  name = "vpc"
}
`,
				"vm.tf": `resource "google_compute_instance" "web1" {
  name = "web1"
}

resource "google_sql_database_instance" "db1" {
  name = "db1"
}
`,
			},
			wantAssets: 2,
			wantBounds: 1,
		},
		{
			name: "unmapped_type_warns_and_is_skipped",
			files: map[string]string{
				"bucket.tf": `resource "google_storage_bucket" "b1" {
  name = "b1"
}

resource "google_compute_network" "vpc" {
  name = "vpc"
}
`,
			},
			wantAssets:   0,
			wantBounds:   1,
			wantWarnings: []string{`unmapped resource type "google_storage_bucket"`},
		},
		{
			name: "duplicate_address_warns_and_keeps_first",
			files: map[string]string{
				"a.tf": `resource "google_compute_network" "vpc" { name = "a" }
`,
				"b.tf": `resource "google_compute_network" "vpc" { name = "b" }
`,
			},
			wantAssets:   0,
			wantBounds:   1,
			wantWarnings: []string{"duplicate resource address google_compute_network.vpc"},
		},
		{
			name: "non_resource_blocks_are_ignored",
			files: map[string]string{
				"misc.tf": `locals { env = "prod" }

provider "google" { region = "us-central1" }

data "google_compute_network" "lookup" { name = "n" }

terraform { required_version = ">= 1.0" }

resource "google_compute_network" "vpc" { name = "vpc" }
`,
			},
			wantAssets: 0,
			wantBounds: 1,
		},
		{
			name: "syntax_error_aborts",
			files: map[string]string{
				"broken.tf": `resource "google_compute_network" { name = }
`,
			},
			wantErr: "syntax error",
		},
		{
			name: "empty_dir_yields_empty_report",
			files: map[string]string{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir, profilePath := writeImportFixture(t, tc.files)

			report, warnings, err := ImportFromDir(dir, profilePath)
			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("expected error containing %q, got: %v", tc.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(report.TechnicalAssets) != tc.wantAssets {
				t.Errorf("assets = %d, want %d: %+v", len(report.TechnicalAssets), tc.wantAssets, report.TechnicalAssets)
			}
			if len(report.TrustBoundaries) != tc.wantBounds {
				t.Errorf("boundaries = %d, want %d: %+v", len(report.TrustBoundaries), tc.wantBounds, report.TrustBoundaries)
			}

			for _, want := range tc.wantWarnings {
				found := false
				for _, w := range warnings {
					if strings.Contains(w, want) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("missing warning containing %q; got: %v", want, warnings)
				}
			}
		})
	}
}

func TestImportFromDirPopulatesFields(t *testing.T) {
	dir, profilePath := writeImportFixture(t, map[string]string{
		"all.tf": `resource "google_compute_network" "vpc" { name = "vpc" }
resource "google_compute_instance" "web1" { name = "web1" }
`,
	})

	report, _, err := ImportFromDir(dir, profilePath)
	if err != nil {
		t.Fatalf("ImportFromDir: %v", err)
	}

	if len(report.TrustBoundaries) != 1 {
		t.Fatalf("boundaries = %d, want 1", len(report.TrustBoundaries))
	}
	b := report.TrustBoundaries[0]
	if b.Id != "vpc" || b.Name != "VPC network" || b.Technology != "vpc" || b.TerraformRef != "google_compute_network.vpc" {
		t.Errorf("boundary = %+v, want Id=vpc Name=%q Technology=vpc TerraformRef=google_compute_network.vpc", b, "VPC network")
	}

	if len(report.TechnicalAssets) != 1 {
		t.Fatalf("assets = %d, want 1", len(report.TechnicalAssets))
	}
	a := report.TechnicalAssets[0]
	if a.Id != "web1" || a.Name != "Web server" || a.Technology != "web-server" || a.TerraformRef != "google_compute_instance.web1" {
		t.Errorf("asset = %+v, want Id=web1 Name=%q Technology=web-server TerraformRef=google_compute_instance.web1", a, "Web server")
	}
}

func TestImportFromDirErrors(t *testing.T) {
	// A valid profile for the cases where the directory is the problem.
	_, goodProfile := writeImportFixture(t, nil)

	cases := []struct {
		name    string
		dir     string
		profile string
		wantErr string
	}{
		{
			name:    "missing_profile",
			dir:     t.TempDir(),
			profile: filepath.Join(t.TempDir(), "nope.yaml"),
			wantErr: "load profile",
		},
		{
			name:    "missing_dir",
			dir:     filepath.Join(t.TempDir(), "nope"),
			profile: goodProfile,
			wantErr: "stat",
		},
		{
			name:    "dir_is_a_file",
			dir:     filepath.Join(t.TempDir(), "file.tf"),
			profile: goodProfile,
			wantErr: "is not a directory",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name == "dir_is_a_file" {
				if err := os.WriteFile(tc.dir, []byte("x"), 0644); err != nil {
					t.Fatalf("seed file: %v", err)
				}
			}

			_, _, err := ImportFromDir(tc.dir, tc.profile)
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("expected error containing %q, got: %v", tc.wantErr, err)
			}
		})
	}
}
