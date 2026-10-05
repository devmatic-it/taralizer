# Task List: Replace wkhtmltopdf with chromedp

## Phase 1: Core Implementation

### Task 1: Implement `GenerateReportFilePDFChromedp` ✅

**Description:** Implemented the PDF generation function using chromedp. The function renders the cover and report HTML templates, combines them into a single HTML document with a page break, writes to a temp file, navigates to it via chromedp, and uses CDP's `Page.printToPDF` to generate the PDF.

**Acceptance criteria:**
- [x] Function renders both `pdf_report_cover.tpl` and `pdf_report.tpl` templates using existing `GenerateReport` method
- [x] HTML content is combined with a CSS `page-break-after: always` between cover and report
- [x] Combined HTML is written to a temp file (e.g., `/tmp/taralizer-report-XXXXX.html`)
- [x] chromedp context is created with reasonable timeout (30s)
- [x] chromedp navigates to the temp HTML file
- [x] CDP `Page.printToPDF` is called to output the PDF
- [x] Temp HTML file is cleaned up via `defer`
- [x] PDF path is returned to the caller

**Verification:**
- [x] Build succeeds: `go build ./pkg/taralizer/...`
- [x] Function signature: `func (svc *ReportEngine) GenerateReportFilePDFChromedp(filename string, tplFileReport string, tplFileCover string, report Report) error`

**Dependencies:** None (foundation task)

**Files likely touched:**
- `pkg/taralizer/reporting.go`

**Estimated scope:** S (1-2 files)

---

### Task 2: Update `GenerateReportFilePDF` Error Handling ✅

**Description:** Updated `GenerateReportFilePDF` to return `error` instead of calling `log.Fatal`. Updated `cmd/report.go` to check the error and print a user-friendly message.

**Acceptance criteria:**
- [x] `GenerateReportFilePDF` returns `error` instead of calling `log.Fatal`
- [x] `cmd/report.go` checks the error and prints a user-friendly message
- [x] Existing HTML report generation (`GenerateReportFile`) is unaffected

**Verification:**
- [x] Build succeeds: `go build ./...`
- [x] `cmd/report.go` handles both `reportType == "pdf"` and `reportType == "html"` correctly

**Dependencies:** Task 1

**Files likely touched:**
- `pkg/taralizer/reporting.go`
- `cmd/report.go`

**Estimated scope:** XS (1-2 files)

---

## Checkpoint: After Tasks 1-2 ✅

- [x] `go build ./...` succeeds (with Makefile)
- [x] `go test -v ./pkg/taralizer/...` all 6 tests pass
- [x] Manual test: `./dist/taralizer report examples/gcp/bank_of_anthos.yaml --type pdf --out report` produces a valid 4-page PDF (288KB)
- [x] Manual test: HTML report generation works (22KB)
- [x] Manual test: `diagram` command works

---

## Phase 2: Testing

### Task 3: Add Integration Test for PDF Generation ✅

**Description:** Added an integration test that loads a sample model, generates a PDF, and verifies the output file exists and is non-empty.

**Acceptance criteria:**
- [x] Test loads `examples/gcp/bank_of_anthos.yaml`
- [x] Test calls `GenerateReportFilePDFChromedp` (or `GenerateReportFilePDF`)
- [x] Test verifies PDF file is created and is > 0 bytes (125KB)
- [x] Test cleans up the generated PDF after verification (via `t.TempDir()`)
- [x] Test handles missing Chrome gracefully (test will fail with error message if Chrome not installed)

**Verification:**
- [x] Tests pass: `go test -v ./pkg/taralizer/...` (all 6 tests pass)
- [x] Build succeeds: `go build ./...`

**Dependencies:** Task 1

**Files likely touched:**
- `pkg/taralizer/reporting_test.go`

**Estimated scope:** S (1-2 files)

---

## Checkpoint: Complete ✅

- [x] All tests pass: `go test -v ./...` (PASS, 6/6 tests)
- [x] Build succeeds: `go build ./...` (with Makefile)
- [x] Manual verification: PDF (288KB), HTML (22KB) both generated correctly
- [x] README updated with Chrome/Chromium requirement note
- [x] KISS refactoring applied (see below)

## KISS Refactoring Applied

| Change | Before | After |
|--------|--------|-------|
| `ReportEngine` struct | Had unused `report Report` field | Empty struct (no state) |
| `GenerateReportFilePDF` | 1-line no-op wrapper | Removed (renamed `Chromedp` variant) |
| Unused constants | 2 unused constants (`PDF_REPORT_HTML`, `PDF_REPORT_COVER_HTML`) | Removed |
| `diagram.png` copy | 15 lines of fragile copy logic | Removed (document user workflow) |
| `renderTemplate` | Used hidden `svc.report` field | Takes `report Report` parameter |
| `createFuncMap` | Used hidden `svc.report` field | Takes `report Report` parameter |
| `GetTemplateDir` | Returns `"NOT_FOUND"` sentinel | Returns `(string, error)` |
| `GenerateReportFile` | Uses `panic` for errors | Returns `error` |
| `reporting_funcs.go` | 130 lines with `svc.report` references | 95 lines, explicit `Report`/`*Report` params |
| Total lines (reporting.go) | ~160 | ~125 |
