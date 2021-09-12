# Bank of Anthos Example

The following example provides a **Taralizer** model of the [Bank of Anthos](<https://github.com/GoogleCloudPlatform/bank-of-anthos>) example solution for retrail banking.
This solution demonstrates the application of the [Google Cloud Security Foundation Guide](<https://services.google.com/fh/files/misc/google-cloud-security-foundations-guide.pdf>).
We introducted the following issue into the model to demonstrate the rule checking:

- communication between on-premise HSM and GCP KMS uses FTP and is not encrypted

## Getting started

Taralizer creates dataflow diagrams based on your model file. Currently, the rendering engines plantuml and graphviz dot are supported.
The generated dataflow diagram is stored in the default file *diagram.png* that will be also referenced for HTML and PDF reports.
The dataflow diagram helps you to visualize the deloyment of your modeled solution.

- Create a graphviz dataflow diagram and stores it in the *diagram.png* file: `taralizer diagram bank_of_anthos.yaml`:
![GraphViz dot Dataflow Diagram](https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/diagram_graphviz.png)
- Create a plantuml dataflow diagram and stores it in the *diagram.png* file: `taralizer diagram bank_of_anthos.yaml --engine plantuml`
![GraphViz dot Dataflow Diagram](https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/diagram.png)

Once the dataflow diagram is created, a **Security Thread and Risk Analysis** can be conducted using the supported validation rules.
We started the implementation of the [OWASP Application Security Verification Standard](<https://github.com/OWASP/ASVS/raw/v4.0.2/4.0/OWASP%20Application%20Security%20Verification%20Standard%204.0.2-en.pdf>), additional security standards to come.
You can find out more about the currently implemented rules using the command: `taralizer rules`
The following commands will create you reports:

- Create an HTML report referencing *diagram.png* file:  `taralizer report bank_of_anthos.yaml`:
[HTML Report](<https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/report.html>)
- Create a PDF report emedding *diagram.png* file: `taralizer report bank_of_anthos.yaml --type pdf`:
[PDF Report](<https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/report.pdf>)
