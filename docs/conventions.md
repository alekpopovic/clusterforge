---
title: Conventions
permalink: /conventions/
---

# ClusterForge Conventions

## Module Structure

Reusable Terraform modules live under `modules/`. Every module must include:

- `main.tf` for resources, data sources, and locals.
- `variables.tf` for typed inputs with descriptions.
- `outputs.tf` for real outputs only.
- `versions.tf` for Terraform/OpenTofu and provider constraints.
- `README.md` with purpose, status, future resources, and usage.

Placeholder modules must remain valid Terraform but must not create fake
resources or fake outputs.

## Naming Conventions

Use lowercase, hyphen-separated names for module inputs and generated resource
names. Prefer `name`, `environment`, and a normalized local name prefix before
introducing module-specific naming inputs.

Use `locals` for shared name construction, common tags, and common labels so
resource blocks stay readable.

## Environment Separation

Real environment compositions live under `live/`. Development, staging, and
production roots must stay separate so state, provider configuration, and
plans are isolated by environment.

Examples live under `examples/` and should remain copy-paste friendly. They are
not substitutes for production `live/` roots.

## Provider Placement

Reusable child modules should not configure providers. Provider blocks belong
in root modules under `live/` or in runnable examples.

Every root configuration must declare its required providers in `versions.tf`.
Use safe provider version constraints and explicit aliases when multiple
clusters, regions, or cloud accounts are involved.

## State Separation

Keep Terraform/OpenTofu state separate per environment and per independently
managed stack. Do not share one state file across unrelated orchestrators,
platform layers, or production and non-production environments.

Remote state configuration belongs in root modules. Reusable modules must not
define backend configuration.

## Secret Handling

Never commit credentials, kubeconfig files, private keys, `tfstate`, or real
secret values. Do not generate real secrets into `tfvars`.

Prefer references to external secret stores such as AWS Secrets Manager,
SSM Parameter Store, Vault, Kubernetes External Secrets, or cloud-native secret
services. Avoid putting sensitive values into Terraform state unless there is
no practical alternative.

## Production Safety Rules

Production changes require explicit review and confirmation. Do not auto-apply
production changes from CI.

Production apply workflows must use an existing reviewed plan file. Production
destroy operations are blocked by default and require an explicit break-glass
workflow before they can be enabled.
