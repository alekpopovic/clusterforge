# Prompt 230 — Emergency break-glass workflow

```text
Design and implement emergency break-glass workflow.

Goal:
Allow controlled emergency override of locks, freeze windows, incident mode, or approval restrictions.

Break-glass principles:
- rare
- time-limited
- heavily audited
- requires strongest role
- requires reason
- optional second approval
- not available to runner by default

Control Plane:
- break_glass_events table:
  - id
  - actor
  - environment_id
  - reason
  - scope
  - expires_at
  - created_at
  - used_at
  - metadata_json

API:
- POST /api/v1/break-glass
- GET /api/v1/break-glass
- POST /api/v1/break-glass/{id}/revoke

CLI:
- cf break-glass request prod --reason "restore service"
- cf break-glass approve <id>
- cf break-glass revoke <id>
- cf break-glass list

Behavior:
- can temporarily override specific blocks:
  - freeze window
  - environment lock
  - incident apply block
- cannot bypass policy checks for destructive changes unless explicitly allowed
- every use audited and notified

Tests:
- break-glass requires reason
- expires automatically
- override works only within scope
- revoked break-glass denied
- audit and notification emitted

Docs:
- docs/control-plane/break-glass.md

Rules:
- Do not make break-glass easy.
- No silent use.
- No permanent break-glass.
```
