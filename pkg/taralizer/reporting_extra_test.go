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
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateReportFileHTML(t *testing.T) {
	report, err := Load("../../examples/gcp/bank_of_anthos.yaml")
	require.NoError(t, err)

	engine := NewReportEngine()
	tmpDir := t.TempDir()
	outPath := filepath.Join(tmpDir, "report.html")

	require.NoError(t, engine.GenerateReportFile(outPath, "../../templates/html.tpl", report))
	info, err := os.Stat(outPath)
	require.NoError(t, err)
	assert.Greater(t, info.Size(), int64(0))

	// Unwritable output path -> create error.
	badPath := filepath.Join(tmpDir, "no-such-dir", "report.html")
	assert.Error(t, engine.GenerateReportFile(badPath, "../../templates/html.tpl", report))
}

func TestGenerateReportFileMarkdownError(t *testing.T) {
	report := Report{Title: "t"}
	engine := NewReportEngine()

	badPath := filepath.Join(t.TempDir(), "no-such-dir", "report.md")
	assert.Error(t, engine.GenerateReportFileMarkdown(badPath, "../../templates/markdown_report.tpl", report))
}

func TestGenerateReportTemplateErrors(t *testing.T) {
	engine := NewReportEngine()
	report := Report{Title: "t"}

	// Missing template file -> open error.
	assert.Error(t, engine.GenerateReport(os.Stdout, filepath.Join(t.TempDir(), "missing.tpl"), report))

	// Malformed template -> parse error.
	badTpl := filepath.Join(t.TempDir(), "bad.tpl")
	require.NoError(t, os.WriteFile(badTpl, []byte("{{.Title"), 0o644))
	assert.Error(t, engine.GenerateReport(os.Stdout, badTpl, report))
}

func TestGenerateReportFilePDFRenderErrors(t *testing.T) {
	report, err := Load("../../examples/gcp/bank_of_anthos.yaml")
	require.NoError(t, err)

	engine := NewReportEngine()
	outPath := filepath.Join(t.TempDir(), "out.pdf")

	// Missing cover template -> render error before Chrome is involved.
	assert.Error(t, engine.GenerateReportFilePDF(outPath, "../../templates/pdf_report.tpl", filepath.Join(t.TempDir(), "missing_cover.tpl"), report))

	// Valid cover, missing report template -> second render error.
	assert.Error(t, engine.GenerateReportFilePDF(outPath, filepath.Join(t.TempDir(), "missing_report.tpl"), "../../templates/pdf_report_cover.tpl", report))
}

// withTemplatesDir chdirs into a temp dir containing templates/console.tpl so
// GetTemplateDir resolves "./templates/".
func withTemplatesDir(t *testing.T) {
	t.Helper()
	dir := t.TempDir()
	require.NoError(t, os.Mkdir(filepath.Join(dir, "templates"), 0o755))
	src, err := os.ReadFile("../../templates/console.tpl")
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, "templates", "console.tpl"), src, 0o644))
	t.Chdir(dir)
}

func TestGetTemplateDir(t *testing.T) {
	engine := NewReportEngine()

	// From the package directory no candidate exists -> error.
	_, err := engine.GetTemplateDir()
	assert.Error(t, err)

	withTemplatesDir(t)
	dir, err := engine.GetTemplateDir()
	require.NoError(t, err)
	assert.Equal(t, "./templates/", dir)
}

func TestGenerateConsoleReportTemplateDir(t *testing.T) {
	report := Report{Title: "Console Test", Customer: "ACME"}
	engine := NewReportEngine()

	// No templates directory resolvable -> error.
	assert.Error(t, engine.GenerateConsoleReport(report))

	withTemplatesDir(t)
	// Renders to stdout; only the absence of an error is asserted.
	assert.NoError(t, engine.GenerateConsoleReport(report))
}

func TestGenerateConsoleStyledReport(t *testing.T) {
	report := Report{Title: "Styled Console Test", Customer: "ACME"}
	engine := NewReportEngine()

	// No templates directory resolvable -> error.
	assert.Error(t, engine.GenerateConsoleStyledReport(report))

	withTemplatesDir(t)
	// Renders through glamour to stdout; only the absence of an error is asserted.
	assert.NoError(t, engine.GenerateConsoleStyledReport(report))
}
