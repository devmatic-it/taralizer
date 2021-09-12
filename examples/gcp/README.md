# Bank of Anthos Example

The following example provides a Taralizer model of the Bank of Anthos <https://github.com/GoogleCloudPlatform/bank-of-anthos>  example solution for retrail banking.
This solution demonstrates the application of the Google Cloud Security Foundation Guide <https://services.google.com/fh/files/misc/google-cloud-security-foundations-guide.pdf>.

## Getting started

- Create a graphviz dataflow diagram: `taralizer diagram bank_of_anthos.yaml`
- Create a plantuml dataflow diagram: `taralizer diagram bank_of_anthos.yaml --engine plantuml`
- Create a HTML report with graphviz diagrams: `taralizer report bank_of_anthos.yaml`
- Create an HTML report with plantuml diagram:  `taralizer report bank_of_anthos.yaml --engine plantuml`
- Create a PDF report with graphviz diagrams: `taralizer report bank_of_anthos.yaml --type pdf`

## Example Data Flow Diagrams

### Generated PlantUML Diagram

<img src="https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/diagram.png" width="100%">

### Generated Graphviz Dot Diagram

![GraphViz dot Dataflow Diagram](https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/diagram_graphviz.png)

