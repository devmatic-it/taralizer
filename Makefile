# Copyright 2021 taralizer authors
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

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/dist
PKG := "github.com/devmatic-it/taralizer"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
COVERAGE_THRESHOLD := 85

all: lint build test

.PHONY: all build test lint check-coverage enforce-rules clean init check-tf

# Build the binary
build:
	@echo "Building binary..."
	CGO_ENABLED=0 go build  -json -v -ldflags="-X 'github.com/devmatic-it/taralizer/cmd.ProductVersion=${VERSION}'" -o dist/taralizer
	cp -R templates dist/templates
	cp -R profiles dist/profiles

# Run tests with race detector and generate coverage
test:
	@echo "Running tests..."
	go test -json -race -coverprofile=coverage.txt -covermode=atomic ./...

# Validate exported Terraform scaffolds with the terraform CLI (requires terraform).
# Exports the Bank of Anthos example model and runs `terraform init` +
# `terraform validate` on the result, so scaffold regressions (bad GCP names,
# missing required arguments, unknown resource types) fail the build.
check-tf: build
	@echo "Validating exported Terraform scaffolds..."
	sh scripts/check-tf.sh ./dist/taralizer examples/gcp/bank_of_anthos.yaml

# Hard constraints that fail the build if the agent cheats
enforce-rules:
	@echo "Enforcing anti-pattern rules..."
	sh scripts/enforce-rules.sh

# Run linters (Hard failures only)
lint: enforce-rules
	@echo "Running go vet..."
	go vet -json ./...
	@echo "Running staticcheck..."
	staticcheck ./...
	@echo "Running govulncheck..."
	govulncheck -json ./...

# Evaluate coverage against the threshold
check-coverage: test
	@echo "Checking coverage threshold..."
	sh scripts/check-coverage.sh $(COVERAGE_THRESHOLD)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf dist/*
	rm -f cover.out coverage.txt .cov

# Initialize development environment
init:
	@echo "Installing dependencies..."
	go mod download
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
