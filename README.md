# ClusterForge

ClusterForge is an opinionated, readable Terraform/OpenTofu framework for
building container platforms. It is designed to keep infrastructure logic
visible while still giving teams a consistent module layout for common
orchestrators and workloads.

The first production target is AWS EKS. The repository structure leaves room
for ECS/Fargate, Nomad, Docker Swarm, AKS, GKE, K3s, RKE2, and generic
Kubernetes without forcing them into one giant module.

## Design Principles

- Keep modules small: one module, one responsibility.
- Keep provider configuration in `live/` roots and examples.
- Keep Terraform/OpenTofu readable; generated files must still be reviewable.
- Use typed variables, validation, clear outputs, and consistent tags/labels.
- Do not put secrets in plain text `*.tfvars`.
- Never auto-apply production changes from CI.

## Layers

ClusterForge is split into four layers:

1. **Foundation**: cloud networking, IAM, DNS, storage, registry, firewalling.
2. **Orchestrator**: EKS, ECS, Nomad, Docker Swarm, and future targets.
3. **Platform**: ingress, TLS, DNS automation, observability, logging, secrets,
   and GitOps.
4. **Workload**: web apps, workers, cronjobs, services, scheduled tasks,
   Nomad jobs, and Docker services.

## Repository Layout

```text
modules/
  core/
  cloud/
  orchestrators/
  platform/
  workloads/
live/
  dev/
  staging/
  prod/
examples/
policies/
scripts/
cli/
```

## Current Status

This initial scaffold includes:

- Core naming, tag, and Kubernetes label helper modules.
- An AWS VPC/network module.
- An AWS EKS orchestrator module.
- A Kubernetes app workload module.
- A `live/dev/aws-eks` root showing provider configuration at the root.
- Basic formatting, validation, and CI workflows.

## Quick Start

Format all Terraform files:

```bash
./scripts/lint.sh
```

Validate every Terraform root and module:

```bash
./scripts/validate.sh
```

Try the development EKS root:

```bash
cd live/dev/aws-eks
terraform init
terraform plan
```

Use `tofu` instead of `terraform` by setting `TERRAFORM_BIN`:

```bash
TERRAFORM_BIN=tofu ./scripts/validate.sh
```
