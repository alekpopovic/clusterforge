## Prompt 89 — EKS production hardening options

```text
Harden the EKS module for production use.

Target:
- modules/orchestrators/kubernetes/eks

Add inputs:
- endpoint_public_access
- endpoint_private_access
- public_access_cidrs
- enable_cluster_encryption
- kms_key_arn
- create_kms_key
- cluster_log_retention_days
- enabled_cluster_log_types
- deletion_protection or documented equivalent if supported
- node_group_update_config
- node_group_ami_type
- node_group_release_version
- node_group_force_update_version
- node_group_remote_access optional and disabled by default

Add resources where appropriate:
- aws_kms_key optional
- aws_kms_alias optional
- aws_cloudwatch_log_group for EKS control plane logs if needed
- encryption_config on cluster
- node group update_config

Behavior:
- Default should remain developer-friendly but safe.
- Production example should use private endpoint or restricted CIDR.
- Encryption should be easy to enable.
- Do not enable SSH access by default.
- Do not create overly permissive security groups.

Docs:
- Update EKS README.
- Add docs/aws-eks-production.md.
- Add example:
  examples/aws-eks-production-hardened

Validation:
- private endpoint only config example.
- encryption enabled example.
- restricted public CIDR example.

Run:
- terraform fmt -recursive
- validation where possible
```

---
