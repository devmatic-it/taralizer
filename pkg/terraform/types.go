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

// GCPResource represents a Google Cloud Platform resource
type GCPResource struct {
	Name string
	Type string
}

// ParseTerraformDir parses Terraform files in a directory and returns the parsed content
func ParseTerraformDir(tfDir string) ([]GCPResource, error) {
	// This function should parse Terraform files and return a list of GCP resources
	// For now, we'll implement it as a stub that returns an empty slice
	return []GCPResource{}, nil
}
