---
title: Backends
permalink: /backends/
---

# Backends

Terraform/OpenTofu backends decide where state is stored. State can contain
sensitive infrastructure data, generated passwords, secret metadata, resource
IDs, and provider-specific attributes, so treat it as protected data.

## Supported Backend Templates

ClusterForge can generate backend templates for:

- `local`
- `s3`
- `azurerm` placeholder
- `gcs` placeholder

The first real remote backend implementation is S3.

## Configuration

Backends are configured in `clusterforge.yaml` by environment:

```yaml
backends:
  dev:
    type: local
  prod:
    type: s3
    bucket: my-terraform-state-bucket
    region: eu-central-1
    dynamodb_table: my-terraform-locks
    key_prefix: clusterforge/prod
```

Use the CLI to configure or inspect a backend:

```bash
cf backend configure dev --backend local
cf backend configure prod \
  --backend s3 \
  --bucket my-terraform-state-bucket \
  --region eu-central-1 \
  --dynamodb-table my-terraform-locks \
  --key-prefix clusterforge/prod
cf backend show prod
```

`cf generate <env>` uses the configured backend when writing `backend.tf`.

## S3 Backend Bootstrap Flow

Do not create the backend bucket in the same Terraform root that uses it.
Bootstrap it separately:

```bash
cd examples/aws-tfstate-backend
cp terraform.tfvars.example terraform.tfvars
terraform init
terraform plan
terraform apply
```

Then configure ClusterForge to use the created bucket and lock table:

```bash
cf backend configure prod \
  --backend s3 \
  --bucket example-clusterforge-terraform-state \
  --region eu-central-1 \
  --dynamodb-table example-clusterforge-terraform-locks \
  --key-prefix clusterforge/prod
cf generate prod --force
```

## Stacked Environments

For stacked environments, ClusterForge generates one backend file per stack and
uses separate state keys:

```text
clusterforge/prod/network/terraform.tfstate
clusterforge/prod/cluster/terraform.tfstate
clusterforge/prod/platform/terraform.tfstate
clusterforge/prod/apps/terraform.tfstate
```

Downstream stacks should consume upstream outputs through
`terraform_remote_state` or another explicit handoff.

## Safety Notes

- Remote state matters because it centralizes locking, history, and team access.
- State can contain sensitive data, even if `tfvars` files do not.
- Do not commit `terraform.tfstate`, plan files, credentials, or kubeconfigs.
- Production environments should use a remote backend with locking.
- ClusterForge warns when `prod` or `production` uses the local backend.
