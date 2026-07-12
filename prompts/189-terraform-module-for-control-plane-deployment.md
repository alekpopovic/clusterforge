# Prompt 189 — Terraform module for Control Plane deployment

```text
Create Terraform module for deploying ClusterForge Control Plane to Kubernetes.

Path:
- modules/platform/kubernetes/clusterforge-control-plane

Purpose:
Install the Control Plane Helm chart into an existing Kubernetes cluster.

Inputs:
- namespace default "clusterforge-system"
- create_namespace default true
- chart_version default ""
- chart_repository optional
- chart_path optional for local chart
- values list(string)
- api_ingress_enabled default false
- api_host default ""
- existing_secret_name default ""
- labels map(string)

Resources:
- helm_release
- namespace optional

Outputs:
- namespace
- release_name
- service_name

Docs:
- README with examples
- existing Kubernetes deployment
- EKS deployment
- local kind deployment

Rules:
- Do not create database in this module unless explicitly enabled later.
- Do not store admin token in Terraform values by default.
- Document secret handling.

Run:
- terraform fmt -recursive
```
