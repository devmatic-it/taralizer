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
	reportType string
	reportFile string
	reportCmd  = &cobra.Command{
		Use:   "report <model>",
		Short: "creates a report",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			t := taralizer.NewTaralizer(ruleSet)

			messages := t.Validate(args[0])
			for _, msg := range messages {
				fmt.Printf("WARNING %s\n", msg)
			}

			report := t.Evaluate(args[0])

			r := taralizer.NewReportEngine()
			if reportType == "pdf" {
				r.GenerateReportFilePDF(reportFile+".pdf",
					r.GetTemplateDir()+"pdf_report.tpl",
					r.GetTemplateDir()+"pdf_report_cover.tpl", report)
			} else {
				r.GenerateReportFile(reportFile+".html",
					r.GetTemplateDir()+"html.tpl", report)
			}
		},
	}
)

// initializes arguments for version command
func init() {
	reportCmd.Flags().StringVar(&reportFile, "out", "report", "output file name")
	reportCmd.Flags().StringVar(&reportType, "type", "html", "type of report")
	rootCmd.AddCommand(reportCmd)
}
