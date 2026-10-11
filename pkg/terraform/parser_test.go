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
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTerraform(t *testing.T) {
	files, err := ParseTerraform("../../examples/gcp/terraform")
	assert.NoError(t, err)
	assert.NotEmpty(t, files)
}

func TestParseTerraformDir(t *testing.T) {
	resources, err := ParseTerraformDir("../../examples/gcp/terraform")
	assert.NoError(t, err)
	// examples/gcp/terraform/main.tf declares exactly two resources
	assert.Len(t, resources, 2)

	assert.Equal(t, "google_compute_network", resources[0].Type)
	assert.Equal(t, "vpc_network", resources[0].Name)

	assert.Equal(t, "google_compute_instance", resources[1].Type)
	assert.Equal(t, "vm_instance", resources[1].Name)
}

func TestParseTerraformDirWithValidFiles(t *testing.T) {
	// Create a temporary directory with test Terraform files
	tempDir, err := os.MkdirTemp("", "terraform-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a simple Terraform file with resources
	testFile := filepath.Join(tempDir, "test.tf")
	terraformContent := `
resource "google_compute_network" "vpc_network" {
  name = "my-vpc-network"
}

resource "google_compute_instance" "vm_instance" {
  name = "my-vm-instance"
}
`
	err = os.WriteFile(testFile, []byte(terraformContent), 0644)
	assert.NoError(t, err)

	// Test parsing the directory
	resources, err := ParseTerraformDir(tempDir)
	assert.NoError(t, err)

	// Verify we found the resources
	assert.Len(t, resources, 2)

	// Check first resource
	assert.Equal(t, "google_compute_network", resources[0].Type)
	assert.Equal(t, "vpc_network", resources[0].Name)

	// Check second resource
	assert.Equal(t, "google_compute_instance", resources[1].Type)
	assert.Equal(t, "vm_instance", resources[1].Name)
}

func TestParseTerraformDirWithEmptyDirectory(t *testing.T) {
	// Create a temporary directory with no Terraform files
	tempDir, err := os.MkdirTemp("", "terraform-test-empty")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Test parsing the empty directory
	resources, err := ParseTerraformDir(tempDir)
	assert.NoError(t, err)

	// Should return empty slice
	assert.Empty(t, resources)
}

func TestParseTerraformDirWithNonExistentDirectory(t *testing.T) {
	// Test parsing a non-existent directory
	resources, err := ParseTerraformDir("/non/existent/directory")
	assert.Error(t, err)
	assert.Empty(t, resources)
}
