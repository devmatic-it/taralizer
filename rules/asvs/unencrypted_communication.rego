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
import data.rules.different_trust_boundaries
import data.rules.calc_impact
import data.rules.is_unencrypted_protocol


# METADATA
# title: Insecure Protocol
# description: Technical assets should use secure protocols.
# custom:
#   rule: "insecure-proto"
#   mitigation: "apply a secure protocol to the technical asset."
#   stride: "Elevation of Privilege"
#   likelihood: 3
#   impact: 1

violation[{
    "id":id,
     "msg": msg, 
     "likelihood":likelihood,
     "impact": impact}] {
    server := input.technical_assets[_]
    conn := server.communication_links[_]
    is_unencrypted_protocol(conn.protocol)
    different_trust_boundaries(server.id, conn.target)
    

    msg := sprintf("asset '%v' communicating to '%v' uses insecure protocol '%v'", [server.id, conn.target, conn.protocol])
	id := sprintf("insecure-proto@%v>%v", [server.id, conn.target])
    likelihood := 1 #unlikely
    impact := calc_impact(2)
}
