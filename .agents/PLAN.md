# Spec: `docs/ideas/tara-first-terraform-pipeline.md` · Quality bar: `CONSTRAINTS.md` (command-backed, reviewed 2026-10-11)

## Overview

Make the Taralizer YAML model the source of truth for GCP infrastructure: **export** a reviewable Terraform scaffold from a TARA model (with round-trip markers), **import** a real model from an existing Terraform directory, and **detect structural drift** between the two as a markdown report with recommendations. GCP only; architecture stays extensible to other providers via profiles.

Direction A from the one-pager, in three phases: Task 0 (green suite) → A1 export + markers → A2 import fix = drift engine foundation → A3 structural drift report.

## Architecture Decisions

1. **hcl/v2 is the parser foundation.** `pkg/terraform/parser.go` (`ParseTerraform`) already walks a directory and parses every `.tf` with `hashicorp/hcl/v2` — HashiCorp's own Terraform parser. All new work (Task 0, A2, A3) builds on it. This de-risks one-pager assumption #2 (parser robustness vs real-world HCL): we no longer depend on the custom ANTLR grammar for directory parsing.
2. **ANTLR single-file path becomes legacy.** `ImportFromFile` + `TreeShapeListener` (pkg/terraform/terraform.go) stay in place and tested, but the CLI stops using them after A2. Deletion is a future deprecation task (out of v1 scope).
3. **Round-trip markers (no manifest file — user constraint):**
   - Authoritative: per-asset `terraform_ref` field in the taralizer YAML holding the HCL address (e.g. `google_compute_instance.frontend`), written back by export.
   - Provenance/fallback: single-line `# taralizer:asset=<id>` comment above each generated resource block.
   - Failure mode is explicit: renamed resource → stale ref → drift reports **broken ref**. Never silent fuzzy matching.
4. **Profile-driven export.** A technology is *exportable* iff it has both a `terraform` mapping and a `scaffold` template in `profiles/google.yaml`. Scaffold = minimal HCL body (with `{name}` placeholder) that passes `terraform validate`. This keeps the CONSTRAINTS.md invariant "exported `.tf` passes `terraform validate`" clean and makes new resource types a data change, not a code change (multi-cloud seam).
5. **CLI shape.** `terraform` becomes a parent command with subcommands: `import`, `export`, `drift`. Existing invocation `taralizer terraform import …` keeps working (now as a subcommand).
6. **Resolved open questions from the one-pager:** per-asset `terraform_ref` field (not a top-level section); single-line comment format; export into non-empty dir refuses by default (`--force` to proceed); write-back happens in place on the model file by default (`--no-write-back` to skip).
7. **Known limitation (documented, not fixed in v1):** write-back re-marshals the model via struct-based `yaml.Marshal`, which drops YAML comments from the original file. Data is preserved; a comment-preserving (yaml.v3 Node-based) writer is a future improvement.

## Task List

### Phase 0: Green suite (prerequisite — approved as "1a", no exception)

- [ ] **Task 0: Implement `ParseTerraformDir`**
  - Description: Replace the stub in `pkg/terraform/types.go`. Call existing `ParseTerraform(tfDir)` (hcl/v2), walk each file's body blocks, and for every `resource` block return `GCPResource{Name: labels[1], Type: labels[0]}`. Non-existent dir → error (existing test expects this).
  - Acceptance criteria:
    - `TestParseTerraformDirWithValidFiles` passes (2 resources, correct type/name) — no panic
    - `TestParseTerraformDir` updated: it currently asserts Empty (documents the stub); after this task it must assert the 2 real resources in `examples/gcp/terraform/main.tf`
    - Empty dir → empty slice; non-existent dir → error (existing tests unchanged)
  - Verification: `go test ./pkg/terraform` green, then full `make test` green
  - Dependencies: none
  - Files likely touched: `pkg/terraform/types.go`, `pkg/terraform/parser_test.go` (2 files)
  - Estimated scope: S

### Checkpoint 0: `make test` fully green (suite was RED at baseline)

### Phase A1: Export — YAML → Terraform scaffold + markers

- [ ] **Task A1.1: Add `terraform_ref` to the model**
  - Description: Add `TerraformRef string \`yaml:"terraform_ref,omitempty"\`` to `TechnicalAsset` and `TrustBoundary` in `pkg/taralizer/model.go`.
  - Acceptance criteria: YAML round-trip preserves the field; existing models (no `terraform_ref`) load and re-marshal unchanged (`omitempty`).
  - Verification: unit test in `pkg/taralizer` (marshal → unmarshal → field intact); `go build ./... && go vet ./...`
  - Dependencies: none (parallel with A1.2)
  - Files likely touched: `pkg/taralizer/model.go`, one test file (2 files)
  - Estimated scope: XS

- [ ] **Task A1.2: Expand `profiles/google.yaml` (Option A — approved, implement without review)**
  - Description: Add high-confidence GCP mappings + `scaffold` templates for the technologies used by `examples/gcp/bank_of_anthos.yaml`. Baseline finding: only 1 of its 18 technologies (`vpc`) maps today; the whole profile has 5 non-empty mappings.
    - New mappings (approximate ones marked): `project → google_project`, `subnet → google_compute_subnetwork`, `database-sql → google_sql_database_instance`, `kubernetes-cluster → google_container_cluster`, `waf → google_cloud_armor_security_policy` (approx), `database-nosql → google_firestore_database` (approx), `logging → google_logging_project_sink` (approx), `hsm → google_kms_key_ring` (approx).
    - Intentionally unmapped (workload-level, no google-provider TF resource): `kubernetes-pod`, `kubernetes-ingress`, `kubernetes-network-policies` — add explanatory comments in the profile.
    - Add a `scaffold:` HCL template (minimal body that passes `terraform validate`, `{name}` placeholder) to every exportable technology, including the 5 pre-existing mappings.
    - Add `Scaffold string \`yaml:"scaffold,omitempty"\`` to the Technology structs (both `pkg/taralizer/model.go` and `pkg/terraform/mapping.go`).
  - Acceptance criteria: profile loads via both `LoadProfileSet` and `LoadGoogleProfile`; every technology with a non-empty `terraform` field also has a non-empty `scaffold`; bank_of_anthos.yaml has ≥ 8 exportable assets/boundaries.
  - Verification: unit test asserting mapping+scaffold completeness; manual export in A1.3 produces ≥ 8 resource blocks
  - Dependencies: none (parallel with A1.1)
  - Files likely touched: `profiles/google.yaml`, `pkg/taralizer/model.go`, `pkg/terraform/mapping.go`, one test file (4 files)
  - Estimated scope: M

- [ ] **Task A1.3: Export engine**
  - Description: New `pkg/terraform/export.go`. `ExportToTerraform(report taralizer.Report, profilePath string, outDir string) (refs map[string]string, warnings []string, err error)`:
    - For each TechnicalAsset/TrustBoundary: look up technology in profile; exportable (mapping + scaffold) → emit `# taralizer:asset=<id>` comment + `resource "<type>" "<name>" { <scaffold with name> }`; not exportable → warning (never silent drop — floor rule).
    - Emit `provider.tf` (`google` provider with `var.project`/`var.region`) and `variables.tf`.
    - Refuse non-empty target dir by default; `force bool` param to proceed.
    - Return asset-id → HCL-address map for write-back.
  - Acceptance criteria: exporting bank_of_anthos.yaml produces provider.tf + variables.tf + ≥ 8 resource blocks, each preceded by its `# taralizer:asset=` comment; non-empty dir refused without force; unmapped technologies appear in warnings.
  - Verification: unit tests (temp dir); `go build ./... && go vet ./...`
  - Dependencies: A1.1, A1.2
  - Files likely touched: `pkg/terraform/export.go` (new), `pkg/terraform/export_test.go` (new) (2 files)
  - Estimated scope: M

- [ ] **Task A1.4: `terraform_ref` write-back**
  - Description: After scaffold generation, set `TerraformRef` on each exported asset/boundary and re-marshal the report to the model file (in place by default; `noWriteBack bool` param to skip). Document the comment-loss limitation (decision 7) in the command help.
  - Acceptance criteria: after export, reloading the model YAML shows correct `terraform_ref` values (HCL addresses) on exported assets only; unexported assets have empty refs; `--no-write-back` leaves the file byte-identical.
  - Verification: unit test (export → reload → assert refs); manual run on a copy of bank_of_anthos.yaml
  - Dependencies: A1.3
  - Files likely touched: `pkg/terraform/export.go`, `pkg/terraform/export_test.go` (2 files)
  - Estimated scope: S

- [ ] **Task A1.5: CLI — `terraform` parent + `export` subcommand**
  - Description: Restructure `cmd/terraform.go`: `terraformCmd` becomes a parent (no Run); move existing import behavior into an `import` subcommand unchanged; add `export <model.yaml> -o <dir> [--force] [--no-write-back]`.
  - Acceptance criteria: `taralizer terraform export examples/gcp/bank_of_anthos.yaml -o /tmp/tf` writes the scaffold and updates refs; `taralizer terraform import …` still behaves as before (ANTLR path); help text documents write-back + comment-loss limitation.
  - Verification: `go build ./...`; manual runs of both commands
  - Dependencies: A1.4
  - Files likely touched: `cmd/terraform.go` (1 file)
  - Estimated scope: S

- [ ] **Task A1.6: `terraform validate` CI gate**
  - Description: Add a CI job (`.github/workflows/test.yml`) using `hashicorp/setup-terraform` that exports a fixture model and runs `terraform init -backend=false && terraform validate` on the output. Locally: same sequence as a make target or documented manual step.
  - Acceptance criteria: CI job fails if exported scaffold does not validate; passes on current scaffolds.
  - Verification: run the sequence locally against the bank_of_anthos export; push and watch CI
  - Dependencies: A1.5
  - Files likely touched: `.github/workflows/test.yml`, `Makefile` (2 files)
  - Estimated scope: S

### Checkpoint A1: export demo — `terraform export` on bank_of_anthos.yaml → scaffold passes `terraform validate`, refs written back, suite green

### Phase A2: Import fix = drift engine foundation

- [ ] **Task A2.1: Repoint `import` at hcl/v2 path; actually write the model**
  - Description: New `ImportFromDir(tfDir, profilePath string) (taralizer.Report, []string /*warnings*/, error)` in `pkg/terraform` — uses `ParseTerraformDir` + profile mapping, builds a real `taralizer.Report` struct (replaces the string-concatenated template for this path). Unknown resource types → collected warnings, never a panic (fixes the `findTechnologyByTerraform` nil-pointer class). CLI: `terraform import <model.yaml> <tf_dir>` writes the marshaled report to `<model.yaml>`.
  - Acceptance criteria: `taralizer terraform import out.yaml examples/gcp/terraform` writes a model with the vpc trust boundary + vm technical asset; unknown types printed as warnings on stderr, exit code 0; no panic on any resource type.
  - Verification: unit tests (temp dirs incl. unknown-type fixture); manual run; `make test`
  - Dependencies: Task 0
  - Files likely touched: `pkg/terraform/import.go` (new), `cmd/terraform.go`, one test file (3 files)
  - Estimated scope: M

- [ ] **Task A2.2: Parser robustness fixture corpus (validates one-pager assumption #2)**
  - Description: Add `pkg/terraform/testdata/corpus/*.tf` fixtures covering real-world GCP HCL: `for_each`, `dynamic` blocks, `count`, nested/inline blocks, provider aliases, `locals`, `data` sources. `TestParserCorpus`: every fixture parses without error via `ParseTerraform`; resource extraction returns the symbolic name (labels[1]) for each `resource` block.
  - Acceptance criteria: ≥ 5 fixtures, all parse cleanly; extraction assertions pass for for_each/dynamic/count resources.
  - Verification: `go test ./pkg/terraform -run TestParserCorpus` (this is the CONSTRAINTS.md "parser robustness" row)
  - Dependencies: Task 0
  - Files likely touched: `pkg/terraform/testdata/corpus/*.tf` (new), one test file (2 files + fixtures)
  - Estimated scope: S

### Checkpoint A2: import writes a real model; corpus green; suite green

### Phase A3: Structural drift report

- [ ] **Task A3.1: Drift engine**
  - Description: New `pkg/terraform/drift.go`. `DetectDrift(report taralizer.Report, resources []GCPResource) DriftReport` with four structural categories:
    - **Missing asset**: `terraform_ref` set, address not among parsed resources (type+name)
    - **Broken ref**: `terraform_ref` set, malformed or type exists but name doesn't — reported explicitly, never fuzzy-matched
    - **Unmodeled resource**: parsed resource matched by no `terraform_ref` (comment-based provenance noted when a `# taralizer:asset=` comment is present but ref missing/stale)
    - **Stale marker**: `# taralizer:asset=<id>` comment present but no matching ref in the model
    - Each finding carries a recommendation string (e.g. "add `terraform_ref` to asset X", "model this resource or mark it out_of_scope").
  - Acceptance criteria: unit tests with synthetic model+TF pairs covering all four categories + a clean (no-drift) case.
  - Verification: `go test ./pkg/terraform -run TestDetectDrift`; `make test`
  - Dependencies: A2.1 (needs comment extraction from parsed files — extend `GCPResource` with an optional marker field or a parallel map)
  - Files likely touched: `pkg/terraform/drift.go` (new), `pkg/terraform/types.go`, one test file (3 files)
  - Estimated scope: M

- [ ] **Task A3.2: Markdown drift report**
  - Description: Render `DriftReport` to markdown: summary counts, one section per category, table of findings (address/asset id, detail, recommendation). Clean case renders "no structural drift detected".
  - Acceptance criteria: rendering test asserts section presence + row content for a sample report; clean case output.
  - Verification: unit test; `make test`
  - Dependencies: A3.1
  - Files likely touched: `pkg/terraform/drift.go` (or `drift_report.go`), test file (2 files)
  - Estimated scope: S

- [ ] **Task A3.3: CLI `drift` subcommand**
  - Description: `taralizer terraform drift <model.yaml> <tf_dir> [-o report.md]` — default output stdout, `-o` writes markdown file.
  - Acceptance criteria: round-trip demo — export bank_of_anthos.yaml → `drift` reports clean; add one unmodeled resource to the TF dir → drift lists it under Unmodeled with a recommendation; rename a resource → Broken ref reported.
  - Verification: manual round-trip (documented in task completion); `make test`
  - Dependencies: A3.2
  - Files likely touched: `cmd/terraform.go` (1 file)
  - Estimated scope: XS

### Checkpoint A3 / Final: full round-trip demo (export → clean drift → mutate TF → correct findings); `make all` + `govulncheck ./...`; coverage ratchet ≥ 57.0%

## Parallel Opportunities

- Task A1.1 ∥ A1.2 (independent: model field vs profile data)
- Task A2.1 ∥ A2.2 (both depend only on Task 0)
- Everything else is sequentially dependent as listed.

## Risks and Mitigations

| Risk | Impact | Mitigation |
| ------ | -------- | ----------- |
| Scaffold defaults fail `terraform validate` for some resource types | High — breaks the core invariant | CI validate gate (A1.6) catches per-type failures; iterate scaffold data until green before Checkpoint A1 |
| Approximate profile mappings (waf→Cloud Armor, hsm→KMS) mislead users | Medium — wrong infra scaffolded | Approximations marked with comments in the profile; export warnings list every mapping used; expert review at SHIP |
| Write-back drops YAML comments from the model file | Medium — user-visible data loss (comments only) | Documented in command help + plan decision 7; `--no-write-back` escape hatch; Node-based writer as future task |
| Resource naming collisions (two assets → same HCL address) | Medium — ambiguous refs | Export detects duplicate addresses and errors before writing anything |
| hcl/v2 extraction edge cases (for_each/dynamic) | Low — HashiCorp's own parser | A2.2 corpus pins the behavior; drift keys on symbolic name which is stable under for_each |
| ANTLR legacy path diverges from hcl/v2 behavior | Low — not CLI-wired after A2.1 | Kept tested; deletion tracked as future deprecation task |

## Open Questions (post-v1)

- Delete the ANTLR single-file import path (`ImportFromFile`, generated parser, `TreeShapeListener`) once A2.1 is proven in use — deprecation task with expand/contract discipline
- Comment-preserving YAML writer for write-back (yaml.v3 Node-based)
- Attribute-level drift (beyond structural) — explicitly Not Doing for v1
- CI gate packaging of the whole pipeline (one-pager direction B) — thin follow-up after A3
