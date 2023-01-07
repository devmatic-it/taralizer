// Package terraform import/export of terraform files
// Copyright 2021 taralizer authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package terraform

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/devmatic-it/taralizer/pkg/taralizer"
	"log"
)

type Terraform struct {
}

// New creates a new instance of the Taralizer engine.
func NewTerraform() *Terraform {
	instance := Terraform{}
	return &instance
}

// convertMapToRisk converts an untyped instance Risk into a typed one.
func (svc *Terraform) ImportFromFile(terraformFile string) taralizer.Report {
	report := taralizer.Report{}
	input, err := antlr.NewFileStream(terraformFile)
	if err != nil {
		log.Fatalf("terraform.ImportFromFile: cannot open file %v", err)
	}
	lexer := NewTerraformLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewTerraformParser(stream)

	//parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.File_()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(&report), tree)
	return report
}

type TreeShapeListener struct {
	*BaseTerraformListener
	report              *taralizer.Report
	profile             taralizer.ProfileSet
	currentResourceType string
	currentName         string
}

func NewTreeShapeListener(report *taralizer.Report) *TreeShapeListener {
	listener := new(TreeShapeListener)
	listener.report = report
	return listener
}

func (listener *TreeShapeListener) ExitProvider(ctx *ProviderContext) {
	fmt.Printf("Provider: %s\n", listener.currentResourceType)
	listener.profile = taralizer.LoadProfileSet(fmt.Sprintf("%s.yaml", listener.currentResourceType))
}

func (listener *TreeShapeListener) ExitResource(ctx *ResourceContext) {
	fmt.Printf("Resource %s: %s\n", listener.currentResourceType, listener.currentName)
	technology := listener.findTechnologyByTerraform(listener.currentResourceType)
	if technology.Type == "technical_asset" {
		asset := taralizer.TechnicalAsset{}
		asset.Id = listener.currentName
		asset.Technology = technology.Id
		listener.report.TechnicalAssets = append(listener.report.TechnicalAssets, asset)
	} else if technology.Type == "trust_boundary" {
		boundary := taralizer.TrustBoundary{}
		boundary.Id = listener.currentName
		boundary.Technology = technology.Id
		listener.report.TrustBoundaries = append(listener.report.TrustBoundaries, boundary)
	}
}

func (listener *TreeShapeListener) ExitResourcetype(ctx *ResourcetypeContext) {
	str := ctx.GetText()
	listener.currentResourceType = str[1 : len(str)-1]
}

func (listener *TreeShapeListener) ExitName(ctx *NameContext) {
	str := ctx.GetText()
	listener.currentName = str[1 : len(str)-1]
}

// findTechnologyByTerraform locates the technology by terraform
func (listener *TreeShapeListener) findTechnologyByTerraform(terraform string) *taralizer.Technology {
	for i := 0; i < len(listener.profile.Technologies); i++ {
		if listener.profile.Technologies[i].Terraform == terraform {
			return &listener.profile.Technologies[i]
		}
	}
	return nil
}
