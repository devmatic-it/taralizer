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
import data.rules.is_web_access_protocol
import data.rules.calc_impact

# METADATA
# title: Authentication Failure
# description: Technical assets should implement authentication failure controls.
# custom:
#   rule: "authentication-failure"
#   mitigation: "apply authentication failure controls to the technical asset."
#   cwe: 306
#   likelihood: 2
#   impact: 1

violation[{
    "id":id,
    "msg": msg,
    "likelihood": likelihood,
    "impact": impact
}] {
    server := input.technical_assets[_]
    conn := server.communication_links[_]
    is_web_access_protocol(conn.protocol)
    conn.authentication == "none"
    target := technical_asset_by_id(conn.target)
    target.technology == "web-application"
    id := sprintf("auth-failure@%v>%v", [server.id, conn.target])
    msg := sprintf("asset '%v' communicates to '%v' without authentication", [server.id, conn.target])
    likelihood := 3
    impact := calc_impact(3)
}
