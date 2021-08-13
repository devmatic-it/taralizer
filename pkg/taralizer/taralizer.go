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
	"fmt"
	"github.com/open-policy-agent/opa/rego"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Taralzer struct
type Taralizer struct {
	ctx     context.Context
	ruleset RuleSet
}

// New creates a new instance of the Taralizer engine.
func NewTaralizer(ruleset string) *Taralizer {
	instance := Taralizer{}
	instance.ctx = context.TODO()
	instance.ruleset = instance.RuleSet()
	return &instance
}

// convertMapToRisk converts an untyped instance Risk into a typed one.
func (svc *Taralizer) convertMapToRisk(input interface{}) Risk {
	if input == nil {
		log.Fatal("interface cannot be nil")
	}

	data := input.(map[string]interface{})
	msg := Risk{
		Id: data["id"].(string),
		//		Title:       data["title"].(string),
		//		Description: data["description"].(string),
		//		Mitigation:  data["mitigation"].(string),
		//		Url:         data["url"].(string),
		Message: data["msg"].(string),
	}

	msg.Likelihood = GetMapIntValue(data, "likelihood")
	msg.Impact = GetMapIntValue(data, "impact")

	//	msg.Cwe, _ = data["cwe"].(json.Number).Int64()
	// TODO: fetch metadata from ruleset

	rule := svc.findRule(msg.Id)
	if rule != nil {
		msg.Title = rule.Title
		msg.Description = rule.Description
		msg.Mitigation = rule.Mitigation
		msg.Url = rule.Url
		msg.Cwe = rule.Cwe
	}
	return msg
}

// findRule searches metadata
func (svc *Taralizer) findRule(id string) *Rule {
	for _, rule := range svc.ruleset.Rules {
		if strings.HasPrefix(id, rule.Id) {
			return &rule
		}
	}
	return nil
}

// convertMapToRisk converts an untyped instance Rule into a typed one.
func (svc *Taralizer) convertMapToRule(input interface{}) Rule {
	if input == nil {
		log.Fatal("interface cannot be nil")
	}

	data := input.(map[string]interface{})
	msg := Rule{
		Id:          data["id"].(string),
		Title:       data["title"].(string),
		Description: data["description"].(string),
		Mitigation:  data["mitigation"].(string),
		Url:         data["url"].(string),
	}

	msg.Cwe = GetMapIntValue(data, "cwe")
	msg.Likelihood = GetMapIntValue(data, "likelihood")
	msg.Impact = GetMapIntValue(data, "impact")
	return msg
}

// convertMapToRuleSet converts an untyped instance RuleSet
func (svc *Taralizer) convertMapToRuleSet(input interface{}) RuleSet {
	if input == nil {
		log.Fatal("interface cannot be nil")
	}

	data := input.(map[string]interface{})
	ruleSet := RuleSet{
		Name:        data["name"].(string),
		Title:       data["title"].(string),
		Description: data["description"].(string),
		Version:     data["version"].(string),
		Url:         data["url"].(string),
		Rules:       []Rule{},
	}

	rules := data["rules"].([]interface{})
	for _, v := range rules {
		ruleSet.Rules = append(ruleSet.Rules, svc.convertMapToRule(v))
	}

	return ruleSet
}

// Evaluate executes an Open Policy Agent (OPA) query against the rule sets
// and stores the resulting risks into the returned report.
func (svc *Taralizer) Evaluate(fileName string) Report {

	results := svc.query(fileName, fmt.Sprintf("data.rules.%s.violation[msg]", svc.ruleset.Name))

	// load model into structured report
	report := Load(fileName)
	report.RuleSet = svc.ruleset
	for i := 0; i < len(results); i++ {
		msg := results[i].Bindings["msg"]
		if msg != nil {
			item := svc.convertMapToRisk(msg)
			item.Severity = item.Likelihood * item.Impact
			report.Risks = append(report.Risks, item)
		}
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

// RulesSet returns the rules of the given rulset
func (svc *Taralizer) RuleSet() RuleSet {

	results := svc.queryString("data.asvs.ruleset")

	// load model into structured report
	ruleset := RuleSet{}
	if len(results) == 1 {
		data := results[0].Expressions[0].Value.(map[string]interface{})
		ruleset = svc.convertMapToRuleSet(data)
	}

	return ruleset
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

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	defaultRulesDir := []string{"./rules/", "/etc/taralizer/rules/", exPath + "/rules/", "../../rules/"}
	rules := []string{}
	for _, v := range defaultRulesDir {
		if _, err := os.Stat(v); !os.IsNotExist(err) {
			rules = append(rules, v)
		}
	}

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

// // Evaluate executes an Open Policy Agent (OPA) query against the rule sets calling the given query 'queryStr'
func (svc *Taralizer) queryString(queryStr string) rego.ResultSet {

	defaultRulesDir := []string{"./rules/", "/etc/taralizer/rules/", "../../rules/"}
	rules := []string{}
	for _, v := range defaultRulesDir {
		if _, err := os.Stat(v); !os.IsNotExist(err) {
			rules = append(rules, v)
		}
	}

	query, err := rego.New(rego.Query(queryStr), rego.Load(rules, nil)).PrepareForEval(svc.ctx)
	if err != nil {
		log.Fatalf("cannot load model file into rego engine: %v", err)
	}

	results, err := query.Eval(svc.ctx)

	if err != nil {
		log.Fatalf("cannot evaluate rules against input: %v", err)
	}
	return results
}
