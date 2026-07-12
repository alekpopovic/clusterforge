# Prompt 171 — Approval workflow MVP

```text
Implement approval workflow MVP.

Goal:
Require approval before any apply job can run through Control Plane.

Resources:
- approvals
- apply_requests

Approval states:
- pending
- approved
- rejected
- expired
- canceled

Control Plane endpoints:
- POST /api/v1/apply-requests
- GET /api/v1/apply-requests
- GET /api/v1/apply-requests/{id}
- POST /api/v1/apply-requests/{id}/approve
- POST /api/v1/apply-requests/{id}/reject

Rules:
- Apply request must reference a successful plan request.
- Prod apply requires approval from someone other than requester.
- Optional policy:
  require_two_person_approval_for_prod: true
- Rejected apply cannot be executed.
- Expired approval cannot be executed.
- Approval decision must be audit logged.

CLI:
- cf apply request --plan <plan-id>
- cf approval list
- cf approval approve <apply-id>
- cf approval reject <apply-id>
- cf approval status <apply-id>

Runner:
- Do not execute apply yet unless apply request is approved.
- Runner should ignore unapproved apply jobs.

Tests:
- prod requester cannot self-approve if policy enabled
- viewer cannot approve
- operator can request apply
- admin can approve
- rejection blocks apply
- audit event created

Docs:
- docs/approval-workflow.md

Rules:
- No automatic approval.
- No apply without successful plan.
- No apply without policy checks.
- Be conservative for prod.
```
