<!DOCTYPE html>
<html lang="en">
<head>
 <title>Threat and Risk Analysis for {{.Title}}</title>
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
width:100%;
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
<span class="title">TARALIZER</span>

<h2>System Description</h2>
<p>
The Data Flow Diagram below provides an overview of the analyzed architecture.
<img src="diagram.png" alt="Data Flow Diagram"/>
</p>

<h3>Trust Boundaries</h3>
<table aria-describedby="Trust Boundaries">
<tr>
  <th scope="col">Name</th>
  <th scope="col">Technology</th>
  <th scope="col">Description</th>
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
<table aria-describedby="Technical Assets">
<tr>
  <th scope="col">Name</th>
  <th scope="col">Technology</th>
  <th scope="col">Description</th>
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

<table aria-describedby="Data Assets">
<tr> 
  <th scope="col">Name</th> 
  <th scope="col">Description</th>
  <th scope="col">C</th>
  <th scope="col">I</th>
  <th scope="col">A</th>
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
<table aria-describedby="Threat Agents">
<tr>
  <th scope="col">Name</th>  
  <th scope="col">Description</th>
</tr>
{{range .ThreatAgents}}  
    <tr>
    <td>{{.Name}}</td>  <td> {{.Description}} </td>
    </tr>
{{end}}
</table>
</p>


<h2>Risk Assessment</h2>

<h3>Risk Matrix</h3>
We follow the <a href="https://owasp.org/www-community/OWASP_Risk_Rating_Methodology">OWASP Risk Rating Methodology</a>.
<table aria-describedby="Risk Matrix">
<tbody><tr>
<th scope="col" colspan="5" align="center">Overall Risk Severity = Impact x Likelihood</th>
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

<h3>Identified Risks</h3>
<p>
The following risks have been identified:
<table aria-describedby="Identified Risks">
<tr>
    <th scope="col">ID</th>
    <th scope="col">Likelihood</th>
    <th scope="col">Impact</th>
    <th scope="col">Severity</th>
    <th scope="col">Risk</th>
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

</body>
</html>
