# Prompt 216 — Environment locks and freeze windows

```text
Implement environment locks and freeze windows.

Goal:
Prevent unsafe changes during incidents, releases, or maintenance freezes.

Control Plane:
- environment_locks table:
  - id
  - environment_id
  - reason
  - locked_by
  - locked_at
  - expires_at nullable
  - active

- freeze_windows table:
  - id
  - environment_id
  - name
  - cron_or_schedule
  - timezone
  - reason
  - active

API:
- POST /api/v1/environments/{id}/lock
- POST /api/v1/environments/{id}/unlock
- GET /api/v1/environments/{id}/locks
- POST /api/v1/environments/{id}/freeze-windows
- GET /api/v1/environments/{id}/freeze-windows

CLI:
- cf env lock prod --reason "incident"
- cf env unlock prod
- cf env locks prod
- cf freeze list prod
- cf freeze create prod

Behavior:
- plan allowed during lock unless policy says otherwise
- apply blocked during active lock
- apply blocked during freeze window unless override permission
- all lock/freeze actions audited

Tests:
- lock blocks apply request
- unlock allows apply
- freeze window blocks apply
- admin override audited
- expired lock ignored

Docs:
- docs/control-plane/environment-locks.md

Rules:
- Lock/freeze must never be silently ignored.
- Override requires high privilege.
```
