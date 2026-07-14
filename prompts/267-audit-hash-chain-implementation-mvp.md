# Prompt 267 — Audit hash chain implementation MVP

```text
Implement audit hash chain MVP.

Goal:
Make audit events tamper-evident within the database.

Database changes:
- audit_events:
  - sequence_number
  - previous_hash
  - event_hash
  - hash_algorithm

Behavior:
- event_hash computed from canonical event fields
- previous_hash links to prior event in same organization
- sequence_number increments per organization
- audit event updates/deletes remain disallowed
- migration initializes hash chain for existing events if possible

API:
- GET /api/v1/audit-events/verify
- GET /api/v1/audit-events/checkpoints

CLI:
- cf audit verify
- cf audit checkpoint create
- cf audit checkpoint list

Tests:
- hash chain valid
- modified event fails verification
- missing event fails verification
- per-organization chains separate
- checkpoint created

Docs:
- docs/control-plane/immutable-audit.md

Rules:
- This is tamper-evident, not tamper-proof.
- Do not claim legal WORM compliance.
- Do not allow audit event mutation.
```
