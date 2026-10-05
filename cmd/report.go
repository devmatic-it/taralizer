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

	"github.com/devmatic-it/taralizer/pkg/taralizer"
	"github.com/spf13/cobra"
)

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
			tplDir, err := r.GetTemplateDir()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			if reportType == "pdf" {
				err := r.GenerateReportFilePDF(reportFile+".pdf",
					tplDir+"pdf_report.tpl",
					tplDir+"pdf_report_cover.tpl", report)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error generating PDF report: %v\n", err)
					os.Exit(1)
				}
			} else if reportType == "markdown" {
				err := r.GenerateReportFileMarkdown(reportFile+".md",
					tplDir+"markdown_report.tpl", report)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error generating markdown report: %v\n", err)
					os.Exit(1)
				}
			} else {
				err := r.GenerateReportFile(reportFile+".html",
					tplDir+"html.tpl", report)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error generating HTML report: %v\n", err)
					os.Exit(1)
				}
			}
		},
	}
)

func init() {
	reportCmd.Flags().StringVar(&reportFile, "out", "report", "output file name")
	reportCmd.Flags().StringVar(&reportType, "type", "html", "type of report")
	rootCmd.AddCommand(reportCmd)
}
