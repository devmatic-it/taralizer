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

# helpers for new rules

# Direct inbound connections from end-user technologies (browser, mobile-app, etc.)
has_direct_end_user_inbound(server) {
    source := input.technical_assets[_]
    source.id != server.id
    is_end_user_technology(source.technology)
    source.communication_links[_].target == server.id
}

# Direct inbound from end-user over an insecure protocol
has_direct_unencrypted_end_user_inbound(server) {
    source := input.technical_assets[_]
    source.id != server.id
    is_end_user_technology(source.technology)
    conn := source.communication_links[_]
    conn.target == server.id
    not is_encrypted_protocol(conn.protocol)
}

# Direct inbound from end-user over an encrypted protocol
has_direct_encrypted_end_user_inbound(server) {
    source := input.technical_assets[_]
    source.id != server.id
    is_end_user_technology(source.technology)
    conn := source.communication_links[_]
    conn.target == server.id
    is_encrypted_protocol(conn.protocol)
}

# Server receives traffic from a WAF (trusted intermediary handles TLS)
has_waf_in_inbound_path(server) {
    source := input.technical_assets[_]
    source.id != server.id
    source.technology == "waf"
    source.communication_links[_].target == server.id
}

# Server receives traffic from a WAF or API-gateway (trusted intermediary handles validation)
has_waf_or_api_gateway_in_inbound_path(server) {
    source := input.technical_assets[_]
    source.id != server.id
    {
        source.technology == "waf"
    }
    source.communication_links[_].target == server.id
}

# Server receives direct inbound from end-user technologies (browser, mobile-app, etc.)
has_direct_end_user_inbound(server) {
    source := input.technical_assets[_]
    source.id != server.id
    is_end_user_technology(source.technology)
    source.communication_links[_].target == server.id
}

# Server receives direct inbound from outside its trust boundary
has_cross_boundary_inbound(server) {
    source := input.technical_assets[_]
    source.id != server.id
    source.communication_links[_].target == server.id
    different_trust_boundaries(server.id, source.id)
}

# Server receives traffic from a WAF or API-gateway (trusted intermediary handles validation)
has_waf_or_api_gateway_in_inbound_path(server) {
    source := input.technical_assets[_]
    source.id != server.id
    {
        source.technology == "waf"
    }
    source.communication_links[_].target == server.id
}

# Server receives direct inbound from end-user technologies (browser, mobile-app, etc.)
has_direct_end_user_inbound(server) {
    source := input.technical_assets[_]
    source.id != server.id
    is_end_user_technology(source.technology)
    source.communication_links[_].target == server.id
}

# Server receives direct inbound from outside its trust boundary
has_cross_boundary_inbound(server) {
    source := input.technical_assets[_]
    source.id != server.id
    source.communication_links[_].target == server.id
    different_trust_boundaries(server.id, source.id)
}

has_security_headers(server){
    server.security_headers == true
}

has_external_system_review(server){
    server.external_system_review == true
}

has_security_context(server){
    server.security_context == true
}

has_input_validation(server){
    server.input_validation == true
}

has_cryptographic_controls(server){
    server.cryptographic_controls == true
}

has_access_control(server){
    server.access_control == true
}

has_dependency_scanning(server){
    server.dependency_scanning == true
}

has_software_integrity(server){
    server.software_integrity == true
}

has_ssrf_protection(server){
    server.ssrf_protection == true
}

has_monitoring(server){
    server.monitoring == true
}

has_data_classification(server){
    server.data_classification == true
}
