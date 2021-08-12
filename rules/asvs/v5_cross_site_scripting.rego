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

violation[{
    "id":id,
     "msg": msg, 
     "likelihood":likelihood,
     "impact": impact}] {
    server := input.technical_assets[_]

    # a WAF has some rules againest XSS
    count({x | input.technical_assets[x] ; input.technical_assets[x].technology == "waf"} ) == 0 
    server.technology == "kubernetes-pod"
    
	id := sprintf("cross-site-scripting@%v", [server.id])
    msg := sprintf("asset '%v' has risk of Cross Site Scripting (XSS)", [server.id])
    likelihood := 2 
    impact := calc_impact(2) # medium
}
