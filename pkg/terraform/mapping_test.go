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
	"testing"

	"github.com/devmatic-it/taralizer/pkg/taralizer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const googleProfilePath = "../../profiles/google.yaml"

func TestLoadGoogleProfile(t *testing.T) {
	profile, err := LoadGoogleProfile(googleProfilePath)
	if err != nil {
		t.Fatalf("Failed to load google profile: %v", err)
	}

	assert.Equal(t, "Google Cloud", profile.Name)
	assert.Equal(t, "google", profile.TerraformProvider)
	assert.Greater(t, len(profile.Technologies), 0)
}

func TestGoogleProfileDualLoaderAgreement(t *testing.T) {
	profile, err := LoadGoogleProfile(googleProfilePath)
	if err != nil {
		t.Fatalf("Failed to load google profile: %v", err)
	}

	profileSet, err := taralizer.LoadProfileSet("google.yaml")
	if err != nil {
		t.Fatalf("Failed to load google profile set: %v", err)
	}

	assert.Equal(t, len(profile.Technologies), len(profileSet.Technologies))
	assert.Equal(t, profile.TerraformProvider, profileSet.TerraformProvider)

	// Both loaders must agree on the technology IDs.
	ids := make(map[string]bool, len(profile.Technologies))
	for _, tech := range profile.Technologies {
		ids[tech.ID] = true
	}
	for _, tech := range profileSet.Technologies {
		assert.True(t, ids[tech.Id], "technology %q missing from LoadProfileSet result", tech.Id)
	}
}

func TestScaffoldCoversAllMappedTechnologies(t *testing.T) {
	profile, err := LoadGoogleProfile(googleProfilePath)
	if err != nil {
		t.Fatalf("Failed to load google profile: %v", err)
	}

	for _, tech := range profile.Technologies {
		if tech.Terraform == "" {
			continue
		}
		assert.NotEmpty(t, tech.Scaffold, "technology %q maps to %s but has no scaffold", tech.ID, tech.Terraform)
	}
}

func TestLoadGoogleProfileErrors(t *testing.T) {
	// Non-existent file.
	_, err := LoadGoogleProfile(filepath.Join(t.TempDir(), "missing.yaml"))
	require.Error(t, err)

	// Malformed YAML.
	bad := filepath.Join(t.TempDir(), "bad.yaml")
	require.NoError(t, os.WriteFile(bad, []byte(": : not: valid: yaml: ["), 0o644))
	_, err = LoadGoogleProfile(bad)
	require.Error(t, err)
}

func TestMapResourceToTechnology(t *testing.T) {
	profile, err := LoadGoogleProfile(googleProfilePath)
	require.NoError(t, err)

	// Known technical asset mapping.
	tech := profile.MapResourceToTechnology("google_sql_database_instance")
	require.NotNil(t, tech)
	assert.Equal(t, "database-sql", tech.ID)
	assert.Equal(t, "technical_asset", tech.Type)

	// Known trust boundary mapping.
	tech = profile.MapResourceToTechnology("google_compute_network")
	require.NotNil(t, tech)
	assert.Equal(t, "vpc", tech.ID)
	assert.Equal(t, "trust_boundary", tech.Type)

	// Unknown resource type.
	assert.Nil(t, profile.MapResourceToTechnology("google_storage_bucket"))
}

func TestResourceToAssetMapping(t *testing.T) {
	profile, err := LoadGoogleProfile(googleProfilePath)
	require.NoError(t, err)

	resources := []GCPResource{
		{Name: "db-1", Type: "google_sql_database_instance"},
		{Name: "vpc-1", Type: "google_compute_network"},
		{Name: "bucket-1", Type: "google_storage_bucket"}, // unmapped -> warning only
	}

	assets, boundaries := profile.ResourceToAssetMapping(resources)

	require.Len(t, assets, 1)
	assert.Equal(t, "db-1", assets[0].ID)
	assert.Equal(t, "database-sql", assets[0].Technology)

	require.Len(t, boundaries, 1)
	assert.Equal(t, "vpc-1", boundaries[0].ID)
	assert.Equal(t, "vpc", boundaries[0].Technology)
}

func TestBankOfAnthosExportableCount(t *testing.T) {
	report, err := taralizer.Load("../../examples/gcp/bank_of_anthos.yaml")
	if err != nil {
		t.Fatalf("Failed to load bank of anthos model: %v", err)
	}

	profile, err := LoadGoogleProfile(googleProfilePath)
	if err != nil {
		t.Fatalf("Failed to load google profile: %v", err)
	}

	mapped := make(map[string]bool, len(profile.Technologies))
	for _, tech := range profile.Technologies {
		if tech.Terraform != "" {
			mapped[tech.ID] = true
		}
	}

	exportable := 0
	for _, asset := range report.TechnicalAssets {
		if mapped[asset.Technology] {
			exportable++
		}
	}
	for _, boundary := range report.TrustBoundaries {
		if mapped[boundary.Technology] {
			exportable++
		}
	}

	assert.GreaterOrEqual(t, exportable, 8)
}
