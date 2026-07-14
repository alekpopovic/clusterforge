# Prompt 260 — Policy exceptions and waivers

```text
Implement policy exception and waiver workflow.

Goal:
Allow controlled temporary exceptions for policy findings.

Use cases:
- temporary use of public ingress
- temporary latest tag in dev
- known security exception
- migration period
- legacy module exception

Database:
- policy_exceptions
  - id
  - organization_id
  - project_id nullable
  - environment_id nullable
  - policy_id
  - scope_type
  - scope_id
  - reason
  - created_by
  - approved_by nullable
  - expires_at
  - status
  - created_at
  - updated_at

Statuses:
- requested
- approved
- rejected
- expired
- revoked

API:
- POST /api/v1/policy-exceptions
- GET /api/v1/policy-exceptions
- POST /api/v1/policy-exceptions/{id}/approve
- POST /api/v1/policy-exceptions/{id}/reject
- POST /api/v1/policy-exceptions/{id}/revoke

CLI:
- cf policy exception request --policy <id> --reason "..."
- cf policy exception list
- cf policy exception approve <id>
- cf policy exception revoke <id>

Behavior:
- exceptions are time-limited
- high severity exceptions require approval
- expired exceptions ignored
- policy results show waived status
- all exceptions audited
- dashboard displays active exceptions

Tests:
- exception waives policy
- expired exception does not waive
- high severity requires approval
- exception scoped to correct environment only
- audit events created

Docs:
- docs/control-plane/policy-exceptions.md

Rules:
- No permanent exceptions by default.
- Reason required.
- Approval required for prod/blocking policies.
```
