# AWS Terraform State Backend

This example bootstraps the AWS resources commonly used by Terraform/OpenTofu
S3 backends:

- S3 bucket for state
- S3 bucket versioning
- S3 server-side encryption
- S3 public access block
- DynamoDB table for state locking

Run this bootstrap root before configuring an environment to use the S3 backend.
Do not create the backend bucket in the same root that uses it for state.

## Usage

```bash
cp terraform.tfvars.example terraform.tfvars
terraform init
terraform plan
terraform apply
```

Then configure ClusterForge:

```bash
cf backend configure prod \
  --backend s3 \
  --bucket example-clusterforge-terraform-state \
  --region eu-central-1 \
  --dynamodb-table example-clusterforge-terraform-locks \
  --key-prefix clusterforge/prod
```

The example values use `example-` names. Replace them with globally unique,
non-secret names before applying.
