// Package taralizer Threat and Risk Analyzer
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
	"fmt"
	"log"
	"strings"
	"text/template"
)

const REPORT_FMT_STRING = "%s(%d)"

// createFuncMap returns template functions bound to the given report.
func (svc *ReportEngine) createFuncMap(report Report) template.FuncMap {
	rp := &report
	return template.FuncMap{
		"findTrustedBoundary": func(id string) *TrustBoundary { return findTrustBoundary(rp, id) },
		"findTechnicalAsset":  func(id string) *TechnicalAsset { return findTechnicalAsset(rp, id) },
		"findThreatAgent":     func(id string) *ThreatAgent   { return findThreatAgent(rp, id) },
		"isRootTrustBoundary": func(id string) bool           { return isRootTrustBoundary(rp, id) },
		"likelihood":          func(s int64) string           { return likelihoodimpact(s) },
		"impact":              func(s int64) string           { return likelihoodimpact(s) },
		"severity":            func(s int64) string           { return severity(s) },
		"dataAssetNames":      func(ids []string) string      { return getDataAssetNames(report, ids) },
		"replaceAll":          func(old, new, s string) string      { return strings.ReplaceAll(s, old, new) },
		"sanitizeMermaidID":   func(name string) string           { return sanitizeMermaidID(name) },
		"mermaidDFD":          func() string { return mermaidDFD(report) },
		"markdownDFD":         func() string { return mermaidDFD(report) },
		"repeat":              func(s string, count int) string { return strings.Repeat(s, count) },
		"add":                 func(a, b int) int              { return a + b },
	}
}

func severity(severity int64) string {
	switch {
	case severity == 9:
		return fmt.Sprintf(REPORT_FMT_STRING, "CRITICAL", severity)
	case severity >= 6:
		return fmt.Sprintf(REPORT_FMT_STRING, "HIGH", severity)
	case severity >= 4:
		return fmt.Sprintf(REPORT_FMT_STRING, "MEDIUM", severity)
	case severity >= 1:
		return fmt.Sprintf(REPORT_FMT_STRING, "LOW", severity)
	case severity >= 0:
		return fmt.Sprintf(REPORT_FMT_STRING, "NONE", severity)
	default:
		return ""
	}
}

func likelihoodimpact(severity int64) string {
	switch {
	case severity == 0:
		return fmt.Sprintf(REPORT_FMT_STRING, "NONE", severity)
	case severity == 1:
		return fmt.Sprintf(REPORT_FMT_STRING, "LOW", severity)
	case severity == 2:
		return fmt.Sprintf(REPORT_FMT_STRING, "MEDIUM", severity)
	case severity == 3:
		return fmt.Sprintf(REPORT_FMT_STRING, "HIGH", severity)
	case severity >= 4:
		return fmt.Sprintf(REPORT_FMT_STRING, "VERY HIGH", severity)
	default:
		return ""
	}
}

func findTrustBoundary(report *Report, id string) *TrustBoundary {
	for i := range report.TrustBoundaries {
		if report.TrustBoundaries[i].Id == id {
			return &report.TrustBoundaries[i]
		}
	}
	log.Printf("WARN Trust Boundary %s not found.\n", id)
	return nil
}

func getDataAssetNames(report Report, ids []string) string {
	if len(ids) == 0 {
		return ""
	}
	result := "["
	for i, id := range ids {
		da := findDataAsset(&report, id)
		if da != nil {
			result += da.Name
			if i < len(ids)-1 {
				result += ", "
			}
		}
	}
	return result + "]"
}

func findTechnicalAsset(report *Report, id string) *TechnicalAsset {
	for i := range report.TechnicalAssets {
		if report.TechnicalAssets[i].Id == id {
			return &report.TechnicalAssets[i]
		}
	}
	log.Printf("WARN Technical Asset %s not found.\n", id)
	return nil
}

func findThreatAgent(report *Report, id string) *ThreatAgent {
	for i := range report.ThreatAgents {
		if report.ThreatAgents[i].Id == id {
			return &report.ThreatAgents[i]
		}
	}
	log.Printf("WARN Threat Agent %s not found.\n", id)
	return nil
}

func isRootTrustBoundary(report *Report, id string) bool {
	for i := range report.TrustBoundaries {
		for j := range report.TrustBoundaries[i].TrustBoundariesNested {
			if report.TrustBoundaries[i].TrustBoundariesNested[j] == id {
				return false
			}
		}
	}
	return true
}

// sanitizeMermaidID converts a string to a valid Mermaid subgraph ID by
// replacing spaces and special characters with underscores.
func sanitizeMermaidID(name string) string {
	result := strings.ReplaceAll(name, " ", "_")
	result = strings.ReplaceAll(result, "-", "_")
	// Replace any remaining special characters with underscores
	var sanitized strings.Builder
	for _, r := range result {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
			sanitized.WriteRune(r)
		} else {
			sanitized.WriteRune('_')
		}
	}
	return sanitized.String()
}

// mermaidDFD generates the mermaid DFD source for the report.
func mermaidDFD(report Report) string {
	var b strings.Builder
	b.WriteString("flowchart TD\n")

	// Collect all unique threat agents and technical assets first,
	// then define them once at the root level (mermaid does not allow
	// duplicate node definitions across subgraphs).
	seenTA := make(map[string]bool)
	seenAsset := make(map[string]bool)

	// First pass: collect all nodes
	for _, tb := range report.TrustBoundaries {
		for _, taID := range tb.ThreatAgentsInside {
			seenTA[taID] = true
		}
		for _, assetID := range tb.TechnicalAssetsInside {
			seenAsset[assetID] = true
		}
	}

	// Second pass: generate root-level node definitions
	for _, ta := range report.ThreatAgents {
		if seenTA[ta.Id] {
			b.WriteString(fmt.Sprintf("    %s[\"%s\"]\n", ta.Id, ta.Name))
		}
	}
	for _, asset := range report.TechnicalAssets {
		if seenAsset[asset.Id] {
			b.WriteString(fmt.Sprintf("    %s[\"%s\"]\n", asset.Id, asset.Name))
		}
	}

	// Helper to generate a trust boundary subgraph with proper nesting depth.
	var genBoundary func(*TrustBoundary, int)
	genBoundary = func(tb *TrustBoundary, depth int) {
		indent := strings.Repeat("    ", depth)
		subgraphID := sanitizeMermaidID(tb.Name)
		b.WriteString(fmt.Sprintf("%ssubgraph %s [\"%s\"]\n", indent, subgraphID, tb.Name))
		contentIndent := strings.Repeat("    ", depth+1)
		for _, taID := range tb.ThreatAgentsInside {
			b.WriteString(fmt.Sprintf("%s%s\n", contentIndent, taID))
		}
		for _, nestedID := range tb.TrustBoundariesNested {
			for _, nested := range report.TrustBoundaries {
				if nested.Id == nestedID {
					genBoundary(&nested, depth+1)
					break
				}
			}
		}
		// Skip nodes whose ID matches the subgraph ID to avoid
		// mermaid conflicts between subgraph and node identifiers.
		for _, assetID := range tb.TechnicalAssetsInside {
			if assetID != subgraphID {
				b.WriteString(fmt.Sprintf("%s%s\n", contentIndent, assetID))
			}
		}
		b.WriteString(fmt.Sprintf("%send\n", indent))
	}

	// Generate root trust boundaries
	for _, tb := range report.TrustBoundaries {
		if isRootTrustBoundary(&report, tb.Id) {
			genBoundary(&tb, 1)
		}
	}

	// Generate communication links
	for _, asset := range report.TechnicalAssets {
		for _, conn := range asset.CommunicationLinks {
			b.WriteString(fmt.Sprintf("    %s --> %s\n", asset.Id, conn.Target))
		}
	}

	return b.String()
}
