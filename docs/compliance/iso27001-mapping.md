# ISO 27001 control mapping

This mapping is not ISO/IEC 27001 certification or legal/compliance advice.

| Control area | ClusterForge feature or policy | Status | Evidence source | Limitations |
|---|---|---|---|---|
| Access control | IAM/workload identity and Kubernetes RBAC modules | partial | Terraform config, policy result | Joiner/mover/leaver and periodic review processes are not covered. |
| Configuration management | Versioned Terraform, module checks, drift | implemented | Terraform config, CI result | Asset ownership and approved baselines remain organization-specific. |
| Logging and monitoring | Local audit and observability modules | partial | Audit log, Terraform config | Central retention, correlation and response SLAs are external. |
| ICT readiness | Backup, restore and DR documentation | documented | Manual procedure, CI result | Business impact analysis and exercised recovery evidence are not provided. |
