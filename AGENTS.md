# Agent Governance

`CONSTRAINTS.md` is the absolute source of truth. If a conflict arises, constraints win.

## Quality Control

Bypassing checks is strictly forbidden unless an Exception ID is documented in `CONSTRAINTS.md`.

- **No silencing linters** (`//nolint`).
- **No suppressing errors** (no blank identifiers `_ = ...` for error handling).
- **No skipping tests** (`t.Skip()`) to force a green build.

## Operating Rules

1. **State Assumptions:** Before implementing non-trivial changes, list assumptions and ask for confirmation.
2. **Clarify Confusion:** If requirements or code are inconsistent, **STOP** and ask for clarification.
3. **Maintain Scope:** Touch only what is required. Do not refactor adjacent code without explicit instruction.
4. **Edit tool Usage:** When modifying code files, prefer creating or overwriting the entire file using the write tool, or ensure exact line-by-line matches when using the edit tool.

## Audit Requirement

Every PR must include a **Quality Compliance Report**:

- Build Status (`go build ./...`)
- Lint Status (`golangci-lint run ./...`)
- Race Safety (`go test -race ./...`)
- Coverage Delta (e.g., +2.5%)
- Constraint Check (No floor violations)
