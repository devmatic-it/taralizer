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

# Authentication-related data asset IDs
auth_data_asset("session-id")
auth_data_asset("id-token")
auth_data_asset("access-token")
auth_data_asset("refresh-token")
auth_data_asset("user-credentials")

# A server processes authentication data
has_auth_data_processing(server) {
    data := server.data_assets_processed[_]
    auth_data_asset(data)
}

# A server stores authentication data
has_auth_data_storing(server) {
    data := server.data_assets_stored[_]
    auth_data_asset(data)
}

# A server handles authentication data (processes or stores)
has_auth_data_handling(server) {
    has_auth_data_processing(server)
}

has_auth_data_handling(server) {
    has_auth_data_storing(server)
}

# A server has audit logging configured (manual override or exempt)
has_audit_logging(server) {
    server.audit_logging == true
}

# Servers that handle auth data but are exempt (WAF, databases, logging services)
has_audit_logging(server) {
    has_auth_data_handling(server)
    server.technology == "waf"
}

has_audit_logging(server) {
    has_auth_data_handling(server)
    server.technology == "logging"
}

has_audit_logging(server) {
    has_auth_data_handling(server)
    server.technology == "monitoring"
}

# METADATA
# title: Missing Audit Logging
# description: Technical assets should log security-relevant events.
# custom:
#   rule: "missing-audit-logging"
#   mitigation: "add audit logging to the technical asset."
#   cwe: 778
#   likelihood: 3
#   impact: 2

# Violation: Web-application handles authentication data without audit logging
violation[{
    "id":id,
    "msg": msg,
    "likelihood": likelihood,
    "impact": impact
}] {
    server := input.technical_assets[_]
    server.technology == "web-application"
    has_auth_data_handling(server)
    not has_audit_logging(server)
    id := sprintf("audit-logging@%v", [server.id])
    msg := sprintf("asset '%v' handles authentication data but lacks audit logging", [server.id])
    likelihood := 1
    impact := calc_impact(3)
}
