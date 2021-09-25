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
import data.rules.different_trust_boundaries
import data.rules.calc_impact
import data.rules.is_web_access_protocol
import data.rules.is_web_application_technology

violation[{
    "id":id,
     "msg": msg, 
     "likelihood":likelihood,
     "impact": impact}] {
    server := input.technical_assets[_]    
    conn := server.communication_links[_]
    is_web_access_protocol(conn.protocol)    
    different_trust_boundaries(server.id, conn.target)

    some k
    input.technical_assets[k].id == conn.target
    target:= input.technical_assets[k]    
    is_web_application_technology(target.technology)

	id := sprintf("cross-site-request-forgery@%v", [conn.target])
    msg := sprintf("asset '%v' has risk of Cross Site Request Forgery(CSRF)", [conn.target])
    likelihood := 3
    impact := calc_impact(1) # low
}
