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
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

// initializes arguments for version commmand
func init() {
	cobra.OnInitialize(initConfig)
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

// Execute runs the root cobra command for the help
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
