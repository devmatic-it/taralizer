# Spec: Replace wkhtmltopdf with chromedp for PDF Generation

## Objective

Replace the external `wkhtmltopdf` CLI dependency with the native Go library `chromedp` for generating PDF reports. The existing codebase already imports `github.com/chromedp/chromedp` and `github.com/chromedp/cdp`, but the target function `GenerateReportFilePDFChromedp` is **not yet implemented** — it is called from `GenerateReportFilePDF` but does not exist, causing a compile error.

**User story:** As a taralizer user, I want to generate PDF reports without installing `wkhtmltopdf` on my system, so that deployment and local development are simpler (single Go binary, no system dependencies).

**Success criteria:**
- `GenerateReportFilePDFChromedp` generates a PDF from the existing `pdf_report.tpl` and `pdf_report_cover.tpl` templates
- The generated PDF matches the visual output of the current `wkhtmltopdf`-based reports (verified against `examples/gcp/report.pdf`)
- No `wkhtmltopdf` binary is required at runtime
- The Go build remains CGO-free (or CGO is clearly documented as the cost)

## Tech Stack

- **Language:** Go 1.27 (already in `go.mod`)
- **PDF engine:** `github.com/chromedp/chromedp` (already imported in `reporting.go`)
- **CDP protocol:** `github.com/chromedp/cdp` (already imported)
- **Template engine:** existing `text/template` with custom funcMap
- **Output:** PDF files via Chrome headless `--print-to-pdf`

## Commands

```
Build:  make build
Test:   go test -v ./pkg/taralizer/... ./cmd/...
Lint:   golangci-lint run ./...
Manual: dist/taralizer report examples/gcp/bank_of_anthos.yaml --type pdf --out report
```

## Project Structure

```
pkg/taralizer/
  reporting.go          ← Add GenerateReportFilePDFChromedp() here
  reporting_test.go    ← Add integration test for PDF generation
cmd/
  report.go            ← No changes needed (already calls the method)
templates/
  pdf_report.tpl       ← No changes needed (already produces valid HTML)
  pdf_report_cover.tpl ← No changes needed (already produces valid HTML)
```

## Code Style

Key conventions from the existing codebase:

```go
// Use named return errors, log.Fatal only at CLI entry points
func (svc *ReportEngine) GenerateReportFilePDFChromedp(
    filename string,
    tplFileReport string,
    tplFileCover string,
    report Report,
) error {
    // 1. Render report HTML to temp file
    // 2. Render cover HTML to temp file
    // 3. Use chromedp to open the report HTML and print to PDF
    // 4. Clean up temp files
    return nil
}
```

- Errors are returned (not swallowed), callers decide how to handle them
- Temp files are cleaned up via `defer os.Remove(...)`
- The `ReportEngine` struct holds `report Report` (already exists)
- Existing funcMap helpers (`severity`, `likelihood`, etc.) are reused

## Testing Strategy

- **Unit test:** Render HTML from template → verify HTML string is valid and contains expected sections
- **Integration test:** Full end-to-end: model → report data → HTML → PDF → verify PDF file exists and is non-empty
- **Visual regression (manual):** Compare generated PDF against `examples/gcp/report.pdf`

## Boundaries

- **Always do:** Clean up temp files, return errors properly, use existing funcMap, handle missing `diagram.png` gracefully
- **Ask first:** Changing the PDF template format, adding new CLI flags, modifying the ReportEngine struct beyond what's needed
- **Never do:** Install system binaries at runtime, use CGO unless absolutely unavoidable, modify existing template files

## Success Criteria

- [ ] `go build ./...` succeeds with no errors
- [ ] `GenerateReportFilePDFChromedp` produces a valid PDF file from `examples/gcp/bank_of_anthos.yaml`
- [ ] The PDF output visually matches the existing `examples/gcp/report.pdf` (tables, colors, images, cover page)
- [ ] No `wkhtmltopdf` binary is required (verified by running without it installed)
- [ ] Existing `html` report generation is unaffected
- [ ] Integration test exists and passes

## Open Questions

1. **Chrome/Chromium availability:** chromedp requires a Chrome or Chromium binary on the system. Is this acceptable (it's far more common than `wkhtmltopdf`), or do we need a fallback strategy for headless server environments?

2. **CGO requirement:** chromedp may require CGO on some platforms. Is CGO acceptable, or must the binary remain `CGO_ENABLED=0` (as the Makefile currently does)?

3. **Image embedding:** The templates reference `diagram.png` as a relative path. Should we:
   - (a) Require the user to run `diagram` first and place `diagram.png` next to the output (current behavior), or
   - (b) Accept a `--diagram` flag to auto-generate the diagram as part of PDF generation?

4. **Multi-page cover + report:** The cover page is a separate HTML file. Should we merge them into a single PDF (cover page first, then report) or keep them as separate files?

5. **Page size / margins:** `wkhtmltopdf` has default margins and page size. What are acceptable defaults for chromedp (A4, letter, margins)?
