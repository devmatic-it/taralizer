# Examples

This directory contains example Terraform configurations that can be used for end-to-end (E2E) integration testing of Taralizer.

## Structure

- `gcp/` - Google Cloud Platform examples
  - `terraform/` - Terraform configuration files for GCP resources
- `aws/` - Amazon Web Services examples (to be added)
- `azure/` - Microsoft Azure examples (to be added)

## Usage

These examples can be used to test Taralizer's functionality with real Terraform configurations. They serve as:
1. Integration test cases for the Taralizer tool
2. Examples of how to structure Terraform configurations for use with Taralizer
3. Baseline configurations for testing threat and risk analysis capabilities

## Running Tests

To run E2E tests with these examples, use the following command:

```bash
taralizer analyze ./examples/gcp/terraform
```

This will process the Terraform configuration files and generate reports based on the defined threat models.