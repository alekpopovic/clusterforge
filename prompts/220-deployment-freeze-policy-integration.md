# Prompt 220 — Deployment freeze policy integration

```text
Integrate environment locks and freeze windows with policy engine.

Goal:
Make lock/freeze behavior visible in policy results and apply workflow.

Policy rules:
- env.locked.blocks_apply
- env.freeze_window.blocks_apply
- prod.requires_approval
- prod.requires_no_active_incident
- destroy.requires_override

Behavior:
- policy check includes active locks/freezes
- apply request is blocked if policy says blocked
- dashboard shows environment locked/frozen status
- notifications sent when apply is blocked by lock/freeze

CLI:
- cf policy check prod should show lock/freeze policies
- cf apply request should fail clearly when blocked
- cf env status shows lock/freeze

Tests:
- locked env policy result BLOCKED
- freeze window policy result BLOCKED
- dashboard status API includes lock
- notification event emitted

Docs:
- docs/control-plane/environment-locks.md
- docs/policy-engine.md

Rules:
- No silent overrides.
- Overrides require explicit RBAC permission and audit event.
```
