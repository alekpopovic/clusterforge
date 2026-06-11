# Roadmap

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
