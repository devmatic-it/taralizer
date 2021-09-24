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
	"text/template"
)

const REPORT_FMT_STRING = "%s(%d)"

// initalizes additional TPL functions here
func (svc *ReportEngine) createFuncMap() template.FuncMap {
	return template.FuncMap{
		"findTrustedBoundary": svc.findTrustedBoundary,
		"findTechnicalAsset":  svc.findTechnicalAsset,
		"findThreatAgent":     svc.findThreatAgent,
		"isRootTrustBoundary": svc.isRootTrustBoundary,
		"likelihood":          svc.likelihood,
		"impact":              svc.impact,
		"severity":            svc.severity,
	}

}

// severity is a helper method to convert severity ids into strings.
func (svc *ReportEngine) severity(severity int64) string {
	result := "NONE"
	if severity == 9 {
		result = "CRITICAL"
	} else if severity >= 6 {
		result = "HIGH"
	} else if severity >= 4 {
		result = "MEDIUM"
	} else if severity >= 1 {
		result = "LOW"
	}

	return fmt.Sprintf(REPORT_FMT_STRING, result, severity)
}

// severity is a helper method to convert severity ids into strings.
func (svc *ReportEngine) likelihood(severity int64) string {
	result := "NONE"
	if severity == 1 {
		result = "LOW"
	} else if severity == 2 {
		result = "MEDIUM"
	} else if severity == 3 {
		result = "HIGH"
	} else if severity >= 4 {
		result = "VERY HIGH"
	}

	return fmt.Sprintf(REPORT_FMT_STRING, result, severity)
}

// severity is a helper method to convert severity ids into strings.
func (svc *ReportEngine) impact(severity int64) string {
	result := "NONE"
	if severity == 1 {
		result = "LOW"
	} else if severity == 2 {
		result = "MEDIUM"
	} else if severity == 3 {
		result = "HIGH"
	} else if severity >= 4 {
		result = "VERY HIGH"
	}

	return fmt.Sprintf(REPORT_FMT_STRING, result, severity)
}

// findTrustedBoundary  is a TPL function that searches a trust boundary by Id
func (svc *ReportEngine) findTrustedBoundary(id string) *TrustBoundary {
	for i := 0; i < len(svc.report.TrustBoundaries); i++ {
		if svc.report.TrustBoundaries[i].Id == id {
			return &svc.report.TrustBoundaries[i]
		}
	}

	log.Printf("WARN Trust Boundary %s not found.\n", id)
	return nil
}

// findTrustTechnicalAsset  is a TPL function that searches a technical asset by Id
func (svc *ReportEngine) findTechnicalAsset(id string) *TechnicalAsset {
	for i := 0; i < len(svc.report.TechnicalAssets); i++ {
		if svc.report.TechnicalAssets[i].Id == id {
			return &svc.report.TechnicalAssets[i]
		}
	}

	log.Printf("WARN Technical Asset %s not found.\n", id)
	return nil
}

// findThreatAgent  is a TPL function that searches threat agents by Id
func (svc *ReportEngine) findThreatAgent(id string) *ThreatAgent {
	for i := 0; i < len(svc.report.ThreatAgents); i++ {
		if svc.report.ThreatAgents[i].Id == id {
			return &svc.report.ThreatAgents[i]
		}
	}

	log.Printf("WARN Threat Agent %s not found.\n", id)
	return nil
}

// isRootTrustBoundary is a TPL function that determines if a trust boundary is a top-level trust boundary
func (svc *ReportEngine) isRootTrustBoundary(id string) bool {
	for i := 0; i < len(svc.report.TrustBoundaries); i++ {
		for j := 0; j < len(svc.report.TrustBoundaries[i].TrustBoundariesNested); j++ {
			if svc.report.TrustBoundaries[i].TrustBoundariesNested[j] == id {
				return false
			}
		}
	}
	return true
}
