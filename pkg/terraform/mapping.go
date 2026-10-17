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
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// TechnologyAttribute is one security-relevant attribute of a technology
// (Phase A4). Name is the key under which a model entry declares its value;
// Path is the Terraform argument path (dot-separated, e.g.
// "settings.ip_configuration.ipv4_enabled") whose static literal value is
// compared against the model-declared one.
type TechnologyAttribute struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

// Technology represents a technology mapping in the taralizer model
type Technology struct {
	ID        string                `yaml:"id"`
	Name      string                `yaml:"name"`
	Type      string                `yaml:"type"`
	Terraform string                `yaml:"terraform"`
	Scaffold  string                `yaml:"scaffold,omitempty"`
	// Attributes lists the security-relevant attributes compared during
	// attribute-level drift detection (Phase A4). Empty for technologies with
	// no declared relevance.
	Attributes []TechnologyAttribute `yaml:"attributes,omitempty"`
}

// MappingProfile represents the structure of google.yaml profile
type MappingProfile struct {
	Name              string       `yaml:"name"`
	Description       string       `yaml:"description"`
	TerraformProvider string       `yaml:"terraform_provider"`
	Technologies      []Technology `yaml:"technologies"`
}

// LoadGoogleProfile loads the google.yaml profile to map GCP resources to taralizer assets
func LoadGoogleProfile(profilePath string) (*MappingProfile, error) {
	// Read the YAML file
	yamlContent, err := os.ReadFile(profilePath)
	if err != nil {
		return nil, err
	}

	// Parse the YAML content
	var profile MappingProfile
	err = yaml.Unmarshal(yamlContent, &profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

// MapResourceToTechnology maps a GCP resource to its taralizer technology mapping
func (p *MappingProfile) MapResourceToTechnology(resourceType string) *Technology {
	for _, tech := range p.Technologies {
		if tech.Terraform == resourceType {
			return &tech
		}
	}
	return nil
}

// TechnologyByID returns the profile technology with the given id, or nil if
// no technology has that id. Used by attribute-level drift (Phase A4) to look
// up the relevance declared for a model entry's technology.
func (p *MappingProfile) TechnologyByID(id string) *Technology {
	for i := range p.Technologies {
		if p.Technologies[i].ID == id {
			return &p.Technologies[i]
		}
	}
	return nil
}

// ResourceToAssetMapping maps GCP resources to taralizer assets and boundaries
func (p *MappingProfile) ResourceToAssetMapping(resources []GCPResource) ([]TaralizerAsset, []TrustBoundary) {
	var assets []TaralizerAsset
	var boundaries []TrustBoundary

	for _, resource := range resources {
		if tech := p.MapResourceToTechnology(resource.Type); tech != nil {
			switch tech.Type {
			case "technical_asset":
				asset := TaralizerAsset{
					ID:         resource.Name,
					Name:       tech.Name,
					Technology: tech.ID,
				}
				assets = append(assets, asset)
			case "trust_boundary":
				boundary := TrustBoundary{
					ID:         resource.Name,
					Name:       tech.Name,
					Technology: tech.ID,
				}
				boundaries = append(boundaries, boundary)
			}
		} else {
			log.Printf("Warning: Unknown GCP resource type: %s", resource.Type)
		}
	}

	return assets, boundaries
}

// TaralizerAsset represents a technical asset in taralizer
type TaralizerAsset struct {
	ID         string `yaml:"id"`
	Name       string `yaml:"name"`
	Technology string `yaml:"technology"`
}

// TrustBoundary represents a trust boundary in taralizer
type TrustBoundary struct {
	ID         string `yaml:"id"`
	Name       string `yaml:"name"`
	Technology string `yaml:"technology"`
}
