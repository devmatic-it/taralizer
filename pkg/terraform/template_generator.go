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
)

// GenerateTaralizerTemplate generates a taralizer YAML template with the parsed assets
func GenerateTaralizerTemplate(assets []TaralizerAsset, boundaries []TrustBoundary) string {
	template := `# Taralizer YAML Template Generated from Terraform
name: "Generated from Terraform"
description: "Automatically generated template based on Google Cloud resources"

# Default values for best practices
owner: ""
reviewer: ""
version: "0.1"
status: "draft"

# Technical assets
technical_assets:
`

	for _, asset := range assets {
		template += fmt.Sprintf("  - id: %s\n", asset.ID)
		template += fmt.Sprintf("    name: %s\n", asset.Name)
		template += fmt.Sprintf("    technology: %s\n", asset.Technology)
		template += "    description: \"\"\n"
		template += "    criticality: high\n"
		template += "    data_assets:\n"
		template += "      - id: \"\"\n"
		template += "        name: \"\"\n"
		template += "    out_of_scope: false\n\n"
	}

	template += "# Trust boundaries\ntrust_boundaries:\n"

	for _, boundary := range boundaries {
		template += fmt.Sprintf("  - id: %s\n", boundary.ID)
		template += fmt.Sprintf("    name: %s\n", boundary.Name)
		template += fmt.Sprintf("    technology: %s\n", boundary.Technology)
		template += "    description: \"\"\n"
		template += "    out_of_scope: false\n\n"
	}

	template += "# Data assets (filled with defaults)\ndata_assets:\n  - id: \"\"\n    name: \"\"\n    description: \"\"\n    classification: \"\"\n    storage:\n      - id: \"\"\n        name: \"\"\n    flow_source: false\n\n"

	// Add some default system elements
	template += "# Default values for other sections\n"
	template += "data_flows:\n  - id: \"\"\n    name: \"\"\n    description: \"\"\n"
	template += "    source: \"\"\n    destination: \"\"\n\n"

	return template
}

// GenerateFromTerraform generates a complete taralizer YAML from terraform files
func GenerateFromTerraform(tfDir, profilePath string) (string, error) {
	// Parse Terraform files
	resources, err := ParseTerraformDir(tfDir)
	if err != nil {
		return "", fmt.Errorf("failed to parse terraform files: %w", err)
	}

	// Load the google.yaml profile
	profile, err := LoadGoogleProfile(profilePath)
	if err != nil {
		return "", fmt.Errorf("failed to load google profile: %w", err)
	}

	// Map GCP resources to taralizer assets
	assets, boundaries := profile.ResourceToAssetMapping(resources)

	// Generate the taralizer YAML template
	yamlContent := GenerateTaralizerTemplate(assets, boundaries)

	return yamlContent, nil
}
