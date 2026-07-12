# Prompt 192 — Control Plane backup and restore

```text
Create backup and restore support for ClusterForge Control Plane.

Docs:
- docs/control-plane-backup-restore.md

Cover:
- PostgreSQL backup
- database restore
- audit event retention
- artifact storage if implemented
- configuration backup
- runner token rotation
- dashboard statelessness
- disaster recovery steps

CLI:
- cf api backup plan
- cf api backup check

For MVP:
- read-only checks:
  - database configured
  - backup docs exist
  - backup evidence file exists
  - recent backup evidence is present

Evidence file:
control-plane-backup-evidence.yaml

Tests:
- evidence parsing
- missing backup warning
- backup report JSON

Rules:
- Do not dump database in CLI MVP.
- Do not store backup credentials.
- No destructive restore automation.
```
