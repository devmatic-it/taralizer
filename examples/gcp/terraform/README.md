# GCP Terraform Examples

This directory contains example Terraform configurations for Google Cloud Platform resources that can be used for end-to-end (E2E) integration testing of Taralizer.

## Structure

- `main.tf` - Main Terraform configuration file
- `variables.tf` - Variables definition file

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