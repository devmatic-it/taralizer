// Package cmd cobra command line interface
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
package cmd

import (
	"fmt"
	"github.com/devmatic-it/taralizer/pkg/taralizer"
	"github.com/spf13/cobra"
)

// config variables
var (
	validateCmd = &cobra.Command{
		Use:   "validate <model>",
		Short: "validates the model and returns potential model errors",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			t := taralizer.NewTaralizer(ruleSet)

			messages := t.Validate(args[0])
			if len(messages) == 0 {
				fmt.Println("Model validated successfully. No Validation errors found.")
			} else {
				for _, msg := range messages {
					fmt.Printf("WARNING %s\n", msg)
				}
			}
		},
	}
)

// initializes arguments for version command
func init() {
	rootCmd.AddCommand(validateCmd)
}
