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
    conn.protocol!= "https"
    conn.protocol!= "ssh" 
    conn.protocol!= "kek"
    conn.protocol!= "sql_encrypted" 
    conn.protocol!= "nosql_encrypted"     
    conn.protocol!= "ssh"
    conn.protocol!= "ftps"
    conn.protocol!= "sftp"
    conn.protocol!= "scp"
    conn.protocol!= "smtps"

    different_trust_boundaries(server.id, conn.target)
    
    msg := sprintf("asset '%v' communicating to '%v' uses insecure protocol '%v'", [server.id, conn.target, conn.protocol])
	id := sprintf("insecure-proto@%v>%v", [server.id, conn.target])
	title:= "Unencrypted Communication"
    description:= "Data at transition should be encrypted."
    mitigation:= "Apply transport layer encryption to the communication link."
	url:= "https://cheatsheetseries.owasp.org/cheatsheets/Transport_Layer_Protection_Cheat_Sheet.html"
    cwe := 319
    likelihood := 1 #unlikely
    impact := 2 # medium
}
