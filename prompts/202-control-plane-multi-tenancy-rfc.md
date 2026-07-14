# Prompt 202 — Control Plane multi-tenancy RFC

```text
Create the multi-tenancy RFC for ClusterForge Control Plane.

Goal:
Design safe organization/workspace/project isolation for enterprise use.

Create:
- docs/rfcs/019-control-plane-multi-tenancy.md
- docs/control-plane/multi-tenancy.md

Cover:

1. Goals
   - multiple organizations
   - multiple workspaces
   - multiple projects
   - team-based access
   - isolated runners
   - isolated audit events
   - isolated policy results
   - isolated plan/apply workflows

2. Non-goals for first implementation
   - public SaaS-grade tenant isolation
   - billing
   - marketplace
   - cross-tenant sharing
   - untrusted plugins
   - shared cloud credential storage

3. Resource hierarchy
   Organization
     Workspace
       Project
         Environment
           Cluster
           Stack
           App

4. Isolation boundaries
   - database rows
   - API authorization
   - runner permissions
   - audit log scoping
   - artifact scoping
   - dashboard filtering

5. Security model
   - organization admins
   - workspace admins
   - project operators
   - environment viewers
   - runner service accounts
   - approval roles

6. Data model changes
   - organization_id on every scoped table
   - workspace_id where appropriate
   - project_id where appropriate
   - environment_id where appropriate

7. API behavior
   - every request is scoped
   - actor context includes organization/workspace/project permissions
   - cross-organization access blocked by default

8. Migration strategy
   - existing single-tenant data migrates to default organization
   - default workspace created
   - existing projects attached to default workspace

9. Risks
   - accidental cross-tenant access
   - missing authorization checks
   - runner token overreach
   - audit visibility gaps

10. Testing strategy
   - authorization unit tests
   - multi-tenant API tests
   - cross-tenant access denial tests
   - runner scoping tests

Do not implement code in this prompt.
Update:
- ROADMAP_V0.6.md if it exists
- docs/control-plane/security.md
```
