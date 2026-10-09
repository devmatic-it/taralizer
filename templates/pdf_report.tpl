<!DOCTYPE html>
<html>
<head>
<style>
  /* ── Reset & Base ──────────────────────────────── */
  *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

  body {
    font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
    font-size: 13px;
    line-height: 1.55;
    color: #2d3748;
    max-width: 100%;
    margin: 0 auto;
    padding: 24px 20px;
  }

  /* ── Header / Branding ─────────────────────────── */
  .report-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding-bottom: 16px;
    margin-bottom: 24px;
    border-bottom: 3px solid #1a365d;
  }
  .report-header .brand {
    font-size: 20px;
    font-weight: 700;
    color: #1a365d;
    letter-spacing: 0.4px;
  }
  .report-header .meta {
    font-size: 11px;
    color: #718096;
    text-align: right;
  }

  /* ── Headings ──────────────────────────────────── */
  h1 {
    font-size: 22px;
    font-weight: 700;
    color: #1a365d;
    margin: 32px 0 12px;
  }
  h2 {
    font-size: 17px;
    font-weight: 600;
    color: #2d3748;
    margin: 24px 0 10px;
    padding-bottom: 4px;
    border-bottom: 1px solid #e2e8f0;
  }
  h3 {
    font-size: 14px;
    font-weight: 600;
    color: #4a5568;
    margin: 16px 0 8px;
  }
  h4 {
    font-size: 13px;
    font-weight: 600;
    color: #4a5568;
    margin: 12px 0 6px;
  }

  /* ── Paragraphs ────────────────────────────────── */
  p { margin-bottom: 10px; }

  /* ── Tables ────────────────────────────────────── */
  table {
    width: 100%;
    border-collapse: collapse;
    margin: 12px 0 18px;
    font-size: 12px;
  }
  thead th {
    background: #1a365d;
    color: #ffffff;
    padding: 8px 10px;
    text-align: left;
    font-weight: 600;
    font-size: 11px;
    text-transform: uppercase;
    letter-spacing: 0.3px;
  }
  thead th:first-child { border-radius: 4px 0 0 0; }
  thead th:last-child  { border-radius: 0 4px 0 0; }
  tbody td {
    padding: 8px 10px;
    border-bottom: 1px solid #e2e8f0;
    vertical-align: top;
  }
  tbody tr:hover { background: #f7fafc; }
  tbody tr:nth-child(even) { background: #fafbfc; }
  tbody tr:nth-child(even):hover { background: #f0f4f8; }

  /* ── Severity / Status Badges ──────────────────── */
  /* ── Identified Risks Table ──────────────────────── */
  .risks-table {
    width: 100%;
    table-layout: fixed;
    border-collapse: collapse;
    margin: 12px 0 18px;
    font-size: 9px;
    line-height: 1.4;
  }
  .risks-table th {
    background: #2d3748;
    color: #ffffff;
    padding: 6px 4px;
    text-align: left;
    font-weight: 600;
    font-size: 9px;
    border: 1px solid #cbd5e0;
    white-space: nowrap;
  }
  .risks-table td {
    padding: 6px 4px;
    border: 1px solid #cbd5e0;
    vertical-align: top;
    word-wrap: break-word;
    overflow-wrap: break-word;
    hyphens: auto;
  }
  .risks-table .col-id { width: 10%; }
  .risks-table .col-likelihood { width: 6%; }
  .risks-table .col-impact { width: 6%; }
  .risks-table .col-severity { width: 6%; }
  .risks-table .col-risk { width: 24%; }
  .risks-table .col-action { width: 8%; }
  .risks-table .col-mitigation { width: 15%; }
  .risks-table .col-res-impact { width: 6%; }
  .risks-table .col-res-likelihood { width: 6%; }
  .risks-table .col-res-severity { width: 6%; }
  .risks-table .col-status { width: 7%; }
  .risks-table a {
    color: #2b6cb0;
    text-decoration: none;
    font-size: 9px;
  }
  .risks-table p {
    margin: 3px 0 0 0;
    font-size: 9px;
    color: #4a5568;
    line-height: 1.3;
  }
  .badge {
    display: inline-block;
    padding: 1px 5px;
    border-radius: 8px;
    font-size: 8px;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.2px;
  }
  .sev-low      { background: #c6f6d5; color: #276749; }
  .sev-medium   { background: #fefcbf; color: #975a16; }
  .sev-high     { background: #fed7d7; color: #c53030; }
  .sev-critical { background: #b91c1c; color: #ffffff; }
  .sev-tbd      { background: #e2e8f0; color: #718096; }

  /* ── Risk Matrix ───────────────────────────────── */
  .risk-matrix {
    width: 100%;
    border-collapse: collapse;
    margin: 12px 0 18px;
    font-size: 12px;
  }
  .risk-matrix th,
  .risk-matrix td {
    padding: 10px 8px;
    text-align: center;
    border: 1px solid #cbd5e0;
  }
  .risk-matrix thead th {
    background: #2d3748;
    color: #ffffff;
    border-color: #a0aec0;
  }
  .risk-matrix .axis-label {
    background: #f7fafc;
    font-weight: 600;
    color: #4a5568;
    text-align: right;
    padding-right: 12px;
  }

  /* ── Mermaid (static for PDF) ──────────────────── */
  pre.mermaid {
    background: #f7fafc;
    padding: 12px;
    border-radius: 4px;
    margin: 12px 0 18px;
    font-size: 11px;
    overflow-x: auto;
  }

  /* ── Footer / Disclaimer ───────────────────────── */
  .disclaimer {
    margin-top: 36px;
    padding: 14px 16px;
    background: #f7fafc;
    border-left: 3px solid #a0aec0;
    font-size: 11px;
    color: #718096;
    line-height: 1.65;
  }
  .disclaimer h3 {
    font-size: 13px;
    color: #2d3748;
    margin: 0 0 6px;
  }

  /* ── Links ─────────────────────────────────────── */
  a { color: #2b6cb0; text-decoration: none; }
  a:hover { text-decoration: underline; }

  /* ── Methodology / Rules ───────────────────────── */
  .rule-table td { vertical-align: top; }
  .rule-table th {
    background: #e2e8f0;
    color: #2d3748;
    width: 150px;
  }
</style>
</head>

<body>

<!-- ── Header ─────────────────────────────────────── -->
<div class="report-header">
  <span class="brand">TARALIZER — Threat &amp; Risk Analysis</span>
  <span class="meta">
    {{.Title}}<br>
    Generated: {{.Date}}
  </span>
</div>

<!-- ── Scope & Assumptions ────────────────────────── -->
<h1>Scope and Assumptions</h1>

<h2>System Description</h2>
<p>The Data Flow Diagram below provides an overview of the analyzed architecture.</p>
<pre class="mermaid">{{mermaidDFD}}</pre>

<h3>Trust Boundaries</h3>
<table>
<tr><th>Name</th><th>Technology</th><th>Description</th></tr>
{{range .TrustBoundaries}}
<tr><td>{{.Name}}</td><td>{{.Technology}}</td><td>{{.Description}}</td></tr>
{{end}}
</table>

<h3>Technical Assets</h3>
<table>
<tr><th>Name</th><th>Technology</th><th>Description</th></tr>
{{range .TechnicalAssets}}
<tr><td>{{.Name}}</td><td>{{.Technology}}</td><td>{{.Description}}</td></tr>
{{end}}
</table>

<!-- ── Problem Description ────────────────────────── -->
<h2>Problem Description</h2>

<h3>Data Assets</h3>
<table>
<tr><th>Name</th><th>Description</th><th>C</th><th>I</th><th>A</th></tr>
{{range .DataAssets}}
<tr>
  <td>{{.Name}}</td>
  <td>{{.Description}}</td>
  <td>{{dataAssetLabel .Confidentiality}}</td>
  <td>{{dataAssetLabel .Integrity}}</td>
  <td>{{dataAssetLabel .Availability}}</td>
</tr>
{{end}}
</table>

<h3>Threat Agents</h3>
<table>
<tr><th>Name</th><th>Description</th></tr>
{{range .ThreatAgents}}
<tr><td>{{.Name}}</td><td>{{.Description}}</td></tr>
{{end}}
</table>

<!-- ── Risk Assessment ────────────────────────────── -->
<h1>Risk Assessment</h1>

<h2>Risk Matrix</h2>
<p>We follow the <a href="https://owasp.org/www-community/OWASP_Risk_Rating_Methodology">OWASP Risk Rating Methodology</a>.</p>

<table class="risk-matrix">
<tr>
  <th colspan="4" style="text-align:center;">Overall Risk Severity = Impact × Likelihood</th>
</tr>
<tr>
  <td rowspan="3" class="axis-label">Impact ↓ <br> Likelihood →</td>
  <td class="matrix-cell" style="background:#c6f6d5;">Low</td>
  <td class="matrix-cell" style="background:#fefcbf;">Medium</td>
  <td class="matrix-cell" style="background:#fed7d7;">High</td>
</tr>
<tr>
  <td class="matrix-cell" style="background:#c6f6d5;">Low</td>
  <td class="matrix-cell" style="background:#fefcbf;">Medium</td>
  <td class="matrix-cell" style="background:#fed7d7;">High</td>
</tr>
<tr>
  <td class="matrix-cell" style="background:#c6f6d5;">Low</td>
  <td class="matrix-cell" style="background:#fefcbf;">Medium</td>
  <td class="matrix-cell" style="background:#b91c1c;color:#fff;">Critical</td>
</tr>
<tr>
  <td class="axis-label">Likelihood (columns)</td>
  <td style="text-align:center;font-weight:600;">LOW</td>
  <td style="text-align:center;font-weight:600;">MEDIUM</td>
  <td style="text-align:center;font-weight:600;">HIGH</td>
</tr>
</table>

<h3>Identified Risks</h3>
<table class="risks-table">
<tr>
  <th class="col-id">ID</th>
  <th class="col-likelihood">Like.</th>
  <th class="col-impact">Impact</th>
  <th class="col-severity">Severity</th>
  <th class="col-risk">Risk</th>
  <th class="col-action">Action</th>
  <th class="col-mitigation">Mitigation</th>
  <th class="col-res-impact">R. Imp.</th>
  <th class="col-res-likelihood">R. Like.</th>
  <th class="col-res-severity">R. Sev.</th>
  <th class="col-status">Status</th>
</tr>
{{range $index, $risk :=.Risks}}
<tr>
  <td class="col-id">{{$risk.Id}}</td>
  <td class="col-likelihood"><span class="badge badge-{{lower $risk.Likelihood}}">{{likelihood $risk.Likelihood}}</span></td>
  <td class="col-impact"><span class="badge badge-{{lower $risk.Impact}}">{{impact $risk.Impact}}</span></td>
  <td class="col-severity"><span class="badge sev-{{lower $risk.Severity}}">{{severity $risk.Severity}}</span></td>
  <td class="col-risk">
    <a href="https://cwe.mitre.org/data/definitions/{{$risk.Cwe}}">CWE-{{$risk.Cwe}}</a>
    {{if gt $risk.Cwe 0}}<a href="https://cwe.mitre.org/data/definitions/{{$risk.Cwe}}">CWE-{{$risk.Cwe}}</a> {{end}}{{$risk.Title}}: {{$risk.Message}}
    <p>{{$risk.Description}}</p>
  </td>
  <td class="col-action">{{$risk.Action}}</td>
  <td class="col-mitigation">{{$risk.Mitigation}}</td>
  <td class="col-res-impact"><span class="badge badge-{{lower $risk.ResidualImpact}}">{{impact $risk.ResidualImpact}}</span></td>
  <td class="col-res-likelihood"><span class="badge badge-{{lower $risk.ResidualLikelihood}}">{{likelihood $risk.ResidualLikelihood}}</span></td>
  <td class="col-res-severity"><span class="badge sev-{{lower $risk.ResidualSeverity}}">{{severity $risk.ResidualSeverity}}</span></td>
  <td class="col-status">{{$risk.Status}}</td>
</tr>
{{end}}
</table>

<!-- ── Methodology ────────────────────────────────── -->
<h1>Methodology</h1>
<h2>STRIDE</h2>
<h2>Likelihood Scale</h2>
<h2>Impact Scale</h2>

<!-- ── About Taralizer ────────────────────────────── -->
<h1>About Taralizer</h1>
<h2>Risk rules checked by Taralizer</h2>

<h3>{{.RuleSet.Title}} — {{.RuleSet.Version}}</h3>
<p>The {{.RuleSet.Title}} is specified <a href="{{.RuleSet.Url}}">HERE</a>.</p>
<p>The following list provides supported rules:</p>

{{range .RuleSet.Rules}}
<h4>Rule {{.Id}}</h4>
<table class="rule-table">
  <tr><th>Title</th><td>{{.Title}}</td></tr>
  <tr><th>Description</th><td>{{.Description}}</td></tr>
  <tr><th>CWE</th><td><a href="https://cwe.mitre.org/data/definitions/{{.Cwe}}">{{.Cwe}}</a></td></tr>
  <tr><th>Mitigation</th><td>{{.Mitigation}}</td></tr>
  <tr><th>URL</th><td>{{.Url}}</td></tr>
  <tr><th>Base Likelihood</th><td>{{likelihood .Likelihood}}</td></tr>
  <tr><th>Base Impact</th><td>{{impact .Impact}}</td></tr>
</table>
{{end}}

<!-- ── Disclaimer ─────────────────────────────────── -->
<div class="disclaimer">
  <h3>Disclaimer</h3>
  {{.Author.Name}} conducted this threat analysis using the open-source TARALIZER toolkit on the applications and systems that were modeled as of this report's date.
  Information security threats are continually changing, with new vulnerabilities discovered on a daily basis, and no application can ever be 100% secure no matter how much threat modeling is conducted. It is recommended to execute threat modeling and also penetration testing on a regular basis (for example yearly) to ensure a high ongoing level of security and constantly check for new attack vectors.
  This report cannot and does not protect against personal or business loss as the result of use of the applications or systems described.
  {{.Author.Name}} and the TARALIZER toolkit offers no warranties, representations or legal certifications concerning the applications or systems it tests.
  All software includes defects: nothing in this document is intended to represent or warrant that threat modeling was complete and without error, nor does this document represent or warrant that the architecture analyzed is suitable to task, free of other defects than reported, fully compliant with any industry standards, or fully compatible with any operating system, hardware, or other application.
  Threat modeling tries to analyze the modeled architecture without having access to a real working system and thus cannot and does not test the implementation for defects and vulnerabilities.
  These kinds of checks would only be possible with a separate code review and penetration test against a working system and not via a threat model.
  By using the resulting information you agree that John Doe and the Threagile toolkit shall be held harmless in any event.
  This report is confidential and intended for internal, confidential use by the client.
  The recipient is obligated to ensure the highly confidential contents are kept secret.
  The recipient assumes responsibility for further distribution of this document.
  In this particular project, a timebox approach was used to define the analysis effort.
  This means that the author allotted a prearranged amount of time to identify and document threats.
  Because of this, there is no guarantee that all possible threats and risks are discovered.
  Furthermore, the analysis applies to a snapshot of the current state of the modeled architecture (based on the architecture information provided by the customer) at the examination time.
</div>

<h3>Report Distribution</h3>
<p style="font-size:11px;color:#718096;">
  Distribution of this report (in full or in part like diagrams or risk findings) requires that this disclaimer as well as the chapter about the TARALIZER toolkit and method used is kept intact as part of the distributed report or referenced from the distributed parts.
</p>

</body>
</html>
