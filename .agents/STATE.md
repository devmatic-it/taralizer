# Agent State Tracker: TARA-First Terraform Pipeline

## Phase 0: Green Suite (Prerequisite)

- [X] Task 0: Implement `ParseTerraformDir` in `pkg/terraform/types.go` (replace stub, call `ParseTerraform`, extract `GCPResource`). Update `TestParseTerraformDir` in `pkg/terraform/parser_test.go`.

## Phase A1: Export (YAML → Terraform scaffold + markers)

- [X] Task A1.1: Add `terraform_ref` to `TechnicalAsset` and `TrustBoundary` in `pkg/taralizer/model.go` with `omitempty`. Add unit tests.
- [ ] Task A1.2: Expand `profiles/google.yaml` with mappings and `scaffold` templates for Bank of Anthos components. Add `Scaffold` string to Technology structs in `pkg/taralizer/model.go` and `pkg/terraform/mapping.go`.
- [ ] Task A1.3: Create `pkg/terraform/export.go` and `pkg/terraform/export_test.go`. Implement `ExportToTerraform` matching the acceptance criteria (markers, provider, refuse non-empty dir).
- [ ] Task A1.4: Implement `terraform_ref` write-back logic in `pkg/terraform/export.go`.
- [ ] Task A1.5: Restructure `cmd/terraform.go` to have `terraform` parent command and `export` subcommand. Update help docs.
- [ ] Task A1.6: Add `.github/workflows/test.yml` for `terraform validate` CI gate. Update `Makefile` if necessary.
- [ ] **DOCUMENT:** Update `docs/specs.md` and `docs/prd.md` to reflect Phase A1 export architecture, scaffolding rules, and CLI changes.

## Phase A2: Import Fix (Drift engine foundation)

- [ ] Task A2.1: Create `pkg/terraform/import.go`. Implement `ImportFromDir` building a real `taralizer.Report`. Update `cmd/terraform.go` to use this for the `import` subcommand.
- [ ] Task A2.2: Add parser robustness fixtures in `pkg/terraform/testdata/corpus/*.tf` and implement `TestParserCorpus` in Go tests.
- [ ] **DOCUMENT:** Update `docs/specs.md` to reflect the new `import` behavior, directory parsing, and unknown-type handling.

## Phase A3: Structural Drift Report

- [ ] Task A3.1: Create `pkg/terraform/drift.go`. Implement `DetectDrift` categorizing missing assets, broken refs, unmodeled resources, and stale markers. Add tests.
- [ ] Task A3.2: Implement markdown rendering for `DriftReport` in `pkg/terraform/drift.go` (or `drift_report.go`). Add tests.
- [ ] Task A3.3: Add `drift` subcommand to `cmd/terraform.go` (`drift <model.yaml> <tf_dir> [-o report.md]`).
- [ ] **DOCUMENT:** Update `docs/specs.md` and `docs/prd.md` to reflect the final drift CLI command and structural report format.
