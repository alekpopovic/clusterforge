# Prompt 227 — Environment promotion through Control Plane

```text
Implement environment promotion workflow through Control Plane.

Goal:
Track and approve promotion from dev to staging to prod.

Control Plane:
- promotion_requests table:
  - id
  - project_id
  - app_id
  - from_environment_id
  - to_environment_id
  - source_version
  - target_version
  - status
  - requested_by
  - approved_by
  - created_at
  - updated_at

API:
- POST /api/v1/promotions
- GET /api/v1/promotions
- GET /api/v1/promotions/{id}
- POST /api/v1/promotions/{id}/approve
- POST /api/v1/promotions/{id}/reject

CLI:
- cf promote request --app api --from staging --to prod --version 1.2.3
- cf promote list
- cf promote approve <id>
- cf promote reject <id>

Behavior:
- generates app manifest patch or PR guidance
- does not apply automatically
- prod promotion requires approval
- promotion result links to plan/apply request if executed later

Dashboard:
- promotions page
- pending promotion approvals

Tests:
- create promotion
- approve promotion
- reject promotion
- prod approval required
- patch generated
- audit events

Docs:
- docs/control-plane/promotions.md

Rules:
- Do not copy secrets.
- Do not mutate prod without approved workflow.
```
