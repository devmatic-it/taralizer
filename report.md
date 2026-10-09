# Threat and Risk Analysis: Bank of Anthos

**Customer:** Anthos Bank Investment Ltd.
**Version:** 1.1
**Date:** 2020-07-02
**Author:** Ferenc Bator

---

## System Description

### Data Flow Diagram

```mermaid
flowchart TD
    unauthorized-external-agent["TA1"]
    authorized-employee["TA2"]
    unauthorized-insider["TA3"]
    browser["User"]
    developer_client["DevOps"]
    waf["WAF"]
    ingress["istio-ingress-gateway"]
    frontend["frontend"]
    user_service["User Service"]
    contacts["Contacts"]
    user_service["User Service"]
    accounts_db["accounts-db"]
    ledger_writer["ledger writer"]
    balance_reader["balance reader"]
    trans_history["trans history"]
    ledger_db["ledger-db"]
    kms["KMS"]
    hsm["HSM"]
    logging["Cloud Logging"]
    monitoring["Cloud Monitoring"]
    subgraph Internet ["Internet"]
        unauthorized-external-agent
        subgraph Google_Cloud_Platform ["Google Cloud Platform"]
            unauthorized-external-agent
            subgraph project_boa_gke ["project boa-gke"]
                subgraph VPC ["VPC"]
                    subgraph public_subnet ["public subnet"]
                        waf
                    end
                    subgraph privat_subnet ["privat subnet"]
                        subgraph GKE ["GKE"]
                            subgraph istion_system ["istion-system"]
                                ingress
                            end
                            subgraph frontend ["frontend"]
                            end
                            subgraph accounts ["accounts"]
                                user_service
                                contacts
                            end
                            subgraph transactions ["transactions"]
                                ledger_writer
                                balance_reader
                                trans_history
                            end
                        end
                    end
                end
            end
            subgraph project_boa_sql ["project boa-sql"]
                accounts_db
                ledger_db
            end
            subgraph project_boa_sec ["project boa-sec"]
                kms
            end
            subgraph project_boa_ops ["project boa-ops"]
                logging
                monitoring
            end
        end
        browser
    end
    subgraph Company_Network ["Company Network"]
        authorized-employee
        unauthorized-insider
        developer_client
        hsm
    end
    browser -->|https / Link to waf| waf
    developer_client -->|https / Link to waf| waf
    waf -->|https / Link to app| ingress
    ingress -->|https| frontend
    frontend -->|https| user_service
    frontend -->|https| contacts
    frontend -->|https| trans_history
    frontend -->|https| balance_reader
    frontend -->|https| ledger_writer
    user_service -->|sql_encrypted / Link to database| accounts_db
    contacts -->|sql_encrypted / Link to database| accounts_db
    user_service -->|sql_encrypted / Link to database| accounts_db
    ledger_writer -->|sql_encrypted / Link to database| ledger_db
    balance_reader -->|sql_encrypted / Link to database| ledger_db
    trans_history -->|sql_encrypted / Link to database| ledger_db
    hsm -->|kek / key transfer| kms

```

### Trust Boundaries

| Name | Technology | Description |
|------|------------|-------------|
| Internet | internet | Internet |
| Google Cloud Platform | internet | Google Cloud |
| project boa-gke | project | Google Project |
| VPC | vpc | Virtual Private Cloud |
| public subnet | subnet | internet facing zone |
| privat subnet | subnet | Application Network |
| GKE | kubernetes-cluster | kubernetes cluster |
| frontend | kubernetes-network-policies | kubernetes frontend namespace |
| accounts | kubernetes-network-policies | kubernetes default namespace |
| transactions | kubernetes-network-policies | kubernetes default namespace |
| istion-system | kubernetes-network-policies | kubernetes istio-system |
| project boa-sql | cloud-services | Azure Shared Services |
| project boa-sec | project | Contains Secret Manager and KMS instances for secrets that are specific to the Bank of Anthos application. |
| project boa-ops | project | Used for storing environment logs as well as monitoring the environment instance of the Bank of Anthos application. |
| Company Network | on-premise | trusted on-premise company network |


### Technical Assets

| Name | Technology | Description |
|------|------------|-------------|
| User | browser | The browser used by the end customer |
| DevOps | devops-client | laptop used by developers and operators to manage the system |
| WAF | waf | Google Cloud Armor Web Application Firewall |
| istio-ingress-gateway | kubernetes-ingress | ISTIO ingress gateway |
| frontend | kubernetes-pod | Exposes an HTTP server to serve the website. Contains a login page, a signup page, and a home page. |
| User Service | kubernetes-pod |  |
| Contacts | kubernetes-pod | Stores a list of additional accounts that are associated with a user. These accounts are listed in the application's Send Payment and Deposit forms. |
| User Service | kubernetes-pod | Manages user accounts and authentication. The service signs JWTs that are used for authentication by other services. |
| accounts-db | database-sql | SQL database |
| ledger writer | kubernetes-pod | Accepts and validates incoming transactions before writing them to the ledger. |
| balance reader | kubernetes-pod | Provides an efficient readable cache of user balances, as read from ledger-db. |
| trans history | kubernetes-pod | Provides an efficient readable cache of past transactions, as read from ledger-db. |
| ledger-db | database-nosql | CloudSQL for hyperledger |
| KMS | hsm | Google Key Management Service |
| HSM | hsm | on-premise Harware Security Module (HSM) |
| Cloud Logging | logging | Cloud Monitoring |
| Cloud Monitoring | monitoring | KMS |


---

## Problem Description

### Data Assets

| Name | Description | C | I | A |
|------|-------------|---|---|---|
| A1 | Angular and other client-side code delivered by the application. | Public | Public | Public |
| A2 | OIDC identity token | Internal | Public | Public |
| A3 | OAuth2 access token | Internal | Public | Public |
| A4 | OAuth2 refresh token | Internal | Public | Public |
| A5 | active/ongoing bank transactions | Restricted | Internal | Internal |
| A6 | bank transaction history | Restricted | Internal | Internal |
| A7 | account balance of client | Restricted | Internal | Internal |
| A8 | root certificates | Restricted | Internal | Internal |
| A9 | Session ID | Internal | Internal | Internal |
| A10 | personal-related end user information like e-mail to identity user | Unknown | Unknown | Unknown |
| A11 | passwort of a user | Restricted | Internal | Internal |
| A12 | contacts of the customer from past transactions | Internal | Internal | Internal |
| A13 | credentials to access ledger db | Restricted | Internal | Internal |
| A14 | credentials to access ledger db | Restricted | Internal | Internal |


### Threat Agents

| Name | Description |
|------|-------------|
| TA1 | external threat agent |
| TA2 | authorized employee |
| TA3 | Employee/Contractor of the company that is not authorized to access the system |


---

## Risk Assessment

### Risk Matrix

We follow the [OWASP Risk Rating Methodology](https://owasp.org/www-community/OWASP_Risk_Rating_Methodology).

| Likelihood →<br>Impact ↓ | LOW | MEDIUM | HIGH |
|--------------------------|-----|--------|------|
| **LOW** | Low | Low | Medium |
| **MEDIUM** | Low | Medium | High |
| **HIGH** | Medium | High | Critical |

### Identified Risks

| ID | Likelihood | Impact | Severity | Risk |
|----|------------|--------|----------|------|
| cross-site-request-forgery@balance_reader | HIGH (3) | MED (2) | HIGH (6) | [CWE-352](https://cwe.mitre.org/data/definitions/352) **Cross-Site Request Forgery**: asset 'balance_reader' has risk of Cross Site Request Forgery(CSRF)<br><br><br>**Mitigation:**  |
| cross-site-request-forgery@contacts | HIGH (3) | MED (2) | HIGH (6) | [CWE-352](https://cwe.mitre.org/data/definitions/352) **Cross-Site Request Forgery**: asset 'contacts' has risk of Cross Site Request Forgery(CSRF)<br><br><br>**Mitigation:**  |
| cross-site-request-forgery@frontend | HIGH (3) | MED (2) | HIGH (6) | [CWE-352](https://cwe.mitre.org/data/definitions/352) **Cross-Site Request Forgery**: asset 'frontend' has risk of Cross Site Request Forgery(CSRF)<br><br><br>**Mitigation:**  |
| cross-site-request-forgery@ledger_writer | HIGH (3) | MED (2) | HIGH (6) | [CWE-352](https://cwe.mitre.org/data/definitions/352) **Cross-Site Request Forgery**: asset 'ledger_writer' has risk of Cross Site Request Forgery(CSRF)<br><br><br>**Mitigation:**  |
| cross-site-request-forgery@trans_history | HIGH (3) | MED (2) | HIGH (6) | [CWE-352](https://cwe.mitre.org/data/definitions/352) **Cross-Site Request Forgery**: asset 'trans_history' has risk of Cross Site Request Forgery(CSRF)<br><br><br>**Mitigation:**  |
| cross-site-request-forgery@user_service | HIGH (3) | MED (2) | HIGH (6) | [CWE-352](https://cwe.mitre.org/data/definitions/352) **Cross-Site Request Forgery**: asset 'user_service' has risk of Cross Site Request Forgery(CSRF)<br><br><br>**Mitigation:**  |
| container-baseimage-backdooring@balance_reader | LOW (1) | MED (2) | LOW (2) | [CWE-513](https://cwe.mitre.org/data/definitions/513) **Container Baseimage Backdooring**: asset 'balance_reader' has risk of container image backdooring through base images<br><br><br>**Mitigation:** Google distroless image is used to mitigate risks.  <p>The product owner accepted the residual risks.</p>
 |
| container-baseimage-backdooring@contacts | LOW (1) | MED (2) | LOW (2) | [CWE-513](https://cwe.mitre.org/data/definitions/513) **Container Baseimage Backdooring**: asset 'contacts' has risk of container image backdooring through base images<br><br><br>**Mitigation:** Google distroless image is used to mitigate risks.  <p>The product owner accepted the residual risks.</p>
 |
| container-baseimage-backdooring@frontend | LOW (1) | MED (2) | LOW (2) | [CWE-513](https://cwe.mitre.org/data/definitions/513) **Container Baseimage Backdooring**: asset 'frontend' has risk of container image backdooring through base images<br><br><br>**Mitigation:** Google distroless image is used to mitigate risks.  <p>The product owner accepted the residual risks.</p>
 |
| container-baseimage-backdooring@ledger_writer | LOW (1) | MED (2) | LOW (2) | [CWE-513](https://cwe.mitre.org/data/definitions/513) **Container Baseimage Backdooring**: asset 'ledger_writer' has risk of container image backdooring through base images<br><br><br>**Mitigation:** Google distroless image is used to mitigate risks.  <p>The product owner accepted the residual risks.</p>
 |
| container-baseimage-backdooring@trans_history | LOW (1) | MED (2) | LOW (2) | [CWE-513](https://cwe.mitre.org/data/definitions/513) **Container Baseimage Backdooring**: asset 'trans_history' has risk of container image backdooring through base images<br><br><br>**Mitigation:** Google distroless image is used to mitigate risks.  <p>The product owner accepted the residual risks.</p>
 |
| container-baseimage-backdooring@user_service | LOW (1) | MED (2) | LOW (2) | [CWE-513](https://cwe.mitre.org/data/definitions/513) **Container Baseimage Backdooring**: asset 'user_service' has risk of container image backdooring through base images<br><br><br>**Mitigation:** Google distroless image is used to mitigate risks.  <p>The product owner accepted the residual risks.</p>
 |
| software-integrity@balance_reader | MED (2) | VERY HIGH (4) | HIGH (8) | [CWE-697](https://cwe.mitre.org/data/definitions/697) **Software Integrity**: asset 'balance_reader' lacks software integrity verification (signing, SBOM, etc.)<br><br><br>**Mitigation:**  |
| software-integrity@contacts | MED (2) | VERY HIGH (4) | HIGH (8) | [CWE-697](https://cwe.mitre.org/data/definitions/697) **Software Integrity**: asset 'contacts' lacks software integrity verification (signing, SBOM, etc.)<br><br><br>**Mitigation:**  |
| software-integrity@frontend | MED (2) | VERY HIGH (4) | HIGH (8) | [CWE-697](https://cwe.mitre.org/data/definitions/697) **Software Integrity**: asset 'frontend' lacks software integrity verification (signing, SBOM, etc.)<br><br><br>**Mitigation:**  |
| software-integrity@ledger_writer | MED (2) | VERY HIGH (4) | HIGH (8) | [CWE-697](https://cwe.mitre.org/data/definitions/697) **Software Integrity**: asset 'ledger_writer' lacks software integrity verification (signing, SBOM, etc.)<br><br><br>**Mitigation:**  |
| software-integrity@trans_history | MED (2) | VERY HIGH (4) | HIGH (8) | [CWE-697](https://cwe.mitre.org/data/definitions/697) **Software Integrity**: asset 'trans_history' lacks software integrity verification (signing, SBOM, etc.)<br><br><br>**Mitigation:**  |
| software-integrity@user_service | MED (2) | VERY HIGH (4) | HIGH (8) | [CWE-697](https://cwe.mitre.org/data/definitions/697) **Software Integrity**: asset 'user_service' lacks software integrity verification (signing, SBOM, etc.)<br><br><br>**Mitigation:**  |
| vulnerable-dependencies@balance_reader | MED (2) | HIGH (3) | HIGH (6) | [CWE-1101](https://cwe.mitre.org/data/definitions/1101) **Vulnerable Dependencies**: asset 'balance_reader' uses vulnerable dependencies without scanning<br><br><br>**Mitigation:**  |
| vulnerable-dependencies@contacts | MED (2) | HIGH (3) | HIGH (6) | [CWE-1101](https://cwe.mitre.org/data/definitions/1101) **Vulnerable Dependencies**: asset 'contacts' uses vulnerable dependencies without scanning<br><br><br>**Mitigation:**  |
| vulnerable-dependencies@frontend | MED (2) | HIGH (3) | HIGH (6) | [CWE-1101](https://cwe.mitre.org/data/definitions/1101) **Vulnerable Dependencies**: asset 'frontend' uses vulnerable dependencies without scanning<br><br><br>**Mitigation:**  |
| vulnerable-dependencies@ledger_writer | MED (2) | HIGH (3) | HIGH (6) | [CWE-1101](https://cwe.mitre.org/data/definitions/1101) **Vulnerable Dependencies**: asset 'ledger_writer' uses vulnerable dependencies without scanning<br><br><br>**Mitigation:**  |
| vulnerable-dependencies@trans_history | MED (2) | HIGH (3) | HIGH (6) | [CWE-1101](https://cwe.mitre.org/data/definitions/1101) **Vulnerable Dependencies**: asset 'trans_history' uses vulnerable dependencies without scanning<br><br><br>**Mitigation:**  |
| vulnerable-dependencies@user_service | MED (2) | HIGH (3) | HIGH (6) | [CWE-1101](https://cwe.mitre.org/data/definitions/1101) **Vulnerable Dependencies**: asset 'user_service' uses vulnerable dependencies without scanning<br><br><br>**Mitigation:**  |


---

## Methodology

### STRIDE

- **S**poofing — Impersonating another entity
- **T**ampering — Modifying data or code
- **R**epudiation — Denying actions occurred
- **I**nformation Disclosure — Leaking data
- **D**enial of Service — Disrupting availability
- **E**levation of Privilege — Gaining unauthorized access

### Likelihood Scale

| Score | Level |
|-------|-------|
| 0 | NONE |
| 1 | LOW |
| 2 | MEDIUM |
| 3 | HIGH |
| ≥4 | VERY HIGH |

### Impact Scale

| Score | Level |
|-------|-------|
| 1 | LOW |
| 2 | MEDIUM |
| 3 | HIGH |
| 6 | HIGH |
| 9 | CRITICAL |

---

## About Taralizer

### Risk Rules: OWASP Application Security Verification Standard — 4.0.2

The OWASP Application Security Verification Standard is specified [here](https://github.com/OWASP/ASVS/raw/v4.0.2/4.0/OWASP%20Application%20Security%20Verification%20Standard%204.0.2-en.pdf).



### Disclaimer

Ferenc Bator conducted this threat analysis using the open-source TARALIZER toolkit on the applications and systems that were modeled as of this report's date. Information security threats are continually changing, with new vulnerabilities discovered on a daily basis, and no application can ever be 100% secure no matter how much threat modeling is conducted. It is recommended to execute threat modeling and also penetration testing on a regular basis (for example yearly) to ensure a high ongoing level of security and constantly check for new attack vectors.

This report cannot and does not protect against personal or business loss as the result of use of the applications or systems described. Ferenc Bator and the TARALIZER toolkit offers no warranties, representations or legal certifications concerning the applications or systems it tests. All software includes defects: nothing in this document is intended to represent or warrant that threat modeling was complete and without error, nor does this document represent or warrant that the architecture analyzed is suitable to task, free of other defects than reported, fully compliant with any industry standards, or fully compatible with any operating system, hardware, or other application.

Threat modeling tries to analyze the modeled architecture without having access to a real working system and thus cannot and does not test the implementation for defects and vulnerabilities. These kinds of checks would only be possible with a separate code review and penetration test against a working system and not via a threat model.

By using the resulting information you agree that Ferenc Bator and the TARALIZER toolkit shall be held harmless in any event.

This report is confidential and intended for internal, confidential use by the client. The recipient is obligated to ensure the highly confidential contents are kept secret. The recipient assumes responsibility for further distribution of this document.

In this particular project, a timebox approach was used to define the analysis effort. This means that the author allotted a prearranged amount of time to identify and document threats. Because of this, there is no guarantee that all possible threats and risks are discovered. Furthermore, the analysis applies to a snapshot of the current state of the modeled architecture (based on the architecture information provided by the customer) at the examination time.

### Report Distribution

Distribution of this report (in full or in part like diagrams or risk findings) requires that this disclaimer as well as the chapter about the TARALIZER toolkit and method used is kept intact as part of the distributed report or referenced from the distributed parts.
