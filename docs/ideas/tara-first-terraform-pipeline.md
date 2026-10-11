# TARA-First Terraform Pipeline (taralizer ↔ GCP)

## Problem Statement
How might we make the TARA model (taralizer YAML) the living source of truth for GCP infrastructure — generating a starting Terraform scaffold from it, and later detecting when the deployed reality drifts from what was assessed?

## Recommended Direction
**Direction A: TARA-first pipeline, phased.** Workflow: YAML (written first) → export TF scaffold with markers → engineer finishes + deploys → drift detection reports structural mismatches → human updates the YAML.

Linkage without a manifest file: (1) per-asset `terraform_ref` in the taralizer YAML — authoritative; (2) `# taralizer:asset=<id>` comment above each generated resource — provenance + fallback. The mapping lives where the source of truth lives; engineers see provenance inline in their TF; nothing extra to keep around.

Phases: **A1** export + markers → **A2** import fix (the drift engine — today `terraform import` discards its own result) → **A3** structural drift report.

## Key Assumptions to Validate
- [ ] **Profile richness:** `profiles/google.yaml` can carry per-resource-type defaults for a usable scaffold → test by exporting `examples/gcp/bank_of_anthos.yaml` and judging the output
- [ ] **Parser robustness:** ANTLR HCL parser survives real-world GCP modules (`for_each`, `dynamic` blocks, functions) → run against 3–5 real modules *before* building drift on top
- [ ] **Matching fidelity:** YAML ref + TF comments cover common cases; renamed/moved resources surface as "broken ref", never silent mismatches
- [ ] **Structural-only drift is acceptable for v1** (assets added/removed; no attribute-level comparison)
- [ ] **Stable asset IDs:** renaming IDs in YAML breaks drift → documented convention

## MVP Scope
**In:**
- `taralizer terraform export model.yaml -o tf/` — GCP scaffold: provider block, one resource per technical asset via profile mapping, `variables.tf` for engineer knobs (region/sizes), generated-by header + per-resource comments, writes `terraform_ref` back into the model
- Fix `taralizer terraform import` — make it actually write a YAML model to file (currently discarded; also fixes the nil-pointer panic on unknown resource types)
- `taralizer terraform drift model.yaml ./tf/` — structural diff: unmodeled resources, missing assets, broken refs; markdown report with concrete recommendations
- GCP only (`profiles/google.yaml`)

## Not Doing (and Why)
- **Two-way sync / auto-updating YAML from TF** — human updates the model; auto-merge is a trap
- **Attribute-level drift** (instance size, region changes) — needs per-resource-type schemas in profiles; big lift
- **AWS/Azure** — keep the `ProfileSet` seam, build GCP first
- **State-based drift** (`terraform show -json`) — v2; source-level matches the PR-review workflow
- **Separate manifest file** — user constraint; comments + YAML extension replace it
- **CI gate packaging** — thin follow-up after A3 (a `--fail-on-unmodeled` flag + one GitHub Action)

## Open Questions
- Exact YAML shape for the ref: per-asset `terraform_ref` field vs. a top-level `terraform:` section — decide at A1 design
- Comment format: single line above the resource, or block markers around it?
- Export into a non-empty TF dir: refuse by default, `--force` to overwrite?

## Context (codebase state at ideation time)
- Working: ANTLR HCL parser (`pkg/terraform/terraform_parser.go`, ~5.6k lines), GCP profile mapping (`pkg/terraform/mapping.go` + `profiles/google.yaml`)
- Broken: `taralizer terraform import` CLI discards its result (`cmd/terraform.go` — validates 1 arg, usage declares 2, `Run` only reads `args[0]`, Report never written); nil-pointer panic in `findTechnologyByTerraform` on unknown resource types (`pkg/terraform/terraform.go`); `ParseTerraformDir` is a stub returning an empty slice (`pkg/terraform/types.go`)
- Missing: any YAML → TF export path; any drift detection
