digraph diagram{
rankdir=LR;
labelloc="c";
//nodesep=0.01;
//ranksep=0.01;
node [shape=box];

{{ define "generateTrustBoundary" }}
subgraph cluster_{{.Id}} {
label="{{.Name}}";
style=dashed;
color=blue;
fontcolor=blue;

{{ range $taName := .ThreatAgentsInside }}
{{ $ta := findThreatAgent $taName}}{{$ta.Name}}[shape=ellipse];
{{ end }}
{{ range $boundaryName := .TrustBoundariesNested }}
{{ $boundary := findTrustedBoundary $boundaryName}}{{ template "generateTrustBoundary" $boundary }}
{{ end }}
{{ range $assetName := .TechnicalAssetsInside }}
{{ $asset := findTechnicalAsset $assetName}}
{{$asset.Id}}[label="{{$asset.Name}}\n{{dataAssetNames $asset.DataAssetsProcessed}}\n{{dataAssetNames $asset.DataAssetsStored}}"];
{{ end }}}
{{ end }}
{{range $boundary := .TrustBoundaries}}
{{if isRootTrustBoundary $boundary.Id}}
{{ template "generateTrustBoundary" $boundary }}
{{end}}{{end}}

{{range $asset := .TechnicalAssets}}{{ $conns := .CommunicationLinks }}
{{range $conn := $conns}} {{$asset.Id}} -> {{$conn.Target}} [fontcolor="red", label="{{$conn.Protocol}}",  headlabel="{{dataAssetNames $conn.DataAssetsSent}}", taillabel="{{dataAssetNames $conn.DataAssetsReceived}}"];
{{end}}{{end}}
}