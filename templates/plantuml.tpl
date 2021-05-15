@startuml Data Flow Diagram
'All Kubernetes symbols
!define KubernetesPuml https://raw.githubusercontent.com/dcasati/kubernetes-PlantUML/master/dist
!includeurl KubernetesPuml/kubernetes_Common.puml
!includeurl KubernetesPuml/kubernetes_Context.puml
!includeurl KubernetesPuml/kubernetes_Simplified.puml
!includeurl KubernetesPuml/OSS/all.puml

'All GCP symbols
!define GCPPuml https://raw.githubusercontent.com/davidholsgrove/gcp-icons-for-plantuml/master/dist
!includeurl GCPPuml/GCPCommon.puml
!includeurl GCPPuml/GCPSimplified.puml
!includeurl GCPPuml/DeveloperTools/all.puml
!includeurl GCPPuml/Storage/all.puml
!includeurl GCPPuml/Security/all.puml
!includeurl GCPPuml/Networking/all.puml
!includeurl GCPPuml/ManagementTools/all.puml
!includeurl GCPPuml/Databases/all.puml
!includeurl GCPPuml/Compute/all.puml

'All AWS symbols
!include <awslib/AWSCommon>
!include <awslib/AWSSimplified.puml>
!include <awslib/Compute/all.puml>
!include <awslib/mobile/all.puml>
!include <awslib/general/all.puml>

'All Azure symbols
!define AzurePuml https://raw.githubusercontent.com/plantuml-stdlib/Azure-PlantUML/release/2-1/dist
!includeurl AzurePuml/AzureCommon.puml
!includeurl AzurePuml/AzureSimplified.puml
!includeurl AzurePuml/Databases/all.puml
!includeurl AzurePuml/Networking/all.puml
!includeurl AzurePuml/Compute/all.puml

left to right direction

{{ define "generateTrustBoundary" }}
{{.Puml}} {
{{ range $taName := .ThreatAgentsInside }}
{{ $ta := findThreatAgent $taName}}actor {{$ta.Name}}
{{ end }}
{{ range $boundaryName := .TrustBoundariesNested }}
{{ $boundary := findTrustedBoundary $boundaryName}}{{ template "generateTrustBoundary" $boundary }}
{{ end }}
{{ range $assetName := .TechnicalAssetsInside }}
{{ $asset := findTechnicalAsset $assetName}}{{$asset.Puml}}
{{ end }}

}
{{ end }}
{{range $boundary := .TrustBoundaries}}
{{if isRootTrustBoundary $boundary.Id}}
    {{ template "generateTrustBoundary" $boundary }}
{{end}}
{{end}}
{{range $asset := .TechnicalAssets}}{{ $conns := .CommunicationLinks }}
{{range $conn := $conns}} {{$asset.Id}} --> {{$conn.Target}} : {{$conn.Protocol}} 
{{end}}{{end}}
@enduml
