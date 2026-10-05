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
package taralizer

import (
	"context"
	"fmt"
	"github.com/open-policy-agent/opa/rego"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const RULSET_YAML = "ruleset.yaml"
const RULE_REGO = "rego rule"

// Taralzer struct
type Taralizer struct {
	ctx     context.Context
	ruleset RuleSet
}

// New creates a new instance of the Taralizer engine.
func NewTaralizer(ruleset string) *Taralizer {
	instance := Taralizer{}
	instance.ctx = context.TODO()
	// Note: This assumes RuleSet(string) exists and returns a RuleSet. 
	// Based on previous errors, it seems the user might have intended to call a method or field.
	// However, I will fix only the compilation errors identified.
	instance.ruleset = instance.RuleSet(ruleset)
	return &instance
}

// convertMapToRisk converts an untyped instance Risk into a typed one.
func (svc *Taralizer) convertMapToRisk(input interface{}) (Risk, error) {
	if input == nil {
		return Risk{}, fmt.Errorf("convertMapToRisk: interface cannot be nil")
	}

	data := input.(map[string]interface{})
	id, err := GetMapStringValue(data, "id", RULE_REGO)
	if err != nil {
		return Risk{}, err
	}
	msg, err := GetMapStringValue(data, "msg", RULE_REGO)
	if err != nil {
		return Risk{}, err
	}

	risk := Risk{
		Id:      id,
		Message: msg,
	}

	likelihood, err := GetMapIntValue(data, "likelihood", RULE_REGO)
	if err != nil {
		return Risk{}, err
	}
	risk.Likelihood = likelihood

	impact, err := GetMapIntValue(data, "impact", RULE_REGO)
	if err != nil {
		return Risk{}, err
	}
	risk.Impact = impact

	rule := svc.findRule(risk.Id)
	if rule != nil {
		risk.Title = rule.Title
		risk.Description = rule.Description
		risk.Mitigation = rule.Mitigation
		risk.Url = rule.Url
		risk.Cwe = rule.Cwe
	}
	return risk, nil
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

// convertMapToRule converts an untyped instance Rule into a typed one.
func (svc *Taralizer) convertMapToRule(input interface{}) (Rule, error) {
	if input == nil {
		return Rule{}, fmt.Errorf("convertMapToRule: interface cannot be nil")
	}

	loc := svc.ruleset.Name + "/" + RULSET_YAML
	data := input.(map[string]interface{})
	rule := Rule{
		Id:          "",
		Title:       "",
		Description: "",
		Mitigation:  "",
		Url:         "",
	}

	id, err := GetMapStringValue(data, "id", loc)
	if err != nil {
		return Rule{}, err
	}
	rule.Id = id

	title, err := GetMapStringValue(data, "title", loc)
	if err != nil {
		return Rule{}, err
	}
	rule.Title = title

	description, err := GetMapStringValue(data, "description", loc)
	if err != nil {
		return Rule{}, err
	}
	rule.Description = description

	mitigation, err := GetMapStringValue(data, "mitigation", loc)
	if err != nil {
		return Rule{}, err
	}
	rule.Mitigation = mitigation

	url, err := GetMapStringValue(data, "url", loc)
	if err != nil {
		return Rule{}, err
	}
	rule.Url = url

	cwe, err := GetMapIntValue(data, "cwe", loc)
	if err != nil {
		return Rule{}, err
	}
	rule.Cwe = cwe

	likelihood, err := GetMapIntValue(data, "likelihood", loc)
	if err != nil {
		return Rule{}, err
	}
	rule.Likelihood = likelihood

	impact, err := GetMapIntValue(data, "impact", loc)
	if err != nil {
		return Rule{}, err
	}
	rule.Impact = impact

	return rule, nil
}

// convertMapToRuleSet converts an untyped instance RuleSet
func (svc *Taralizer) convertMapToRuleSet(input interface{}) (RuleSet, error) {
	if input == nil {
		return RuleSet{}, fmt.Errorf("convertMapToRuleSet: interface cannot be nil")
	}

	loc := svc.ruleset.Name + "/" + RULSET_YAML
	data := input.(map[string]interface{})
	ruleSet := RuleSet{
		Name:        "",
		Title:       "",
		Description: "",
		Version:     "",
		Url:         "",
		Rules:       []Rule{},
	}

	name, err := GetMapStringValue(data, "name", loc)
	if err != nil {
		return RuleSet{}, err
	}
	ruleSet.Name = name

	title, err := GetMapStringValue(data, "title", loc)
	if err != nil {
		return RuleSet{}, err
	}
	ruleSet.Title = title

	description, err := GetMapStringValue(data, "description", loc)
	if err != nil {
		return RuleSet{}, err
	}
	ruleSet.Description = description

	version, err := GetMapStringValue(data, "version", loc)
	if err != nil {
		return RuleSet{}, err
	}
	ruleSet.Version = version

	url, err := GetMapStringValue(data, "url", loc)
	if err != nil {
		return RuleSet{}, err
	}
	ruleSet.Url = url

	rulesData, ok := data["rules"].([]interface{})
	if !ok {
		return RuleSet{}, fmt.Errorf("rules field is not an array")
	}

	for _, v := range rulesData {
		rule, err := svc.convertMapToRule(v)
		if err != nil {
			return RuleSet{}, err
		}
		ruleSet.Rules = append(ruleSet.Rules, rule)
	}

	return ruleSet, nil
}

// mitigateRisk reads out the measures from the model to fill in the risk mitigations
func (report *Report) addRisk(risk Risk) {
	for _, measure := range report.RiskTracking {
		match, _ := regexp.MatchString(measure.Id, risk.Id)
		if match {
			risk.Action = strings.ToUpper(measure.Action)
			risk.Mitigation = measure.Justification
			risk.Status = strings.ToUpper(measure.Status)
			risk.ResidualLikelihood = measure.ResidualLikelihood
			risk.ResidualImpact = measure.ResidualImpact
			risk.ResidualSeverity = risk.ResidualImpact * risk.ResidualLikelihood
		} else {
			risk.Action = "TBD"
			risk.Status = "OPEN"
			risk.ResidualLikelihood = -1
			risk.ResidualImpact = -1
			risk.ResidualSeverity = -1
		}
	}

	report.Risks = append(report.Risks, risk)
}

// Evaluate executes an Open Policy Agent (OPA) query against the rule sets
// and stores the resulting risks into the returned report.
func (svc *Taralizer) Evaluate(fileName string) Report {

	results := svc.query(fileName, fmt.Sprintf("data.rules.%s.violation[msg]", svc.ruleset.Name))

	// load model into structured report
	report, err := Load(fileName)
	if err != nil {
		log.Printf("Error loading model: %v", err)
	} else {
		report.RuleSet = svc.ruleset
		for i := 0; i < len(results); i++ {
			msg := results[i].Bindings["msg"]
			if msg != nil {
				item, err := svc.convertMapToRisk(msg)
				if err != nil {
					log.Printf("Error converting risk: %v", err)
					continue
				}
				item.Severity = int64(item.Likelihood) * int64(item.Impact)
				report.addRisk(item)
			}
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
		data, ok := results[i].Bindings["msg"].(map[string]interface{})
		if !ok {
			continue
		}
		id, _ := data["id"].(string)
		msg, _ := data["msg"].(string)
		item := fmt.Sprintf("%s: %s", id, msg)
		messages = append(messages, item)
	}

	return messages
}

// RulesSet returns the rules of the given rulset
func (svc *Taralizer) RuleSet(rs string) RuleSet {

	results := svc.queryString("data." + rs + ".ruleset")

	// load model into structured report
	svc.ruleset = RuleSet{}
	svc.ruleset.Name = rs

	if len(results) == 1 {
		data, ok := results[0].Expressions[0].Value.(map[string]interface{})
		if ok {
			ruleSet, err := svc.convertMapToRuleSet(data)
			if err == nil {
				svc.ruleset = ruleSet
			}
		}
	}

	return svc.ruleset
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

	data, err := io.ReadAll(jsonFile)
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