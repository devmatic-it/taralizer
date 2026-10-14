// Package terraform provides Terraform parsing capabilities for taralizer
//
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

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// ParseTerraform parses Terraform files in a directory and returns the parsed content
func ParseTerraform(tfDir string) ([]*hcl.File, error) {
	var files []*hcl.File
	err := filepath.Walk(tfDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil // continue walking
		}

		// Only process .tf files
		if strings.HasSuffix(path, ".tf") {
			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", path, err)
			}

			file, diags := hclsyntax.ParseConfig(content, path, hcl.Pos{Line: 1, Column: 1})
			if diags.HasErrors() {
				return fmt.Errorf("failed to parse HCL in file %s: %w", path, diags)
			}

			files = append(files, file)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory %s: %w", tfDir, err)
	}

	return files, nil
}
