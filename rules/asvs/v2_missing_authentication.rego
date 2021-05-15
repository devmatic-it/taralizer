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

violation[{
    "id":id,
     "msg": msg, 
     "cwe":cwe,
     "title":title,
     "description":description,
     "mitigation":mitigation,
     "likelihood":likelihood,
     "impact": impact,
     "url":url}] {
    server := input.technical_assets[_]
    conn := server.communication_links[_]
    conn.authentication == "none"
    
    # add additional exceptions here
    target := technical_asset_by_id(conn.target)
    target.technology != "waf"
    target.technology != "ids"
    target.technology != "ips"
    target.technology != "load-balancer"
    target.technology != "reverse-proxy"

    different_trust_boundaries(conn.target, server.id)

    msg := sprintf("asset '%v' should authenticate incomming request from '%v' ", [conn.target, server.id])
	id := sprintf("missing-authentication@%v>%v", [server.id, conn.target])
	title:= "Missing Authentication"
    description:= "Technical assets should autheticate incoming requests. "
    mitigation:= "apply an authentication method to the technical asset."
	url:= "https://cheatsheetseries.owasp.org/cheatsheets/Transport_Layer_Protection_Cheat_Sheet.html"
    cwe := 306
    likelihood := 1 #unlikely
    impact := 2 # medium
}
