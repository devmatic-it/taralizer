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

.PHONY: all build test lint check-coverage enforce-rules clean init

# Build the binary
build:
	@echo "Building binary..."
	CGO_ENABLED=0 go build -v -ldflags="-X 'github.com/devmatic-it/taralizer/cmd.ProductVersion=${VERSION}'" -o dist/taralizer
	cp -R templates dist/templates
	cp -R profiles dist/profiles

# Run tests with race detector and generate coverage
test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# Hard constraints that fail the build if the agent cheats
enforce-rules:
	@echo "Enforcing anti-pattern rules..."
	@if grep -rnw . --include=\*.go -e '//nolint'; then echo "FAIL: //nolint is forbidden"; exit 1; fi
	@if grep -rnw . --include=\*.go -e '_ = err' -e '_ = fmt.Errorf'; then echo "FAIL: Blank identifier error suppression is forbidden"; exit 1; fi
	@if grep -rnw . --include=\*.go -e 't.Skip()'; then echo "FAIL: Skipping tests is forbidden"; exit 1; fi

# Run linters (Hard failures only)
lint: enforce-rules
	@echo "Running go vet..."
	go vet ./...
	@echo "Running staticcheck..."
	staticcheck ./...
	@echo "Running govulncheck..."
	govulncheck ./...

# Evaluate coverage against the threshold
check-coverage: test
	@echo "Checking coverage threshold..."
	@go tool cover -func=coverage.txt | tail -n 1 | awk '{ print $$3 }' | sed 's/%//' > .cov
	@COVERAGE=$$(cat .cov); \
	echo "Current coverage: $$COVERAGE%"; \
	if [ 1 -eq $$(echo "$$COVERAGE < $(COVERAGE_THRESHOLD)" | bc -l) ]; then \
		echo "Coverage $$COVERAGE% is below threshold of $(COVERAGE_THRESHOLD)%"; \
		rm .cov; \
		exit 1; \
	fi
	@rm .cov
	@echo "Coverage OK."

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
