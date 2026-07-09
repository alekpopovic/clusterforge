# cloud/aws/tfstate-backend

## Purpose

Creates the AWS resources needed by Terraform/OpenTofu S3 remote state:

- S3 bucket for state storage
- Optional S3 bucket versioning
- Optional S3 default server-side encryption
- S3 public access block
- DynamoDB table with `LockID` hash key for state locking

## Status

Implemented.

## Usage

```hcl
module "tfstate_backend" {
  source = "../../../modules/cloud/aws/tfstate-backend"

  name                = "clusterforge-tfstate"
  environment         = "bootstrap"
  bucket_name         = "example-clusterforge-terraform-state"
  dynamodb_table_name = "example-clusterforge-terraform-locks"
  kms_key_arn         = module.state_key.key_arn

  tags = {
    Project = "clusterforge"
  }
}
```

When `kms_key_arn` is empty, the bucket uses S3-managed `AES256` encryption.
Set `kms_key_arn` to use a customer-managed KMS key.

## Notes

Do not configure the backend inside this module. Apply this module from a
separate bootstrap root first, then point ClusterForge environments at the
created bucket and lock table.

State can contain sensitive data. Restrict access to the bucket, enable
versioning and encryption, and keep `force_destroy = false` for production.

Terraform does not have a native lifecycle setting that prevents all accidental
deletes from outside Terraform. For production, combine this module with IAM
guardrails, bucket policies, AWS Backup or retention policies where appropriate,
and reviewed change workflows.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
