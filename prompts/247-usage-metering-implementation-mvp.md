# Prompt 247 — Usage metering implementation MVP

```text
Implement usage metering MVP.

Goal:
Record and report non-sensitive usage events.

Database:
- usage_events
  - id
  - organization_id
  - workspace_id nullable
  - project_id nullable
  - environment_id nullable
  - event_type
  - quantity
  - metadata_json
  - created_at

- usage_rollups_daily
  - id
  - organization_id
  - date
  - event_type
  - quantity
  - metadata_json
  - created_at

Events:
- api.request
- job.created
- job.completed
- plan.requested
- plan.completed
- apply.requested
- apply.completed
- policy.checked
- drift.checked
- cost.scanned
- artifact.uploaded
- artifact.downloaded
- preview.created
- preview.deleted
- runner.heartbeat

API:
- GET /api/v1/usage/events
- GET /api/v1/usage/summary
- POST /api/v1/usage/rollup/run

CLI:
- cf usage summary
- cf usage export --format json|csv

Dashboard:
- optional usage summary page

Requirements:
- tenant-scoped
- RBAC protected
- no secrets
- metadata sanitized
- retention controlled
- rollup job idempotent

Tests:
- usage event recorded for plan request
- usage event recorded for artifact upload
- summary aggregates by event type
- cross-tenant usage denied
- metadata redaction works

Docs:
- docs/control-plane/usage-metering.md

Run:
- cd control-plane && go test ./...
- cd cli && go test ./...
```
