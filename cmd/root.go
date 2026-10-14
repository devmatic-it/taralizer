// Package cmd cobra command line interface
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
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// Add import for our terraform package
	"github.com/devmatic-it/taralizer/pkg/terraform"
)

// config variables
var (
	cfgFile = ""
	ruleSet = "asvs"
	rootCmd = &cobra.Command{
		Use:   "taralizer <model>",
		Short: "Threat and Risk Analyser",
		Long: `TARAlizer - Threat and Risk Analyzer


Taralizer is a Threat and Risk Analysis tool.
		`,
	}
)

// terraformToYamlCmd represents the terraform-to-yaml command
var terraformToYamlCmd = &cobra.Command{
	Use:   "terraform-to-yaml",
	Short: "Generate taralizer YAML from Google Cloud Terraform files",
	Long: `Generate a taralizer YAML template by parsing Google Cloud Terraform files
and mapping the resources to appropriate taralizer assets using the google.yaml profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		tfDir, _ := cmd.Flags().GetString("tf-dir")
		outputFile, _ := cmd.Flags().GetString("output")

		if tfDir == "" || outputFile == "" {
			fmt.Println("Usage: taralizer terraform-to-yaml --tf-dir <path> --output <file>")
			return
		}

		// Generate the YAML template from Terraform files
		yamlContent, err := terraform.GenerateFromTerraform(tfDir, "./profiles/google.yaml")
		if err != nil {
			fmt.Printf("Error generating template: %v\n", err)
			return
		}

		// Write to output file
		err = os.WriteFile(outputFile, []byte(yamlContent), 0644)
		if err != nil {
			fmt.Printf("Error writing output file: %v\n", err)
			return
		}

		fmt.Printf("Successfully generated taralizer template at %s\n", outputFile)
	},
}

// initializes arguments for version command
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(terraformToYamlCmd)

	// Add flags for terraform-to-yaml command
	terraformToYamlCmd.Flags().String("tf-dir", "", "Directory containing Terraform files")
	terraformToYamlCmd.Flags().String("output", "", "Output YAML file path")

	rootCmd.Flags().StringVar(&cfgFile, "config", "", "configuration file. If not specified, 'config.yaml' is searched in path.")
	rootCmd.Flags().StringVar(&ruleSet, "ruleset", "asvs", "compliance standard to use. Currently, only the OWASP ASVS 'asvs' has been parially implemented.")

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("config")           // name of config file (without extension)
		viper.SetConfigType("yaml")             // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath("/etc/taralizer/")  // path to look for the config file in
		viper.AddConfigPath("$HOME/.taralizer") // call multiple times to add many search paths
		viper.AddConfigPath(".")                // optionally look for config in the working directory
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
