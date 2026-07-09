## Prompt 90 — AWS KMS reusable module

```text
Create reusable AWS KMS module.

Path:
- modules/cloud/aws/kms-key

Purpose:
Create a KMS key and alias for encryption use cases:
- EKS secrets encryption
- S3 state bucket encryption
- EBS encryption
- application secrets encryption

Inputs:
- name
- environment
- description
- alias_name
- deletion_window_in_days default 30
- enable_key_rotation default true
- policy_json default ""
- tags map(string)

Resources:
- aws_kms_key
- aws_kms_alias

Outputs:
- key_id
- key_arn
- alias_name
- alias_arn

README:
- Basic key example.
- EKS encryption example.
- S3 backend encryption example.
- Warning about key deletion and recovery window.
- Warning about IAM permissions.

Update:
- EKS module docs to reference this module.
- tfstate backend module to optionally accept kms_key_arn if not already supported.

Rules:
- Do not create overly broad key policy unless documented.
- Do not hardcode account IDs.
- Rotation enabled by default.

Run:
- terraform fmt -recursive
```

---
