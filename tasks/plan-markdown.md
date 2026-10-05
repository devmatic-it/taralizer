# Implementation Plan: Markdown Report with Embedded Mermaid

## Overview

Add a `markdown` report type to taralizer. The report outputs a single `.md` file with an embedded mermaid DFD code block (no external images needed) and all existing report sections (trust boundaries, assets, risks, methodology).

## Architecture Decisions

- **Single template:** One `markdown_report.tpl` (no cover page needed for markdown)
- **Embedded mermaid:** The DFD is rendered as a ` ```mermaid ` code block, which renders natively in GitHub, GitLab, Obsidian, VS Code, etc.
- **No new dependencies:** Mermaid rendering happens in the markdown viewer, not in Go
- **Use `flowchart TD`:** Modern mermaid syntax (same semantics as `graph TD`, but preferred by mermaid maintainers)
- **Text labels for severity:** Use existing `severity()` / `likelihood()` functions (no CSS in markdown)

## Task List

### Phase 1: Template

- [ ] Task 1: Create `templates/markdown_report.tpl`

**Description:** Create a markdown template that embeds the mermaid DFD and includes all report sections (trust boundaries, technical assets, data assets, threat agents, risk assessment, methodology/rules).

**Acceptance criteria:**
- [ ] Header with title, version, date, author, customer
- [ ] Embedded mermaid DFD as ` ```mermaid ` code block using `flowchart TD`
- [ ] Trust boundaries, technical assets, data assets, threat agents as markdown tables
- [ ] Risk assessment table with severity/likelihood labels
- [ ] Methodology/rules section
- [ ] Uses existing template functions (`severity`, `likelihood`, `findTrustedBoundary`, etc.)

**Verification:**
- [ ] Template renders without errors: `go test -run TestMarkdownTemplate`

**Dependencies:** None

**Files likely touched:**
- `templates/markdown_report.tpl`

**Estimated scope:** S (1 file)

---

### Phase 2: Report Engine

- [ ] Task 2: Implement `GenerateReportFileMarkdown` in `reporting.go`

**Description:** Add a new method to `ReportEngine` that generates a markdown report from a single template file.

**Acceptance criteria:**
- [ ] Function signature: `func (svc *ReportEngine) GenerateReportFileMarkdown(filename string, tplFile string, report Report) error`
- [ ] Reuses existing `renderTemplate` and `createFuncMap` (no duplication)
- [ ] Returns `error` (consistent with KISS refactoring)

**Verification:**
- [ ] Build succeeds: `go build ./pkg/taralizer/...`

**Dependencies:** Task 1

**Files likely touched:**
- `pkg/taralizer/reporting.go`

**Estimated scope:** XS (1 file, ~10 lines)

---

### Phase 3: CLI Integration

- [ ] Task 3: Update `cmd/report.go` to handle `--type markdown`

**Description:** Add a `markdown` branch to the report command that calls `GenerateReportFileMarkdown` and outputs a `.md` file.

**Acceptance criteria:**
- [ ] `taralizer report model.yaml --type markdown --out report` generates `report.md`
- [ ] Error handling consistent with existing PDF/HTML branches
- [ ] No changes to existing HTML or PDF report types

**Verification:**
- [ ] Build succeeds: `go build ./...`
- [ ] Manual test: `./dist/taralizer report examples/gcp/bank_of_anthos.yaml --type markdown --out report`

**Dependencies:** Task 2

**Files likely touched:**
- `cmd/report.go`

**Estimated scope:** XS (1 file, ~10 lines)

---

## Checkpoint: After Tasks 1-3

- [ ] `go build ./...` succeeds
- [ ] Manual test: `./dist/taralizer report examples/gcp/bank_of_anthos.yaml --type markdown --out report` generates valid `.md`
- [ ] Review with human before proceeding

---

### Phase 4: Testing

- [ ] Task 4: Add integration test for markdown report

**Description:** Add a test that generates a markdown report and verifies the output contains expected sections and a mermaid code block.

**Acceptance criteria:**
- [ ] Test loads `examples/gcp/bank_of_anthos.yaml`
- [ ] Test calls `GenerateReportFileMarkdown`
- [ ] Test verifies `.md` file is created and is > 0 bytes
- [ ] Test verifies output contains ` ```mermaid ` block
- [ ] Test verifies output contains expected tables (trust boundaries, risks, etc.)

**Verification:**
- [ ] Tests pass: `go test -v ./pkg/taralizer/...`
- [ ] Build succeeds: `go build ./...`

**Dependencies:** Task 2

**Files likely touched:**
- `pkg/taralizer/reporting_test.go`

**Estimated scope:** S (1 file)

---

## Checkpoint: Complete

- [ ] All tests pass: `go test -v ./...`
- [ ] Build succeeds: `go build ./...`
- [ ] Manual verification: Generated markdown opens and renders correctly in GitHub/GitLab/Obsidian
- [ ] README updated with markdown report documentation
- [ ] Ready for review

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Mermaid DFD syntax changes between versions | Low — mermaid is stable | Pin mermaid version in documentation; use `flowchart TD` (current standard) |
| Markdown viewers don't support mermaid | Low — GitHub, GitLab, VS Code all support it | Document supported viewers; provide HTML/PDF as alternatives |
| Large risk tables overflow terminal | Low — markdown viewers handle this | Use reasonable table formatting |

## Open Questions (Resolved)

1. **Mermaid syntax:** `flowchart TD` (modern, preferred by mermaid maintainers)
2. **Color coding:** Text labels via existing functions (no CSS in markdown)
3. **Cover page:** Not needed for markdown (single template)
4. **Methodology section:** Included (valuable reference content)

## Files Likely Touched

- `templates/markdown_report.tpl` (new)
- `pkg/taralizer/reporting.go` (~10 lines)
- `cmd/report.go` (~10 lines)
- `pkg/taralizer/reporting_test.go` (new test)
- `README.md` (minor update)
