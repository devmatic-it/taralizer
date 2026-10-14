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

	"github.com/spf13/cobra"

	"github.com/devmatic-it/taralizer/pkg/taralizer"
	"github.com/devmatic-it/taralizer/pkg/terraform"
)

// config variables
var (
	// terraformCmd is the parent command grouping all Terraform integration
	// subcommands (export, import). It has no direct action: running it prints
	// the help for its subcommands.
	terraformCmd = &cobra.Command{
		Use:   "terraform",
		Short: "Terraform integration commands",
		Long: `Terraform integration for taralizer models.

Subcommands:
  export   Render a model's technical assets and trust boundaries into
           Terraform scaffolds and record each HCL address in the model's
           terraform_ref field.
  import   Build a taralizer report from Terraform files (legacy single-file
           mode; directory-based import is planned).`,
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

	// terraformImportCmd keeps the legacy single-file import behavior. Phase A2
	// replaces its body with directory-based ImportFromDir.
	terraformImportCmd = &cobra.Command{
		Use:   "import <terraform_file>",
		Short: "Import the model from a terraform file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tf := terraform.NewTerraform()
			tf.ImportFromFile(args[0])
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

	terraformCmd.AddCommand(terraformExportCmd, terraformImportCmd)
	rootCmd.AddCommand(terraformCmd)
}
