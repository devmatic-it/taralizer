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
package taralizer

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Defines a risk identifed in model
type Risk struct {
	Id          string `yaml:"id"`
	Cwe         int64  `yaml:"cwe"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Mitigation  string `yaml:"mitigation"`
	Message     string `yaml:"message"`
	Url         string `yaml:"url"`
	Impact      int64  `yaml:"impact"`
	Likelihood  int64  `yaml:"likelihood"`
	Severity    int64  `yaml:"severity"`
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
	Puml                string              `yaml:"puml,omitempty"`
	UsedAsClientByHuman bool                `yaml:"used_as_client_by_human,omitempty"`
	OutOfScope          bool                `yaml:"out_of_scope,omitempty"`
	Internet            bool                `yaml:"internet,omitempty"`
	DataAssetsStored    []string            `yaml:"data_assets_stored,omitempty"`
	CommunicationLinks  []CommunicationLink `yaml:"communication_links,omitempty"`
}

type TrustBoundary struct {
	Id                    string   `yaml:"id"`
	Name                  string   `yaml:"name"`
	Description           string   `yaml:"description,omitempty"`
	Technology            string   `yaml:"technology,omitempty"`
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
	Title           string           `yaml:"title,omitempty"`
	Version         string           `yaml:"version,omitempty"`
	Customer        string           `yaml:"customer,omitempty"`
	Date            string           `yaml:"date,omitempty"`
	Author          Author           `yaml:"author,omitempty"`
}

// Load opens the model file and loads it into the Report model
// Please, node that the Risks property is empty.
func Load(fileName string) Report {
	report := Report{}

	/* #nosec G304 */
	jsonFile, err := os.Open(fileName)
	/* #nosec G307 */
	defer jsonFile.Close()

	if err != nil {
		log.Fatalf("cannot load model file: %v", err)
	}

	data, err := ioutil.ReadAll(jsonFile)

	// unmarshal for report
	err = yaml.Unmarshal([]byte(data), &report)
	if err != nil {
		log.Fatalf("cannot unmarshal report file: %v", err)
	}

	// replace variables in Puml
	for i := 0; i < len(report.TechnicalAssets); i++ {
		report.TechnicalAssets[i].Puml = strings.Replace(report.TechnicalAssets[i].Puml, "$id", report.TechnicalAssets[i].Id, -1)
		report.TechnicalAssets[i].Puml = strings.Replace(report.TechnicalAssets[i].Puml, "$name", report.TechnicalAssets[i].Name, -1)

		dataAssets := ""
		for j := 0; j < len(report.TechnicalAssets[i].DataAssetsStored); j++ {
			dataAssets += findDataAsset(&report, report.TechnicalAssets[i].DataAssetsStored[j]).Name
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

	return report
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
