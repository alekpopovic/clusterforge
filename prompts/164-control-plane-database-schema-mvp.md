# Prompt 164 — Control Plane database schema MVP

```text
Implement the database schema MVP for ClusterForge Control Plane.

Goal:
Store inventory, audit events, policy results, drift results, plan/apply requests, and runners.

Database:
- PostgreSQL preferred for production.
- SQLite optional for local development if easy.
- Use a migration tool or simple SQL migrations.

Create migrations for tables:

1. organizations
   - id
   - name
   - slug
   - created_at
   - updated_at

2. workspaces
   - id
   - organization_id
   - name
   - slug
   - description
   - created_at
   - updated_at

3. projects
   - id
   - workspace_id
   - name
   - repo_url
   - default_branch
   - created_at
   - updated_at

4. environments
   - id
   - project_id
   - name
   - cloud
   - orchestrator
   - region
   - path
   - status
   - created_at
   - updated_at

5. clusters
   - id
   - environment_id
   - name
   - cloud
   - orchestrator
   - region
   - status
   - metadata_json
   - created_at
   - updated_at

6. apps
   - id
   - project_id
   - name
   - owner
   - image
   - metadata_json
   - created_at
   - updated_at

7. policy_results
   - id
   - project_id
   - environment_id nullable
   - policy_id
   - severity
   - status
   - message
   - metadata_json
   - created_at

8. drift_results
   - id
   - environment_id
   - stack
   - status
   - summary_json
   - created_at

9. cost_reports
   - id
   - environment_id
   - summary_json
   - created_at

10. audit_events
   - id
   - actor
   - action
   - resource_type
   - resource_id
   - metadata_json
   - created_at

11. runners
   - id
   - name
   - status
   - last_seen_at
   - metadata_json
   - created_at
   - updated_at

12. plan_requests
   - id
   - project_id
   - environment_id
   - stack
   - status
   - requested_by
   - summary_json
   - created_at
   - updated_at

13. apply_requests
   - id
   - plan_request_id
   - status
   - requested_by
   - approved_by nullable
   - created_at
   - updated_at

Create:
- internal/db
- internal/models
- migration command:
  clusterforge-server migrate up

Tests:
- migrations run on test database or SQLite if supported
- repositories can create/list projects
- repositories can create/list environments
- audit event insert works

Rules:
- Do not store Terraform state.
- Do not store cloud credentials.
- Do not store secret values.
- JSON metadata must be sanitized by service layer later.

Run:
- cd control-plane && go test ./...
```
