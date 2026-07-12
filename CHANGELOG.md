# Changelog

All notable changes to ClusterForge will be documented in this file.

ClusterForge follows semantic versioning while the public interfaces stabilize.

## [0.4.0] - 2026-07-12

### Changed

- Rebuilt the GitHub Pages documentation around the ClusterForge visual identity,
  with responsive navigation and light, dark, and system color modes.
- Corrected the canonical Go module and all internal imports to
  `github.com/alekpopovic/clusterforge/cli` so source builds and release version
  metadata match the repository location.
- Updated release, installer, CLI, and documentation links to the canonical
  `alekpopovic/clusterforge` repository.

### Validation

- The complete CLI unit, command, end-to-end, golden, formatting, vet, and build
  checks pass with the corrected module path.
- Release artifacts remain checksum verified for Linux amd64/arm64, macOS
  amd64/arm64, and Windows amd64.

### Known limitations

- The open risks recorded in the release assessment still apply; this release
  does not claim production-cloud validation for every supported target.

## [0.4.0-rc.1] - 2026-07-12

### Added

- Local plugin MVP, versioned template pack registry, policy engine v2, execution
  profiles, Terraform Cloud metadata, AWS account/region configuration, upgrade
  planners, fleet health and inventory exports.
- Service catalog and Backstage generation, dashboard data export, audit/SIEM
  export, secret rotation planning, backup evidence checks, compliance mapping,
  cross-cluster GitOps rendering, offline bundles and migration analysis.
- Enterprise and operations documentation covering multi-account, FinOps,
  incident/DR, admission security, compliance, air-gapped environments and v0.4
  release scope.

### Changed

- CLI project scaffolding now supports a summary-first wizard and safe
  non-interactive defaults.
- Kubernetes admission examples default to audit/dry-run, and the ECS service
  module exposes an experimental task runtime platform while retaining
  Linux/X86_64 defaults.
- Generated/exported operational data uses explicit versioning and stronger
  redaction/exclusion rules.

### Security

- Plugins remain disabled by default and require explicit trust in CI.
- Offline bundles reject state, kubeconfig, credential-like files, symlinks,
  checksum modifications and unlisted additions.
- Audit, dashboard, migration and secret workflows avoid secret values and cloud
  mutations; compliance mappings make no certification claim.

### Known limitations

- This release candidate has no current production-cloud apply evidence. AWS,
  Azure and GCP workflows still require environment-specific review and testing.
- Cluster federation, edge lifecycle, automatic failover/remediation, hosted SaaS,
  plugin marketplace and production Windows workload support are excluded.
- Static inventory/migration/offline discovery can miss dynamic or transitive
  dependencies and does not prove runtime state.
- A critical GitHub Dependabot finding is visible on the default branch and must
  be triaged/resolved or formally shown outside release scope before v0.4.0.

## [0.1.0] - Unreleased

### Added

- Initial repository structure for the four ClusterForge layers: foundation,
  orchestrator, platform, and workload.
- Core Terraform metadata modules:
  - `modules/core/naming`
  - `modules/core/tags`
  - `modules/core/labels`
- AWS foundation modules for networking, Route53 DNS, Terraform state backend
  resources, IRSA roles, and Karpenter IAM foundations.
- AWS EKS module with managed node groups, optional add-ons, OIDC provider
  creation, and EBS CSI IRSA support.
- AWS ECS/Fargate cluster and service modules, plus Application Load Balancer
  support.
- Kubernetes platform modules for Helm-based add-ons, bootstrap composition,
  External Secrets Operator, Argo CD app-of-apps bootstrap, Karpenter,
  cert-manager issuers, observability, Pod Security labels, and baseline
  NetworkPolicies.
- Kubernetes workload modules for web/service apps, workers, CronJobs, and Helm
  applications.
- Nomad service workload module.
- Docker container and Docker Swarm service workload modules.
- Initial Go CLI, `cf`, with project, environment, generation, backend, app
  manifest, policy, doctor, Terraform runner, JSON output, completion, and
  release-build support.
- App manifest schema and renderer for Kubernetes-family and ECS targets.
- Stacked environment layout support for `network`, `cluster`, `platform`, and
  `apps` stacks.
- Policy/risk summary support for Terraform plan JSON.
- Local developer automation with `Makefile` and scripts for formatting,
  validation, linting, CLI tests, docs, and security scans.
- GitHub Actions workflows for Terraform validation, CLI tests, security scans,
  documentation pages, and CLI release artifacts.
- Practical documentation under `docs/`, including architecture, CLI usage,
  app manifests, environment layouts, backends, security, GitOps, validation,
  observability, autoscaling, and roadmap guidance.

### Changed

- Generated Terraform roots now keep provider configuration visible and avoid
  hiding infrastructure logic behind the CLI.
- Production safety rules are documented and enforced in CLI policy paths:
  production apply requires a plan file, production destroy is blocked by
  default, and destructive plan summaries are surfaced before apply.
- Terraform validation scripts prefer explicit skip reasons over silent
  failures when providers, credentials, or backend initialization are not
  available.

### Security

- Added explicit repository rules against committing credentials, kubeconfigs,
  private keys, `tfstate`, and plan files.
- Added recommended secret workflow using cloud secret managers, External
  Secrets Operator, and Kubernetes Secret references instead of secret values in
  Terraform variables.
- Added Terraform remote state backend bootstrap module with private S3 bucket,
  encryption, versioning, public access blocking, and DynamoDB locking.
- Added security scanning configuration for TFLint, Checkov, and Trivy when
  those tools are installed locally or in CI.

### Known Limitations

- v0.1.0 is an MVP release candidate. It has not been proven by a real
  production cloud apply in this repository.
- AWS EKS and ECS modules need account-specific review before production use,
  especially IAM permissions, cluster add-ons, node sizing, and upgrade
  policies.
- Helm chart versions should be pinned and values should be tuned before
  production rollout.
- AKS, GKE, K3s, RKE2, and generic Kubernetes orchestrator support are planned
  but not production implemented in this release.
- Some Terraform validation paths are limited without initialized providers,
  cloud credentials, or real backend configuration.
- Nomad and Docker targets are intentionally lightweight and should be hardened
  before use as primary production platforms.
