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
	"runtime"

	"github.com/spf13/cobra"
)

// ProductVersion contains the product version injected by the build system
var ProductVersion string = "0.0.0"

// version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "product version",
	Long:  `provides version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("taralizer version: %s\n", ProductVersion)
		fmt.Printf("go version: %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	},
}

// initializes arguments for version commmand
func init() {
	rootCmd.AddCommand(versionCmd)

}
