# CIS AWS Foundations Benchmark mapping

This is not a CIS certification or claim of AWS account compliance.

| Control area | ClusterForge feature or policy | Status | Evidence source | Limitations |
|---|---|---|---|---|
| IAM | Workload identity and no-static-secret guidance | partial | Terraform config, manual procedure | Root, users, Identity Center and access reviews are outside project scope. |
| Logging | Security policy guidance and observability integrations | documented | Terraform config, manual procedure | Organization trails and every enabled region/account need separate proof. |
| Monitoring | Alerting/incident integration guidance | partial | CI result, manual procedure | Metric filters, routing and response operation are organization-owned. |
| Networking | VPC, endpoints and security group modules | partial | Terraform config, policy result | Account-wide default VPC and unrestricted rule inventory is not automatic. |
