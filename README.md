<p align="center">
  <img src="docs/assets/clusterforge-logo.svg" alt="ClusterForge" width="420">
</p>

<p align="center">
  Readable Terraform/OpenTofu platform engineering for Kubernetes, ECS, Nomad, and Docker.
</p>

<p align="center">
  <a href="https://alekpopovic.github.io/clusterforge/">Documentation</a> ·
  <a href="https://github.com/alekpopovic/clusterforge/releases/tag/v0.4.0-rc.1">v0.4.0-rc.1</a> ·
  <a href="docs/architecture.md">Architecture</a> ·
  <a href="docs/security.md">Security</a>
</p>

ClusterForge is a Terraform/OpenTofu framework and CLI for Kubernetes, ECS,
Nomad, and Docker-based container platforms.

ClusterForge standardizes container-platform infrastructure without hiding the
Terraform. The CLI helps generate projects, environments, and workload files,
but the generated infrastructure remains readable and reviewable.

Contributors and AI agents must follow the repository rules in
[`AGENTS.md`](AGENTS.md).

Project documentation lives in [`docs/`](docs/) and is published with the
GitHub Pages workflow.

## CI Status

Status badges will be enabled once workflows are active on GitHub:

- Terraform Validate
- CLI Test
- Security Scan

## Why It Exists

- Standardize infrastructure layout across teams and environments.
- Avoid copy-paste Terraform chaos.
- Keep Terraform/OpenTofu readable and directly reviewable.
- Support multiple container orchestrators through clear adapters.
- Provide a CLI for project generation, repeatable validation, and safer
  workflows.

## Architecture

ClusterForge is split into four layers:

- **Foundation layer**: cloud networking, IAM, DNS, storage, registry, and
  firewall/security groups.
- **Orchestrator layer**: EKS, ECS, Nomad, Docker Engine, Docker Swarm, and
  future Kubernetes variants.
- **Platform layer**: ingress, TLS, external-dns, observability, logging,
  secrets integrations, and GitOps.
- **Workload layer**: web apps, workers, cronjobs, services, scheduled tasks,
  Nomad jobs, Docker containers, and Docker services.

See [docs/architecture.md](docs/architecture.md) for more detail.

## Repository Layout

```text
clusterforge/
  modules/    Reusable Terraform modules
  live/       Real environment compositions
  examples/   Copy-paste friendly examples
  cli/        Go CLI source and templates
  policies/   Conftest and Checkov policy locations
  scripts/    Repeatable local validation scripts
  docs/       Practical project documentation
```

## Quickstart

Install Terraform or OpenTofu first. Install the latest released CLI on Linux or
macOS:

```bash
curl -fsSL https://github.com/alekpopovic/clusterforge/releases/latest/download/install.sh | bash
```

The installer verifies the release SHA256 checksum and writes to
`$HOME/.local/bin` by default. To pin a version or choose a directory:

```bash
curl -fsSL https://github.com/alekpopovic/clusterforge/releases/download/v0.4.0/install.sh \
  | VERSION=v0.4.0 INSTALL_DIR=/usr/local/bin bash
```

For development, build from source:

```bash
cd cli
go build -o cf .
cd ..
```

Initialize a project:

```bash
./cli/cf project init demo
```

Create an AWS EKS environment:

```bash
./cli/cf env create dev --cloud aws --orchestrator eks --region eu-central-1
```

Generate readable Terraform files:

```bash
./cli/cf generate dev
```

Run Terraform/OpenTofu through the CLI:

```bash
./cli/cf init dev
./cli/cf plan dev
```

## Example Workflows

### AWS EKS

```bash
./cli/cf env create dev --cloud aws --orchestrator eks --region eu-central-1
./cli/cf generate dev
./cli/cf init dev
./cli/cf plan dev --out .cf/plans/dev.tfplan --risk-summary
```

The generated EKS root composes tags, AWS networking, and the EKS module. A
commented platform bootstrap block is included for later add-ons.

### AWS ECS

```bash
./cli/cf env create dev-ecs --cloud aws --orchestrator ecs --region eu-central-1
./cli/cf generate dev-ecs
./cli/cf init dev-ecs
./cli/cf plan dev-ecs
```

The generated ECS root composes tags, AWS networking, and an ECS/Fargate
cluster. ECS services can be rendered from app manifests.

### Kubernetes App

```bash
./cli/cf app add api \
  --image ghcr.io/company/api:1.0.0 \
  --port 8080 \
  --replicas 2 \
  --host api.dev.example.com

./cli/cf app render api --env dev
```

For Kubernetes-family targets, the app renderer generates a module call to
`modules/workloads/kubernetes/app`.

### App Manifest Render

Application manifests live in `apps/<name>.yaml`. Rendered Terraform goes to
`env.path/apps/<name>.tf`.

```bash
./cli/cf app list
./cli/cf app render api --env dev
```

For ECS targets, the renderer uses `modules/workloads/ecs/service` and emits
comments where IAM permissions or external secret references must be provided.

## Safety Model

- Production changes are never auto-applied by ClusterForge.
- Production apply requires an existing reviewed plan file.
- Production destroy is blocked by default.
- Plan risk summaries highlight creates, updates, deletes, and replacements.
- Production plans with delete actions require `--allow-destroy`.
- Real secrets must not be stored in `tfvars`.
- Secret inputs should reference external secret stores or existing platform
  secrets.

See [docs/security.md](docs/security.md) for the detailed safety model.
The repository also maintains an explicit [security threat
model](docs/security-threat-model.md), a practical [security
checklist](docs/security-checklist.md), and private vulnerability reporting
guidance in [SECURITY.md](SECURITY.md). These controls reduce risk but are not a
compliance claim; operators must review state, IAM, network exposure, provider
sources, plans, and Kubernetes permissions for their environment.

## Operations And DR

Operational runbooks live in [docs/operations.md](docs/operations.md). Disaster
recovery procedures start with [docs/dr/state-recovery.md](docs/dr/state-recovery.md)
and include EKS, ECS, Kubernetes, cluster restore, app restore, and DNS failover
guidance. Restore procedures must be tested; ClusterForge does not provide
automatic DR or RTO/RPO guarantees.

## Module Development Guide

Reusable modules live under `modules/`. Each module must do one thing and must
include:

- `main.tf`
- `variables.tf`
- `outputs.tf`
- `versions.tf`
- `README.md`

When adding a module:

- Keep provider configuration in root/live environments.
- Use typed variables with descriptions.
- Add validation for important inputs.
- Use clear outputs only for real values.
- Include a README with purpose, usage, inputs, outputs, and notes.
- Run `terraform fmt -recursive` and validation before committing.

See [docs/module-conventions.md](docs/module-conventions.md).

## CLI Development Guide

The CLI is written in Go with Cobra:

```text
cli/
  main.go
  cmd/
  internal/
    app/
    config/
    generator/
    policy/
    terraform/
    ui/
  templates/
```

Run CLI checks locally:

```bash
make test-cli
```

For a reproducible setup, open the repository in the optional
[devcontainer](docs/development-environment.md). It includes the main Go,
Terraform/OpenTofu, Kubernetes, documentation, lint, and security tools.
Contributors who prefer a local toolchain can optionally use the pinned
[asdf versions](docs/tool-versions.md); neither workflow is required by the
Makefile.

Install the optional fast pre-commit checks:

```bash
./scripts/install-hooks.sh
pre-commit run --all-files
```

See [docs/pre-commit.md](docs/pre-commit.md) for tool requirements and the
documented emergency bypass process.

Add new commands under `cli/cmd/`, keep business logic under `cli/internal/`,
and add tests for parsing, generation, policy, and command construction where
practical.

See [docs/cli.md](docs/cli.md).

## Validation

## GitLab CI

GitLab users can compose credential-free CLI/security checks, Terraform plan
jobs, scheduled drift checks, and manual production apply from the templates in
[`ci/gitlab/`](ci/gitlab/). See [docs/gitlab-ci.md](docs/gitlab-ci.md). No
automatic production apply or destroy job is included.

## Validation

Run local checks before opening a pull request:

```bash
make fmt-check
make test
make ci
```

Use OpenTofu instead of Terraform:

```bash
TERRAFORM_BIN=tofu make validate
```

Common developer commands:

```bash
make help
make fmt
make lint
make validate
make build-cli
make security
make clean
```

Additional workflows:

- [Version support](docs/version-support.md)
- [Module release](docs/module-release.md)
- [Registry strategy](docs/registry-strategy.md)
- [Drift detection](docs/drift-detection.md)
- [State helpers](docs/state.md)
- [Upgrade workflow](docs/upgrades.md)
- [Policy packs](docs/policies.md)
- [Template packs](docs/template-packs.md)
- [Cost awareness](docs/cost-awareness.md)
- [Integration testing](docs/testing-integration.md)

## Tutorials

- [Local existing Kubernetes](docs/tutorials/01-local-existing-kubernetes.md)
- [AWS ECS first service](docs/tutorials/02-aws-ecs-first-service.md)
- [AWS EKS platform](docs/tutorials/03-aws-eks-platform.md)
- [GitOps with Argo CD](docs/tutorials/04-gitops-with-argocd.md)
- [Production safety](docs/tutorials/05-production-safety.md)

## Roadmap

- **Phase 1**: AWS EKS
- **Phase 2**: Kubernetes platform
- **Phase 3**: Kubernetes app workloads
- **Phase 4**: ECS
- **Phase 5**: CLI hardening
- **Phase 6**: Nomad
- **Phase 7**: Docker Swarm
- **Phase 8**: AKS, GKE, K3s, and RKE2

See [docs/roadmap.md](docs/roadmap.md).

Planning docs:

- [v0.2 roadmap](ROADMAP_V0.2.md)
- [v0.3 roadmap](ROADMAP_V0.3.md)
- [Backlog](BACKLOG.md)
