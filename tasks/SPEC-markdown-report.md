# Spec: Markdown Report with Embedded Mermaid Diagrams

## Objective

Add a `markdown` report type that outputs a single `.md` file with embedded mermaid DFD diagrams. Unlike HTML (which references `diagram.png`) and PDF (which uses chromedp), the markdown report embeds the mermaid source code directly, making it natively renderable in GitHub, GitLab, Obsidian, VS Code, and other markdown viewers.

**User story:** As a security analyst, I want a portable, viewable-anywhere markdown report that shows the threat model diagram without needing external image files, so I can share it via email, paste into documentation, or commit to a repo.

**Success criteria:**
- `taralizer report --type markdown` generates a valid `.md` file
- The mermaid DFD is embedded as a ` ```mermaid ` code block (no external images)
- All existing report sections (trust boundaries, assets, risks, methodology) are present
- The output is a single file with no external dependencies
- Existing HTML and PDF report types are unaffected

## Tech Stack

- **Language:** Go 1.27 (already in `go.mod`)
- **Template engine:** existing `text/template` (no new deps)
- **Diagram rendering:** mermaid code blocks (rendered by the markdown viewer, not taralizer)
- **Output:** single `.md` file

## Commands

```
Build:  make build
Test:   go test -v ./pkg/taralizer/...
Report: ./dist/taralizer report model.yaml --type markdown --out report
```

## Project Structure

```
templates/
  markdown_report.tpl  ← NEW: markdown template with embedded mermaid
pkg/taralizer/
  reporting.go         ← Add GenerateReportFileMarkdown()
  reporting_test.go    ← Add integration test
cmd/
  report.go            ← Add "markdown" branch
```

## Code Style

```go
// Follow existing KISS pattern: no hidden state, explicit params, return errors.
func (svc *ReportEngine) GenerateReportFileMarkdown(filename string, tplFile string, report Report) error {
    fo, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("create %s: %w", filename, err)
    }
    defer fo.Close()
    return svc.GenerateReport(fo, tplFile, report)
}
```

- No new dependencies (mermaid renders in the viewer, not in Go)
- Single template file (no cover page needed for markdown)
- Use existing `renderTemplate` and `createFuncMap` (pass report explicitly)

## Testing Strategy

- **Unit test:** Render markdown template → verify it contains ` ```mermaid ` block and expected tables
- **Integration test:** Full flow: model → report data → markdown → verify file exists and is non-empty
- **Manual:** Open in GitHub/GitLab/Obsidian to verify mermaid renders correctly

## Boundaries

- **Always do:** Single file output, no external dependencies, use existing template system
- **Ask first:** Changing the report data model, adding new CLI flags
- **Never do:** Require external tools (mermaid CLI, image converters), break existing report types

## Success Criteria

- [ ] `go build ./...` succeeds
- [ ] `taralizer report --type markdown` generates a valid `.md` file
- [ ] The mermaid DFD is embedded as a code block (no `diagram.png` reference)
- [ ] All report sections present: trust boundaries, technical assets, data assets, threat agents, risk assessment, methodology
- [ ] Existing HTML and PDF reports unaffected
- [ ] Integration test passes

## Open Questions

1. **Mermaid DFD syntax:** The existing `mermaid.tpl` generates a graphviz-style `graph TD` DFD. Should we keep this format (compatible with mermaid) or use `flowchart TD` (newer mermaid syntax)? → **Decision:** Use `flowchart TD` (modern mermaid, same semantics).

2. **Risk table color coding:** HTML uses CSS classes for severity colors. Markdown has no CSS. Options:
   - (a) Use emoji indicators (🔴🟠🟡🟢)
   - (b) Use text labels only (current `severity()` / `likelihood()` functions)
   - (c) Use markdown badges (not universally supported)
   → **Decision:** Use text labels (existing functions already produce "HIGH(9)" etc.), optionally add emoji as a template-level enhancement.

3. **Cover page:** PDF has a separate cover page. Markdown reports typically don't need one. → **Decision:** Single template, no cover page.

4. **Methodology section:** The PDF template includes a methodology/rules section. Should the markdown report include it? → **Decision:** Yes, include it — it's valuable reference content.
