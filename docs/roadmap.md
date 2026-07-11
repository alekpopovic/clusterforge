---
title: Roadmap
permalink: /roadmap/
---

# Roadmap

See also:

- [v0.2 roadmap](../ROADMAP_V0.2.md)
- [v0.3 roadmap](../ROADMAP_V0.3.md)
- [v0.4 roadmap](../ROADMAP_V0.4.md)
- [v0.4 release plan](../RELEASE_PLAN_V0.4.md)
- [v0.4 backlog](../BACKLOG_V0.4.md)
- [backlog](../BACKLOG.md)
- [AKS RFC](rfcs/002-aks-support.md)
- [GKE RFC](rfcs/003-gke-support.md)
- [AKS production hardening](azure-aks-production.md)
- [GKE production hardening](gcp-gke-production.md)
- [self-hosted Kubernetes RFC](rfcs/004-self-hosted-kubernetes.md)
- [plugin architecture RFC](rfcs/006-cli-plugin-architecture.md)
- [workload identity RFC](rfcs/007-workload-identity.md)
- [service mesh RFC](rfcs/008-service-mesh.md)
- [cluster federation RFC](rfcs/013-cluster-federation.md)
- [edge deployments RFC](rfcs/014-edge-deployments.md)

ClusterForge is being built in practical phases. The order favors a useful AWS
Kubernetes path first, then broader orchestrator coverage.

## Phase 1: AWS EKS

- AWS network module
- EKS module with managed node groups
- Dev/staging/prod live examples
- Basic validation and CI

## Phase 2: Kubernetes Platform

- ingress-nginx
- cert-manager
- external-dns
- metrics-server
- prometheus-stack
- Loki
- Argo CD

## Phase 3: Kubernetes App Workloads

- Generic app module
- CronJob module
- Helm app module
- App manifest rendering

## Phase 4: ECS

- ECS cluster module
- Fargate service module
- Scheduled task module
- ECS environment generation

## Phase 5: CLI Hardening

- Project and environment generation
- App manifest workflows
- Risk summaries
- Policy checks
- Better packaged template handling

## Phase 6: Nomad

- Nomad cluster patterns
- Service and batch jobs
- Consul and ingress integration

## Phase 7: Docker Swarm

- Docker Engine support
- Swarm service support
- Simple self-hosted examples

## Phase 8: AKS, GKE, K3s, RKE2

- Azure AKS
- Google GKE
- Lightweight Kubernetes targets
- Generic Kubernetes adapter improvements
- AKS and GKE remain experimental until private networking, identity,
  observability, backup/restore, upgrades, and real-cloud evidence pass their
  production checklists

## Phase 9: Optional Advanced Kubernetes Platform

- Argo Rollouts progressive delivery remains opt-in
- Istio is the recommended first service mesh implementation
- Linkerd and Consul service mesh remain evaluated alternatives
- No mesh, injection, strict mTLS, or ingress gateway is enabled by default
- Implementation starts only after disposable-cluster upgrade, rollback, and
  cleanup acceptance criteria from RFC 008 are automated

## Phase 10: Explicit Multi-cluster Placement

- Maintain multi-cluster inventory, read-only fleet health, GitOps rendering, and
  documented DNS failover as the initial supported boundary
- Evaluate Argo CD ApplicationSet after deterministic two-cluster rollout and
  rollback evidence; evaluate native Flux output separately
- Treat Cluster API and multi-cluster service mesh as optional future integrations
- Do not provide an automatic global scheduler, secret/data replication, or hidden
  failover; see RFC 013

## Phase 11: Edge Evaluation

- Define a K3s-first hardware, OS, architecture and resource profile; retain RKE2
  as the hardened experimental alternative
- Design signed offline bundles, a local registry mirror, lightweight
  observability, edge workload defaults, backup targets, and GitOps pull behavior
- Exercise disconnection, reconnect, power loss, disk pressure, upgrade, device
  revocation, backup and full rebuild on disposable matching hardware
- Do not implement or advertise `cf edge` commands as supported until RFC 014
  acceptance evidence exists
