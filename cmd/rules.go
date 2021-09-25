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

// rules command
var rulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "shows rules",
	Long:  `provides a list of installed rules.`,
	Run: func(cmd *cobra.Command, args []string) {
		t := taralizer.NewTaralizer(ruleSet)
		rs := t.RuleSet(ruleSet)
		fmt.Printf("%s (%s) %s\n", rs.Title, rs.Name, rs.Version)
		for _, v := range rs.Rules {
			fmt.Printf("  %s - %s: %s\n", v.Id, v.Title, v.Description)
		}
	},
}

// initializes arguments for rules command
func init() {
	rootCmd.AddCommand(rulesCmd)

}
