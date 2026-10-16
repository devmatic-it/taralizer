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
package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// cliProfilePath is the google.yaml profile relative to this package's test
// working directory (cmd/).
const cliProfilePath = "../profiles/google.yaml"

// cliModelFixture is a minimal model with one exportable asset and one
// unmapped boundary (which must produce a warning, not an error).
const cliModelFixture = `title: CLI Test
technical_assets:
  - id: vpc
    name: VPC
    technology: vpc

trust_boundaries:
  - id: cloud
    name: Cloud boundary
    technology: internet
`

// runRoot executes the root command with args and returns combined output.
func runRoot(t *testing.T, args ...string) (string, error) {
	t.Helper()
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	return buf.String(), err
}

// writeCLIModel writes the model fixture into a temp dir and returns its path.
func writeCLIModel(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "model.yaml")
	if err := os.WriteFile(path, []byte(cliModelFixture), 0644); err != nil {
		t.Fatalf("write model: %v", err)
	}
	return path
}

func TestTerraformExportCommand(t *testing.T) {
	modelPath := writeCLIModel(t)
	outDir := filepath.Join(t.TempDir(), "tf")

	out, err := runRoot(t, "terraform", "export", modelPath, "-o", outDir, "--profile", cliProfilePath)
	if err != nil {
		t.Fatalf("terraform export: %v\noutput:\n%s", err, out)
	}

	for _, base := range []string{"provider.tf", "variables.tf", "vpc.tf"} {
		if _, err := os.Stat(filepath.Join(outDir, base)); err != nil {
			t.Errorf("expected %s to be written: %v", base, err)
		}
	}

	if !strings.Contains(out, "Exported 1 resource(s)") {
		t.Errorf("expected success summary in output, got:\n%s", out)
	}
	if !strings.Contains(out, "WARNING") || !strings.Contains(out, `boundary "cloud"`) {
		t.Errorf("expected warning for unmapped boundary cloud, got:\n%s", out)
	}

	// Write-back must have recorded the ref in the model file.
	model, err := os.ReadFile(modelPath)
	if err != nil {
		t.Fatalf("read model: %v", err)
	}
	if !strings.Contains(string(model), "terraform_ref: google_compute_network.vpc") {
		t.Errorf("model missing terraform_ref write-back; got:\n%s", model)
	}
	if !strings.Contains(string(model), "terraform_ref: google_compute_network.vpc\n    name: VPC") {
		t.Errorf("terraform_ref not inserted below the id line; got:\n%s", model)
	}
}

func TestTerraformExportNoWriteBack(t *testing.T) {
	modelPath := writeCLIModel(t)
	outDir := filepath.Join(t.TempDir(), "tf")

	if _, err := runRoot(t, "terraform", "export", modelPath, "-o", outDir, "--profile", cliProfilePath, "--write-back=false"); err != nil {
		t.Fatalf("terraform export --write-back=false: %v", err)
	}

	model, err := os.ReadFile(modelPath)
	if err != nil {
		t.Fatalf("read model: %v", err)
	}
	if strings.Contains(string(model), "terraform_ref") {
		t.Errorf("model modified despite --no-write-back; got:\n%s", model)
	}
	if _, err := os.Stat(filepath.Join(outDir, "vpc.tf")); err != nil {
		t.Errorf("expected vpc.tf to be written: %v", err)
	}
}

func TestTerraformExportRefusesNonEmptyDir(t *testing.T) {
	modelPath := writeCLIModel(t)
	outDir := filepath.Join(t.TempDir(), "tf")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		t.Fatalf("mkdir outDir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "existing.tf"), []byte("# user\n"), 0644); err != nil {
		t.Fatalf("seed existing.tf: %v", err)
	}

	out, err := runRoot(t, "terraform", "export", modelPath, "-o", outDir, "--profile", cliProfilePath)
	if err == nil {
		t.Fatalf("expected error for non-empty outDir, output:\n%s", out)
	}
	if !strings.Contains(out, "not empty") {
		t.Errorf("expected 'not empty' error in output, got:\n%s", out)
	}

	if _, err := os.Stat(filepath.Join(outDir, "provider.tf")); !os.IsNotExist(err) {
		t.Error("expected no provider.tf after refused export")
	}

	// --force must succeed.
	if _, err := runRoot(t, "terraform", "export", modelPath, "-o", outDir, "--profile", cliProfilePath, "--force"); err != nil {
		t.Fatalf("terraform export --force: %v", err)
	}
	if _, err := os.Stat(filepath.Join(outDir, "provider.tf")); err != nil {
		t.Errorf("expected provider.tf after forced export: %v", err)
	}
}

func TestTerraformExportRequiresOutDir(t *testing.T) {
	modelPath := writeCLIModel(t)

	// Flag state persists across Execute calls on the shared command tree, so
	// reset out-dir to its empty default before asserting the required check.
	if err := terraformExportCmd.Flags().Set("out-dir", ""); err != nil {
		t.Fatalf("reset out-dir flag: %v", err)
	}

	out, err := runRoot(t, "terraform", "export", modelPath, "--profile", cliProfilePath)
	if err == nil {
		t.Fatalf("expected error when -o is missing, output:\n%s", out)
	}
	if !strings.Contains(out, "out-dir") {
		t.Errorf("expected required-flag error mentioning out-dir, got:\n%s", out)
	}
}

func TestTerraformExportBadModel(t *testing.T) {
	badPath := filepath.Join(t.TempDir(), "missing.yaml")

	out, err := runRoot(t, "terraform", "export", badPath, "-o", t.TempDir(), "--profile", cliProfilePath)
	if err == nil {
		t.Fatalf("expected error for missing model file, output:\n%s", out)
	}
	if !strings.Contains(out, "load model") {
		t.Errorf("expected load-model error in output, got:\n%s", out)
	}
}

func TestTerraformParentShowsHelp(t *testing.T) {
	out, err := runRoot(t, "terraform")
	if err != nil {
		t.Fatalf("bare terraform command should not error: %v", err)
	}
	for _, want := range []string{"Available Commands", "export", "import", "drift"} {
		if !strings.Contains(out, want) {
			t.Errorf("terraform help missing %q; got:\n%s", want, out)
		}
	}
}

// TestTerraformExportCommandWrongArgCount guards the Args validator.
func TestTerraformExportCommandWrongArgCount(t *testing.T) {
	out, err := runRoot(t, "terraform", "export")
	if err == nil {
		t.Fatalf("expected error with no args, output:\n%s", out)
	}
	if !strings.Contains(out, "accepts 1 arg") {
		t.Errorf("expected ExactArgs error, got:\n%s", out)
	}

	out, err = runRoot(t, "terraform", "export", "a.yaml", "b.yaml")
	if err == nil {
		t.Fatalf("expected error with two args, output:\n%s", out)
	}
	if !strings.Contains(out, "accepts 1 arg") {
		t.Errorf("expected ExactArgs error, got:\n%s", out)
	}
}

func TestTerraformImportCommand(t *testing.T) {
	// Export the model fixture, then import the exported directory back into
	// a new model: the round trip must preserve id, technology and ref.
	modelPath := writeCLIModel(t)
	tfDir := filepath.Join(t.TempDir(), "tf")

	if out, err := runRoot(t, "terraform", "export", modelPath, "-o", tfDir, "--profile", cliProfilePath); err != nil {
		t.Fatalf("terraform export: %v\noutput:\n%s", err, out)
	}

	outModel := filepath.Join(t.TempDir(), "imported.yaml")
	out, err := runRoot(t, "terraform", "import", tfDir, "-o", outModel, "--profile", cliProfilePath)
	if err != nil {
		t.Fatalf("terraform import: %v\noutput:\n%s", err, out)
	}

	// The fixture's only exportable object (vpc) maps to a trust boundary.
	if !strings.Contains(out, "Imported 0 asset(s), 1 boundary(ies)") {
		t.Errorf("expected import summary, got:\n%s", out)
	}

	imported, err := os.ReadFile(outModel)
	if err != nil {
		t.Fatalf("read imported model: %v", err)
	}
	for _, want := range []string{
		"trust_boundaries:",
		"id: vpc",
		"technology: vpc",
		"terraform_ref: google_compute_network.vpc",
	} {
		if !strings.Contains(string(imported), want) {
			t.Errorf("imported model missing %q; got:\n%s", want, imported)
		}
	}
}

func TestTerraformImportMissingDir(t *testing.T) {
	out, err := runRoot(t, "terraform", "import", filepath.Join(t.TempDir(), "nope"),
		"-o", filepath.Join(t.TempDir(), "model.yaml"), "--profile", cliProfilePath)
	if err == nil {
		t.Fatalf("expected error for missing dir, output:\n%s", out)
	}
	if !strings.Contains(out, "import:") {
		t.Errorf("expected import error in output, got:\n%s", out)
	}
}

// exportCLIModel exports the CLI model fixture into a fresh tf dir and
// returns both paths; it is the shared setup for the drift command tests.
func exportCLIModel(t *testing.T) (modelPath, tfDir string) {
	t.Helper()
	modelPath = writeCLIModel(t)
	tfDir = filepath.Join(t.TempDir(), "tf")

	// --write-back=true is passed explicitly: flag state persists across
	// Execute calls on the shared command tree, and an earlier test left it
	// false. Drift detection needs the refs written back into the model.
	if out, err := runRoot(t, "terraform", "export", modelPath, "-o", tfDir, "--profile", cliProfilePath, "--write-back=true"); err != nil {
		t.Fatalf("terraform export: %v\noutput:\n%s", err, out)
	}
	return modelPath, tfDir
}

func TestTerraformDriftCommandClean(t *testing.T) {
	modelPath, tfDir := exportCLIModel(t)

	out, err := runRoot(t, "terraform", "drift", modelPath, tfDir)
	if err != nil {
		t.Fatalf("drift on in-sync model+tf must exit 0: %v\noutput:\n%s", err, out)
	}
	if !strings.Contains(out, "No drift detected") {
		t.Errorf("expected 'No drift detected' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "No drift detected.") {
		t.Errorf("markdown report should contain the clean line, got:\n%s", out)
	}
}

func TestTerraformDriftCommandDetects(t *testing.T) {
	modelPath, tfDir := exportCLIModel(t)

	// Remove the exported vpc resource (its model ref becomes missing) and add
	// an unmodeled resource with no marker.
	if err := os.Remove(filepath.Join(tfDir, "vpc.tf")); err != nil {
		t.Fatalf("remove vpc.tf: %v", err)
	}
	extra := "resource \"google_compute_disk\" \"data\" {\n  name = \"d\"\n}\n"
	if err := os.WriteFile(filepath.Join(tfDir, "extra.tf"), []byte(extra), 0644); err != nil {
		t.Fatalf("write extra.tf: %v", err)
	}

	reportPath := filepath.Join(t.TempDir(), "report.md")
	out, err := runRoot(t, "terraform", "drift", modelPath, tfDir, "-o", reportPath)
	if err == nil {
		t.Fatalf("expected non-zero exit on drift, output:\n%s", out)
	}
	if !strings.Contains(err.Error(), "drift detected: 2 finding(s)") {
		t.Errorf("expected drift error with finding count, got: %v", err)
	}

	report, rerr := os.ReadFile(reportPath)
	if rerr != nil {
		t.Fatalf("read report: %v", rerr)
	}
	for _, want := range []string{
		"# Taralizer Drift Report",
		"- Findings: 2",
		"## Missing assets (1)",
		"google_compute_network.vpc",
		"## Unmodeled resources (1)",
		"google_compute_disk.data",
	} {
		if !strings.Contains(string(report), want) {
			t.Errorf("report missing %q; got:\n%s", want, report)
		}
	}

	// The gate must not print the usage block on an expected failure.
	if strings.Contains(out, "Usage:") {
		t.Errorf("drift failure should suppress usage output, got:\n%s", out)
	}
}

func TestTerraformDriftCommandWrongArgCount(t *testing.T) {
	out, err := runRoot(t, "terraform", "drift")
	if err == nil {
		t.Fatalf("expected error with no args, output:\n%s", out)
	}
	if !strings.Contains(out, "accepts 2 arg") {
		t.Errorf("expected ExactArgs error, got:\n%s", out)
	}

	out, err = runRoot(t, "terraform", "drift", "a.yaml", "b", "c")
	if err == nil {
		t.Fatalf("expected error with three args, output:\n%s", out)
	}
	if !strings.Contains(out, "accepts 2 arg") {
		t.Errorf("expected ExactArgs error, got:\n%s", out)
	}
}

func TestTerraformImportRequiresOut(t *testing.T) {
	// Flag state persists across Execute calls on the shared command tree, so
	// reset out to its empty default before asserting the required check.
	if err := terraformImportCmd.Flags().Set("out", ""); err != nil {
		t.Fatalf("reset out flag: %v", err)
	}

	out, err := runRoot(t, "terraform", "import", t.TempDir(), "--profile", cliProfilePath)
	if err == nil {
		t.Fatalf("expected error when -o is missing, output:\n%s", out)
	}
	if !strings.Contains(out, "out") {
		t.Errorf("expected required-flag error mentioning out, got:\n%s", out)
	}
}

// TestTerraformCommandTreeShape asserts the parent/subcommand wiring.
func TestTerraformCommandTreeShape(t *testing.T) {
	if terraformCmd.HasParent() && terraformCmd.Parent() != rootCmd {
		t.Errorf("terraform parent should be root, got %v", terraformCmd.Parent())
	}
	var names []string
	for _, sub := range terraformCmd.Commands() {
		names = append(names, sub.Name())
	}
	if len(names) != 3 || names[0] != "drift" || names[1] != "export" || names[2] != "import" {
		t.Errorf("terraform subcommands = %v, want [drift export import] (cobra sorts alphabetically)", names)
	}

	var found *cobra.Command
	for _, sub := range terraformCmd.Commands() {
		if sub.Name() == "export" {
			found = sub
		}
	}
	if found == nil {
		t.Fatal("export subcommand not found")
	}
	for _, flag := range []string{"out-dir", "profile", "force", "write-back"} {
		if found.Flags().Lookup(flag) == nil {
			t.Errorf("export subcommand missing flag %q", flag)
		}
	}

	var drift *cobra.Command
	for _, sub := range terraformCmd.Commands() {
		if sub.Name() == "drift" {
			drift = sub
		}
	}
	if drift == nil {
		t.Fatal("drift subcommand not found")
	}
	if drift.Flags().Lookup("out") == nil {
		t.Error("drift subcommand missing flag \"out\"")
	}
}
