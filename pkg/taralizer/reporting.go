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
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

	funcMap := svc.createFuncMap()
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
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	defaultTemplatesDir := []string{"./templates/", "/etc/taralizer/templates/", exPath + "/templates/", "../templates/"}
	templateDir := "NOT_FOUND"
	for _, v := range defaultTemplatesDir {
		if _, err := os.Stat(v); !os.IsNotExist(err) {
			templateDir = v
			break
		}
	}
	return templateDir
}
