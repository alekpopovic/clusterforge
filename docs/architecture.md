---
title: Architecture
permalink: /architecture/
---

# Architecture

ClusterForge is organized around four layers. The goal is to keep each layer
readable, composable, and independently reviewable.

## Foundation Layer

Foundation modules create shared infrastructure that orchestrators and
workloads depend on:

- Cloud networking
- IAM and service roles
- DNS
- Storage
- Registries
- Firewall and security group rules

Examples:

- `modules/cloud/aws/network`
- `modules/core/tags`
- `modules/core/labels`
- `modules/core/naming`

Provider configuration belongs in root environments, not reusable child
modules.

## Orchestrator Layer

Orchestrator modules create or connect to the container platform:

- AWS EKS
- AWS ECS/Fargate
- Nomad
- Docker Engine and Docker Swarm
- Future AKS, GKE, K3s, and RKE2 modules

These modules should focus on platform control-plane and scheduling resources.
They should not deploy business applications.

## Platform Layer

Platform modules install cluster-level add-ons:

- Ingress controllers
- cert-manager and TLS automation
- external-dns
- Metrics and observability
- Logging
- GitOps tooling such as Argo CD

For Kubernetes, the bootstrap module composes Helm-based child modules. It
assumes Kubernetes and Helm providers are configured by the root environment.

## Workload Layer

Workload modules deploy applications or jobs:

- Kubernetes apps and cronjobs
- ECS services and scheduled tasks
- Nomad services and batch jobs
- Docker containers and Swarm services

The CLI app manifest renderer targets this layer. It generates readable module
calls rather than hiding workload definitions inside the CLI.

## Live Environments

Real environment compositions live under `live/`. Each environment root owns:

- Provider configuration
- Backend configuration
- Module composition
- Environment-specific variables
- State boundary

Keep production, staging, and development roots separate.
