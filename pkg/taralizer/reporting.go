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
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"text/template"
)

// Taralzer struct
type ReportEngine struct {
	report Report
}

// creates a new reporting engine
func NewReportEngine() ReportEngine {
	return ReportEngine{}
}

// severity is a helper method to convert severity ids into strings.
func (svc *ReportEngine) severity(severity int64) string {
	result := "NONE"
	if severity > 0 {
		result = "LOW"
	} else if severity >= 4 {
		result = "MEDIUM"
	} else if severity >= 9 {
		result = "HIGH"
	}

	return fmt.Sprintf("%s(%d)", result, severity)
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

	return fmt.Sprintf("%s(%d)", result, severity)
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

	return fmt.Sprintf("%s(%d)", result, severity)
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

// GenerateReportFilePDF creates a report to the file 'filename' on the local file system
// It uses the 'wkhtmltopdf' command line tool that should be available in the path
func (svc *ReportEngine) GenerateReportFilePDF(filename string, tplFileReport string, tplFileCover string, report Report) {
	_, err := exec.LookPath("wkhtmltopdf")
	if err != nil {
		log.Fatal(err)
	}

	svc.GenerateReportFile("pdf_report.html", tplFileReport, report)
	svc.GenerateReportFile("pdf_report_cover.html", tplFileCover, report)
	if err != nil {
		log.Fatal(err)
	}

	// #nosec G204 we intentionally call wkhtmltopdf to create PDF reports
	wkhtmltopdf := exec.Command("wkhtmltopdf", "--enable-local-file-access",
		"--footer-font-size", "8",
		"--footer-line",
		"--footer-left", "powered by Taralizer",
		"--footer-right", "[page] / [topage]",
		"--footer-center", "--confidential--",
		"--header-line",
		"--header-font-size", "8",
		"--header-center", "Threat and Risk Analysis - "+report.Title,
		"cover", "pdf_report_cover.html",
		"toc",
		"pdf_report.html",
		filename)
	err = wkhtmltopdf.Run()
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("pdf_report.html")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("pdf_report_cover.html")
	if err != nil {
		log.Fatal(err)
	}
}

// GenerateReportFile creates a report to the file 'filename' on the local file system
func (svc *ReportEngine) GenerateReportFile(filename string, tplFile string, report Report) {
	fo, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	svc.GenerateReport(fo, tplFile, report)
}

// GenerateReport uses the golang template file 'tplFile' to generate a
// text report. Several templates have been defined and stored in the 'templates'directory'
func (svc *ReportEngine) GenerateReport(wr io.Writer, tplFile string, report Report) {
	funcMap := template.FuncMap{
		"findTrustedBoundary": svc.findTrustedBoundary,
		"findTechnicalAsset":  svc.findTechnicalAsset,
		"findThreatAgent":     svc.findThreatAgent,
		"isRootTrustBoundary": svc.isRootTrustBoundary,
		"likelihood":          svc.likelihood,
		"impact":              svc.impact,
		"severity":            svc.severity,
	}

	svc.report = report
	/* #nosec G304 */
	f, err := os.Open(tplFile)
	if err != nil {
		log.Fatalf("GenerateReport: cannot load template file: %v", err)
	}

	/* #nosec G307 */
	defer f.Close()

	templateText, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("GenerateReport: ReadAll error: %s", err)
	}

	tpl, err := template.New("tpl").Funcs(funcMap).Parse(string(templateText))
	if err != nil {
		log.Fatalf("GenerateReport: template.New error: %s", err)
	}

	err = tpl.Execute(wr, report)
	if err != nil {
		log.Fatalf("GenerateReport: Execute error: %s", err)
	}
}

// GetTemplateDir returns the directory of the template files
func (svc *ReportEngine) GetTemplateDir() string {
	defaultTemplatesDir := []string{"./templates/", "/etc/taralizer/templates/", "../templates/", "../../templates/"}
	templateDir := "NOT_FOUND"
	for _, v := range defaultTemplatesDir {
		if _, err := os.Stat(v); !os.IsNotExist(err) {
			templateDir = v
			break
		}
	}
	return templateDir
}
