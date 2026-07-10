---
title: Roadmap
permalink: /roadmap/
---

# Roadmap

See also:

- [v0.2 roadmap](../ROADMAP_V0.2.md)
- [v0.3 roadmap](../ROADMAP_V0.3.md)
- [backlog](../BACKLOG.md)
- [AKS RFC](rfcs/002-aks-support.md)
- [GKE RFC](rfcs/003-gke-support.md)
- [self-hosted Kubernetes RFC](rfcs/004-self-hosted-kubernetes.md)
- [plugin architecture RFC](rfcs/006-cli-plugin-architecture.md)
- [workload identity RFC](rfcs/007-workload-identity.md)
- [service mesh RFC](rfcs/008-service-mesh.md)

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

## Phase 9: Optional Advanced Kubernetes Platform

- Argo Rollouts progressive delivery remains opt-in
- Istio is the recommended first service mesh implementation
- Linkerd and Consul service mesh remain evaluated alternatives
- No mesh, injection, strict mTLS, or ingress gateway is enabled by default
- Implementation starts only after disposable-cluster upgrade, rollback, and
  cleanup acceptance criteria from RFC 008 are automated
