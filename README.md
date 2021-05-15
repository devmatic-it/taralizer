# Taralizer - The Threat and Risk Analyzer

[![Go Report Card](https://goreportcard.com/badge/github.com/devmatic-it/taralizer)](https://goreportcard.com/report/github.com/devmatic-it/tarlizer)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/devmatic-it/taralizer/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/devmatic-it/taralizer/branch/master/graph/badge.svg)](https://codecov.io/gh/devmatic-it/taralizer)

The following project performs a Threat and Risk Analysis based on an architecture model defined through  YAML file.

## Motivation

The taralizer project was highly inspired by the Threagile (<https://threagile.io>) project which is a first class threat modelling tool for OWASP ASVP.
Unfortunately, threagile has some limits in the customization of reports and the extendability with custom rules.
Taralizer tries to overcome these limitations with the following approach:

- Using Golang templating (<https://golang.org/pkg/text/template/>) for all reports and diagrams
- Use the Open Policy Agent (OPA) engine (<https://www.openpolicyagent.org>) to allow extentabilty
- use plantuml or graphviz dot to generate compelling diagrams
- use of wkhtmltopdf to create PDF reports

## Installation

### Binary

1. Download latest release for your platform: <https://github.com/devmatic-it/taralizer/releases/latest>
2. extract archive: `tar xvfz taralizer_X.Y.Z_linux_amd64.tgz`

## Contribute

### New Issues

1. Use the search tool before opening a new issue: <https://github.com/devmatic-it/taralizer/issues>
2. Please provide source code and commit fix if you found a bug.
3. Review existing issues and provide feedback or react to them.

### Pull requests

1. Open your pull request against master:  <https://github.com/devmatic-it/taralizer/pulls>
2. Your pull request should have no more than two commits, if not you should squash them.
3. It should pass all tests in the available continuous integrations systems such as TravisCI.
4. You should add/modify tests to cover your proposed code changes.
5. If your pull request contains a new feature, please document it on the <https://github.com/devmatic-it/taralizer/blob/master/README.md>

## Credits

This work has ben inspired by the following open source projects:

- Threagile - Agile Threat Modelling (<https://threagile.io>)
- Open Policy Agent (<https://www.openpolicyagent.org>)
- PlantUML (<https://plantuml.com>)
- GraphViz (<https://graphviz.org>)
- WKhtmltoPDF (<https://wkhtmltopdf.org>)
- GoRleaser Builder Image (<https://github.com/goreleaser/goreleaser>)
- Building a basic CI/CD pipeline for a Golang application using GitHub Actions
(<https://dev.to/brpaz/building-a-basic-ci-cd-pipeline-for-a-golang-application-using-github-actions-icj>)
