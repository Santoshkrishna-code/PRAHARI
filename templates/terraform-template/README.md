# Terraform Infrastructure Module Template

This template establishes the design structure and security compliance rules for all Terraform resource modules (e.g. VPC, EKS, RDS, Cognito) in the **PRAHARI** platform.

## Quality & Compliance Standards

1. **Explicit Providers**: Declare required providers and version ranges in `main.tf`.
2. **Mandatory Encryption**: Resources that write data at rest (S3, RDS, OpenSearch, DynamoDB) MUST configure KMS Customer Managed Key encryption.
3. **Zero Trust Access Control**: Enforce private ingress controls (e.g. VPC private subnets, blocking public S3 buckets).
4. **Standard Tagging**: Apply the tag variables dynamically on all components.

## How to Initialize a New Module

1. Copy the contents of this folder into `infrastructure/terraform/modules/<new-module-name>`.
2. Modify `main.tf` to declare the specific resources needed.
3. Define safe input controls in `variables.tf` and outputs in `outputs.tf`.
4. Validate format:
   ```bash
   terraform fmt
   ```
5. Check security compliance using tfsec / checkov:
   ```bash
   tfsec .
   ```
