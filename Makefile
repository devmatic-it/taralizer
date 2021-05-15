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
VERSION := 1.2.3.4
all:  compile test-coverage

compile: build

init: get compile

get:
	@echo "Downloading dependencies..."	
	GOBIN=$(GOBIN) go get -u -v github.com/xuri/xgen/cmd/...
	GOBIN=$(GOBIN) go get github.com/securego/gosec/v2/cmd/gosec
	GOBIN=$(GOBIN) go get

security:
	@echo "Gosec security scan..."
	dist/gosec ./...

test: compile
	go test -v ${PKG_LIST} 

build:
	@echo "Building binary..."
	GOBIN=$(GOBIN) go build -v -ldflags="-X 'cmd.ProductVersion=v1.0.0'"  -o dist/taralizer
	
test-coverage:
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST} 
	@cat cover.out >> coverage.txt

clean:
	@echo "Cleanup dependencies..."	
	rm -Rf ./src/github.com 
	rm -Rf ./src/golang.org
	rm -Rf ./src/gopkg.in
	rm -Rf ./dist/*
	rm -f cover.out coverage.txt gosec_report.html
