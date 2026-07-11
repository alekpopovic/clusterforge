## Prompt 144 — Platform API RFC

```text
Create an RFC for a future ClusterForge platform API.

Create:
- docs/rfcs/011-platform-api.md

Goal:
Design a future API server without implementing it yet.

Cover:
1. Goals
   - central inventory
   - environment status
   - plan requests
   - policy results
   - audit events
   - service catalog
   - fleet operations
   - UI backend

2. Non-goals
   - replacing Terraform
   - storing cloud credentials directly
   - automatic remediation by default
   - multi-tenant SaaS in first version

3. API resources:
   - Organization
   - Workspace
   - Project
   - Environment
   - Cluster
   - Stack
   - App
   - Plan
   - PolicyResult
   - AuditEvent
   - Service
   - Runbook

4. Execution model options:
   - local CLI only
   - GitOps/PR workflow
   - remote runner
   - Terraform Cloud integration

5. Security:
   - authentication
   - authorization
   - audit logging
   - secret handling
   - approval workflow

6. Storage:
   - PostgreSQL
   - object storage for artifacts
   - no Terraform state storage in API by default

7. Future CLI:
   - cf login
   - cf api status
   - cf plan request
   - cf approval list

8. Risks:
   - scope creep
   - credential security
   - concurrency
   - state locking
   - enterprise requirements

Do not implement API server in this prompt.
Update ROADMAP_V0.4.md if it exists.
```


---
