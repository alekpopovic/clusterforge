# CIS Kubernetes Benchmark mapping

This is not a CIS certification or claim of benchmark conformance.

| Control area | ClusterForge feature or policy | Status | Evidence source | Limitations |
|---|---|---|---|---|
| Control plane | Managed Kubernetes module options | partial | Terraform config | Provider-managed flags may be inaccessible; benchmark version varies. |
| RBAC/service accounts | Tenant RBAC and workload identity | partial | Terraform config, policy result | Effective live permissions require cluster inspection. |
| Pod security | PSA, Kyverno and Gatekeeper policies | implemented | Policy result, Terraform config | Enforcement is opt-in and exemptions require review. |
| Network policy | Baseline NetworkPolicy module | implemented | Terraform config, policy result | Allowed application flows are environment-specific. |
