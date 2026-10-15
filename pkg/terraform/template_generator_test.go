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
)

func TestGenerateTaralizerTemplate(t *testing.T) {
	assets := []TaralizerAsset{
		{ID: "vm-1", Name: "Web Server", Technology: "web-server"},
	}
	boundaries := []TrustBoundary{
		{ID: "net-1", Name: "VPC", Technology: "vpc"},
	}

	out := GenerateTaralizerTemplate(assets, boundaries)

	assert.Contains(t, out, `name: "Generated from Terraform"`)
	// Asset entry with all fields.
	assert.Contains(t, out, "  - id: vm-1\n    name: Web Server\n    technology: web-server")
	assert.Contains(t, out, "criticality: high")
	// Boundary entry.
	assert.Contains(t, out, "  - id: net-1\n    name: VPC\n    technology: vpc")
	// Static sections are always present.
	assert.Contains(t, out, "trust_boundaries:")
	assert.Contains(t, out, "data_assets:")
	assert.Contains(t, out, "data_flows:")
}

func TestGenerateTaralizerTemplateEmpty(t *testing.T) {
	out := GenerateTaralizerTemplate(nil, nil)

	// No asset or boundary entries were emitted: the sections are empty.
	assert.Contains(t, out, "technical_assets:\n# Trust boundaries\ntrust_boundaries:")
	assert.NotContains(t, out, "criticality: high")
}

func TestGenerateFromTerraform(t *testing.T) {
	tfDir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(tfDir, "main.tf"), []byte(`
resource "google_compute_network" "vpc" {
  name                    = "test-vpc"
  auto_create_subnetworks = false
}

resource "google_sql_database_instance" "db" {
  name = "test-db"
}

resource "google_storage_bucket" "unmapped" {
  name = "test-bucket"
}
`), 0o644))

	out, err := GenerateFromTerraform(tfDir, googleProfilePath)
	require.NoError(t, err)

	// Mapped resources appear as assets/boundaries (names come from the profile).
	assert.Contains(t, out, "  - id: vpc\n    name: vpc\n    technology: vpc")
	assert.Contains(t, out, "  - id: db\n    name: Cloud SQL\n    technology: database-sql")
	// The unmapped resource must not produce an entry.
	assert.NotContains(t, out, "unmapped")
}

func TestGenerateFromTerraformErrors(t *testing.T) {
	// Non-existent Terraform directory.
	_, err := GenerateFromTerraform(filepath.Join(t.TempDir(), "missing"), googleProfilePath)
	require.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "failed to parse terraform files"),
		"unexpected error: %v", err)

	// Non-existent profile.
	tfDir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(tfDir, "main.tf"), []byte(`
resource "google_compute_network" "vpc" {
  name = "test-vpc"
}
`), 0o644))
	_, err = GenerateFromTerraform(tfDir, filepath.Join(t.TempDir(), "missing.yaml"))
	require.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "failed to load google profile"),
		"unexpected error: %v", err)
}
