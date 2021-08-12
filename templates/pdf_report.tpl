<!DOCTYPE html>
<html>
<head>
<style>
body{
  font-family: Arial, Helvetica, sans-serif;
  border-collapse: collapse;
  width: 800px;
}

table {
  border-collapse: collapse;
}

td, th {
  border: 1px solid #ddd;
  padding: 2px;
}

img {
width:90%;
}

tr:nth-child(even){background-color: #f2f2f2;}

th {
  padding-top: 4px;
  padding-bottom: 4px;
  text-align: left;
  background-color: #4CAF50;
  color: white;
}

.title{
  font-size:32px;
}
</style>
</head>

<body>

<h1>Scope and Assumptions</h1>
<h2>System Description</h2>
<p>
The Data Flow Diagram below provides an overview of the analyzed architecture.
<img src="diagram.png"/>
</p>

<h3>Trust Boundaries</h3>
<table>
<tr>
  <th>Name</th>
  <th>Technology</th>
  <th>Description</th>
</tr>
{{range .TrustBoundaries}}  
    <tr>
      <td> {{.Name}}</td>  
      <td> {{.Technology}}</td>  
      <td> {{.Description}} </td>
    </tr>
{{end}}
</table>
</p>

<h3>Technical Assets</h3>
<table>
<tr>
  <th>Name</th>
  <th>Technology</th>
  <th>Description</th>
</tr>
{{range .TechnicalAssets}}  
    <tr>
      <td> {{.Name}}</td>  
      <td> {{.Technology}}</td>  
      <td> {{.Description}} </td>
    </tr>
{{end}}
</table>
</p>



<h2>Problem Description</h2>
<h3>Data Assets</h3>
<p>
The following data assets are used:

<table>
<tr> 
  <th>Name</th> 
  <th>Description</th>
  <th>C</th>
  <th>I</th>
  <th>A</th>
</tr>
{{range .DataAssets}}  
    <tr>
      <td>{{.Name}}</td>  
      <td> {{.Description}}</td>
      <td> {{.Confidentiality}}</td>
      <td> {{.Integrity}}</td>
      <td> {{.Availability}}</td>
    </tr>
{{end}}
</table>
</p>

<h3>Threat Agents</h3>
<p>
The following threat agents are used:
<table>
<th>Name</th>  <th>Description</th></tr>
{{range .ThreatAgents}}  
    <tr>
    <td>{{.Name}}</td>  <td> {{.Description}} </td>
    </tr>
{{end}}
</table>
</p>


<h1>Risk Assessment</h1>

<h2>Risk Matrix</h2>
We follow the <a href="https://owasp.org/www-community/OWASP_Risk_Rating_Methodology">OWASP Risk Rating Methodology</a>.
<table>
<tbody><tr>
<th colspan="5" align="center">Overall Risk Severity = Impact x Likelihood</th>
</tr>
<tr>
<td rowspan="4" width="15%" align="center">Impact</t>
<td bgcolor="#f2f2f2" align="center">HIGH</td> <!--impact legend-->
<td bgcolor="orange" align="center">Medium</td>
<td bgcolor="red" align="center">High</td>
<td bgcolor="darkred" align="center">Critical</td>
</tr>
<tr>
<td bgcolor="#f2f2f2" align="center">MEDIUM</td> <!--impact legend-->
<td bgcolor="yellow" align="center">Low</td>
<td bgcolor="orange" align="center">Medium</td>
<td bgcolor="red" align="center">High</td>
</tr>
<tr>
<td bgcolor="#f2f2f2" align="center">LOW</td> <!--impact legend-->
<td bgcolor="lightgreen" align="center">Low</td>
<td bgcolor="yellow" align="center">Low</td>
<td bgcolor="orange" align="center">Medium</td>
</tr>
<tr>

<!-- likelihood legend -->
<td bgcolor="#f2f2f2" align="center">&nbsp;</td>
<td bgcolor="#f2f2f2" align="center">LOW</td>
<td bgcolor="#f2f2f2" align="center">MEDIUM</td>
<td bgcolor="#f2f2f2" align="center">HIGH</td>
</tr>
<tr>
<td align="center">&nbsp;</td>
<td colspan="5" align="center">Likelihood</td>
</tr>
</tbody>
</table>

<h2>Identified Risks</h2>
<p>
The following risks have been identified:
<table>
<tr>
    <th>ID</th>
    <th>Likelihood</th>
    <th>Impact</th>
    <th>Severity</th>
    <th>Risk</th>
</tr>
{{range $index, $risk :=.Risks}}
<tr>
    <td>{{$index}}</td>  
    <td>{{likelihood $risk.Likelihood}}</td>  
    <td>{{impact $risk.Impact}}</td>  
    <td>{{severity $risk.Severity}}</td>  
    <td>
        <p><a href="https://cwe.mitre.org/data/definitions/{{$risk.Cwe}}">CWE-{{$risk.Cwe}}</a> {{$risk.Title}}: {{$risk.Message}}</p>
        <p>{{$risk.Description}}</p>
        <p>Mitigation: {{$risk.Mitigation}}</p>
    </td>
</tr>
{{end}}
</table>
</p>

  <h1>Methodology</h1>
  <h2>STRIDE</h2>
  <h2>Likelihood Scale</h2>
  <h2>Impact Scale</h2>

  <h1>About Taralizer</h1>
  <h2>Risk rules checked by Taralizer</h2>

 <h3>{{.RuleSet.Title}} - {{.RuleSet.Version}}</h3>
 The {{.RuleSet.Title}} is specified <a href="{{.RuleSet.Url}}">HERE</a>
 <p>The following list provides supported rules:<p>
{{range .RuleSet.Rules}}  
    <h4>Rule {{.Id}}</h4>
    <table>
        <tr><th>Title</th>  <td> {{.Title}} </td></tr>
        <tr><th>Description</th>  <td> {{.Description}} </td></tr>
        <tr><th>CWE</th>  <td> <a href="https://cwe.mitre.org/data/definitions/{{.Cwe}}">{{.Cwe}}</a> </td></tr>
        <tr><th>Mitigration</th>  <td> {{.Mitigation}} </td></tr>
        <tr><th>URL</th>  <td> {{.Url}} </td></tr>
        <tr><th>Base Likelihood</th>  <td> {{likelihood .Likelihood}} </td></tr>
        <tr><th>Base Impact</th>  <td> {{impact .Impact}} </td></tr>

    </table>
    </tr>
{{end}}

  <h2>Disclaimer</h2>
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
  
  <h3>Report Distribution</h3>
  Distribution of this report (in full or in part like diagrams or risk findings) requires that this disclaimer as well as the chapter about the TARALIZER toolkit and method used is kept intact as part of the distributed report or referenced from the distributed parts.

</body>
</html>
