# ClusterForge

ClusterForge is an early-development Terraform/OpenTofu framework for
container orchestrators. It is intended to provide an opinionated but readable
repository structure for composing infrastructure, platform services, and
application workloads without hiding Terraform behind a black box.

The CLI will be a wrapper and generator for repeatable workflows. It must
generate readable Terraform/OpenTofu files and must not replace direct
Terraform review, planning, or validation.

Contributors and AI agents must follow the repository rules in
[`AGENTS.md`](AGENTS.md).

## Architecture Layers

ClusterForge is organized into four layers:

- **Foundation**: networking, IAM, DNS, storage, registries, and firewalls.
- **Orchestrator**: Kubernetes, ECS/Fargate, Nomad, Docker Engine, and Docker
  Swarm.
- **Platform**: ingress, TLS, DNS automation, observability, logging, secrets,
  and GitOps.
- **Workload**: apps, services, workers, cronjobs, scheduled tasks, Nomad jobs,
  Docker containers, and Docker services.

## Supported Orchestrators

The repository skeleton includes module locations for:

- Kubernetes: EKS first, plus generic Kubernetes for future targets.
- AWS ECS/Fargate.
- Nomad.
- Docker Engine and Docker Swarm.

Future module families may include AKS, GKE, K3s, and RKE2.

## Planned CLI Workflow

The planned Go CLI will help teams:

- Generate new live environment roots from templates.
- Generate workload module usage files.
- Run formatting, validation, policy checks, and tests.
- Require explicit confirmation for destructive operations.
- Require an existing plan file before production apply.

The CLI source lives under `cli/`; generated infrastructure must remain
readable Terraform/OpenTofu.

## Repository Layout

```text
modules/    Reusable Terraform modules
live/       Real environment compositions
examples/   Copy-paste friendly examples
cli/        Go CLI source and templates
policies/   Conftest and Checkov policy locations
scripts/    Repeatable local validation scripts
docs/       Project conventions and design notes
```

## Current Status

ClusterForge is in early development. The current repository focuses on clean
structure and valid placeholders. Modules intentionally create no real cloud or
orchestrator resources yet.

See [docs/conventions.md](docs/conventions.md) for module, provider, state,
secret, and production safety conventions.

## Validation

Format Terraform files:

```bash
./scripts/lint.sh
```

Validate Terraform roots and modules:

```bash
./scripts/validate.sh
```

Use OpenTofu instead of Terraform:

```bash
TERRAFORM_BIN=tofu ./scripts/validate.sh
```
