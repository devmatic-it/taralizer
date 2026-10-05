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
	"log"
	"os"
	"os/exec"

	"github.com/devmatic-it/taralizer/pkg/taralizer"
	"github.com/spf13/cobra"
)

var (
	engine    string
	outFile   string
	imageType string

	diagramCmd = &cobra.Command{
		Use:   "diagram <model>",
		Short: "generates a PlantML diagram ",
		Long:  `Generates a PlantML Data Flow diagram.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			_, err := exec.LookPath(engine)
			if err != nil {
				log.Fatal(err)
			}

			report, err := taralizer.Load(args[0])
			if err != nil {
				log.Fatalf("Failed to load model: %v", err)
			}

			r := taralizer.NewReportEngine()
			tplDir, err := r.GetTemplateDir()
			if err != nil {
				log.Fatal(err)
			}

			var engineCmd *exec.Cmd
			switch engine {
			case "plantuml":
				engineCmd = exec.Command(engine, "-pipe", "-t"+imageType)
			case "dot":
				engineCmd = exec.Command(engine, "-T"+imageType)
			case "mermaid":
				engineCmd = exec.Command("mmdc", "-i", "-", "-o", fmt.Sprintf("%s.%s", outFile, imageType))
			default:
				log.Fatal("Unsupported engine: ", engine)
			}

			stdin, err := engineCmd.StdinPipe()
			if err != nil {
				log.Fatal(err)
			}

			go func() {
				defer stdin.Close()
				r.GenerateReport(stdin, fmt.Sprintf("%s%s.tpl", tplDir, engine), report)
			}()

			fout, err := os.Create(fmt.Sprintf("%s.%s", outFile, imageType))
			if err != nil {
				log.Fatal(err)
			}
			engineCmd.Stdout = fout
			if err := engineCmd.Run(); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	diagramCmd.Flags().StringVar(&engine, "engine", "dot", "default command to generate graph. Currently 'dot' and 'plantuml' are supported.")
	diagramCmd.Flags().StringVar(&outFile, "out", "diagram", "output file name")
	diagramCmd.Flags().StringVar(&imageType, "type", "png", "type of output image")
	rootCmd.AddCommand(diagramCmd)
}
