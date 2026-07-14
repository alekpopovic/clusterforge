# Prompt 233 — Control Plane backup automation integration

```text
Add Control Plane backup automation integration.

Goal:
Provide concrete backup automation examples for production.

Create:
- examples/control-plane-backup/
  postgres-backup-cronjob.yaml
  s3-backup-policy-example.md
  restore-runbook.md

Optional Kubernetes module:
- modules/platform/kubernetes/clusterforge-control-plane-backup

Inputs:
- namespace
- schedule
- database_secret_name
- backup_bucket
- retention_days
- image
- resources

Behavior:
- CronJob runs pg_dump to object storage if configured
- secrets referenced by name only
- no credentials in Terraform values

Docs:
- docs/control-plane/backup-automation.md

Rules:
- Do not include real credentials.
- Do not make backup module default.
- Clearly state restore testing is required.
- Avoid pretending backup is complete without restore validation.

Run:
- terraform fmt -recursive if module added
```
