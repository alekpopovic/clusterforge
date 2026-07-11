# SOC 2 control mapping

This is an implementation aid, not a SOC 2 report, certification, or attestation.

| Control area | ClusterForge feature or policy | Status | Evidence source | Limitations |
|---|---|---|---|---|
| CC6 logical access | Workload identity, RBAC, secret references | partial | Terraform config, policy result | Organization IAM governance and access reviews are external. |
| CC7 system operations | Audit, drift, health, incident runbooks | implemented | Audit log, CI result, manual procedure | Local audit is mutable; monitoring and retention need external systems. |
| CC8 change management | Plans, production plan-file guard, Git workflows | implemented | Terraform config, CI result | Repository approvals and segregation of duties are external. |
| Availability | Backup/DR modules and runbooks | partial | CI result, manual procedure | Recovery objectives and successful exercises need organization evidence. |
