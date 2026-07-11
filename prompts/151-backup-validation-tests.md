## Prompt 151 — Backup validation tests

```text
Add backup validation workflow for Kubernetes backups.

Goal:
Ensure backup documentation includes restore testing, not only backup installation.

Docs:
- docs/backup-validation.md

CLI:
- cf backup check <env>
- cf backup plan <env>
- cf backup report <env>

For MVP:
- Read-only checks:
  - Velero module enabled in config/Terraform if detectable
  - backup bucket module exists if AWS
  - Velero namespace expected
  - runbook exists
  - restore test evidence file exists

Evidence file:
- backup-evidence.yaml

Example:
backup_tests:
  prod:
    last_backup_test: "2026-07-01"
    last_restore_test: "2026-07-02"
    result: passed
    notes: "Restored namespace demo-restore"

Docs:
- Explain backup test evidence.
- Explain restore test environments.
- Explain schedule.
- Explain not to test restore directly in production.

Rules:
- Do not run destructive restore automatically.
- Do not delete backups.
- Do not call cloud APIs unless explicitly enabled.

Tests:
- Evidence file parsing.
- Missing restore test warns.
- Backup report JSON.

Run:
- gofmt
- go test ./...
```


---
