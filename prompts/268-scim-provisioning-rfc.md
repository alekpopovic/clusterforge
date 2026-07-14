# Prompt 268 — SCIM provisioning RFC

```text
Create SCIM provisioning RFC.

Goal:
Design automated user and group provisioning from enterprise identity providers.

Create:
- docs/rfcs/033-scim-provisioning.md
- docs/control-plane/scim.md

Cover:
1. Goals
   - user provisioning
   - user deprovisioning
   - group sync
   - group membership sync
   - RBAC group mapping
   - audit events

2. Non-goals
   - replacing OIDC login
   - storing identity provider credentials
   - supporting every SCIM edge case in MVP

3. Resources
   - Users
   - Groups
   - Group memberships

4. API endpoints
   - /scim/v2/Users
   - /scim/v2/Groups

5. Security
   - SCIM bearer token
   - token rotation
   - scoped to organization
   - audit all provisioning changes
   - deny login for deactivated users

6. Deprovisioning
   - deactivate user
   - remove group memberships
   - retain audit events
   - transfer ownership guidance

7. Testing
   - SCIM contract tests
   - create/update/delete user
   - group membership tests

Do not implement code in this prompt.
Update roadmap.
```
