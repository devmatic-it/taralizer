# Bank of Anthos Example

The following example provides a **Taralizer** model of the [Bank of Anthos](<https://github.com/GoogleCloudPlatform/bank-of-anthos>) example solution for retrail banking.
This solution demonstrates the application of the [Google Cloud Security Foundation Guide](<https://services.google.com/fh/files/misc/google-cloud-security-foundations-guide.pdf>).
We introducted the following issue into the model to demonstrate the rule checking:

- communication between on-premise HSM and GCP KMS uses FTP and is not encrypted

## Getting started

Taralizer generates self-contained reports with embedded Data Flow Diagrams (DFD) rendered via mermaid.js.
No external tools (graphviz, plantuml, mermaid-cli) are required.

The following commands will create you reports:

- Create an HTML report with inline DFD: `taralizer report bank_of_anthos.yaml`:
[HTML Report](<https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/report.html>)
- Create a PDF report with inline DFD: `taralizer report bank_of_anthos.yaml --type pdf`:
[PDF Report](<https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/report.pdf>)
- Create a Markdown report with embedded DFD (renders in GitHub/GitLab/VS Code): `taralizer report bank_of_anthos.yaml --type markdown`:
[Markdown Report](<https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/report.md>)
