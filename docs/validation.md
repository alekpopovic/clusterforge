---
title: Validation
permalink: /validation/
---

# Validation

ClusterForge validation is designed to catch formatting, Terraform syntax,
provider schema, and lightweight module test issues without requiring real
cloud credentials or applying infrastructure.

## What Gets Validated

Run the default validation workflow from the repository root:

```bash
./scripts/validate.sh
```

The script performs these checks:

- `terraform fmt -check -recursive`
- Terraform module contract checks for every module under `modules/`
- Terraform root discovery with `scripts/list-terraform-roots.sh`
- `terraform init -backend=false` for roots that are safe to initialize locally
- `terraform validate` for safe roots
- `terraform test` for modules with `modules/**/tests/*.tftest.hcl`

Terraform roots are directories that contain at least one `.tf` file, excluding
`.terraform`, `.git`, `.cf`, temporary, build, and distribution directories.

## Why Some Roots Are Skipped

Validation should be explicit about skipped directories. A root is skipped when
it cannot be safely initialized or validated without external access, such as:

- missing `versions.tf`
- Terraform Cloud remote operations
- a `backend "remote"` configuration

ClusterForge uses `terraform init -backend=false` for normal roots so local
validation does not initialize real remote state backends. Provider plugins may
still need to be downloaded from the Terraform registry.

## What Validation Does Not Prove

Local validation does not:

- authenticate to AWS, Kubernetes, Nomad, or Docker
- run `terraform plan`
- run `terraform apply`
- verify cloud quotas, IAM permissions, or live endpoint availability
- prove Helm chart installability against a real cluster

It is a fast quality gate, not a substitute for reviewed plans in real
environments.

## Validate With Real Credentials

To validate and plan against a real environment:

1. Configure credentials outside the repository.
2. Copy the safe example variables to a local ignored `terraform.tfvars` file.
3. Initialize the target root.
4. Run plan and save the plan file for review.

Example:

```bash
cd live/dev/aws-eks
cp terraform.tfvars.example terraform.tfvars
terraform init
terraform validate
terraform plan -out .cf-plan.tfplan
```

Never commit `terraform.tfvars`, plan files, state files, kubeconfigs, private
keys, or credentials.

## OpenTofu

Use OpenTofu by setting `TERRAFORM_BIN`:

```bash
TERRAFORM_BIN=tofu ./scripts/validate.sh
```
