# Bank of Anthos Example

The following example provides a Taralizer model of the Bank of Anthos <https://github.com/GoogleCloudPlatform/bank-of-anthos>  example solution for retrail banking.
This solution demonstrates the application of the Google Cloud Security Foundation Guide <https://services.google.com/fh/files/misc/google-cloud-security-foundations-guide.pdf>.

## Getting started

Taralizer creates dataflow diagrams based on your model file. Currently, the rendering engines plantuml and graphviz dot are supported.
The generated dataflow diagram is stored in the default file *diagram.png* that will be also referenced for HTML and PDF reports.
The dataflow diagram helps you to visualize the deloyment of your modeled solution.

- Create a graphviz dataflow diagram and stores it in file *diagram.png*: `taralizer diagram bank_of_anthos.yaml`:
![GraphViz dot Dataflow Diagram](https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/diagram_graphviz.png)
- Create a plantuml dataflow diagram and stores it in file *diagram.png*: `taralizer diagram bank_of_anthos.yaml --engine plantuml`
![GraphViz dot Dataflow Diagram](https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/diagram.png)

Once the dataflow diagram is created, a security thread and risk analysis can be conducted using the supported validation rules.
Currently, only the OWASP Application Security Verification Standard <https://github.com/OWASP/ASVS/raw/v4.0.2/4.0/OWASP%20Application%20Security%20Verification%20Standard%204.0.2-en.pdf> is supported, however Taralizer has been designed to support you on the creation on your own security standard.

- Create an HTML report referencing *diagram.png* file:  `taralizer report bank_of_anthos.yaml`:
[HTML Report](https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/report.html)
- Create a PDF report emedding *diagram.png* file: `taralizer report bank_of_anthos.yaml --type pdf`:
[PDF Report](https://github.com/devmatic-it/taralizer/blob/main/examples/gcp/report.pdf)
