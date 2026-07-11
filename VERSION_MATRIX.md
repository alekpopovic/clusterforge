# Version Matrix

| Component | Supported | Tested | Status | Notes |
| --- | --- | --- | --- | --- |
| ClusterForge CLI | `0.1.x` | current commit | supported | Pre-1.0 compatibility may change by minor release. |
| Terraform | `>= 1.6.0` | local toolchain | supported | `cf doctor` warns below minimum. |
| OpenTofu | `>= 1.6.0` | optional | supported | Use `cf --engine tofu`. |
| Go | `>= 1.22` | GitHub Actions | supported | CLI development only. |
| Kubernetes | maintained minors | `1.29`, `1.30`, `1.31` | supported | Smoke tests record actual cluster version. |
| AWS provider | `>= 5.0, < 7.0` | module validation | supported | Root modules pin constraints. |
| Kubernetes provider | `>= 2.20, < 3.0` | examples | supported | Provider config belongs in roots. |
| Helm provider | `>= 2.10, < 3.0` | examples | supported | Used by platform modules. |
| Nomad provider | `>= 2.0, < 3.0` | module validation | experimental | Nomad support is early. |
| Docker provider | `>= 3.0, < 4.0` | module validation | experimental | Docker targets are local/self-managed. |
| EKS | AWS | smoke-test runbook | supported | Real pass requires matrix evidence. |
| ECS | AWS | smoke-test runbook | supported | Real pass requires matrix evidence. |
| Existing Kubernetes | any conformant cluster | smoke-test runbook | supported | Use disposable namespaces. |
| AKS | Azure | module validation only | experimental | MVP lacks several production hardening controls; see AKS production guide. |
| GKE | GCP | module validation only | experimental | MVP lacks several production hardening controls; see GKE production guide. |
| K3s | self-hosted | user-data generation | experimental | No VM provisioning. |
| RKE2 | self-hosted | user-data generation | experimental | No VM provisioning. |
| Nomad | self-managed | module validation | experimental | Production patterns still evolving. |
| Docker Swarm | self-managed | examples | experimental | Local/self-managed target. |
