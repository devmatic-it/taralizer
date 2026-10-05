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
		"replaceAll":          func(old, new, s string) string { return strings.ReplaceAll(s, old, new) },
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
