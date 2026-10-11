# taralizer

[![Go Report Card](https://goreportcard.com/badge/github.com/devmatic-it/taralizer)](https://goreportcard.com/report/github.com/devmatic-it/taralizer)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Taralizer is a Threat and Risk Analysis tool for identifying threats, risks, and vulnerabilities in software architecture.

## Features

- Threat modeling using [OWASP ASVS](https://owasp.org/www-project-application-security-verification-standard/) compliance standard
- YAML-based configuration and threat modeling
- Terraform integration for automatic asset discovery
- Multi-format report generation (YAML, JSON)

## Installation

```bash
go install github.com/devmatic-it/taralizer@latest
```

Or build from source:

```bash
git clone https://github.com/devmatic-it/taralizer.git
cd taralizer
go build -o taralizer
```

## Usage

### Basic Commands

```bash
# Show help
taralizer --help

# Show version
taralizer version

# Validate a model file
taralizer validate model.yaml

# Generate a report from a model
taralizer report model.yaml

# Show available rules
taralizer rules
```

### Terraform Integration

The terraform-to-yaml command can generate a taralizer YAML template by parsing Google Cloud Terraform files and mapping the resources to appropriate taralizer assets.

```bash
# Generate a taralizer YAML template from Terraform files
taralizer terraform-to-yaml --tf-dir <path_to_terraform_directory> --output <output_file>

# Example
taralizer terraform-to-yaml --tf-dir examples/gcp/terraform --output my-model.yaml
```

### Example

An example model is available in `examples/gcp/terraform/main.tf` and can be used to generate a taralizer template:

```bash
taralizer terraform-to-yaml --tf-dir examples/gcp/terraform --output /tmp/test_output.yaml
```

## Contributing

1. Fork it (https://github.com/devmatic-it/taralizer/fork)
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Create a new Pull Request

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.