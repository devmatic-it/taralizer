# Project Constraints

This document defines the quality, security, and performance standards for the Taralizer project. These constraints are the absolute source of truth (Supremacy Clause) and must be met for any code to be considered "done".

## 1. Quality & Build Standards

All changes must maintain or improve the following metrics. Any regression in these values is a violation of project constraints.

### Build & Compilation

- **Status:** The codebase must always compile successfully. 
- **Command:** `go build ./...`
- **Constraint:** No compilation errors allowed in any branch.

### Linting & Static Analysis

- **Tool:** `golangci-lint`
- **Command:** `golangci-lint run ./...`
- **Constraint:** No new linting errors or warnings. The use of `//nolint` is strictly prohibited unless an Exception ID is documented in this file.

### Testing & Race Safety

- **Tool:** `go test` with race detection.
- **Command:** `go test -race ./...`
- **Constraint:** All tests must pass. No new race conditions or deadlocks allowed.

### Code Coverage

- **Tool:** `go test -coverprofile`
- **Baseline (Current):**
  - `pkg/asvs`: 84.2%
  - `pkg/cwe`: 87.0%
- **Constraint:** Coverage must never decrease below the current baseline for any package.

## 2. Security Standards

- **Tool:** `gosec` (Static Analysis)
- **Command:** `gosec ./...`
- **Constraint:** No new high or medium severity security issues. All user input must be validated before being used in sensitive operations (file I/O, command execution, etc.).

## 3. Performance Standards

- **Constraint:** No significant regressions in execution time for core logic (parsing, scanning). 
- **Requirement:** Any new feature must be profiled using `pprof` if it touches hot paths.

## 4. Governance & Audit

Every Pull Request (PR) must include a **Quality Compliance Report** in its description:

```markdown
### Quality Compliance Report
- **Build Status:** [PASS/FAIL] (`go build ./...`)
- **Lint Status:** [PASS/FAIL] (`golangci-lint run ./...`)
- **Race Safety:** [PASS/FAIL] (`go test -race ./...`)
- **Coverage Delta:** [+X.X%]
- **Constraint Check:** [No Floor violations introduced]
```

**The "No-Silence" Mandate:** Bypassing checks via `//nolint`, blank identifiers for error suppression, or skipping tests is strictly forbidden.
