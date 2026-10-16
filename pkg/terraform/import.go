// Package terraform import engine (Phase A2): builds a taralizer report from
// Terraform files.
// Copyright 2021 taralizer authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package terraform

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/devmatic-it/taralizer/pkg/taralizer"
)

// ImportFromDir parses every .tf file under dir and builds a taralizer.Report
// from the resource blocks it finds. Each resource whose type maps to a
// technology in the profile at profilePath becomes a TechnicalAsset or
// TrustBoundary (per the technology's type), with Id set to the HCL resource
// name and TerraformRef set to the full HCL address (type.name) so that export
// write-back and drift detection can round-trip.
//
// Resource types unknown to the profile produce a warning — they are never
// dropped silently. Files with syntax errors abort the import with an error.
func ImportFromDir(dir, profilePath string) (taralizer.Report, []string, error) {
	profile, err := LoadGoogleProfile(profilePath)
	if err != nil {
		return taralizer.Report{}, nil, fmt.Errorf("load profile %s: %w", profilePath, err)
	}

	files, err := listTerraformFiles(dir)
	if err != nil {
		return taralizer.Report{}, nil, err
	}

	report := taralizer.Report{}
	var warnings []string
	imported := map[string]bool{} // HCL address -> already imported

	for _, file := range files {
		resources, err := parseResourceBlocks(file)
		if err != nil {
			return taralizer.Report{}, warnings, fmt.Errorf("parse %s: %w", file, err)
		}

		for _, res := range resources {
			addr := res.typ + "." + res.name
			if imported[addr] {
				warnings = append(warnings, fmt.Sprintf("duplicate resource address %s in %s: keeping the first occurrence", addr, file))
				continue
			}

			tech := profile.MapResourceToTechnology(res.typ)
			if tech == nil {
				warnings = append(warnings, fmt.Sprintf("unmapped resource type %q (resource %s in %s)", res.typ, addr, file))
				continue
			}

			switch tech.Type {
			case "technical_asset":
				report.TechnicalAssets = append(report.TechnicalAssets, taralizer.TechnicalAsset{
					Id:           res.name,
					Name:         tech.Name,
					Technology:   tech.ID,
					TerraformRef: addr,
				})
			case "trust_boundary":
				report.TrustBoundaries = append(report.TrustBoundaries, taralizer.TrustBoundary{
					Id:           res.name,
					Name:         tech.Name,
					Technology:   tech.ID,
					TerraformRef: addr,
				})
			default:
				warnings = append(warnings, fmt.Sprintf("unknown technology type %q for resource %s in profile", tech.Type, addr))
			}
			imported[addr] = true
		}
	}

	return report, warnings, nil
}

// resourceRef is one parsed resource block (type + HCL name + line).
type resourceRef struct {
	typ  string
	name string
	line int // 1-based line of the `resource` keyword, for drift locations
}

// listTerraformFiles returns the .tf files under dir in lexical order.
func listTerraformFiles(dir string) ([]string, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return nil, fmt.Errorf("stat %s: %w", dir, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", dir)
	}

	var files []string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".tf") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk %s: %w", dir, err)
	}
	return files, nil
}

// parseErrorCounter counts syntax errors reported through the ANTLR error
// listener channel.
type parseErrorCounter struct {
	*antlr.DefaultErrorListener
	count int
}

func (l *parseErrorCounter) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.count++
}

// parseResourceBlocks parses one .tf file and returns its resource blocks.
func parseResourceBlocks(path string) ([]resourceRef, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}

	// Normalize heredocs with arbitrary delimiters (<<SCRIPT, ...) to string
	// literals; the ANTLR grammar only knows the fixed <<EOF / <<-EOF /
	// <<DESCRIPTION tags.
	src := normalizeHeredocs(string(content))

	input := antlr.NewInputStream(src)
	lexer := NewTerraformLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewTerraformParser(stream)

	// Replace the default console listener so errors surface via the return
	// value instead of stderr noise.
	parser.RemoveErrorListeners()
	counter := &parseErrorCounter{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	parser.AddErrorListener(counter)

	tree := parser.File_()
	if counter.count > 0 {
		return nil, fmt.Errorf("%d syntax error(s)", counter.count)
	}

	fileCtx, ok := tree.(*File_Context)
	if !ok {
		return nil, fmt.Errorf("unexpected parse tree %T", tree)
	}

	var refs []resourceRef
	for _, res := range fileCtx.AllResource() {
		rc, ok := res.(*ResourceContext)
		if !ok {
			return nil, fmt.Errorf("unexpected resource context %T", res)
		}
		refs = append(refs, resourceRef{
			typ:  unquote(rc.Resourcetype().GetText()),
			name: unquote(rc.Name().GetText()),
			line: rc.GetStart().GetLine(),
		})
	}
	return refs, nil
}

// unquote strips the surrounding quotes from an HCL string literal.
func unquote(s string) string {
	return strings.Trim(s, `"`)
}
