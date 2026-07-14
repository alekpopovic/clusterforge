# Prompt 261 — Risk acceptance workflow

```text
Implement risk acceptance workflow.

Goal:
Allow teams to explicitly accept operational/security risk with audit trail.

Difference from policy exception:
- policy exception waives a specific policy result
- risk acceptance records acknowledged risk for a defined scope and time

Database:
- risk_acceptances
  - id
  - organization_id
  - project_id nullable
  - environment_id nullable
  - title
  - description
  - risk_level
  - scope_type
  - scope_id
  - owner
  - created_by
  - approved_by
  - expires_at
  - status
  - mitigation_plan
  - metadata_json
  - created_at
  - updated_at

API:
- POST /api/v1/risk-acceptances
- GET /api/v1/risk-acceptances
- GET /api/v1/risk-acceptances/{id}
- POST /api/v1/risk-acceptances/{id}/approve
- POST /api/v1/risk-acceptances/{id}/revoke

CLI:
- cf risk accept
- cf risk list
- cf risk show <id>
- cf risk approve <id>
- cf risk revoke <id>
- cf risk report --format markdown|json

Dashboard:
- risk register page

Tests:
- create risk acceptance
- approval required for high risk
- expired risk shown as expired
- risk report includes active risks
- cross-tenant denied

Docs:
- docs/control-plane/risk-acceptance.md

Rules:
- Do not use risk acceptance to bypass destructive safety automatically.
- High risk requires approval.
- Expiration required by default.
```
