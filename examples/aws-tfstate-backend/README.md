# AWS Terraform State Backend

This example uses `modules/cloud/aws/tfstate-backend` to bootstrap the AWS
resources commonly used by Terraform/OpenTofu S3 backends:

- S3 bucket for state
- S3 bucket versioning
- S3 server-side encryption
- S3 public access block
- DynamoDB table for state locking

Run this bootstrap root before configuring an environment to use the S3 backend.
Do not create the backend bucket in the same root that uses it for state.
This bootstrap root intentionally uses the default local backend unless you
decide to migrate it later.

Terraform state may contain sensitive data. Restrict access to the generated
bucket and lock table.

Keep `force_destroy = false` for production so a bucket with objects cannot be
removed accidentally by Terraform.

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

## Migrating an Environment From Local To S3

1. Apply this bootstrap example to create the bucket and lock table.
2. Run `cf backend configure <env> --backend s3 ...`.
3. Regenerate the environment with `cf generate <env> --force`.
4. Run `terraform init -migrate-state` in the generated environment root or
   stack root.
5. Review the migration prompt carefully before confirming.
