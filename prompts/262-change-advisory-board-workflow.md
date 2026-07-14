# Prompt 262 — Change Advisory Board workflow

```text
Implement Change Advisory Board workflow support.

Goal:
Support formal change review for production environments without forcing it on all users.

Database:
- change_windows
- change_requests
- change_request_reviews

Change request fields:
- id
- organization_id
- project_id
- environment_id
- title
- description
- requested_by
- planned_start
- planned_end
- risk_level
- status
- related_plan_request_id
- related_apply_request_id
- metadata_json

Statuses:
- draft
- submitted
- approved
- rejected
- scheduled
- implemented
- canceled

API:
- POST /api/v1/change-requests
- GET /api/v1/change-requests
- GET /api/v1/change-requests/{id}
- POST /api/v1/change-requests/{id}/submit
- POST /api/v1/change-requests/{id}/approve
- POST /api/v1/change-requests/{id}/reject

CLI:
- cf change-request create
- cf change-request submit <id>
- cf change-request approve <id>
- cf change-request list
- cf change-request show <id>

Integration:
- prod apply can require approved change request
- apply request links to change request
- deployment windows checked
- approvals audited

Tests:
- prod apply blocked without change request when policy enabled
- approved change request allows apply request
- expired window blocks apply
- audit events created

Docs:
- docs/control-plane/change-advisory.md

Rules:
- Optional feature.
- Do not make CAB mandatory for all environments.
```
