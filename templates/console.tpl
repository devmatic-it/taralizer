TARALIZER

(1) Threat Agents
The following threat agents are used:
{{range .ThreatAgents}}  {{.Id}} : {{.Description}}
{{end}}

(2) Data Assets
The following data assets are used:
{{range .DataAssets}}  {{.Id}} : {{.Description}}
{{end}}

(3) Technical Assets
The following technical assets have been defined:
{{range .TechnicalAssets}}  {{.Id}} : {{.Description}}
{{end}}

(4) Risks
The following risks have been identified:
{{range .Risks}}  {{.Severity} {{.Id}} : {{.Message}}{{end}}
