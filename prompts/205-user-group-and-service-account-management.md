# Prompt 205 — User, group and service account management

```text
Implement user, group and service account management.

Goal:
Give admins a way to manage identities used by RBAC.

Control Plane API:

Users:
- GET /api/v1/users
- GET /api/v1/users/{id}
- PATCH /api/v1/users/{id}
- DELETE /api/v1/users/{id} or deactivate

Groups:
- GET /api/v1/groups
- POST /api/v1/groups
- GET /api/v1/groups/{id}
- PATCH /api/v1/groups/{id}
- DELETE /api/v1/groups/{id}
- POST /api/v1/groups/{id}/members
- DELETE /api/v1/groups/{id}/members/{user_id}

Service accounts:
- GET /api/v1/service-accounts
- POST /api/v1/service-accounts
- GET /api/v1/service-accounts/{id}
- DELETE /api/v1/service-accounts/{id}
- POST /api/v1/service-accounts/{id}/tokens
- DELETE /api/v1/service-accounts/{id}/tokens/{token_id}

CLI:
- cf user list
- cf user show <id>
- cf group list
- cf group create <name>
- cf group add-member <group> <user>
- cf service-account list
- cf service-account create <name>
- cf service-account token create <name>

Requirements:
- service account tokens are shown only once
- store only token hashes
- token expiration supported
- token last_used_at tracked
- audit all identity changes
- RBAC required for mutations

Dashboard:
- basic admin pages optional:
  - users
  - groups
  - service accounts

Tests:
- create group
- add member
- create service account
- create token
- token not retrievable after creation
- expired token rejected
- unauthorized user blocked

Docs:
- docs/control-plane/identity-management.md

Rules:
- No default admin in production.
- Avoid account lockout by documenting bootstrap admin flow.
- Do not log tokens.

Run:
- go test ./... in control-plane and cli
```
