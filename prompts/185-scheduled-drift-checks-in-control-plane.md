# Prompt 185 — Scheduled drift checks in Control Plane

```text
Implement scheduled drift checks in Control Plane.

Goal:
Allow environments to have scheduled drift detection through runners.

Config:
drift_schedules:
  prod:
    cron: "0 */6 * * *"
    stacks:
      - network
      - cluster
      - platform
      - apps

Control Plane:
- store drift schedules
- create drift jobs on schedule
- track last run
- track next run
- store drift results

API:
- GET /api/v1/drift-schedules
- POST /api/v1/drift-schedules
- GET /api/v1/drift-results

CLI:
- cf drift schedule list
- cf drift schedule create <env>
- cf drift schedule disable <id>

Runner:
- support drift_check job
- run plan with detailed exit code
- upload summary

Rules:
- No apply.
- No auto-remediation.
- Make schedules disabled by default.
- Drift jobs require runner allowed job type.

Tests:
- schedule creation
- due schedule creates job
- runner uploads result
- drift detected result stored

Docs:
- docs/control-plane-drift.md
```
