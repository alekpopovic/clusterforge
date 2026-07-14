# Prompt 250 — Organization offboarding and data deletion workflow

```text
Design and implement organization offboarding workflow.

Goal:
Allow safe deactivation and deletion of organization data.

Create docs:
- docs/control-plane/offboarding.md
- docs/rfcs/028-data-deletion.md

Control Plane states:
- active
- suspended
- deletion_requested
- deletion_in_progress
- deleted

API:
- POST /api/v1/organizations/{id}/suspend
- POST /api/v1/organizations/{id}/deletion-request
- POST /api/v1/organizations/{id}/deletion-confirm
- GET /api/v1/organizations/{id}/deletion-status

CLI:
- cf org suspend <org>
- cf org deletion request <org>
- cf org deletion confirm <org>
- cf org deletion status <org>

Deletion scope:
- projects
- environments
- clusters metadata
- apps metadata
- jobs
- artifacts
- policy results
- drift results
- cost reports
- usage events
- service catalog
- runbooks imported into DB

Retention exceptions:
- audit events may be retained or anonymized based on policy
- legal hold blocks deletion
- backups have separate retention policy

Safety:
- two-step confirmation
- admin-only
- audit event created
- export before deletion option
- dry-run deletion report
- no cloud resource deletion by default

Tests:
- suspend blocks writes
- deletion dry-run reports resources
- deletion request requires admin
- legal hold blocks deletion
- deleted org inaccessible
- audit event created

Rules:
- Do not delete customer cloud infrastructure.
- Do not delete Git repositories.
- Do not delete Terraform state.
- This deletes Control Plane metadata only unless explicitly extended.
```
