## Prompt 94 — Velero backup module

```text
Add Kubernetes backup support using Velero.

Create module:
- modules/platform/kubernetes/velero

Purpose:
Install Velero via Helm for Kubernetes backup and restore.

Inputs:
- namespace default "velero"
- create_namespace default true
- chart_version default ""
- provider string default "aws"
- bucket string
- backup_storage_location_name default "default"
- volume_snapshot_location_name default "default"
- service_account_role_arn default ""
- values list(string) default []
- labels map(string) default {}

Behavior:
- Install Velero Helm chart.
- Support AWS plugin configuration as first implementation.
- Do not create cloud bucket in this module.
- Use IRSA when service_account_role_arn is provided.
- Do not store cloud credentials in Terraform.

Add AWS module:
- modules/cloud/aws/velero-backup-bucket

Resources:
- S3 bucket
- bucket encryption
- bucket versioning
- public access block
- optional lifecycle policy

Docs:
- docs/backup-restore.md
- backup runbook
- restore runbook
- disaster recovery limitations

Example:
- examples/aws-eks-velero

Rules:
- Backup install is not enough; docs must explain testing restore.
- Do not enable deletion of backup bucket by default.
- Do not store credentials in values.

Run:
- terraform fmt -recursive
```

---
