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
	"github.com/devmatic-it/taralizer/pkg/terraform"
	"github.com/spf13/cobra"
)

// config variables
var (
	terraformCmd = &cobra.Command{
		Use:   "terraform import <model> <terraform_file> ",
		Short: "import the model from terraform file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tf := terraform.NewTerraform()
			tf.ImportFromFile(args[0])
		},
	}
)

// initializes arguments for version command
func init() {
	rootCmd.AddCommand(terraformCmd)
}
