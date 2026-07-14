# Prompt 244 — Tenant-scoped data access layer

```text
Harden the Control Plane data access layer for tenant scoping.

Goal:
Make tenant scoping difficult to forget and easy to test.

Tasks:
1. Introduce a scoped query context:
   - organization_id
   - workspace_id optional
   - project_id optional
   - environment_id optional
   - actor_id
   - permissions

2. Update repositories/services to require scoped context for:
   - projects
   - environments
   - clusters
   - apps
   - artifacts
   - audit events
   - jobs
   - plan requests
   - apply requests
   - approvals
   - policy results
   - drift results
   - cost reports
   - service catalog
   - runbooks

3. Add guardrails:
   - repository methods must not expose unscoped list methods except admin/system methods
   - system methods must be clearly named and audited
   - tests should fail if unscoped repository methods are used in API handlers

4. Add middleware:
   - resolves actor scope
   - attaches scope to request context
   - denies missing organization scope for tenant APIs

5. Add static/conformance check:
   - search for repository calls without scope
   - warn or fail in CI

Docs:
- docs/control-plane/scoped-data-access.md

Tests:
- unscoped access denied
- scoped list returns only allowed data
- admin/system access audited
- API handlers require scope

Rules:
- Preserve local development mode.
- Do not weaken RBAC.
- Do not rely on frontend filtering for security.
- Deny by default.

Run:
- cd control-plane && gofmt -w .
- cd control-plane && go test ./...
```
