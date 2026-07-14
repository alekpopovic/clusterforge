# Prompt 269 — SCIM provisioning MVP

```text
Implement SCIM provisioning MVP.

Goal:
Support basic enterprise user and group provisioning.

Endpoints:
- GET /scim/v2/ServiceProviderConfig
- GET /scim/v2/ResourceTypes
- GET /scim/v2/Schemas
- GET /scim/v2/Users
- POST /scim/v2/Users
- GET /scim/v2/Users/{id}
- PATCH /scim/v2/Users/{id}
- DELETE /scim/v2/Users/{id}
- GET /scim/v2/Groups
- POST /scim/v2/Groups
- GET /scim/v2/Groups/{id}
- PATCH /scim/v2/Groups/{id}
- DELETE /scim/v2/Groups/{id}

Features:
- create user
- deactivate user
- update user email/name
- create group
- update group memberships
- delete/deactivate group
- map SCIM groups to ClusterForge groups

Security:
- SCIM token per organization
- token stored hashed
- token expiration
- audit all changes
- no admin by default from SCIM unless mapped explicitly

Tests:
- create user
- update user
- deactivate user
- create group
- add member
- remove member
- invalid token rejected
- deactivated user cannot authenticate

Docs:
- docs/control-plane/scim.md

Rules:
- Keep implementation minimal and interoperable.
- Do not overclaim full SCIM compatibility if partial.
```
