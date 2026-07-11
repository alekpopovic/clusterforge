# NSA/CISA Kubernetes hardening mapping

This implementation aid is not an NSA/CISA endorsement or compliance claim.

| Control area | ClusterForge feature or policy | Status | Evidence source | Limitations |
|---|---|---|---|---|
| Pod security | Restricted PSA and extended admission pack | implemented | Terraform config, policy result | Runtime exceptions and controller health require verification. |
| Network separation | Namespace and NetworkPolicy modules | partial | Terraform config, policy result | Service mesh and application egress rules are environment-specific. |
| Authentication/authorization | Managed control plane, RBAC, workload identity | partial | Terraform config, manual procedure | Human identity provider configuration is outside ClusterForge. |
| Logging/threat detection | Observability modules and audit export | partial | Audit log, Terraform config | Central SIEM, detection content and response operation are external. |
