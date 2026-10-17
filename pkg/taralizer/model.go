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
package taralizer

import (
	"fmt"
	"io"

	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Defines a risk identifed in model
type Risk struct {
	Id                 string `yaml:"id"`
	Cwe                int64  `yaml:"cwe"`
	Title              string `yaml:"title"`
	Description        string `yaml:"description"`
	Message            string `yaml:"message"`
	Url                string `yaml:"url"`
	Impact             int64  `yaml:"impact"`
	Likelihood         int64  `yaml:"likelihood"`
	Severity           int64  `yaml:"severity"`
	Action             string `yaml:"action"`
	Mitigation         string `yaml:"mitigation"`
	ResidualImpact     int64  `yaml:"residual_impact"`
	ResidualLikelihood int64  `yaml:"residual_likelihood"`
	ResidualSeverity   int64  `yaml:"residual_severity"`
	Status             string `yaml:"status"`
}

type Measure struct {
	Id                 string `yaml:"id"`
	Action             string `yaml:"action"`
	Justification      string `yaml:"justification"`
	Ticket             string `yaml:"ticket"`
	CheckedBy          string `yaml:"checked_by"`
	ResidualImpact     int64  `yaml:"residual_impact"`
	ResidualLikelihood int64  `yaml:"residual_likelihood"`
	ResidualSeverity   int64  `yaml:"residual_severity"`
	Status             string `yaml:"status"`
}

type ThreatAgent struct {
	Id          string `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type DataAsset struct {
	Id              string `yaml:"id"`
	Name            string `yaml:"name"`
	Description     string `yaml:"description,omitempty"`
	Confidentiality string `yaml:"confidentiality,omitempty"`
	Integrity       string `yaml:"integrity,omitempty"`
	Availability    string `yaml:"availability,omitempty"`
}

type CommunicationLink struct {
	Target             string   `yaml:"target"`
	Description        string   `yaml:"description,omitempty"`
	Protocol           string   `yaml:"protocol,omitempty"`
	Authenication      string   `yaml:"authenication,omitempty"`
	Authorization      string   `yaml:"authorization,omitempty"`
	DataAssetsSent     []string `yaml:"data_assets_sent,omitempty"`
	DataAssetsReceived []string `yaml:"data_assets_received,omitempty"`
}

type TechnicalAsset struct {
	Id                  string              `yaml:"id"`
	Name                string              `yaml:"name"`
	Description         string              `yaml:"description,omitempty"`
	Confidentiality     string              `yaml:"confidentiality,omitempty"`
	Integrity           string              `yaml:"integrity,omitempty"`
	Availability        string              `yaml:"availability,omitempty"`
	Technology          string              `yaml:"technology,omitempty"`
	TerraformRef        string              `yaml:"terraform_ref,omitempty"`
	Puml                string              `yaml:"puml,omitempty"`
	UsedAsClientByHuman bool                `yaml:"used_as_client_by_human,omitempty"`
	OutOfScope          bool                `yaml:"out_of_scope,omitempty"`
	Internet            bool                `yaml:"internet,omitempty"`
	DataAssetsStored    []string            `yaml:"data_assets_stored,omitempty"`
	DataAssetsProcessed []string            `yaml:"data_assets_processed,omitempty"`
	CommunicationLinks  []CommunicationLink `yaml:"communication_links,omitempty"`
	// Attributes holds model-declared values for the security-relevant
	// attributes of this asset's technology (Phase A4). Keys are the attribute
	// names declared in the mapping profile; values are compared against the
	// static Terraform literals at each attribute's path during drift.
	Attributes map[string]string `yaml:"attributes,omitempty"`
}

type TrustBoundary struct {
	Id                    string   `yaml:"id"`
	Name                  string   `yaml:"name"`
	Description           string   `yaml:"description,omitempty"`
	Technology            string   `yaml:"technology,omitempty"`
	TerraformRef          string   `yaml:"terraform_ref,omitempty"`
	Puml                  string   `yaml:"puml,omitempty"`
	TrustBoundariesNested []string `yaml:"trust_boundaries_nested,omitempty"`
	TechnicalAssetsInside []string `yaml:"technical_assets_inside,omitempty"`
	ThreatAgentsInside    []string `yaml:"threat_agents_inside,omitempty"`
}

type Author struct {
	Name    string `yaml:"name"`
	Webpage string `yaml:"webpage"`
}

// Evaluation Report
type Report struct {
	ThreatAgents    []ThreatAgent    `yaml:"threat_agents,omitempty"`
	DataAssets      []DataAsset      `yaml:"data_assets,omitempty"`
	TechnicalAssets []TechnicalAsset `yaml:"technical_assets,omitempty"`
	TrustBoundaries []TrustBoundary  `yaml:"trust_boundaries,omitempty"`
	Risks           []Risk           `yaml:"risks,omitempty"`
	RiskTracking    []Measure        `yaml:"risk_tracking,omitempty"`
	Title           string           `yaml:"title,omitempty"`
	Version         string           `yaml:"version,omitempty"`
	Customer        string           `yaml:"customer,omitempty"`
	Date            string           `yaml:"date,omitempty"`
	Author          Author           `yaml:"author,omitempty"`
	RuleSet         RuleSet          `yaml:"ruleset,omitempty"`
}

// rules
type Rule struct {
	Id          string `yaml:"id"`
	Cwe         int64  `yaml:"cwe"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Mitigation  string `yaml:"mitigation"`
	Url         string `yaml:"url"`
	Impact      int64  `yaml:"impact"`
	Likelihood  int64  `yaml:"likelihood"`
	Severity    int64  `yaml:"severity"`
}

// RulSet repesents a ruleset
type RuleSet struct {
	Name        string `yaml:"name"`
	Title       string `yaml:"title"`
	Description string `yaml:"description,omitempty"`
	Version     string `yaml:"version,omitempty"`
	Url         string `yaml:"url,omitempty"`
	Rules       []Rule `yaml:"rules,omitempty"`
}

// TechnologyAttribute is one security-relevant attribute of a technology
// (Phase A4): the model key (Name) and the Terraform argument path (Path)
// compared during attribute-level drift detection.
type TechnologyAttribute struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

// Technology represents a mapping to cloud technology
type Technology struct {
	Id        string `yaml:"id"`
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	Terraform string `yaml:"terraform"`
	Scaffold  string `yaml:"scaffold,omitempty"`
	// Attributes lists the security-relevant attributes compared during
	// attribute-level drift detection (Phase A4). Mirrors the field in
	// pkg/terraform's Technology so both profile loaders stay in sync.
	Attributes []TechnologyAttribute `yaml:"attributes,omitempty"`
}

// ProfileSet represents a mapping profile to cloud technologies
type ProfileSet struct {
	Name              string       `yaml:"name"`
	Description       string       `yaml:"description,omitempty"`
	TerraformProvider string       `yaml:"terraform_provider,omitempty"`
	Technologies      []Technology `yaml:"technologies,omitempty"`
}

// Load opens the model file and loads it into the Report model
func Load(fileName string) (Report, error) {
	report := Report{}

	/* #nosec G304 */
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return Report{}, fmt.Errorf("cannot load model file: %w", err)
	}
	/* #nosec G307 */
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return Report{}, fmt.Errorf("Load: ReadAll error: %w", err)
	}

	// unmarshal for report
	err = yaml.Unmarshal([]byte(data), &report)
	if err != nil {
		return Report{}, fmt.Errorf("Load: unmarshal error: %w", err)
	}

	// replace variables in Puml
	for i := 0; i < len(report.TechnicalAssets); i++ {
		report.TechnicalAssets[i].Puml = strings.Replace(report.TechnicalAssets[i].Puml, "$id", report.TechnicalAssets[i].Id, -1)
		report.TechnicalAssets[i].Puml = strings.Replace(report.TechnicalAssets[i].Puml, "$name", report.TechnicalAssets[i].Name, -1)

		dataAssets := ""
		for j := 0; j < len(report.TechnicalAssets[i].DataAssetsStored); j++ {
			asset := findDataAsset(&report, report.TechnicalAssets[i].DataAssetsStored[j])
			if asset != nil {
				dataAssets += asset.Name
			} else {
				return Report{}, fmt.Errorf("DataAsset not found: %s", report.TechnicalAssets[i].DataAssetsStored[j])
			}

			if j < len(report.TechnicalAssets[i].DataAssetsStored)-1 {
				dataAssets += ", "
			}
		}

		report.TechnicalAssets[i].Puml = strings.Replace(report.TechnicalAssets[i].Puml, "$data", dataAssets, -1)
	}

	// replace variables in Puml
	for i := 0; i < len(report.TrustBoundaries); i++ {
		report.TrustBoundaries[i].Puml = strings.Replace(report.TrustBoundaries[i].Puml, "$id", report.TrustBoundaries[i].Id, -1)
		report.TrustBoundaries[i].Puml = strings.Replace(report.TrustBoundaries[i].Puml, "$name", report.TrustBoundaries[i].Name, -1)
	}

	return report, nil
}

// Save writes the report to fileName as YAML.
func Save(report Report, fileName string) error {
	data, err := yaml.Marshal(&report)
	if err != nil {
		return fmt.Errorf("Save: marshal error: %w", err)
	}
	if err := os.WriteFile(fileName, data, 0644); err != nil {
		return fmt.Errorf("Save: write error: %w", err)
	}
	return nil
}

// findDataAsset locates a data asset by ID
func findDataAsset(report *Report, id string) *DataAsset {
	for i := 0; i < len(report.DataAssets); i++ {
		if report.DataAssets[i].Id == id {
			return &report.DataAssets[i]
		}
	}
	return nil
}

// LoadProfileSet opens a profile file and loads it into the ProfileSet model
func LoadProfileSet(fileName string) (ProfileSet, error) {
	profileSet := ProfileSet{}

	/* #nosec G304 */
	jsonFile, err := os.Open(getProfileDir() + fileName)
	if err != nil {
		return ProfileSet{}, fmt.Errorf("cannot load profile file: %w", err)
	}
	/* #nosec G307 */
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return ProfileSet{}, fmt.Errorf("Load: ReadAll error: %w", err)
	}

	// unmarshal for report
	err = yaml.Unmarshal([]byte(data), &profileSet)
	if err != nil {
		return ProfileSet{}, fmt.Errorf("Load: unmarshal error: %w", err)
	}

	return profileSet, nil
}

// getProfileDir returns the directory of the profile files
func getProfileDir() string {
	ex, err := os.Executable()
	if err != nil {
		// Fallback to current directory if executable path cannot be determined
		return "./profiles/"
	}
	exPath := filepath.Dir(ex)
	defaultProfileDir := []string{"./profiles/", "/etc/profiles/", exPath + "/profiles/", "../profiles/", "../../profiles/"}
	profileDir := "NOT_FOUND"
	for _, v := range defaultProfileDir {
		if _, err := os.Stat(v); !os.IsNotExist(err) {
			profileDir = v
			break
		}
	}
	return profileDir
}
