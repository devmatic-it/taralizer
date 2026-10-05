# Implementation Plan: Replace wkhtmltopdf with chromedp

## Overview

Implement the missing `GenerateReportFilePDFChromedp` function in `pkg/taralizer/reporting.go` using the already-imported `chromedp` library. The function will render the cover and report HTML templates, combine them into a single document, and use Chrome's headless CDP print-to-PDF capability to generate the final PDF.

## Architecture Decisions

- **Single combined HTML:** Cover page + report page are merged into one HTML document with a CSS `page-break-after` between them. This produces a single PDF with the cover as page 1.
- **Temp file approach:** Rendered HTML is written to a temp file, chromedp navigates to it, then prints to PDF. Temp files are cleaned up via `defer`.
- **Error propagation:** The function returns `error` instead of calling `log.Fatal`. The caller (`cmd/report.go`) handles the error.
- **Image path handling:** The `diagram.png` reference in `pdf_report.tpl` is resolved relative to the temp HTML file's directory. The user is expected to run `diagram` first (current behavior). No new flags.
- **No CGO changes:** chromedp works without CGO on supported platforms. The Makefile remains unchanged.

## Task List

### Phase 1: Core Implementation

- [ ] Task 1: Implement `GenerateReportFilePDFChromedp` in `reporting.go`
- [ ] Task 2: Update `GenerateReportFilePDF` to propagate errors properly (remove `log.Fatal`)

### Checkpoint 1: Build & Basic Test

### Phase 2: Testing

- [ ] Task 3: Add integration test for PDF generation

### Checkpoint 2: Full Verification

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Chrome/Chromium not installed on target systems | High — PDF generation fails | Document requirement in README; provide error message suggesting installation |
| Image path `diagram.png` not found | Medium — PDF has broken image | Graceful handling: skip image, log warning, continue |
| chromedp CDP API changes | Low — library is stable | Pin chromedp version in go.mod |
| Large reports cause memory issues | Low — unlikely for threat models | Set reasonable timeouts; chromedp handles large pages well |

## Open Questions (Resolved)

1. **Chrome/Chromium:** Acceptable — more common than wkhtmltopdf. Document in README.
2. **CGO:** Not required — chromedp works without CGO on supported platforms.
3. **Image embedding:** Keep current behavior — user runs `diagram` first.
4. **Cover + report:** Merge into single PDF (cover page 1, report page 2+).
5. **Page defaults:** A4, 1cm margins (reasonable defaults for security reports).

## Files Likely Touched

- `pkg/taralizer/reporting.go` — Add `GenerateReportFilePDFChromedp`, update `GenerateReportFilePDF`
- `pkg/taralizer/reporting_test.go` — Add integration test
- `cmd/report.go` — Update error handling
- `README.md` — Document Chrome/Chromium requirement (minor)
