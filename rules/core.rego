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
package rules

technical_asset_by_id (myid) = asset{
    some k
    input.technical_assets[k].id == myid
    asset:= input.technical_assets[k]
}

data_asset_by_id (myid) = asset{
    some k
    input.data_assets[k].id == myid
    asset:= input.data_assets[k]
}


calc_impact (base)= base {
 count({x | input.data_assets[x] ; input.data_assets[x].confidentiality == 3} ) == 0
 count({x | input.data_assets[x] ; input.data_assets[x].integrity == 3} ) == 0
 count({x | input.data_assets[x] ; input.data_assets[x].availability == 3} ) ==0 
} else = base+1


different_trust_boundaries(id1, id2){
    some i, j
    input.trust_boundaries[i].technical_assets_inside[_] == id1
    input.trust_boundaries[j].technical_assets_inside[_] == id2
    i != j
}

same_trust_boundaries(id1, id2){
    some i, j
    input.trust_boundaries[i].technical_assets_inside[_] == id1
    input.trust_boundaries[j].technical_assets_inside[_] == id2
    i == j
}

direct_connection(id1, id2){
    some i
    input.technical_assets[i].id == id1
    input.technical_assets[i].communication_links[_].target == id2
}

# all end user technologies
is_end_user_technology("browser")
is_end_user_technology("mobile-app")
is_end_user_technology("client-system")
is_end_user_technology("desktop")
is_end_user_technology("devops-client")
is_end_user_technology("iot-device")

# applications that understand http/https and are usually accessed from the web
is_web_application_technology("erp")
is_web_application_technology("cms")
is_web_application_technology("identity-provider")
is_web_application_technology("web-server")
is_web_application_technology("web-application")
is_web_application_technology("kubernetes-pod")
is_web_application_technology("kubernetes-deployment")
is_web_application_technology("kubernetes-statefulset")
is_web_application_technology("web-service-rest")
is_web_application_technology("web-service-soap")

# all web access protocols
is_web_access_protocol("http")
is_web_access_protocol("https")

# all encrypted protocols
is_encrypted_protocol("https")
is_encrypted_protocol("ssh")
is_encrypted_protocol("scp")
is_encrypted_protocol("kek")
is_encrypted_protocol("sql_encrypted")
is_encrypted_protocol("nosql_encrypted")
is_encrypted_protocol("ftps")
is_encrypted_protocol("sftp")
is_encrypted_protocol("smtps")

is_unencrypted_protocol(protocol){
    not is_encrypted_protocol(protocol)
}

is_database_protocol("sql")
is_database_protocol("sql_encrypted")
is_database_protocol("nosql")
is_database_protocol("nosql_encrypted")
