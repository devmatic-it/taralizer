# Azure Examples

This directory contains example Terraform configurations for Microsoft Azure that can be used for end-to-end (E2E) integration testing of Taralizer.

## Structure

- `terraform/` - Terraform configuration files for Azure resources

## Usage

These examples can be used to test Taralizer's functionality with real Terraform configurations. They serve as:
1. Integration test cases for the Taralizer tool
2. Examples of how to structure Terraform configurations for use with Taralizer
3. Baseline configurations for testing threat and risk analysis capabilities

## Running Tests

To run E2E tests with these examples, use the following command:

```bash
taralizer analyze ./examples/azure/terraform
```

This will process the Terraform configuration files and generate reports based on the defined threat models.