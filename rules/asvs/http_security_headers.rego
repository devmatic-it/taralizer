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
import data.rules.has_security_headers

# METADATA
# title: HTTP Security Headers
# description: Technical assets should use security headers.
# custom:
#   rule: "http-security-headers"
#   mitigation: "add security headers to the technical asset."
#   cwe: 693
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
    not has_security_headers(server)
    id := sprintf("http-security-headers@%v", [server.id])
    msg := sprintf("asset '%v' missing HTTP security headers (HSTS, CSP, X-Frame-Options, etc.)", [server.id])
    likelihood := 2
    impact := calc_impact(2)
}
