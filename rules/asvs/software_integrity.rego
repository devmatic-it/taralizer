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
import data.rules.has_software_integrity

# METADATA
# title: Software Integrity
# description: Technical assets should verify the integrity of their software.
# custom:
#   rule: "software-integrity"
#   mitigation: "apply software integrity verification to the technical asset."
#   cwe: 697
#   likelihood: 2
#   impact: 2

violation[{
    "id":id,
    "msg": msg,
    "likelihood": likelihood,
    "impact": impact
}] {
    server := input.technical_assets[_]
    server.technology == "kubernetes-pod"
    not has_software_integrity(server)
    id := sprintf("software-integrity@%v", [server.id])
    msg := sprintf("asset '%v' lacks software integrity verification (signing, SBOM, etc.)", [server.id])
    likelihood := 2
    impact := calc_impact(3)
}
