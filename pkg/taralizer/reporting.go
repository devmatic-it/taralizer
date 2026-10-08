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
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/chromedp/chromedp"
	"github.com/charmbracelet/glamour"
)

// ReportEngine generates HTML and PDF reports from a Taralizer model.
type ReportEngine struct{}

// NewReportEngine creates a new ReportEngine.
func NewReportEngine() ReportEngine {
	return ReportEngine{}
}

// GenerateReportFilePDF generates a PDF report from the cover and report
// templates using chromedp (Chrome headless). A Chrome or Chromium binary
// must be installed on the system.
func (svc *ReportEngine) GenerateReportFilePDF(filename string, tplFileReport string, tplFileCover string, report Report) error {
	tmpDir, err := os.MkdirTemp("", "taralizer-report-*")
	if err != nil {
		return fmt.Errorf("create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	var coverBuf, reportBuf bytes.Buffer
	if err := svc.renderTemplate(&coverBuf, tplFileCover, report); err != nil {
		return fmt.Errorf("render cover: %w", err)
	}
	if err := svc.renderTemplate(&reportBuf, tplFileReport, report); err != nil {
		return fmt.Errorf("render report: %w", err)
	}

	// Combine with page break (cover page 1, report starts page 2).
	html := coverBuf.String() + `<div style="page-break-after: always;"></div>` + reportBuf.String()

	htmlPath := filepath.Join(tmpDir, "report.html")
	if err := os.WriteFile(htmlPath, []byte(html), 0644); err != nil {
		return fmt.Errorf("write HTML: %w", err)
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	if err := chromedp.Do(ctx, chromedp.Navigate("file://"+htmlPath)); err != nil {
		return fmt.Errorf("navigate: %w", err)
	}

	// Capture PDF (mermaid.js renders the DFD before this if internet is available)
	pdfData, err := chromedp.Run(ctx, chromedp.PrintToPDF(
		chromedp.PDFMargins(0.4, 0.4, 0.6, 0.6),
		chromedp.PDFPrintBackground(),
	))
	if err != nil {
		return fmt.Errorf("print-to-PDF: %w", err)
	}

	return os.WriteFile(filename, pdfData, 0644)
}

// renderTemplate renders tplFile into wr using the report data.
func (svc *ReportEngine) renderTemplate(wr io.Writer, tplFile string, report Report) error {
	f, err := os.Open(tplFile) // #nosec G304
	if err != nil {
		return fmt.Errorf("open template %s: %w", tplFile, err)
	}
	defer f.Close()

	src, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("read template %s: %w", tplFile, err)
	}

	tpl, err := template.New("tpl").Funcs(svc.createFuncMap(report)).Parse(string(src))
	if err != nil {
		return fmt.Errorf("parse template %s: %w", tplFile, err)
	}

	return tpl.Execute(wr, report)
}

// GenerateReportFile generates an HTML report.
func (svc *ReportEngine) GenerateReportFile(filename string, tplFile string, report Report) error {
	fo, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create %s: %w", filename, err)
	}
	defer fo.Close()

	return svc.GenerateReport(fo, tplFile, report)
}

// GenerateReportFileMarkdown generates a markdown report with embedded mermaid DFD.
func (svc *ReportEngine) GenerateReportFileMarkdown(filename string, tplFile string, report Report) error {
	fo, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create %s: %w", filename, err)
	}
	defer fo.Close()

	return svc.GenerateReport(fo, tplFile, report)
}

// GenerateConsoleReport renders the console template to stdout.
func (svc *ReportEngine) GenerateConsoleReport(report Report) error {
	tplDir, err := svc.GetTemplateDir()
	if err != nil {
		return fmt.Errorf("find templates: %w", err)
	}
	return svc.GenerateReport(os.Stdout, tplDir+"console.tpl", report)
}

// GenerateConsoleStyledReport renders the console template with glamour styling to stdout.
func (svc *ReportEngine) GenerateConsoleStyledReport(report Report) error {
	tplDir, err := svc.GetTemplateDir()
	if err != nil {
		return fmt.Errorf("find templates: %w", err)
	}

	// Render the console template and style it with glamour
	var buf bytes.Buffer
	if err := svc.GenerateReport(&buf, tplDir+"console.tpl", report); err != nil {
		return fmt.Errorf("render template: %w", err)
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(120),
	)
	if err != nil {
		return fmt.Errorf("create glamour renderer: %w", err)
	}
	defer renderer.Close()

	output, err := renderer.Render(buf.String())
	if err != nil {
		return fmt.Errorf("render markdown: %w", err)
	}

	fmt.Print(output)
	return nil
}

// GenerateReport renders tplFile into wr.
func (svc *ReportEngine) GenerateReport(wr io.Writer, tplFile string, report Report) error {
	return svc.renderTemplate(wr, tplFile, report)
}

// GetTemplateDir returns the first existing templates directory, or an error.
func (svc *ReportEngine) GetTemplateDir() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("find executable: %w", err)
	}
	exPath := filepath.Dir(ex)
	for _, dir := range []string{"./templates/", "/etc/taralizer/templates/", exPath + "/templates/", "../templates/"} {
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			return dir, nil
		}
	}
	return "", fmt.Errorf("templates directory not found (searched: %v)", []string{"./templates/", "/etc/taralizer/templates/", exPath + "/templates/", "../templates/"})
}
