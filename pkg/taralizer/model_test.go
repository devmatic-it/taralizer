// Package main Debian CVE Tracker Analyzer
// Copyright 2019 debcvescan authors
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
package taralizer

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestLoad(t *testing.T) {
	report, err := Load("../../examples/gcp/bank_of_anthos.yaml")
	if err != nil {
		t.Fatalf("Failed to load profile: %v", err)
	}
	assert.Greater(t, len(report.DataAssets), 0)
	assert.Greater(t, len(report.TechnicalAssets), 0)
	assert.Greater(t, len(report.TrustBoundaries), 0)
}

func TestSaveRoundTrip(t *testing.T) {
	report := Report{
		Title: "Save Test",
		TechnicalAssets: []TechnicalAsset{{
			Id: "web1", Name: "Web server", Technology: "web-server",
			TerraformRef: "google_compute_instance.web1",
		}},
		TrustBoundaries: []TrustBoundary{{
			Id: "vpc", Name: "VPC network", Technology: "vpc",
			TerraformRef: "google_compute_network.vpc",
		}},
	}

	path := filepath.Join(t.TempDir(), "model.yaml")
	if err := Save(report, path); err != nil {
		t.Fatalf("Save: %v", err)
	}

	roundTripped, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	assert.Equal(t, "Save Test", roundTripped.Title)
	assert.Len(t, roundTripped.TechnicalAssets, 1)
	assert.Equal(t, "google_compute_instance.web1", roundTripped.TechnicalAssets[0].TerraformRef)
	assert.Len(t, roundTripped.TrustBoundaries, 1)
	assert.Equal(t, "google_compute_network.vpc", roundTripped.TrustBoundaries[0].TerraformRef)
}

func TestSaveMissingDir(t *testing.T) {
	err := Save(Report{}, filepath.Join(t.TempDir(), "nope", "model.yaml"))
	assert.Error(t, err)
}

func TestLoadProfileSet(t *testing.T) {
	profileSet, err := LoadProfileSet("default.yaml")
	if err != nil {
		t.Fatalf("Failed to load profile set: %v", err)
	}
	assert.Greater(t, len(profileSet.Technologies), 0)
}

func TestTerraformRefRoundTrip(t *testing.T) {
	report, err := Load("../../examples/gcp/bank_of_anthos.yaml")
	if err != nil {
		t.Fatalf("Failed to load report: %v", err)
	}

	report.TechnicalAssets[0].TerraformRef = "google_compute_instance.frontend"
	report.TrustBoundaries[0].TerraformRef = "google_compute_network.vpc"

	data, err := yaml.Marshal(report)
	if err != nil {
		t.Fatalf("Failed to marshal report: %v", err)
	}

	var roundTripped Report
	if err := yaml.Unmarshal(data, &roundTripped); err != nil {
		t.Fatalf("Failed to unmarshal report: %v", err)
	}

	assert.Equal(t, "google_compute_instance.frontend", roundTripped.TechnicalAssets[0].TerraformRef)
	assert.Equal(t, "google_compute_network.vpc", roundTripped.TrustBoundaries[0].TerraformRef)
}

func TestTerraformRefOmitEmpty(t *testing.T) {
	report := Report{
		Title:           "test",
		TechnicalAssets: []TechnicalAsset{{Id: "a1", Name: "asset one"}},
	}

	data, err := yaml.Marshal(report)
	if err != nil {
		t.Fatalf("Failed to marshal report: %v", err)
	}

	assert.NotContains(t, string(data), "terraform_ref")
}
