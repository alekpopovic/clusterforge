# cloud/aws/rds-postgres

## Purpose

Creates a production-aware private AWS RDS PostgreSQL instance for application
workloads. Storage encryption and deletion protection are enabled by default,
and the database is not publicly accessible.

## Status

Implemented.

## Basic Example

```hcl
module "postgres" {
  source = "../../../modules/cloud/aws/rds-postgres"

  name                       = "clusterforge-prod-api"
  environment                = "prod"
  vpc_id                     = module.network.vpc_id
  subnet_ids                 = module.network.private_subnet_ids
  allowed_security_group_ids = [module.api.security_group_id]
  database_name              = "app"
}
```

## EKS App Connection Pattern

Use AWS-managed master password output `secret_arn` with External Secrets
Operator. Sync the password into a Kubernetes Secret, then reference it from the
workload module with `secret_env`.

## ECS App Connection Pattern

Grant the task execution role permission to read the AWS Secrets Manager secret
and pass the secret ARN through the ECS service module secret configuration.

## Production Warnings

- `deletion_protection` defaults to `true`; keep it enabled for production.
- `storage_encrypted` defaults to `true`.
- `manage_master_user_password` defaults to `true`; avoid plaintext passwords.
- RDS has hourly, storage, backup, and data transfer costs.
- Multi-AZ improves availability but increases cost.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
