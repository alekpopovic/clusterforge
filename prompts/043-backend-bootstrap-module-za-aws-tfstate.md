## Prompt 43 — Backend bootstrap module za AWS tfstate

```text
Create AWS Terraform state backend bootstrap module and example.

Create module:
- modules/cloud/aws/tfstate-backend

Purpose:
Create S3 bucket and DynamoDB table for Terraform remote state locking.

Inputs:
- name: string
- environment: string
- bucket_name: string
- dynamodb_table_name: string
- force_destroy: bool, default false
- enable_versioning: bool, default true
- enable_encryption: bool, default true
- tags: map(string), default {}

Resources:
- aws_s3_bucket
- aws_s3_bucket_versioning
- aws_s3_bucket_server_side_encryption_configuration
- aws_s3_bucket_public_access_block
- aws_dynamodb_table with LockID hash key
- Optional lifecycle protection guidance in docs

Outputs:
- bucket_name
- bucket_arn
- dynamodb_table_name
- dynamodb_table_arn
- backend_config_example

Create example:
- examples/aws-tfstate-backend

README:
- Explain that backend resources must be created before configuring backend.
- Explain migration from local to S3 backend.
- Warn that tfstate may contain sensitive data.
- Warn about force_destroy=false for production.

Rules:
- Do not configure the backend inside this module.
- Do not hardcode account IDs.
- Do not use public bucket settings.
- Encryption and versioning should be enabled by default.

Run:
- terraform fmt -recursive
- validation where possible
```

---
