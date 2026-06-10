# live/dev/aws-eks

Development AWS EKS environment root.

This root composes:

- `modules/core/tags`
- `modules/cloud/aws/network`
- `modules/orchestrators/kubernetes/eks`

Provider configuration lives here in the root module. The reusable modules do
not configure providers.

## Setup

Create a local variable file from the safe example values:

```bash
cp terraform.tfvars.example terraform.tfvars
```

Review and edit `terraform.tfvars` for your AWS account, region, CIDR ranges,
availability zones, and access CIDRs. Do not put credentials or secrets in
`terraform.tfvars`.

## Commands

```bash
terraform init
terraform plan
terraform apply
```

For a real shared development environment, configure the S3 backend in
`backend.tf` before running `terraform init`.

## Production Warning

This is a development environment root. Do not apply production changes
directly or without review. Production workflows should use reviewed plan files
and explicit approval before apply.
