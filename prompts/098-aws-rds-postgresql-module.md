## Prompt 98 — AWS RDS PostgreSQL module

```text
Implement AWS RDS PostgreSQL module.

Path:
- modules/cloud/aws/rds-postgres

Purpose:
Create a production-aware PostgreSQL database for applications.

Inputs:
- name
- environment
- vpc_id
- subnet_ids
- allowed_security_group_ids
- engine_version
- instance_class
- allocated_storage
- max_allocated_storage
- database_name
- master_username
- manage_master_user_password default true
- master_password default ""
- multi_az default false
- backup_retention_period default 7
- deletion_protection default true
- skip_final_snapshot default false
- storage_encrypted default true
- kms_key_arn default ""
- tags

Resources:
- aws_db_subnet_group
- aws_security_group
- aws_db_instance
- optional random password only if appropriate, but avoid putting secret outputs in normal output

Outputs:
- endpoint
- port
- db_name
- security_group_id
- secret_arn if AWS-managed password is used

README:
- Basic example.
- EKS app connection pattern.
- ECS app connection pattern.
- Secret handling with AWS Secrets Manager / External Secrets.
- Production warnings.

Example:
- examples/aws-rds-postgres

Rules:
- Do not output plaintext password.
- deletion_protection true by default.
- encrypted storage true by default.
- Document cost implications.
- Do not make database publicly accessible by default.

Run:
- terraform fmt -recursive
```

---
