# Prompt 273 — Cluster lifecycle blueprint model

```text
Implement cluster lifecycle blueprint model.

Goal:
Define reusable cluster templates independent of specific implementation backend.

Blueprint types:
- terraform_eks
- terraform_aks
- terraform_gke
- existing_kubernetes
- local_kind
- cluster_api future

Create:
- cluster-blueprints/
  aws-eks-small.yaml
  aws-eks-prod.yaml
  existing-kubernetes.yaml
  local-kind.yaml

Blueprint schema:
name: aws-eks-small
description: Small AWS EKS cluster for dev
backend: terraform
cloud: aws
orchestrator: eks
version: 0.1.0
inputs:
  region:
    type: string
    default: eu-central-1
  node_count:
    type: number
    default: 2
platform_addons:
  ingress_nginx: true
  cert_manager: true
  external_secrets: false
policies:
  production_ready: false

CLI:
- cf blueprint list
- cf blueprint show <name>
- cf blueprint validate <name>
- cf env create dev --blueprint aws-eks-small

Control Plane:
- optional blueprint catalog API:
  - GET /api/v1/blueprints

Tests:
- validate blueprint
- generate env from blueprint
- invalid input type fails
- blueprint list

Docs:
- docs/cluster-blueprints.md

Rules:
- Blueprints generate readable Terraform.
- Blueprints do not hide security warnings.
- Production blueprint must require remote backend and approvals.
```
