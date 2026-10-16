// Package cmd cobra command line interface
// Copyright 2022 taralizer authors
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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/devmatic-it/taralizer/pkg/taralizer"
	"github.com/devmatic-it/taralizer/pkg/terraform"
)

// config variables
var (
	// terraformCmd is the parent command grouping all Terraform integration
	// subcommands (export, import, drift). It has no direct action: running it
	// prints the help for its subcommands.
	terraformCmd = &cobra.Command{
		Use:   "terraform",
		Short: "Terraform integration commands",
		Long: `Terraform integration for taralizer models.

Subcommands:
  export   Render a model's technical assets and trust boundaries into
           Terraform scaffolds and record each HCL address in the model's
           terraform_ref field.
  import   Build a taralizer report from the resource blocks in every .tf
           file under a directory.
  drift    Detect structural drift between a model and a Terraform
           directory; exits non-zero when any drift is found (CI gate).`,
	}

	// terraformExportCmd renders a model into Terraform scaffolds (Phase A1).
	terraformExportCmd = &cobra.Command{
		Use:   "export <model.yaml>",
		Short: "Export a taralizer model to Terraform scaffolds",
		Long: `Render every technical asset and trust boundary whose technology maps to a
Terraform resource type in the profile into outDir as scaffold files.

Each exported object gets its own <name>.tf file containing a marker comment
(# taralizer:asset=<id> or # taralizer:boundary=<id>) directly above the
resource block, plus shared provider.tf and variables.tf files. Objects whose
technology is missing or unmapped produce a warning on stderr and are skipped;
nothing is written if the target directory is not empty (without --force) or
if two objects would collide on one HCL resource name.

Unless --write-back=false is given, the HCL address of each exported object
(e.g. google_compute_network.vpc) is written back into the model file as its
terraform_ref field, which later powers drift detection.

Example:
  taralizer terraform export model.yaml -o ./tf --profile ./profiles/google.yaml`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath := args[0]
			profilePath, _ := cmd.Flags().GetString("profile")
			outDir, _ := cmd.Flags().GetString("out-dir")
			force, _ := cmd.Flags().GetBool("force")
			writeBack, _ := cmd.Flags().GetBool("write-back")

			if outDir == "" {
				return fmt.Errorf("out-dir is required (use -o <dir>)")
			}

			report, err := taralizer.Load(modelPath)
			if err != nil {
				return fmt.Errorf("load model: %w", err)
			}

			refs, warnings, err := terraform.ExportToTerraform(report, profilePath, outDir, force)
			for _, w := range warnings {
				fmt.Fprintf(cmd.ErrOrStderr(), "WARNING: %s\n", w)
			}
			if err != nil {
				return fmt.Errorf("export: %w", err)
			}

			cmd.Printf("Exported %d resource(s) to %s\n", len(refs), outDir)

			if !writeBack {
				return nil
			}
			written, err := terraform.WriteBackTerraformRefs(modelPath, refs)
			if err != nil {
				return fmt.Errorf("write-back: %w", err)
			}
			cmd.Printf("Wrote terraform_ref for %d object(s) into %s\n", written, modelPath)
			return nil
		},
	}

	// terraformImportCmd builds a taralizer report from a directory of
	// Terraform files (Phase A2).
	terraformImportCmd = &cobra.Command{
		Use:   "import <terraform_dir>",
		Short: "Import a taralizer model from Terraform files",
		Long: `Build a taralizer report from the resource blocks in every .tf file under
<terraform_dir>. Each resource whose type maps to a technology in the profile
becomes a technical asset or trust boundary, with its HCL address recorded as
terraform_ref. Unmapped resource types produce a warning on stderr and are
skipped; files with syntax errors abort the import.

Example:
  taralizer terraform import ./tf -o model.yaml --profile ./profiles/google.yaml`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			tfDir := args[0]
			profilePath, _ := cmd.Flags().GetString("profile")
			outFile, _ := cmd.Flags().GetString("out")

			if outFile == "" {
				return fmt.Errorf("out is required (use -o <model.yaml>)")
			}

			report, warnings, err := terraform.ImportFromDir(tfDir, profilePath)
			for _, w := range warnings {
				fmt.Fprintf(cmd.ErrOrStderr(), "WARNING: %s\n", w)
			}
			if err != nil {
				return fmt.Errorf("import: %w", err)
			}

			if err := taralizer.Save(report, outFile); err != nil {
				return fmt.Errorf("save model: %w", err)
			}

			cmd.Printf("Imported %d asset(s), %d boundary(ies) from %s into %s\n",
				len(report.TechnicalAssets), len(report.TrustBoundaries), tfDir, outFile)
			return nil
		},
	}

	// terraformDriftCmd detects structural drift between a model and a
	// Terraform directory (Phase A3). It is a CI gate: any finding makes the
	// command exit non-zero.
	terraformDriftCmd = &cobra.Command{
		Use:   "drift <model.yaml> <terraform_dir>",
		Short: "Detect structural drift between a model and Terraform files",
		Long: `Compare the taralizer model with the resource blocks in every .tf file under
<terraform_dir> and report structural drift:

  missing assets    model entries whose terraform_ref has no matching resource
                    in the Terraform directory
  broken refs       model entries whose terraform_ref is not a valid HCL
                    resource address (type.name)
  unmodeled         resource blocks in the Terraform directory that no model
                    entry references via terraform_ref
  stale markers     taralizer export markers (# taralizer:asset= / :boundary=)
                    whose id no longer exists in the model

The markdown report goes to stdout, or to -o <file> when given. The command
exits non-zero whenever any drift is found, so it can gate CI pipelines;
exit code 0 means model and Terraform are in sync.

Example:
  taralizer terraform drift model.yaml ./tf -o report.md`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			modelPath, tfDir := args[0], args[1]
			outFile, _ := cmd.Flags().GetString("out")

			report, err := terraform.DetectDrift(modelPath, tfDir)
			if err != nil {
				return fmt.Errorf("drift: %w", err)
			}

			for _, w := range report.Warnings {
				fmt.Fprintf(cmd.ErrOrStderr(), "WARNING: %s\n", w)
			}

			md := report.Markdown()
			if outFile != "" {
				if err := os.WriteFile(outFile, []byte(md), 0644); err != nil {
					return fmt.Errorf("write report: %w", err)
				}
				cmd.Printf("Wrote drift report to %s\n", outFile)
			} else {
				fmt.Fprint(cmd.OutOrStdout(), md)
			}

			if n := report.Count(); n > 0 {
				// Drift is an expected outcome for a gate, not a usage error:
				// suppress the usage block cobra would otherwise print on failure.
				cmd.SilenceUsage = true
				return fmt.Errorf("drift detected: %d finding(s) between %s and %s", n, modelPath, tfDir)
			}

			cmd.Printf("No drift detected between %s and %s\n", modelPath, tfDir)
			return nil
		},
	}
)

// initializes arguments for terraform commands
func init() {
	terraformExportCmd.Flags().String("profile", "./profiles/google.yaml", "path to the mapping profile (google.yaml)")
	terraformExportCmd.Flags().StringP("out-dir", "o", "", "directory to write the Terraform scaffolds into (required)")
	terraformExportCmd.Flags().Bool("force", false, "allow exporting into a non-empty directory")
	terraformExportCmd.Flags().Bool("write-back", true, "write terraform_ref back into the model file (disable with --write-back=false)")
	_ = terraformExportCmd.MarkFlagRequired("out-dir")

	terraformImportCmd.Flags().String("profile", "./profiles/google.yaml", "path to the mapping profile (google.yaml)")
	terraformImportCmd.Flags().StringP("out", "o", "", "model file to write the imported report to (required)")
	_ = terraformImportCmd.MarkFlagRequired("out")

	terraformDriftCmd.Flags().StringP("out", "o", "", "write the markdown drift report to this file instead of stdout")

	terraformCmd.AddCommand(terraformExportCmd, terraformImportCmd, terraformDriftCmd)
	rootCmd.AddCommand(terraformCmd)
}
