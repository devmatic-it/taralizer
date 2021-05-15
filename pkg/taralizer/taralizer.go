// Package taralizer Threat and Risk Analyzer
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
package taralizer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/open-policy-agent/opa/rego"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

// Taralzer struct
type Taralizer struct {
	ctx     context.Context
	ruleset string
}

// New creates a new instance of the Taralizer engine.
func NewTaralizer(ruleset string) *Taralizer {
	instance := Taralizer{}
	instance.ctx = context.TODO()
	instance.ruleset = ruleset
	return &instance
}

// convertMessage converts an untyped instance Risk into a typed one.
func (svc *Taralizer) convertMessage(input interface{}) Risk {
	if input == nil {
		log.Fatal("interface cannot be nil")
	}

	data := input.(map[string]interface{})
	msg := Risk{
		Id:          data["id"].(string),
		Title:       data["title"].(string),
		Description: data["description"].(string),
		Mitigation:  data["mitigation"].(string),
		Url:         data["url"].(string),
		Message:     data["msg"].(string),
	}

	msg.Cwe, _ = data["cwe"].(json.Number).Int64()
	msg.Likelihood, _ = data["likelihood"].(json.Number).Int64()
	msg.Impact, _ = data["impact"].(json.Number).Int64()
	return msg
}

// Evaluate executes an Open Policy Agent (OPA) query against the rule sets
// and stores the resulting risks into the returned report.
func (svc *Taralizer) Evaluate(fileName string) Report {

	results := svc.query(fileName, fmt.Sprintf("data.rules.%s.violation[msg]", svc.ruleset))

	// load model into structured report
	report := Load(fileName)
	for i := 0; i < len(results); i++ {
		item := svc.convertMessage(results[i].Bindings["msg"])
		item.Severity = item.Likelihood * item.Impact
		report.Risks = append(report.Risks, item)
	}

	return report
}

// Validate executes an Open Policy Agent (OPA) query against the rule sets to perform a model validation/checking for inconsistencies.
func (svc *Taralizer) Validate(fileName string) []string {

	results := svc.query(fileName, "data.rules.validation[msg]")

	// load model into structured report
	messages := []string{}
	for i := 0; i < len(results); i++ {
		data := results[i].Bindings["msg"].(map[string]interface{})
		item := fmt.Sprintf("%s: %s", data["id"].(string), data["msg"].(string))
		messages = append(messages, item)
	}

	return messages
}

// // Evaluate executes an Open Policy Agent (OPA) query against the rule sets calling the given query 'queryStr'
func (svc *Taralizer) query(fileName string, queryStr string) rego.ResultSet {
	/* #nosec G304 */
	jsonFile, err := os.Open(fileName)

	/* #nosec G307 */
	defer jsonFile.Close()
	if err != nil {
		log.Fatalf("cannot load model file: %v", err)
	}

	data, err := ioutil.ReadAll(jsonFile)
	var model interface{}
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}

	// unmmarshal as maps for Rego engine
	err = yaml.Unmarshal([]byte(data), &model)
	if err != nil {
		log.Fatalf("cannot unmarshal model file: %v", err)
	}

	rules := []string{"./rules/"}
	query, err := rego.New(rego.Query(queryStr), rego.Load(rules, nil)).PrepareForEval(svc.ctx)
	if err != nil {
		log.Fatalf("cannot load model file into rego engine: %v", err)
	}

	results, err := query.Eval(svc.ctx, rego.EvalInput(model))

	if err != nil {
		log.Fatalf("cannot evaluate rules against input: %v", err)
	}
	return results
}
