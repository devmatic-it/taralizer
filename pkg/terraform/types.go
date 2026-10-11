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
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// GCPResource represents a Google Cloud Platform resource
type GCPResource struct {
	Name string
	Type string
}

// ParseTerraformDir parses Terraform files in a directory and returns the resources declared in them.
// Every resource block contributes one GCPResource with Type = labels[0] and Name = labels[1].
func ParseTerraformDir(tfDir string) ([]GCPResource, error) {
	files, err := ParseTerraform(tfDir)
	if err != nil {
		return nil, err
	}

	resources := []GCPResource{}
	for _, file := range files {
		body, ok := file.Body.(*hclsyntax.Body)
		if !ok {
			continue
		}
		for _, block := range body.Blocks {
			if block.Type != "resource" || len(block.Labels) < 2 {
				continue
			}
			resources = append(resources, GCPResource{
				Name: block.Labels[1],
				Type: block.Labels[0],
			})
		}
	}

	return resources, nil
}
