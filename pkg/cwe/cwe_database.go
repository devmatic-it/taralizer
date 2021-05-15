// Package cwe Mitre Common Weekness Evaluation
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
package cwe

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// CWEDatabase represents an in-memory representatio of the ASVS CVE guide downloaded
// from https://cwe.mitre.org/data/downloads.html
type CWEDatabase struct {
	weakness map[string]CWEWeakness
}

type CWEWeakness struct {
	Id                  string
	Name                string
	Description         string
	ExtendedDescription string
	Likelihood          string
	Url                 string
}

// NewASVSParser creates a new instance of the parset
func NewCWEDatabase() *CWEDatabase {
	instance := CWEDatabase{}
	instance.load("../../data/1000.csv")
	return &instance
}

// Finds an ASVS requrirement by its id eg 'V2.1.5' for very that users can change password
func (p *CWEDatabase) Find(id string) CWEWeakness {
	return p.weakness[id]
}

// load loads the csv file from disk
func (p *CWEDatabase) load(filename string) {

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

	p.weakness = make(map[string]CWEWeakness)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		line := strings.Split(str, ",")
		req := CWEWeakness{
			Id:                  line[0],
			Name:                line[1],
			Description:         line[4],
			ExtendedDescription: line[5],
			Likelihood:          line[13],
			Url:                 "https://cwe.mitre.org/data/definitions/" + line[0],
		}
		p.weakness[req.Id] = req
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
