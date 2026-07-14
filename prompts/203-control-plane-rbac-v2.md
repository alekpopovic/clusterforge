# Prompt 203 — Control Plane RBAC v2

```text
Implement Control Plane RBAC v2.

Goal:
Replace simple role checks with scoped role-based access control.

Scopes:
- organization
- workspace
- project
- environment
- runner

Roles:
- owner
- admin
- operator
- approver
- viewer
- auditor
- runner

Permissions:
- organization:read
- organization:write
- workspace:read
- workspace:write
- project:read
- project:write
- environment:read
- environment:write
- app:read
- app:write
- plan:request
- plan:read
- apply:request
- apply:approve
- apply:execute
- policy:read
- drift:read
- cost:read
- audit:read
- runner:read
- runner:write
- runner:claim_job
- admin:all

Database:
- users
- groups
- group_members
- roles
- role_bindings
- service_accounts
- service_account_tokens or token hashes

API:
- GET /api/v1/rbac/roles
- GET /api/v1/rbac/permissions
- GET /api/v1/rbac/bindings
- POST /api/v1/rbac/bindings
- DELETE /api/v1/rbac/bindings/{id}
- GET /api/v1/me/permissions

CLI:
- cf rbac roles
- cf rbac bindings list
- cf rbac bind --role approver --user user@example.com --project payments
- cf rbac whoami

Requirements:
- deny by default
- support inherited permissions from organization to workspace to project
- environment-level overrides
- runner tokens limited to allowed projects/environments
- audit every RBAC change

Tests:
- viewer cannot create plan
- operator can request plan
- approver can approve but not self-approve when policy forbids
- auditor can read audit but not mutate resources
- runner can only claim allowed jobs
- cross-tenant access denied

Docs:
- docs/control-plane/rbac.md

Rules:
- Do not break existing local auth mode.
- Provide migration from old roles to new RBAC.
- Do not store raw tokens.
- Be conservative by default.

Run:
- cd control-plane && go test ./...
- cd cli && go test ./...
```
