// Package asvs OWASP Application Security Verification Standard database
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
package asvs

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

// ASVSDatabase represents an in-memory representatio of the ASVS CVE guide downloaded
// from https://owasp.org/www-project-application-security-verification-standard/
type ASVSDatabase struct {
	requirements map[string]ASVSRequirement
}

type ASVSRequirement struct {
	Id          string
	Description string
	Cwe         string
	Nist        string
}

// NewASVSParser creates a new instance of the parset
func NewASVSDatabase() *ASVSDatabase {
	instance := ASVSDatabase{}
	instance.load("../../data/OWASP Application Security Verification Standard 4.0.2-en.csv.txt")
	return &instance
}

// Finds an ASVS requrirement by its id eg 'V2.1.5' for very that users can change password
func (p *ASVSDatabase) Find(id string) ASVSRequirement {
	return p.requirements[strings.ToUpper(id)]
}

// load loads the csv file from disk
func (p *ASVSDatabase) load(filename string) {
	// Open CSV file
	// #nosec G304 file is only parsed and not included directly
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Cannot open file %s: %s", filename, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalf("Cannot read file %s: %s", filename, err)
	}
	p.requirements = make(map[string]ASVSRequirement)
	for _, line := range lines {
		req := ASVSRequirement{
			Id:          line[4],
			Description: line[5],
			Cwe:         line[9],
			Nist:        line[10],
		}
		p.requirements[req.Id] = req
	}
}
