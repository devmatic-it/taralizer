# Taralizer Architecture Specification

## Overview

Taralizer is a tool designed to analyze Terraform configurations and generate architectural diagrams and technical asset inventories. It focuses on extracting information from GCP Terraform files to create structured output that can be used for security analysis, compliance checking, and infrastructure documentation.

## Architecture Diagram

```mermaid
graph TD
    A[Terraform Files] --> B[Parser]
    B --> C[GCPResource Objects]
    C --> D[Terraform-to-YAML Command]
    D --> E[YAML Template]
    F[Main Application] --> G[CLI Interface]
    G --> H[Terraform-to-YAML Command]
    G --> I[Other Commands]
```

## Component Details

### 1. Main Application Structure

```mermaid
graph LR
    A[main.go] --> B[cmd/root.go]
    B --> C[Execute Function]
    C --> D[Terraform-to-YAML Command Handler]
    D --> E[ParseTerraformDir Function]
    E --> F[GCPResource Type]
    F --> G[YAML Generation]
```

### 2. Terraform Parsing Process

```mermaid
graph TD
    A[ParseTerraformDir] --> B[Read Directory]
    B --> C[Iterate Files]
    C --> D[Parse Each File]
    D --> E[HCL v2 Parsing]
    E --> F[Extract Resource Blocks]
    F --> G[GCPResource Creation]
    G --> H[Return Slice of Resources]

```

### 3. Data Flow

```mermaid
sequenceDiagram
    participant M as Main Application
    participant P as Parser
    participant T as Terraform Files
    participant Y as YAML Generator
    
    M->>P: Execute terraform-to-yaml command
    P->>T: Read Terraform directory
    T->>P: Return HCL files content
    P->>P: Parse HCL using HCL v2 API
    P->>P: Extract resource blocks
    P->>P: Create GCPResource objects
    P->>Y: Pass slice of GCPResource objects
    Y->>M: Generate YAML template
```

## Key Components

### 1. GCPResource Type

The `GCPResource` type is defined in `pkg/terraform/types.go` and contains the following fields:

```go
type GCPResource struct {
	Type string
	Name string
	// Additional fields can be added as needed for more complex scenarios
}
```

### 2. Parser Functions

#### ParseTerraformDir Function
- Located in `pkg/terraform/parser.go`
- Takes a directory path as input
- Parses all Terraform files in the directory
- Returns a slice of `GCPResource` objects

#### ParseTerraform Function  
- Located in `pkg/terraform/parser.go`
- Takes a single file path as input
- Parses the specific Terraform file
- Returns a slice of `GCPResource` objects

### 3. CLI Interface

The main application is structured using Cobra for command-line interface management:
- `main.go` - Entry point
- `cmd/root.go` - Root command with Execute function
- Commands are registered in the root command

### 4. YAML Generation

The `terraform-to-yaml` command uses the parsed GCP resources to generate a structured YAML template that can be used for further analysis or documentation.

## Data Flow

1. User executes `taralizer terraform-to-yaml` command
2. CLI interface parses the arguments and calls `ParseTerraformDir`
3. `ParseTerraformDir` reads the specified directory and processes all Terraform files
4. Each file is parsed using HCL v2 API to extract resource blocks  
5. Resource information is converted into `GCPResource` objects
6. The slice of resources is passed to the YAML generation component
7. A structured YAML template is generated and output

## Technology Stack

- **Language**: Go
- **Parsing Library**: HashiCorp HCL v2
- **Command Line Interface**: Cobra
- **Testing Framework**: testify/assert
- **Documentation Generation**: Mermaid diagrams

## File Structure

```
.
├── main.go
├── cmd/
│   └── root.go
├── pkg/
│   ├── terraform/
│   │   ├── parser.go
│   │   ├── types.go
│   │   ├── mapping.go
│   │   ├── template_generator.go
│   │   └── parser_test.go
└── docs/
    └── spec.md
```

## Usage Example

```bash
# Generate YAML from Terraform files in a directory
./taralizer terraform-to-yaml --directory /path/to/terraform/files

# Show help for the command
./taralizer terraform-to-yaml --help
```

## Phase A1: Model → Terraform Export (supersedes the HCL v2 description above)

> Note: the original sections describe the early `terraform-to-yaml` prototype. The parser is now ANTLR-based (grammar in `.antlr`, generated code in `pkg/terraform/*_parser.go` / `*_lexer.go` — excluded from the coverage gate per CONSTRAINTS.md §3), and the Terraform integration is organized as a `terraform` command group. This section documents the current architecture.

### 1. Export Pipeline (`pkg/terraform/export.go`)

```text
taralizer.Report (YAML model)
  → LoadGoogleProfile(profilePath)          # profiles/google.yaml: technology → terraform type + scaffold
  → for each technical asset / trust boundary:
      MapResourceToTechnology(tech)         # nil or empty scaffold ⇒ warning, object skipped (never dropped silently)
      sanitizeHCLName(id)                   # [^a-zA-Z0-9_] → _ ; hyphens normalized to underscores
      gcpResourceName(id)                   # lowercase, hyphenated — GCP names forbid underscores
      render scaffold with {name} → gcpName # HCL label keeps the sanitized name
  → write <hclName>.tf per object + provider.tf + variables.tf
  → return map[id] = HCL address (e.g. "google_compute_network.vpc")
```

Guards: the target directory must not exist or be empty unless `force` is set; duplicate HCL resource names abort the export before anything is written.

### 2. Markers (drift-detection contract, Phase A3)

Each exported resource block is preceded by a marker comment:

```hcl
# taralizer:asset=vpc
resource "google_compute_network" "vpc" {
  name                    = "vpc"
  auto_create_subnetworks = false
}
```

Prefixes are `# taralizer:asset=<id>` and `# taralizer:boundary=<id>` (constants `AssetMarkerPrefix` / `BoundaryMarkerPrefix`). The asset/boundary distinction keeps the mapping unambiguous when an asset id and a boundary id collide.

### 3. Write-back (`WriteBackTerraformRefs`)

After export, each object's HCL address is written into the model file as its `terraform_ref` field (added to `TechnicalAsset` and `TrustBoundary`, `omitempty`). The edit is surgical and line-level (insert below `- id:` or update in place), preserving comments and formatting. It refuses to write if any ref id is missing from the model or if the edited document no longer parses as a `taralizer.Report`. Idempotent: re-running produces no change.

### 4. CLI (`cmd/terraform.go`)

```console
taralizer terraform export <model.yaml> [-o dir] [--profile p] [--force] [--write-back=false]
taralizer terraform import <terraform_dir> -o model.yaml [--profile p]   # Phase A2
```

`terraform` is a parent command (help only). Warnings go to stderr; the exit code reflects export success.

### 5. Scaffolding Rules (profile-driven)

- A technology is exportable iff its profile entry has a non-empty `terraform` resource type **and** a non-empty `scaffold` template.
- The scaffold body is the resource's attribute block; `{name}` renders as the GCP-safe name (lowercase, hyphens), while the HCL resource label uses the sanitized id.
- `provider.tf` (google provider) and `variables.tf` are always written alongside the per-object files.
- Scaffold correctness is gated in CI: `make check-tf` exports the Bank of Anthos example and runs `terraform init -backend=false && terraform validate` (`.github/workflows/test.yml`, job `terraform-validate`).

### 6. Coverage Gate (Phase A1.7)

`make check-coverage` requires ≥ 65% of project code, with generated files (header `// Code generated ... DO NOT EDIT.`) excluded by header detection in `scripts/check-coverage.sh`. Tests must not exist solely to inflate coverage of generated code (CONSTRAINTS.md §3). Current state: ~85% gated.

## Phase A2: Terraform Import (Terraform → Model)

### 1. Import Pipeline (`pkg/terraform/import.go`)

```text
dir/*.tf (walked in lexical order)
  → per file: ANTLR parse (Terraform.g4), syntax errors abort the import
  → per resource block: type + name (labels)
      MapResourceToTechnology(type)   # nil ⇒ warning, skipped (never dropped silently)
      technical_asset  → TechnicalAsset{Id: name, Name: tech.Name, Technology: tech.ID, TerraformRef: type.name}
      trust_boundary   → TrustBoundary{same shape}
  → taralizer.Save(report, outFile)   # YAML model file
```

Guards: the directory must exist and be a directory; duplicate HCL addresses warn and keep the first occurrence. `data`/`locals`/`provider`/`module`/`output` blocks are ignored — only `resource` blocks import. When several technologies share one terraform type (e.g. `google_compute_instance`), the first profile match wins; markers (Phase A3) will disambiguate later.

### 2. Model Persistence (`taralizer.Save`)

`pkg/taralizer/model.go` gained `Save(report, fileName)` — YAML marshal + write, the inverse of `Load`. The import command writes its report through it.

### 3. CLI (`cmd/terraform.go`)

The legacy single-file `import <file>` (which printed to stdout and discarded its report) is replaced by directory-based import: `taralizer terraform import <dir> -o model.yaml [--profile p]`. Warnings go to stderr; the exit code reflects import success. The legacy `Terraform.ImportFromFile` and its `TreeShapeListener` were removed in Phase A2.3: the single-file path was fully superseded by `ImportFromDir` and strictly inferior (debug output on stdout, a per-resource-type profile lookup that could not resolve, and a nil dereference on unmapped types instead of a warning), so no deprecation shim was kept (KISS).

### 4. Parser Robustness (Phase A2.2)

Import parses real-world HCL2, not just the minimal scaffold subset. Two mechanisms cover constructs the ANTLR grammar does not model natively:

- **Heredoc normalization** (`pkg/terraform/heredoc.go`): before ANTLR parsing, `normalizeHeredocs` rewrites arbitrary-tag HCL2 heredocs (`<<TAG`, `<<-TAG`) into a placeholder string literal, preserving the dropped body's newlines so parse-error line numbers stay accurate. Unterminated heredocs are left untouched so the parser reports its normal syntax error. This keeps the grammar minimal (no per-tag lexer rules) while accepting `metadata_startup_script = <<SCRIPT ...` and friends.
- **Corpus** (`pkg/terraform/testdata/corpus/*.tf` + `TestParserCorpus`): nine fixtures covering comments, heredocs, dynamic blocks, nested blocks, functions/operators, collections, identifiers, all declaration types, and a real-world GCP configuration. Each fixture must parse with zero syntax errors; fixtures with resource blocks carry a `.expected` sidecar listing the exact `type.name` addresses in file order (fixtures without one must contain no resources). The corpus is what drove the grammar extensions for object-constructor for-expressions (`{ for k, v in m : k => v }`), unary minus / logical NOT, general postfix indexing (`[10, 20][2]`, `f(x)[0]`), and the `=>` operator.

Dynamic blocks need no special grammar support: in HCL2 they are ordinary labeled blocks (`dynamic "label" { for_each = ... content { ... } }`) and parse as such.

## Phase A3: Structural Drift Gate (Model ↔ Terraform)

### 1. Drift Pipeline (`pkg/terraform/drift.go`)

```text
model.yaml ── taralizer.Load ─────────────┐
tf_dir/*.tf (walked in lexical order)     │
  → per file: ANTLR parse, syntax errors abort (as in import)
  → per resource block: type.name address + file:line
  → per line: taralizer export markers (asset/boundary)   │
                                                     ▼
                              categorized DriftReport
```

`DetectDrift(modelPath, tfDir)` compares the two sides and returns a `DriftReport` with four structural categories:

| Category | Side | Semantics |
| --- | --- | --- |
| `MissingAssets` | model | entry whose `terraform_ref` is well-formed but has no matching resource block in the directory (deleted or renamed) |
| `BrokenRefs` | model | entry whose `terraform_ref` is not a valid HCL address — exactly one dot separating non-empty type and name parts (corrupt model) |
| `Unmodeled` | Terraform | resource block whose address no model entry references via `terraform_ref` (added outside the model) |
| `StaleMarkers` | Terraform | export marker (`# taralizer:asset=<id>` / `# taralizer:boundary=<id>`) whose id no longer exists in the model (assets or boundaries) |

Design decisions:
- **`terraform_ref` is the sole model↔Terraform linkage** for `Unmodeled`. Markers only feed `StaleMarkers`; a resource with a marker but no model ref is still unmodeled. This keeps the two signals independent and the rule trivially explainable.
- Model entries **without** a `terraform_ref` have no Terraform linkage and are not checked (they are not drift).
- Duplicate HCL addresses across files warn (`duplicate resource address … using the first occurrence`); for `MissingAssets`/ref-owner counting the first occurrence wins, while an unmodeled duplicate address reports **each** occurrence with its own `file:line`.
- Findings carry stable, greppable data: model-side items have `ID` + `Ref`, Terraform-side items have the address as `ID` plus a `file:line` location (stale markers append `(asset marker)` / `(boundary marker)`).
- Attribute-level drift is **out of scope** for Phase A3 (see PRD §14.2, Phase A4).

### 2. Markdown Report (`DriftReport.Markdown`)

The report is a markdown document: a header with model path, Terraform directory and total finding count, then one section per **non-empty** category in fixed order (Missing assets, Broken refs, Unmodeled resources, Stale markers), each with a two-column table. Warnings (e.g. duplicate addresses) render last when present. A report with zero findings renders the header followed by `No drift detected.`

```markdown
# Taralizer Drift Report

- Model: `model.yaml`
- Terraform dir: `./tf`
- Findings: 2

## Missing assets (1)

| ID | terraform_ref |
| --- | --- |
| `vm` | `google_compute_instance.vm` |

## Unmodeled resources (1)

| Resource | Location |
| --- | --- |
| `google_compute_disk.data` | `tf/extra.tf:1` |
```

### 3. CLI and Exit-Code Contract (`cmd/terraform.go`)

```console
taralizer terraform drift <model.yaml> <terraform_dir> [-o report.md]
```

- The markdown report goes to **stdout**, or to `-o <file>` when given; warnings always go to stderr.
- **Exit code 0** ⇔ no findings (model and Terraform in sync). **Non-zero exit whenever any drift is found** — the command is a CI gate. On an expected failure (drift detected) the usage block is suppressed; argument and I/O errors behave like every other command.
- The clean round-trip is the acceptance test: a model exported with `terraform export` (write-back on) drifts clean against its own exported directory.

---

This specification outlines the architecture and components of taralizer, providing a clear understanding of how it processes Terraform configurations and generates technical asset inventories.