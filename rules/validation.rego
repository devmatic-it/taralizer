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

###############################################################################
# validation of trust_boundaries
###############################################################################
validation[{"id":id,"msg": msg}] {
    trust_boundary := input.trust_boundaries[_]
    trust_boundary.technology == null
    msg := sprintf("trust_boundary '%v' using unknown technology, please specify it", [trust_boundary.id])
	id := sprintf("core@%v", [trust_boundary.id])
}

validation[{"id":id,"msg": msg}] {
    trust_boundary := input.trust_boundaries[_]
    trust_boundary.technology != "internet"
    trust_boundary.technology != "cloud"
    trust_boundary.technology != "cloud-services"
    trust_boundary.technology != "on-premise"
    trust_boundary.technology != "vpc"
    trust_boundary.technology != "subnet"
    trust_boundary.technology != "project"
    trust_boundary.technology != "kubernetes-cluster"
    trust_boundary.technology != "kubernetes-network-policies"
	msg := sprintf("trust_boundary '%v' using unsupported technology '%v'", [trust_boundary.id, trust_boundary.technology])
    id := sprintf("core@%v", [trust_boundary.id])
}


###############################################################################
# validation of technical_assets
###############################################################################
validation[{"id":id,"msg": msg}] {
    server := input.technical_assets[_]
    server.technology == null
    msg := sprintf("asset '%v' using unknown technology, please specify it", [server.id])
	id := sprintf("core@%v", [server.id])
}

validation[{"id":id,"msg": msg}] {
    server := input.technical_assets[_]

    #end user
    server.technology != "browser"
    server.technology != "desktop"
    server.technology != "devops-client"
    server.technology != "mobile-app"
    server.technology != "client-system"
    server.technology != "iot-device"

    # networking    
    server.technology != "load-balancer"
    server.technology != "reverse-proxy"    
    server.technology != "gateway"

    # security
    server.technology != "waf"
    server.technology != "ips"
    server.technology != "ids"
    server.technology != "hsm"
    server.technology != "vault"
    server.technology != "identity-provider"

    # compute
    server.technology != "web-server"
    server.technology != "web-application"
    server.technology != "local-file-system"
    server.technology != "cms"
    server.technology != "web-service-rest"
    server.technology != "web-service-soap"
    server.technology != "ejb"
    server.technology != "search-index"
    server.technology != "search-engine"
    server.technology != "service-registry"
    server.technology != "event-listener"
    server.technology != "container-platform"
    server.technology != "service-mesh"
    server.technology != "ai"

    #integration
    server.technology != "batch-processing"
    server.technology != "message-queue"
    server.technology != "stream-processing"

    #database and storage
    server.technology != "data-lake"
    server.technology != "big-data-platform"
    server.technology != "database"
    server.technology != "database-sql"
    server.technology != "database-nosql"
    server.technology != "erp"
    server.technology != "block-storage"
    server.technology != "mail-server"
    server.technology != "file-server"
    server.technology != "ldap-server"
    server.technology != "identity-store-database"
    server.technology != "identity-store-ldap"

    #CI/CD
    server.technology != "build-pipeline"
    server.technology != "artifact-registry"
    server.technology != "sourcecode-repository"
    server.technology != "code-inspection-platform"

    # Kubernetes
    server.technology != "kubernetes-pod"
    server.technology != "kubernetes-ingress"
    server.technology != "kubernetes-service"
    server.technology != "kubernetes-deployment"
    server.technology != "kubernetes-configmap"
    server.technology != "kubernetes-statefulset"
    server.technology != "kubernetes-secret"
    server.technology != "kubernetes-daemonset"
    server.technology != "kubernetes-job"
    server.technology != "kubernetes-cronjob"

    #Operations
    server.technology != "monitoring"
    server.technology != "logging"


    msg := sprintf("asset '%v' using unsupported technology '%v'", [server.id, server.technology])
	id := sprintf("core@%v", [server.id])
}


validation[{"id":id,"msg": msg}] {
    server := input.technical_assets[_]
    data_processed := server.data_assets_processed[_]
    count({x | input.data_assets[x] ; input.data_assets[x].id == data_processed} ) == 0
    msg := sprintf("technical asset '%v' processing data that is not defined as data asset. Please define data asset '%v'", [server.id, data_processed])
	id := sprintf("missing-data-asset-processed@%v", [server.id])
}

validation[{"id":id,"msg": msg}] {
    server := input.technical_assets[_]
    data_stored := server.data_assets_stored[_]
    count({x | input.data_assets[x] ; input.data_assets[x].id == data_stored} ) == 0
    msg := sprintf("technical asset '%v' storing data that is not defined as data asset. Please define data asset '%v'", [server.id, data_stored])
	id := sprintf("missing-data-asset-stored@%v", [server.id])
}

###############################################################################
# validation of technical_assets.communication_links
###############################################################################
validation[{"id":id,"msg": msg}] {
    server := input.technical_assets[_]
    conn := server.communication_links[_]
    conn.protocol != "https"
    conn.protocol != "http"
    conn.protocol != "ssh"
    conn.protocol != "tcp"
    conn.protocol != "udp"
    conn.protocol != "sql"
    conn.protocol != "sql_encrypted"
    conn.protocol != "nosql"
    conn.protocol != "nosql_encrypted"
    conn.protocol != "tcp"
    conn.protocol != "smtp"
    conn.protocol != "smtps"
    conn.protocol != "ftp"
    conn.protocol != "sftp"
    conn.protocol != "ftps"
    conn.protocol != "telnet"
    conn.protocol != "rsync"
    conn.protocol != "pop3"
    conn.protocol != "imap"
    conn.protocol != "icmp"
    conn.protocol != "kek"
    msg := sprintf("asset '%v' communicating to '%v' uses unsupported protocol '%v'", [server.id, conn.target, conn.protocol])
	id := sprintf("core-unsupported-protocol@%v>%v", [server.id, conn.target])
}

validation[{"id":id,"msg": msg}] {
    server := input.technical_assets[_]
    conn := server.communication_links[_]
    conn.protocol == null    
    msg := sprintf("asset '%v' communicating to '%v' uses undefined protocol, please specify it", [server.id, conn.target])
	id := sprintf("core-missing-protocol@%v>%v", [server.id, conn.target])
}


validation[{"id":id,"msg": msg}] {
    server := input.technical_assets[_]
    conn := server.communication_links[_]
    data_sent := conn.data_assets_sent[_]

    count({x | input.data_assets[x] ; input.data_assets[x].id == data_sent} ) == 0

    msg := sprintf("technical asset '%v' sends data to '%v' which is not defined. Please define data asset '%v'", [server.id,  conn.target, data_sent])
	id := sprintf("core-missing-data-sent@%v>%v", [server.id, conn.target])
}


validation[{"id":id,"msg": msg}] {
    server := input.technical_assets[_]
    conn := server.communication_links[_]
    data_received := conn.data_assets_received[_]

    count({x | input.data_assets[x] ; input.data_assets[x].id == data_received} ) == 0

    msg := sprintf("technical asset '%v' received data from '%v' which is not defined. Please define data asset '%v'", [server.id,  conn.target, data_received])
	id := sprintf("core-missing-data-received@%v>%v", [server.id, conn.target])
}


validation[{"id":id,"msg": msg}] {
    server := input.technical_assets[_]
    conn := server.communication_links[_]
    data_received := conn.data_assets_received[_]

    some k
    input.technical_assets[k].id == conn.target
    target:= input.technical_assets[k]    

    count({x | server.data_assets_processed[x] ; server.data_assets_processed[x] == data_received} ) == 0
    count({x | server.data_assets_stored[x] ; server.data_assets_stored[x] == data_received} ) == 0


    msg := sprintf("technical asset '%v' received data from '%v' which is not processed or stored. Please process/store data asset '%v'", [target.id,  server.id, data_received])
	id := sprintf("core-missing-data-assets-handling@%v>%v", [server.id, conn.target])
}
