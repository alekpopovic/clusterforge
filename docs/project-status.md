---
title: Project status
description: Current ClusterForge v0.4 stable release capabilities, maturity, evidence, and limitations.
permalink: /project-status/
---

# Project status

ClusterForge is currently published as `v0.4.3`. This stable version label is
not a claim that every target is production-proven. The project has a broad
implemented module and CLI surface, with maturity varying by provider and
workflow.

## Release and validation snapshot

| Signal | Current state |
|---|---|
| GitHub Release | `v0.4.3`, published as a stable release with checksum-verified installers |
| CLI tests | Go unit, command, e2e, generator golden, and build checks pass |
| Terraform validation | 126 real examples, live roots, and modules validate; 7 stored golden snapshots are covered by Go golden tests |
| Formatting | Terraform recursive format and Go formatting checks pass |
| Secret scanning | Gitleaks history scan reported no known leaks in the v0.4 assessment |
| Static security | Checkov findings remain to be triaged; passing Terraform validation is not a security approval |
| Dependency security | Dependabot alert 2 was resolved by upgrading `aquasecurity/trivy-action` to `v0.36.0` in merge commit `67a6294`; future critical findings remain release blockers |
| Cloud smoke tests | No current production-cloud apply evidence is claimed for v0.4.3 |

See the repository `RELEASE_CANDIDATE_V0.4.md` for the recorded v0.4 assessment,
validation evidence, and open risks.

## CLI capability map

| Workflow | Status | Safety boundary |
|---|---|---|
| Project/env/app scaffolding | Implemented | Summary-first wizard; no apply; non-interactive mode requires explicit inputs/defaults |
| Terraform generation | Implemented | Outputs readable `.tf`; refuses unsafe overwrite paths; providers stay visible in roots |
| Init/plan/apply | Implemented | Production apply requires an existing saved plan; engine remains Terraform/OpenTofu |
| Destroy | Implemented with guardrails | Production destroy blocked by default and destructive operations require confirmation |
| Policy engine | Implemented v2 | New/enhanced admission policies default to advisory/audit; enforcement is opt-in |
| Drift, health, cost | Implemented read-only workflows | May invoke explicitly configured tools/providers; does not auto-remediate |
| Fleet and inventory | Implemented read-only | Stable local/export schemas; no global scheduler or automatic failover |
| Upgrade planning | Implemented read-only | Kubernetes/platform compatibility guidance; does not perform control-plane downgrade |
| Audit/SIEM export | Implemented local export | Redacts sensitive arguments; no direct SIEM API push |
| Backup validation | Implemented evidence checks | Never performs restore/delete automatically; production restore testing discouraged |
| Service catalog/Backstage | Implemented file generation | Metadata only; no credential storage or remote Backstage mutation |
| GitOps multi-cluster | Argo CD rendering implemented | Does not register clusters, commit Git, store repo credentials, or schedule globally |
| Compliance reports | Implemented mappings | Implementation aid only; not certification, attestation, or legal advice |
| Offline bundles | Implemented manifest bundle | Excludes state/kubeconfig/credentials and verifies checksums; does not fetch all artifacts |
| Migration analyzer | Implemented static analysis | Does not read state values, call clouds, modify source, or prove semantic equivalence |
| Plugins | MVP, disabled by default | Local code execution; requires explicit trust and remains outside a public marketplace |

## Terraform module coverage

### Foundation

- AWS networking, DNS, state backend, IAM/IRSA foundations, ECR, KMS, RDS,
  ElastiCache, SQS/SNS, VPC endpoints, and Velero backup storage.
- Initial Azure and GCP network modules.
- Core naming, tags, and labels shared across compositions.

### Orchestrators

- AWS EKS with managed node groups and production hardening inputs.
- AWS ECS/Fargate cluster composition.
- AKS and GKE modules with validated examples but experimental production status.
- Generic/existing Kubernetes attachment.
- Experimental K3s and RKE2 cloud-init generation.
- Nomad, Docker Engine, and Docker Swarm focused patterns.

### Platform

- Kubernetes ingress, certificates, DNS, metrics, Prometheus, Loki/Alloy,
  OpenTelemetry, External Secrets, Argo CD, Argo Rollouts, Karpenter, Velero,
  tenancy, quotas, Pod Security, NetworkPolicy, Kyverno, and Gatekeeper.
- ECS ALB, CloudWatch, and CodeDeploy blue/green building blocks.
- Nomad Consul and ingress building blocks.

### Workloads

- Kubernetes app, worker, CronJob, Helm app, and rollout modules.
- ECS Fargate service and scheduled task modules.
- Nomad service/batch and Docker container/Swarm service modules.

## Enterprise configuration

The config model includes organizations, workspaces, teams, AWS accounts,
regions, execution profiles, Terraform Cloud workspaces, service catalog and
Backstage metadata, policy/template packs, GitOps cluster inventory, health/SLO
settings, and local audit configuration. These are explicit metadata and
composition controls—not identity provider, secret store, billing system, or SaaS
control-plane replacements.

## Experimental and deferred

- AKS/GKE production readiness needs current private networking, identity,
  observability, backup/restore, upgrade, cleanup, and real-cloud evidence.
- K3s/RKE2 edge lifecycle, disconnected update delivery, device identity, and
  local registry management remain RFC/evaluation work.
- Windows ECS runtime metadata is experimental; Kubernetes Windows scheduling and
  node lifecycle are not implemented as supported workflows.
- Cluster federation is limited to inventory, GitOps rendering, fleet health, and
  documented DNS patterns. There is no automatic cross-cluster scheduler.
- Hosted SaaS, automatic remediation, automatic failover, data/secret replication,
  public plugin marketplace, and compliance certification are non-goals for v0.4.

## Production adoption rule

Treat validated configuration as a starting point, not proof of production
fitness. Pin versions, review Terraform plans and IAM/network boundaries, use
external secret stores, run credential-scoped disposable smoke tests, record
cleanup, test backup restoration, and preserve operator-specific evidence before
promoting a target to production.

Continue with [installation]({{ '/installation/' | relative_url }}),
[architecture]({{ '/architecture/' | relative_url }}),
[security]({{ '/security/' | relative_url }}), and
[operations]({{ '/operations.html' | relative_url }}).
