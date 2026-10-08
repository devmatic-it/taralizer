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

all: build test

.PHONY: all build test lint test-coverage clean

# Build the binary
build:
	@echo "Building binary..."
	CGO_ENABLED=0 go build -v -ldflags="-X 'github.com/devmatic-it/taralizer/cmd.ProductVersion=${VERSION}'" -o dist/taralizer
	cp -R templates dist/templates
	cp -R profiles dist/profiles

# Run tests with coverage
test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# Run linter
test-coverage:
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@cat cover.out >> coverage.txt

# Run golangci-lint
lint:
	@echo "Running linter..."
	golangci-lint run --config=.golangci.yml

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf dist/*
	rm -f cover.out coverage.txt

# Initialize development environment
init:
	@echo "Installing dependencies..."
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Generate terraform parser (requires antlr4)
antml:
	@echo "Generating terraform parser with ANTLR..."
	antlr4 -Dlanguage=Go -no-visitor -package terraform terraform.g4
