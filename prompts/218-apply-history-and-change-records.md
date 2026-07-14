# Prompt 218 — Apply history and change records

```text
Implement apply history and change records.

Goal:
Track what changed, when, by whom, and with which plan/apply request.

Database:
- change_records table:
  - id
  - organization_id
  - project_id
  - environment_id
  - stack
  - plan_request_id
  - apply_request_id
  - applied_by
  - approved_by
  - git_ref
  - commit_sha
  - summary_json
  - resource_changes_json
  - created_at

API:
- GET /api/v1/change-records
- GET /api/v1/change-records/{id}
- GET /api/v1/environments/{id}/changes

CLI:
- cf change list <env>
- cf change show <id>
- cf change export <env> --format markdown|json

Dashboard:
- environment change history page
- change detail page

Requirements:
- link to plan/apply/audit events
- include resource change counts
- include risk level
- include policy result summary
- include approval metadata
- no secrets

Tests:
- change record created after successful apply
- failed apply does not create successful change record
- list by environment
- export markdown
- RBAC read restrictions

Docs:
- docs/control-plane/change-history.md
```
