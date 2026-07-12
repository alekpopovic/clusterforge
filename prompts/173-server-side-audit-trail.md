# Prompt 173 — Server-side audit trail

```text
Implement server-side audit trail for Control Plane.

Goal:
Record important actions in the API database.

Audit events:
- login/token use if practical
- project created/updated
- environment created/updated
- inventory sync
- policy results uploaded
- drift result uploaded
- plan request created
- plan job claimed
- plan completed/failed
- apply request created
- approval approved/rejected
- apply job started/completed/failed
- runner registered/heartbeat
- user/role changes if implemented

API:
- GET /api/v1/audit-events
- GET /api/v1/audit-events/{id}

Filters:
- actor
- action
- resource_type
- resource_id
- from
- to
- environment_id
- project_id

CLI:
- cf api audit list
- cf api audit show <id>
- cf api audit export --format jsonl|csv

Requirements:
- structured metadata
- sensitive fields redacted
- request ID linked where possible
- immutable events; no update/delete endpoint

Tests:
- audit event written for plan request
- audit event written for approval
- sensitive metadata redacted
- filters work
- export works

Docs:
- docs/control-plane-audit.md

Rules:
- No telemetry.
- Local/self-hosted audit only.
- Do not log secrets.
```
