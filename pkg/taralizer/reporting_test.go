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
	"os"
	"path/filepath"
	"testing"
)

func TestNewReporting(t *testing.T) {
	report, err := Load("../../examples/gcp/bank_of_anthos.yaml")
	if err != nil {
		t.Fatalf("Failed to load model: %v", err)
	}
	engine := NewReportEngine()
	if err := engine.GenerateReport(os.Stdout, "../../templates/html.tpl", report); err != nil {
		t.Fatalf("GenerateReport failed: %v", err)
	}
}

func TestGenerateReportFilePDFChromedp(t *testing.T) {
	report, err := Load("../../examples/gcp/bank_of_anthos.yaml")
	if err != nil {
		t.Fatalf("Failed to load model: %v", err)
	}

	engine := NewReportEngine()
	tmpDir := t.TempDir()
	outputPath := filepath.Join(tmpDir, "test_report.pdf")

	err = engine.GenerateReportFilePDF(outputPath,
		"../../templates/pdf_report.tpl",
		"../../templates/pdf_report_cover.tpl", report)
	if err != nil {
		t.Fatalf("Failed to generate PDF: %v", err)
	}

	info, err := os.Stat(outputPath)
	if err != nil {
		t.Fatalf("PDF file not found: %v", err)
	}
	if info.Size() == 0 {
		t.Fatal("Generated PDF is empty")
	}
	t.Logf("Generated PDF: %d bytes", info.Size())
}
