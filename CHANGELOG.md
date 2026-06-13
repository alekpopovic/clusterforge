# Changelog

All notable changes to ClusterForge will be documented in this file.

ClusterForge follows semantic versioning while the public interfaces stabilize.

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
