# Prompt 210 — Artifact retention and cleanup

```text
Implement artifact retention and cleanup.

Goal:
Automatically clean expired artifacts while keeping audit records.

Config:
artifacts:
  retention_days: 30
  cleanup:
    enabled: true
    interval: 1h
    delete_batch_size: 100

Behavior:
- artifacts get expires_at at creation time
- cleanup worker deletes expired artifact content
- database record may be marked deleted instead of hard-deleted
- audit event created for cleanup
- manual delete remains available with RBAC

API:
- GET /api/v1/artifacts/retention
- POST /api/v1/artifacts/cleanup/run

CLI:
- cf artifact cleanup
- cf artifact retention show

Tests:
- expired artifacts deleted
- non-expired artifacts retained
- deleted artifact cannot be downloaded
- cleanup audit event created
- cleanup idempotent

Docs:
- docs/control-plane/artifact-retention.md

Rules:
- Do not delete audit events.
- Do not delete unexpired artifacts.
- Make cleanup disabled by default for local dev if safer.
```
