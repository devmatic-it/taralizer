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
import data.rules.calc_impact
import data.rules.has_direct_unencrypted_end_user_inbound
import data.rules.has_direct_encrypted_end_user_inbound
import data.rules.has_waf_in_inbound_path

# HIGH: Direct unencrypted from end users, no WAF
violation[{
    "id":id,
    "msg": msg,
    "likelihood": likelihood,
    "impact": impact
}] {
    server := input.technical_assets[_]
    server.technology == "web-application"
    has_direct_unencrypted_end_user_inbound(server)
    not has_waf_in_inbound_path(server)
    id := sprintf("session-management@%v", [server.id])
    msg := sprintf("asset '%v' receives direct traffic from end users over unencrypted connections without WAF protection", [server.id])
    likelihood := 3
    impact := calc_impact(2)
}

# LOW: Direct encrypted from end users, no WAF
violation[{
    "id":id,
    "msg": msg,
    "likelihood": likelihood,
    "impact": impact
}] {
    server := input.technical_assets[_]
    server.technology == "web-application"
    has_direct_encrypted_end_user_inbound(server)
    not has_waf_in_inbound_path(server)
    id := sprintf("session-management@%v", [server.id])
    msg := sprintf("asset '%v' receives direct traffic from end users without WAF protection", [server.id])
    likelihood := 1
    impact := calc_impact(2)
}
