# Contributing to Taralizer

Thank you for your interest in contributing to Taralizer! This document provides
guidelines and information for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Adding Security Rules](#adding-security-rules)
- [Adding Report Templates](#adding-report-templates)
- [Submitting a Pull Request](#submitting-a-pull-request)
- [Style Guide](#style-guide)

## Code of Conduct

This project adheres to a [Code of Conduct](CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this code.

## Getting Started

### Prerequisites

- Go 1.23 or later
- Chrome or Chromium (for PDF report generation)
- Git

### Setup

```bash
# Clone the repository
git clone https://github.com/devmatic-it/taralizer.git
cd taralizer

# Install dependencies
go mod download

# Run tests to verify setup
make test

# Build the binary
make build
```

## Making Changes

### Project Structure

```
taralizer/
├── cmd/              # CLI commands (report, rules, validate, version)
├── pkg/              # Core packages
│   ├── taralizer/    # Main analysis engine
│   ├── asvs/         # OWASP ASVS database
│   ├── cwe/          # CWE database
│   └── terraform/    # Terraform parser
├── rules/            # OPA Rego security rules
│   ├── asvs/         # ASVS-specific rules
│   ├── core.rego     # Shared rule logic
│   └── validation.rego # Validation helpers
├── templates/        # Report templates (HTML, PDF, Markdown)
├── profiles/         # Technology profiles
├── examples/         # Example architecture models
└── dist/             # Distribution artifacts
```

### Adding Security Rules

Taralizer uses [Open Policy Agent (OPA)](https://www.openpolicyagent.org) with
[Rego](https://www.openpolicyagent.org/docs/latest/policy-language/) for
security rules.

#### Rule Structure

Each rule is a `.rego` file in `rules/asvs/`:

```rego
# Copyright 2024 taralizer authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
package rules.asvs

import data.rules.technical_asset_by_id
import data.rules.calc_impact

# METADATA
# title: Rule Title
# description: Rule description
# custom:
#   rule: "rule-id"
#   mitigation: "How to fix"
#   cwe: 123
#   likelihood: 2
#   impact: 2

violation[{
    "id": id,
    "msg": msg,
    "likelihood": likelihood,
    "impact": impact,
}] {
    server := input.technical_assets[_]
    server.technology == "web-application"
    # your condition here
    id := sprintf("rule-id@%v", [server.id])
    msg := sprintf("description of violation for %v", [server.id])
    likelihood := 2
    impact := calc_impact(2)
}
```

#### Metadata Fields

| Field | Required | Description |
|-------|----------|-------------|
| `title` | Yes | Human-readable rule title |
| `description` | Yes | Rule description |
| `custom.rule` | Yes | Unique rule identifier (snake_case) |
| `custom.mitigation` | Yes | How to remediate |
| `custom.cwe` | No | CWE identifier (integer) |
| `custom.stride` | No | STRIDE threat category |
| `custom.likelihood` | Yes | 1-5 scale (1=low, 5=critical) |
| `custom.impact` | Yes | 1-5 scale (1=low, 5=critical) |

#### Testing Rules

```bash
# Run all tests
make test

# Run specific test
go test ./pkg/taralizer -run TestTaralizerRules -v
```

## Adding Report Templates

Report templates are Go templates that process the risk assessment data.

### Template Variables

| Variable | Type | Description |
|----------|------|-------------|
| `.Title` | string | Report title |
| `.Customer` | string | Customer/organization name |
| `.Date` | string | Report generation date |
| `.Author.Name` | string | Report author |
| `.Risks` | []Risk | List of identified risks |

### Risk Object

```go
type Risk struct {
    Id         string
    Title      string
    Message    string
    Likelihood int
    Impact     int
    Severity   int
    Cwe        int
    Url        string
}
```

### Custom Functions

Templates support custom functions:

- `likelihood(n int) string` — Converts likelihood score to label
- `impact(n int) string` — Converts impact score to label
- `severity(n int) string` — Converts severity score to label

## Submitting a Pull Request

### Before You Start

1. Check existing [issues](https://github.com/devmatic-it/taralizer/issues) for related work
2. Fork the repository
3. Create a feature branch: `git checkout -b feature/your-feature`

### Checklist

- [ ] Tests pass (`make test`)
- [ ] Code is linted (`make lint`)
- [ ] Documentation is updated
- [ ] Changes are documented in the PR description
- [ ] New features include example usage
- [ ] Security rules follow the established pattern

### PR Guidelines

1. **Title**: Clear, concise description (e.g., "Add XSS protection rule")
2. **Description**: Explain what and why, not just what code changed
3. **Screenshots**: Include before/after for UI changes
4. **Examples**: Provide example YAML models for new rules
5. **Tests**: Include tests for new functionality

### Commit Guidelines

We use [Conventional Commits](https://www.conventionalcommits.org/):

```
type(scope): description

fix(rules): add SSRF protection rule
feat(report): support custom report templates
docs(readme): update quick start section
test(asvs): add unit tests for auth rules
```

**Types**: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

## Style Guide

### Go Code

- Follow [Effective Go](https://go.dev/doc/effective_go)
- Use `gofmt` and `goimports`
- All public functions must have documentation comments
- Error handling: never ignore errors, use `fmt.Errorf` with context

### Rego Rules

- Include copyright header (see rule template above)
- Use `# METADATA` block for rule metadata
- Keep rules focused on a single security concern
- Use `sprintf` for unique violation IDs
- Reference `calc_impact` for dynamic impact calculation

### Markdown

- Use proper headings hierarchy
- Include tables where appropriate
- Link to external resources
- Use code blocks with language labels

## Getting Help

- [GitHub Issues](https://github.com/devmatic-it/taralizer/issues)
- [GitHub Discussions](https://github.com/devmatic-it/taralizer/discussions)

We're happy to help! Don't hesitate to ask questions.

---

Thank you for contributing to Taralizer! 🎉
