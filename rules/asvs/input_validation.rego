# Copyright 2021 taralizer authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
package rules.asvs
import data.rules.technical_asset_by_id
import data.rules.calc_impact
import data.rules.has_input_validation

# METADATA
# title: Missing Input Validation
# description: Technical assets should validate all inputs.
# custom:
#   rule: "missing-input-validation"
#   mitigation: "apply input validation to the technical asset."
#   cwe: 20
#   likelihood: 2
#   impact: 2

violation[{
    "id":id,
    "msg": msg,
    "likelihood": likelihood,
    "impact": impact
}] contains true if {
    server := input.technical_assets[_]
    server.technology == "web-application"
    not has_input_validation(server)
    id := sprintf("input-validation@%v", [server.id])
    msg := sprintf("asset '%v' lacks input validation for user-supplied data", [server.id])
    likelihood := 3
    impact := calc_impact(3)
}
