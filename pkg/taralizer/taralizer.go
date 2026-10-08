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
	"github.com/open-policy-agent/opa/v1/rego"
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

	// Read metadata directly from the violation object (embedded in rego # METADATA block)
	if title, ok := data["title"].(string); ok {
		risk.Title = title
	}
	if desc, ok := data["description"].(string); ok {
		risk.Description = desc
	}
	if mitigation, ok := data["mitigation"].(string); ok {
		risk.Mitigation = mitigation
	}
	if url, ok := data["url"].(string); ok {
		risk.Url = url
	}
	if cwe, ok := data["cwe"].(float64); ok {
		risk.Cwe = int64(cwe)
	} else if cwe, ok := data["cwe"].(int); ok {
		risk.Cwe = int64(cwe)
	} else if cwe, ok := data["cwe"].(int64); ok {
		risk.Cwe = cwe
	}
	return risk, nil
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

	// Load ruleset.yaml directly (not through OPA) since it's metadata, not Rego code.
	// Find the ruleset.yaml file in the asvs directory.
	defaultRulesDirs := []string{"./rules/asvs/", "/etc/taralizer/rules/asvs/", "../../rules/asvs/"}
	var rulesetPath string
	for _, dir := range defaultRulesDirs {
		path := dir + "ruleset.yaml"
		if _, err := os.Stat(path); err == nil {
			rulesetPath = path
			break
		}
	}

	if rulesetPath == "" {
		return RuleSet{}
	}

	// Read and parse the ruleset.yaml file
	data, err := os.ReadFile(rulesetPath)
	if err != nil {
		return RuleSet{}
	}

	var model interface{}
	if err := yaml.Unmarshal(data, &model); err != nil {
		return RuleSet{}
	}

	modelMap, ok := model.(map[string]interface{})
	if !ok {
		return RuleSet{}
	}

	rulesetData, ok := modelMap["ruleset"].(map[string]interface{})
	if !ok {
		return RuleSet{}
	}

	// Convert to RuleSet struct
	ruleSet := RuleSet{}
	if name, ok := rulesetData["name"].(string); ok {
		ruleSet.Name = name
	}
	if title, ok := rulesetData["title"].(string); ok {
		ruleSet.Title = title
	}
	if desc, ok := rulesetData["description"].(string); ok {
		ruleSet.Description = desc
	}
	if version, ok := rulesetData["version"].(string); ok {
		ruleSet.Version = version
	}
	if url, ok := rulesetData["url"].(string); ok {
		ruleSet.Url = url
	}
	if rulesData, ok := rulesetData["rules"].([]interface{}); ok {
		for _, v := range rulesData {
			ruleMap, ok := v.(map[string]interface{})
			if !ok {
				continue
			}
			// Handle both formats: with and without 'rule:' wrapper
			var ruleData map[string]interface{}
			if rd, ok := ruleMap["rule"].(map[string]interface{}); ok && rd != nil {
				ruleData = rd
			} else {
				ruleData = ruleMap
			}
			rule := Rule{}
			if id, ok := ruleData["id"].(string); ok {
				rule.Id = id
			}
			if title, ok := ruleData["title"].(string); ok {
				rule.Title = title
			}
			if desc, ok := ruleData["description"].(string); ok {
				rule.Description = desc
			}
			if mitigation, ok := ruleData["mitigation"].(string); ok {
				rule.Mitigation = mitigation
			}
			if url, ok := ruleData["url"].(string); ok {
				rule.Url = url
			}
			if likelihood, ok := ruleData["likelihood"].(float64); ok {
				rule.Likelihood = int64(likelihood)
			} else if likelihood, ok := ruleData["likelihood"].(int); ok {
				rule.Likelihood = int64(likelihood)
			} else if likelihood, ok := ruleData["likelihood"].(int64); ok {
				rule.Likelihood = likelihood
			}
			if impact, ok := ruleData["impact"].(float64); ok {
				rule.Impact = int64(impact)
			} else if impact, ok := ruleData["impact"].(int); ok {
				rule.Impact = int64(impact)
			} else if impact, ok := ruleData["impact"].(int64); ok {
				rule.Impact = impact
			}
			if cwe, ok := ruleData["cwe"].(float64); ok {
				rule.Cwe = int64(cwe)
			} else if cwe, ok := ruleData["cwe"].(int); ok {
				rule.Cwe = int64(cwe)
			} else if cwe, ok := ruleData["cwe"].(int64); ok {
				rule.Cwe = cwe
			}
			ruleSet.Rules = append(ruleSet.Rules, rule)
		}
	}

	svc.ruleset = ruleSet
	return ruleSet
}

// query executes an Open Policy Agent (OPA) query against the rule sets
func (svc *Taralizer) query(fileName string, queryStr string) rego.ResultSet {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("cannot load model file: %v", err)
	}
	defer jsonFile.Close()

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
	defaultRulesDir = append(defaultRulesDir, "./rules/asvs/", "/etc/taralizer/rules/asvs/", exPath+"/rules/asvs/", "../../rules/asvs/")
	rules := []string{}
	for _, v := range defaultRulesDir {
		if _, err := os.Stat(v); !os.IsNotExist(err) {
			rules = append(rules, v)
		}
	}

	// Build explicit list of Rego files, excluding YAML metadata files.
	// OPA bundles directories, so we need to list individual Rego files.
	// The ruleset.yaml is metadata used by RuleSet(), not Rego code.
	var files []string
	for _, r := range rules {
		entries, err := os.ReadDir(r)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			name := entry.Name()
			// Skip YAML metadata files (ruleset.yaml)
			if strings.HasSuffix(name, ".yaml") {
				continue
			}
			// Skip non-Rego files
			if !entry.IsDir() && !strings.HasSuffix(name, ".rego") {
				continue
			}
			if !entry.IsDir() {
				files = append(files, r+name)
			}
		}
	}

	if len(files) == 0 {
		// Fallback: use directories if no Rego files found
		files = rules
	}

	query, err := rego.New(rego.Query(queryStr), rego.Load(files, nil)).PrepareForEval(svc.ctx)
	if err != nil {
		log.Printf("WARN: cannot load model file into rego engine: %v", err)
		return nil
	}

	results, err := query.Eval(svc.ctx, rego.EvalInput(model))

	if err != nil {
		log.Printf("WARN: cannot evaluate rules against input: %v", err)
		return nil
	}
	return results
}

// // Evaluate executes an Open Policy Agent (OPA) query against the rule sets calling the given query 'queryStr'
func (svc *Taralizer) queryString(queryStr string) rego.ResultSet {

	defaultRulesDir := []string{"./rules/", "/etc/taralizer/rules/", "../../rules/"}
	defaultRulesDir = append(defaultRulesDir, "./rules/asvs/", "/etc/taralizer/rules/asvs/", "../../rules/asvs/")
	rules := []string{}
	for _, v := range defaultRulesDir {
		if _, err := os.Stat(v); !os.IsNotExist(err) {
			rules = append(rules, v)
		}
	}

	// Build explicit list of Rego files, excluding YAML metadata files.
	// OPA bundles directories, so we need to list individual Rego files.
	// The ruleset.yaml is metadata used by RuleSet(), not Rego code.
	var files []string
	for _, r := range rules {
		entries, err := os.ReadDir(r)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			name := entry.Name()
			// Skip YAML metadata files (ruleset.yaml)
			if strings.HasSuffix(name, ".yaml") {
				continue
			}
			// Skip non-Rego files
			if !entry.IsDir() && !strings.HasSuffix(name, ".rego") {
				continue
			}
			if !entry.IsDir() {
				files = append(files, r+name)
			}
		}
	}

	if len(files) == 0 {
		// Fallback: use directories if no Rego files found
		files = rules
	}

	query, err := rego.New(rego.Query(queryStr), rego.Load(files, nil)).PrepareForEval(svc.ctx)
	if err != nil {
		log.Printf("WARN: cannot load model file into rego engine: %v", err)
		return nil
	}

	results, err := query.Eval(svc.ctx)

	if err != nil {
		log.Printf("WARN: cannot evaluate rules against input: %v", err)
		return nil
	}
	return results
}
