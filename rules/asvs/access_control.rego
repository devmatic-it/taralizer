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
import data.rules.has_access_control

# METADATA
# title: Missing Access Control
# description: Technical assets should enforce access control.
# custom:
#   rule: "missing-access-control"
#   mitigation: "apply access control to the technical asset."
#   cwe: 285
#   likelihood: 2
#   impact: 1

violation[{
    "id":id,
    "msg": msg,
    "likelihood": likelihood,
    "impact": impact
}] {
    server := input.technical_assets[_]
    server.technology == "web-application"
    not has_access_control(server)
    id := sprintf("access-control@%v", [server.id])
    msg := sprintf("asset '%v' lacks proper access control checks", [server.id])
    likelihood := 3
    impact := calc_impact(3)
}
