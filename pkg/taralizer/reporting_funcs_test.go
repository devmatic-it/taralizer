// Package taralizer Threat and Risk Analyzer
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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeverity(t *testing.T) {
	tests := []struct {
		in   int64
		want string
	}{
		{9, "CRITICAL (9)"},
		{8, "HIGH (8)"},
		{6, "HIGH (6)"},
		{5, "MED (5)"},
		{4, "MED (4)"},
		{3, "LOW (3)"},
		{1, "LOW (1)"},
		{0, "NONE (0)"},
		{-1, "TBD (-1)"},
		{-2, ""},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, severity(tt.in), "severity(%d)", tt.in)
	}
}

func TestLevelLabel(t *testing.T) {
	tests := []struct {
		in   int64
		want string
	}{
		{0, "NONE"},
		{1, "LOW"},
		{2, "MED"},
		{3, "HIGH"},
		{4, "VERY HIGH"},
		{5, "VERY HIGH"},
		{-1, "TBD"},
		{-2, ""},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, levelLabel(tt.in), "levelLabel(%d)", tt.in)
	}
}

func TestLikelihoodImpact(t *testing.T) {
	assert.Equal(t, "MED (2)", likelihoodimpact(2))
	assert.Equal(t, "TBD (-1)", likelihoodimpact(-1))
}

func testReport() Report {
	return Report{
		Title: "Test",
		TrustBoundaries: []TrustBoundary{
			{Id: "tb-root", Name: "Root Boundary"},
			{Id: "tb-nested", Name: "Nested Boundary", TrustBoundariesNested: []string{}},
		},
		TechnicalAssets: []TechnicalAsset{
			{Id: "ta-1", Name: "Web Server"},
		},
		ThreatAgents: []ThreatAgent{
			{Id: "agent-1", Name: "Internet Attacker"},
		},
		DataAssets: []DataAsset{
			{Id: "da-1", Name: "Customer Data"},
			{Id: "da-2", Name: "Session Tokens"},
		},
	}
}

func TestFindTrustBoundary(t *testing.T) {
	report := testReport()
	found := findTrustBoundary(&report, "tb-root")
	assert.NotNil(t, found)
	assert.Equal(t, "Root Boundary", found.Name)

	assert.Nil(t, findTrustBoundary(&report, "missing"))
}

func TestFindTechnicalAsset(t *testing.T) {
	report := testReport()
	found := findTechnicalAsset(&report, "ta-1")
	assert.NotNil(t, found)
	assert.Equal(t, "Web Server", found.Name)

	assert.Nil(t, findTechnicalAsset(&report, "missing"))
}

func TestFindThreatAgent(t *testing.T) {
	report := testReport()
	found := findThreatAgent(&report, "agent-1")
	assert.NotNil(t, found)
	assert.Equal(t, "Internet Attacker", found.Name)

	assert.Nil(t, findThreatAgent(&report, "missing"))
}

func TestIsRootTrustBoundary(t *testing.T) {
	report := testReport()
	// tb-nested is nested inside tb-root, so it is not a root boundary.
	report.TrustBoundaries[0].TrustBoundariesNested = []string{"tb-nested"}

	assert.True(t, isRootTrustBoundary(&report, "tb-root"))
	assert.False(t, isRootTrustBoundary(&report, "tb-nested"))
}

func TestGetDataAssetNames(t *testing.T) {
	report := testReport()

	assert.Equal(t, "", getDataAssetNames(report, nil))
	assert.Equal(t, "[Customer Data]", getDataAssetNames(report, []string{"da-1"}))
	// Unknown IDs are skipped; multiple names are comma-separated.
	assert.Equal(t, "[Customer Data, Session Tokens]", getDataAssetNames(report, []string{"da-1", "missing", "da-2"}))
}

func TestSanitizeMermaidID(t *testing.T) {
	assert.Equal(t, "my_boundary", sanitizeMermaidID("my boundary"))
	assert.Equal(t, "a_b_c", sanitizeMermaidID("a-b c"))
	// Special characters are replaced with underscores.
	assert.Equal(t, "a_b_c", sanitizeMermaidID("a.b/c"))
}

func TestBuildLinkLabel(t *testing.T) {
	assert.Equal(t, "", buildLinkLabel(CommunicationLink{}))
	assert.Equal(t, "HTTPS", buildLinkLabel(CommunicationLink{Protocol: "HTTPS"}))
	assert.Equal(t, "HTTPS / user login", buildLinkLabel(CommunicationLink{Protocol: "HTTPS", Description: "user login"}))
}

func TestDataAssetLabel(t *testing.T) {
	assert.Equal(t, "Public", dataAssetLabel("1"))
	assert.Equal(t, "Internal", dataAssetLabel("2"))
	assert.Equal(t, "Restricted", dataAssetLabel("3"))
	assert.Equal(t, "Confidential", dataAssetLabel("4"))
	assert.Equal(t, "Strictly-Confidential", dataAssetLabel("5"))
	assert.Equal(t, "Unknown", dataAssetLabel("9"))
}

// TestCreateFuncMapClosures invokes every template function closure so the
// bound helpers are exercised (templates only call a subset of them).
func TestCreateFuncMapClosures(t *testing.T) {
	engine := NewReportEngine()
	report := testReport()

	fm := engine.createFuncMap(report)

	assert.Equal(t, "Root Boundary", fm["findTrustedBoundary"].(func(string) *TrustBoundary)("tb-root").Name)
	assert.Nil(t, fm["findTrustedBoundary"].(func(string) *TrustBoundary)("missing"))
	assert.Equal(t, "Web Server", fm["findTechnicalAsset"].(func(string) *TechnicalAsset)("ta-1").Name)
	assert.Nil(t, fm["findTechnicalAsset"].(func(string) *TechnicalAsset)("missing"))
	assert.Equal(t, "Internet Attacker", fm["findThreatAgent"].(func(string) *ThreatAgent)("agent-1").Name)
	assert.Nil(t, fm["findThreatAgent"].(func(string) *ThreatAgent)("missing"))
	assert.True(t, fm["isRootTrustBoundary"].(func(string) bool)("tb-root"))
	assert.Equal(t, "[Customer Data]", fm["dataAssetNames"].(func([]string) string)([]string{"da-1"}))
	assert.Equal(t, "b-b", fm["replaceAll"].(func(string, string, string) string)("a", "b", "a-a"))
	assert.Equal(t, "a_b", fm["sanitizeMermaidID"].(func(string) string)("a b"))
	assert.Equal(t, "x", fm["repeat"].(func(string, int) string)("x", 1))
	assert.Equal(t, 5, fm["add"].(func(int, int) int)(2, 3))
	assert.Equal(t, "med", fm["lower"].(func(int64) string)(2))
	assert.NotEmpty(t, fm["mermaidDFD"].(func() string)())
}

// TestMermaidDFDNestedAndUnlabeledLinks exercises nested subgraphs and the
// unlabeled communication-link branch.
func TestMermaidDFDNestedAndUnlabeledLinks(t *testing.T) {
	report := testReport()
	report.TrustBoundaries[0].TrustBoundariesNested = []string{"tb-nested"}
	report.TrustBoundaries[0].ThreatAgentsInside = []string{"agent-1"}
	report.TrustBoundaries[0].TechnicalAssetsInside = []string{"ta-1"}
	// One labeled and one unlabeled link.
	report.TechnicalAssets[0].CommunicationLinks = []CommunicationLink{
		{Target: "ta-2", Protocol: "HTTPS", Description: "api"},
		{Target: "ta-3"}, // no protocol -> unlabeled edge
	}

	out := mermaidDFD(report)

	assert.Contains(t, out, "flowchart TD")
	assert.Contains(t, out, `subgraph Root_Boundary ["Root Boundary"]`)
	assert.Contains(t, out, `subgraph Nested_Boundary ["Nested Boundary"]`)
	assert.Contains(t, out, "ta-1 -->|HTTPS / api| ta-2")
	assert.Contains(t, out, "ta-1 --> ta-3")
	// Root-level node definitions are emitted once.
	assert.Contains(t, out, `agent-1["Internet Attacker"]`)
}

// TestMermaidDFDEmpty verifies the generator on a report without boundaries.
func TestMermaidDFDEmpty(t *testing.T) {
	out := mermaidDFD(Report{})
	assert.Equal(t, "flowchart TD\n", out)
}
