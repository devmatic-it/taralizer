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
	"log"
	"os"
	"os/exec"
)

// initializes arguments for version commmand
func init() {
	diagramCmd.Flags().StringVar(&engine, "engine", "dot", "default command to generate graph. Currently 'dot' and 'plantuml' are supported.")
	diagramCmd.Flags().StringVar(&outFile, "out", "diagram", "output file name")
	diagramCmd.Flags().StringVar(&imageType, "type", "png", "type of output image")
	rootCmd.AddCommand(diagramCmd)

}

// version command
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

			report := taralizer.Load(args[0])
			r := taralizer.NewReportEngine()

			var engineCmd *exec.Cmd
			if engine == "plantuml" {
				// #nosec G204 we intentionally call plantuml to generate graphs
				engineCmd = exec.Command(engine, "-pipe", "-t"+imageType)
			} else if engine == "dot" {
				// #nosec G204 we intentionally call dot to generate graphs
				engineCmd = exec.Command(engine, "-T"+imageType)
			} else {
				log.Fatal("Unsupported engine: ", engine)
			}

			stdin, err := engineCmd.StdinPipe()
			if err != nil {
				log.Fatal(err)
			}

			go func() {
				defer stdin.Close()
				r.GenerateReport(stdin, fmt.Sprintf("%s%s.tpl", r.GetTemplateDir(), engine), report)
			}()

			fout, err := os.Create(fmt.Sprintf("%s.%s", outFile, imageType))
			if err != nil {
				log.Fatal(err)
			}
			engineCmd.Stdout = fout
			err = engineCmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)
