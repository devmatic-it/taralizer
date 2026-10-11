# RAAL Constraints Contract

This file defines the absolute quality thresholds for the project. You must verify your work against these constraints using the provided commands. **A non-zero exit code means you have failed and must fix the code.**

## 1. Absolute Floor (Zero Exceptions)

- No `//nolint` comments.
- No blank identifiers `_ = err` or `_ = fmt.Errorf` for error handling. Handle all errors explicitly.
- No skipped tests (`t.Skip()`).
- No secrets or hardcoded credentials.
- `docs/prd.md` and `docs/specs.md` MUST be patched to reflect new commands, flags, or interfaces before a task is marked complete in `.agents/STATE.md`.

## 2. The Verification Commands

You must run these commands via `run_cmd` to verify your work. If any command fails, you are blocked from advancing to the next task.

| Constraint | Command to Execute | Expected Outcome |
| :--- | :--- | :--- |
| **Compilation** | `make build` | Exit code 0 |
| **Static Analysis & Rules**| `make lint` | Exit code 0 (No warnings, no anti-patterns) |
| **Unit Tests** | `make test` | Exit code 0 (All pass, no race conditions) |
| **Coverage** | `make check-coverage` | Exit code 0 (Coverage ≥ 85%) |
| **TF Parser** | `go test ./pkg/terraform -run TestParserCorpus` | Exit code 0 (No panics on fixtures) |

## 3. Tooling Environment

Assume the following are already installed and available in your `PATH`:
- `go`
- `make`
- `staticcheck`
- `govulncheck`
- `terraform` (v1.16.x)
