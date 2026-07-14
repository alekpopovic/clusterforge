# Prompt 229 — Incident mode

```text
Implement incident mode for environments.

Goal:
Allow teams to mark an environment as under incident and automatically restrict risky actions.

Control Plane:
- incidents table:
  - id
  - environment_id
  - title
  - severity
  - status
  - started_by
  - started_at
  - resolved_by
  - resolved_at
  - metadata_json

API:
- POST /api/v1/environments/{id}/incident/start
- POST /api/v1/incidents/{id}/resolve
- GET /api/v1/incidents

CLI:
- cf incident start prod --severity sev2 --title "Ingress outage"
- cf incident resolve <id>
- cf incident list
- cf env status prod

Policy:
- active incident blocks apply by default
- emergency override requires admin/incident commander role
- plan and read-only operations allowed
- audit everything

Notifications:
- incident started
- incident resolved
- apply blocked due to incident

Dashboard:
- incident banner on environment page
- incident list

Tests:
- incident blocks apply
- incident resolve unblocks
- override audited
- dashboard API includes incident status

Docs:
- docs/control-plane/incident-mode.md
```
