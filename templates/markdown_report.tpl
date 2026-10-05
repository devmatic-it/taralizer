# Threat and Risk Analysis: {{.Title}}

**Customer:** {{.Customer}}
**Version:** {{.Version}}
**Date:** {{.Date}}
**Author:** {{.Author.Name}}{{if .Author.Webpage}} ([{{.Author.Webpage}}]({{.Author.Webpage}})){{end}}

---

## System Description

### Data Flow Diagram

```mermaid
flowchart TD
    {{ define "generateTrustBoundary" }}
    subgraph {{ .Name | replaceAll " " "_" }} ["{{ .Name }}"]
        {{ range $taName := .ThreatAgentsInside }}
            {{ $ta := findThreatAgent $taName }}
                {{ $ta.Id }}["{{ $ta.Name }}"]
        {{ end }}

        {{ range $boundaryName := .TrustBoundariesNested }}
            {{ $boundary := findTrustedBoundary $boundaryName }}
                {{ template "generateTrustBoundary" $boundary }}
        {{ end }}

        {{ range $assetName := .TechnicalAssetsInside }}
            {{ $asset := findTechnicalAsset $assetName }}
                {{ $asset.Id }}["{{ $asset.Name }}"]
        {{ end }}
    end
    {{ end }}

    {{ range $boundary := .TrustBoundaries }}
        {{ if isRootTrustBoundary $boundary.Id }}
            {{ template "generateTrustBoundary" $boundary }}
        {{ end }}
    {{ end }}

    {{ range $asset := .TechnicalAssets }}
        {{ $conns := .CommunicationLinks }}
        {{ range $conn := $conns }}
            {{ $asset.Id }} --> {{ $conn.Target }}
        {{ end }}
    {{ end }}
```

### Trust Boundaries

| Name | Technology | Description |
|------|------------|-------------|
{{range .TrustBoundaries}}| {{.Name}} | {{.Technology}} | {{.Description}} |
{{end}}

### Technical Assets

| Name | Technology | Description |
|------|------------|-------------|
{{range .TechnicalAssets}}| {{.Name}} | {{.Technology}} | {{.Description}} |
{{end}}

---

## Problem Description

### Data Assets

| Name | Description | C | I | A |
|------|-------------|---|---|---|
{{range .DataAssets}}| {{.Name}} | {{.Description}} | {{.Confidentiality}} | {{.Integrity}} | {{.Availability}} |
{{end}}

### Threat Agents

| Name | Description |
|------|-------------|
{{range .ThreatAgents}}| {{.Name}} | {{.Description}} |
{{end}}

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
{{range $index, $risk :=.Risks}}| {{$risk.Id}} | {{likelihood $risk.Likelihood}} | {{impact $risk.Impact}} | {{severity $risk.Severity}} | {{if $risk.Url}}[CWE-{{$risk.Cwe}}]({{$risk.Url}}) {{end}}**{{$risk.Title}}**: {{$risk.Message}}<br><br>{{$risk.Description}}<br>**Mitigation:** {{$risk.Mitigation}} |
{{end}}

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

### Risk Rules: {{.RuleSet.Title}} — {{.RuleSet.Version}}

The {{.RuleSet.Title}} is specified [here]({{.RuleSet.Url}}).

{{range .RuleSet.Rules}}
#### Rule {{.Id}}

- **Title:** {{.Title}}
- **Description:** {{.Description}}
- **CWE:** [{{.Cwe}}](https://cwe.mitre.org/data/definitions/{{.Cwe}})
- **Mitigation:** {{.Mitigation}}
- **URL:** [{{.Url}}]({{.Url}})
- **Base Likelihood:** {{likelihood .Likelihood}}
- **Base Impact:** {{impact .Impact}}

{{end}}

### Disclaimer

{{.Author.Name}} conducted this threat analysis using the open-source TARALIZER toolkit on the applications and systems that were modeled as of this report's date. Information security threats are continually changing, with new vulnerabilities discovered on a daily basis, and no application can ever be 100% secure no matter how much threat modeling is conducted. It is recommended to execute threat modeling and also penetration testing on a regular basis (for example yearly) to ensure a high ongoing level of security and constantly check for new attack vectors.

This report cannot and does not protect against personal or business loss as the result of use of the applications or systems described. {{.Author.Name}} and the TARALIZER toolkit offers no warranties, representations or legal certifications concerning the applications or systems it tests. All software includes defects: nothing in this document is intended to represent or warrant that threat modeling was complete and without error, nor does this document represent or warrant that the architecture analyzed is suitable to task, free of other defects than reported, fully compliant with any industry standards, or fully compatible with any operating system, hardware, or other application.

Threat modeling tries to analyze the modeled architecture without having access to a real working system and thus cannot and does not test the implementation for defects and vulnerabilities. These kinds of checks would only be possible with a separate code review and penetration test against a working system and not via a threat model.

By using the resulting information you agree that {{.Author.Name}} and the TARALIZER toolkit shall be held harmless in any event.

This report is confidential and intended for internal, confidential use by the client. The recipient is obligated to ensure the highly confidential contents are kept secret. The recipient assumes responsibility for further distribution of this document.

In this particular project, a timebox approach was used to define the analysis effort. This means that the author allotted a prearranged amount of time to identify and document threats. Because of this, there is no guarantee that all possible threats and risks are discovered. Furthermore, the analysis applies to a snapshot of the current state of the modeled architecture (based on the architecture information provided by the customer) at the examination time.

### Report Distribution

Distribution of this report (in full or in part like diagrams or risk findings) requires that this disclaimer as well as the chapter about the TARALIZER toolkit and method used is kept intact as part of the distributed report or referenced from the distributed parts.
