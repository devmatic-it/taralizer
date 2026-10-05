# Task List: Markdown Report with Embedded Mermaid

## Phase 1: Template

### Task 1: Create `templates/markdown_report.tpl` ✅

**Description:** Created a markdown template that embeds the mermaid DFD and includes all report sections (trust boundaries, technical assets, data assets, threat agents, risk assessment, methodology/rules).

**Acceptance criteria:**
- [x] Header with title, version, date, author, customer
- [x] Embedded mermaid DFD as ` ```mermaid ` code block using `flowchart TD`
- [x] Trust boundaries, technical assets, data assets, threat agents as markdown tables
- [x] Risk assessment table with severity/likelihood labels
- [x] Methodology/rules section
- [x] Uses existing template functions (`severity`, `likelihood`, `findTrustedBoundary`, etc.)

**Verification:**
- [x] Template renders without errors

**Dependencies:** None

**Files likely touched:**
- `templates/markdown_report.tpl`

**Estimated scope:** S (1 file)

---

### Task 2: Implement `GenerateReportFileMarkdown` ✅

**Description:** Added `GenerateReportFileMarkdown` to `ReportEngine` that generates a markdown report from a single template file.

**Acceptance criteria:**
- [x] Function signature: `func (svc *ReportEngine) GenerateReportFileMarkdown(filename string, tplFile string, report Report) error`
- [x] Reuses existing `renderTemplate` and `createFuncMap` (no duplication)
- [x] Returns `error` (consistent with KISS refactoring)

**Verification:**
- [x] Build succeeds: `go build ./pkg/taralizer/...`

**Dependencies:** Task 1

**Files likely touched:**
- `pkg/taralizer/reporting.go`

**Estimated scope:** XS (1 file, ~10 lines)

---

### Task 3: Update `cmd/report.go` ✅

**Description:** Added a `markdown` branch to the report command that calls `GenerateReportFileMarkdown` and outputs a `.md` file.

**Acceptance criteria:**
- [x] `taralizer report model.yaml --type markdown --out report` generates `report.md`
- [x] Error handling consistent with existing PDF/HTML branches
- [x] No changes to existing HTML or PDF report types

**Verification:**
- [x] Build succeeds: `go build ./...`
- [x] Manual test: `./dist/taralizer report examples/gcp/bank_of_anthos.yaml --type markdown --out report`

**Dependencies:** Task 2

**Files likely touched:**
- `cmd/report.go`

**Estimated scope:** XS (1 file, ~10 lines)

---

## Checkpoint: After Tasks 1-3 ✅

- [x] `go build ./...` succeeds
- [x] Manual test: `./dist/taralizer report examples/gcp/bank_of_anthos.yaml --type markdown --out report` generates valid 26KB `.md` file with embedded mermaid DFD
- [x] Manual test: Existing HTML and PDF reports unaffected
- [x] Review with human before proceeding

---

## Phase 2: Testing

### Task 4: Add Integration Test for Markdown Report ✅

**Description:** Added a test that generates a markdown report and verifies the output contains expected sections and a mermaid code block.

**Acceptance criteria:**
- [x] Test loads `examples/gcp/bank_of_anthos.yaml`
- [x] Test calls `GenerateReportFileMarkdown`
- [x] Test verifies `.md` file is created and is > 0 bytes (12,293 bytes)
- [x] Test verifies output contains ` ```mermaid ` block
- [x] Test verifies output contains expected tables (trust boundaries, risks, etc.)

**Verification:**
- [x] Tests pass: `go test -v ./pkg/taralizer/...` (all 7 tests pass)
- [x] Build succeeds: `go build ./...`

**Dependencies:** Task 2

**Files likely touched:**
- `pkg/taralizer/reporting_test.go`

**Estimated scope:** S (1 file)

---

## Checkpoint: Complete ✅

- [x] All tests pass: `go test -v ./...` (PASS, 7/7 tests)
- [x] Build succeeds: `go build ./...` (with Makefile)
- [x] Manual verification: Markdown report (26KB) generated correctly with embedded mermaid DFD
- [x] README updated with markdown report documentation
- [x] Ready for review
