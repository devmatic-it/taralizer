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
import data.rules.has_monitoring

# METADATA
# title: Insufficient Logging
# description: Technical assets should log security-relevant events.
# custom:
#   rule: "insufficient-logging"
#   mitigation: "add audit logging to the technical asset."
#   cwe: 778
#   likelihood: 3
#   impact: 2

violation[{
    "id":id,
    "msg": msg,
    "likelihood": likelihood,
    "impact": impact
}] contains true if {
    server := input.technical_assets[_]
    server.technology == "web-application"
    not has_monitoring(server)
    id := sprintf("insufficient-logging@%v", [server.id])
    msg := sprintf("asset '%v' lacks monitoring and alerting for security events", [server.id])
    likelihood := 1
    impact := calc_impact(3)
}
