<p align="center">
  <a href="https://alekpopovic.github.io/clusterforge/">
    <img src="docs/assets/clusterforge-logo.svg" alt="ClusterForge" width="430">
  </a>
</p>

<p align="center">
  <strong>Forge consistent container platforms without hiding Terraform.</strong><br>
  Readable Terraform/OpenTofu modules and a safety-focused Go CLI for Kubernetes,
  ECS, Nomad, and Docker.
</p>

<p align="center">
  <a href="https://github.com/alekpopovic/clusterforge/releases/tag/v0.4.3"><img alt="Release v0.4.3" src="https://img.shields.io/badge/release-v0.4.3-f97316?style=for-the-badge&logo=github"></a>
  <a href="https://github.com/alekpopovic/clusterforge/actions/workflows/cli-test.yml"><img alt="CLI tests" src="https://img.shields.io/github/actions/workflow/status/alekpopovic/clusterforge/cli-test.yml?branch=main&style=for-the-badge&label=CLI%20tests&logo=go&logoColor=white"></a>
  <a href="https://github.com/alekpopovic/clusterforge/actions/workflows/terraform-validate.yml"><img alt="Terraform validation" src="https://img.shields.io/github/actions/workflow/status/alekpopovic/clusterforge/terraform-validate.yml?branch=main&style=for-the-badge&label=Terraform&logo=terraform&logoColor=white"></a>
  <a href="https://github.com/alekpopovic/clusterforge/actions/workflows/security-scan.yml"><img alt="Security scan" src="https://img.shields.io/github/actions/workflow/status/alekpopovic/clusterforge/security-scan.yml?branch=main&style=for-the-badge&label=Security&logo=securityscorecard&logoColor=white"></a>
</p>

<p align="center">
  <a href="https://github.com/alekpopovic/clusterforge/actions/workflows/govulncheck.yml"><img alt="Go vulnerability check" src="https://img.shields.io/github/actions/workflow/status/alekpopovic/clusterforge/govulncheck.yml?branch=main&style=flat-square&label=govulncheck&logo=go&logoColor=white"></a>
  <a href="https://github.com/alekpopovic/clusterforge/actions/workflows/module-conformance.yml"><img alt="Module conformance" src="https://img.shields.io/github/actions/workflow/status/alekpopovic/clusterforge/module-conformance.yml?branch=main&style=flat-square&label=module%20conformance&logo=terraform&logoColor=white"></a>
  <a href="https://github.com/alekpopovic/clusterforge/actions/workflows/pages.yml"><img alt="Documentation" src="https://img.shields.io/github/actions/workflow/status/alekpopovic/clusterforge/pages.yml?branch=main&style=flat-square&label=Docs&logo=githubpages&logoColor=white"></a>
  <a href="https://github.com/alekpopovic/clusterforge/wiki"><img alt="ClusterForge Wiki" src="https://img.shields.io/badge/Wiki-Explore-8b5cf6?style=flat-square&logo=github&logoColor=white"></a>
</p>

<p align="center">
  <a href="https://alekpopovic.github.io/clusterforge/"><strong>📚 Documentation</strong></a>
  &nbsp;•&nbsp;
  <a href="https://github.com/alekpopovic/clusterforge/wiki"><strong>🌐 Wiki</strong></a>
  &nbsp;•&nbsp;
  <a href="https://alekpopovic.github.io/clusterforge/installation/"><strong>⚡ Install</strong></a>
  &nbsp;•&nbsp;
  <a href="docs/architecture.md"><strong>🏗️ Architecture</strong></a>
  &nbsp;•&nbsp;
  <a href="docs/security.md"><strong>🛡️ Security</strong></a>
  &nbsp;•&nbsp;
  <a href="https://github.com/alekpopovic/clusterforge/releases"><strong>📦 Releases</strong></a>
</p>

---

## 🔥 What is ClusterForge?

ClusterForge is an open Terraform/OpenTofu framework for building and operating
container platforms. Its CLI generates projects, environments, workload files,
policies, reports, and operational evidence while keeping the resulting `.tf`
files visible, editable, and reviewable.

```text
                       ┌──────────────────────────┐
                       │       cf CLI             │
                       │ generate · plan · policy │
                       └────────────┬─────────────┘
                                    │ readable .tf
       ┌────────────────┬───────────┴───────┬────────────────┐
       │ Foundation     │ Orchestrators     │ Platform       │ Workloads
       │ network · IAM  │ EKS · ECS · AKS   │ GitOps · TLS   │ apps · workers
       │ DNS · storage  │ GKE · Nomad       │ secrets · obs. │ jobs · services
       └────────────────┴───────────────────┴────────────────┘
```

ClusterForge wraps Terraform workflows; it does not replace Terraform, hide
provider configuration, or operate a hosted control plane.

## ✨ Highlights

| | Capability | What it provides |
|---|---|---|
| 🧱 | **Readable modules** | Focused foundation, orchestrator, platform, and workload modules |
| ⚡ | **Go CLI** | Project wizard, generation, init/plan/apply, drift, state, output, and graph workflows |
| 🛡️ | **Production guardrails** | Saved-plan requirement, destroy blocking, confirmations, and risk summaries |
| ☸️ | **Multi-platform** | Kubernetes, AWS ECS/Fargate, Nomad, Docker Engine, and Docker Swarm patterns |
| 🧭 | **Fleet operations** | Inventory, health, drift, upgrades, cost heuristics, audit, and dashboard exports |
| 📋 | **Governance** | Policy engine, compliance mappings, admission packs, and ownership metadata |
| 🚚 | **Delivery workflows** | App manifests, Argo CD rendering, Backstage catalog, template packs, and offline bundles |
| 🔐 | **Security boundaries** | External-secret references, redacted exports, checksum verification, and no automatic remediation |

## 🌈 Platform coverage

| Target | Current scope | Maturity |
|---|---|---|
| **AWS EKS** | Network, cluster, node groups, IAM/IRSA, add-ons, platform and workloads | 🟢 Primary |
| **AWS ECS/Fargate** | Cluster, services, scheduled tasks, ALB, CloudWatch and blue/green building blocks | 🟢 Primary |
| **Existing Kubernetes** | Attach to a reviewed context; render platform and workload modules | 🟢 Implemented |
| **Azure AKS / Google GKE** | Validated modules and examples | 🟡 Experimental production status |
| **K3s / RKE2** | Cloud-init focused generation | 🟡 Experimental |
| **Nomad** | Cluster, service/batch workload, Consul and ingress patterns | 🟡 Focused support |
| **Docker / Swarm** | Engine, container, and Swarm service patterns | 🟡 Focused support |

> [!IMPORTANT]
> A validated module is a starting point, not proof of production fitness.
> Cloud targets still require account-specific IAM, network, backup, upgrade,
> observability, plan, apply, and cleanup evidence.

## 🚀 Install the CLI

### Linux and macOS

The installer detects OS/architecture and verifies the binary SHA256 checksum:

```bash
curl -fsSL https://github.com/alekpopovic/clusterforge/releases/latest/download/install.sh | bash
```

Pin automation to the current release:

```bash
curl -fsSL https://github.com/alekpopovic/clusterforge/releases/download/v0.4.3/install.sh \
  | VERSION=v0.4.3 INSTALL_DIR="$HOME/.local/bin" bash
```

Release artifacts support Linux amd64/arm64, macOS Intel/Apple Silicon, and
Windows amd64. See the complete [installation guide](docs/installation.md) for
PowerShell, manual checksum verification, upgrades, and uninstall instructions.

## ⚡ Credential-free quickstart

Create a local project that targets an existing Kubernetes context (including
kind), generate its Terraform root, initialize providers, and save a plan:

```bash
cf --non-interactive wizard --defaults
cf generate dev
cf init dev
cf plan dev --out .cf/plans/dev.tfplan --risk-summary
```

Create a kind cluster first only when one does not already exist:

```bash
cf local create kind
```

No default command automatically applies infrastructure. Review generated files
and the saved plan before running `cf apply`.

### AWS EKS example

```bash
cf project init demo
cf env create dev --cloud aws --orchestrator eks --region eu-central-1
cf generate dev
cf init dev
cf plan dev --out .cf/plans/dev.tfplan --risk-summary
```

Standalone release binaries generate module sources pinned to the matching
ClusterForge Git tag. Repository checkouts use local module paths for development.

## 🛡️ Safety by design

- Production apply requires an existing reviewed plan file.
- Production destroy is blocked by default.
- Destructive production plans require explicit additional approval.
- Generated infrastructure remains plain Terraform/OpenTofu.
- Plugins are local, disabled by default, and require explicit trust in CI.
- Audit, dashboard, inventory, migration, and secret reports redact or exclude
  sensitive values by design.
- ClusterForge never asks for secret values during project scaffolding.

Read the [security model](docs/security.md), [threat model](docs/security-threat-model.md),
and [production safety tutorial](docs/tutorials/05-production-safety.md).

## 🗺️ Repository map

```text
clusterforge/
├── cli/          Go CLI, Cobra commands, internal packages, and embedded templates
├── modules/      Reusable Terraform/OpenTofu modules
├── live/         Reviewable environment compositions
├── examples/     Copy-paste-friendly validated examples
├── policies/     Built-in and extended policy packs
├── scripts/      Repeatable validation, security, release, and installer scripts
├── docs/         GitHub Pages documentation, tutorials, RFCs, and runbooks
└── ci/           Reusable CI examples, including GitLab templates
```

## 📊 Current project status

| Signal | Current evidence |
|---|---|
| Stable release | **v0.4.3**, checksum-verified GitHub artifacts |
| CLI | Unit, command, e2e, golden, build, vet, race, and release-binary smoke tests pass |
| Terraform | **126** real module/example/live roots validated |
| Native module tests | **15** assertions pass across core labels, naming, and tags |
| CLI surface | **164** command/group help paths traversed in the latest audit |
| Formatting | Recursive Terraform formatting and Go formatting checks pass |
| Production cloud evidence | Not claimed for every target; see [project status](docs/project-status.md) |

Security scan findings and dependency alerts are tracked separately from
functional validation. Passing tests are not a security certification or a
promise that every cloud composition is production-ready.

## 🧰 Development

```bash
git clone https://github.com/alekpopovic/clusterforge.git
cd clusterforge
make test-cli
make fmt-check
make validate
```

Useful commands:

```bash
make help          # discover local automation
make build-cli     # build cli/cf with version metadata
make test          # CLI tests plus Terraform validation
make lint          # configured lint gates
make security      # available local security scanners
```

The optional [devcontainer](docs/development-environment.md) provides the main
Go, Terraform/OpenTofu, Kubernetes, lint, documentation, and security tools.
Contributor and AI-agent rules live in [AGENTS.md](AGENTS.md).

## 📚 Explore

| Learn | Build | Operate |
|---|---|---|
| [Architecture](docs/architecture.md) | [Module catalog](docs/modules.md) | [Operations](docs/operations.md) |
| [CLI guide](docs/cli.md) | [App manifests](docs/app-manifest.md) | [Drift detection](docs/drift-detection.md) |
| [Environment layouts](docs/environments.md) | [Policy packs](docs/policies.md) | [Backup and DR](docs/dr/) |
| [Security](docs/security.md) | [GitOps](docs/gitops.md) | [Upgrade workflow](docs/upgrades.md) |
| [Tutorials](docs/tutorials/) | [Examples](examples/) | [Runbooks](docs/runbooks.md) |

## 🧭 Direction

The v0.4 line delivers the broad CLI, module, policy, enterprise-metadata, and
operational-documentation surface. The v0.5 work is focused on hardening:
real-cloud evidence, security finding triage, explicit module maturity,
provenance-aware releases, stable CLI contracts, and deeper provider
compatibility testing. Credential-gated checks remain explicit and their absence
limits support claims rather than being reported as a pass.

See the [roadmap](docs/roadmap.md), [v0.5 backlog](BACKLOG_V0.5.md),
[v0.4 assessment](RELEASE_CANDIDATE_V0.4.md), and [changelog](CHANGELOG.md).

---

<p align="center">
  Built for platform teams that want <strong>consistency without opacity</strong>.
</p>
