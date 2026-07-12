# Prompt 162 — ClusterForge Control Plane architecture RFC

```text
Create the architecture RFC for ClusterForge Control Plane.

Goal:
Design the future API/server/dashboard layer without replacing Terraform or the CLI.

Create:
- docs/rfcs/016-control-plane-architecture.md
- docs/control-plane.md

Control Plane goals:
- central inventory
- project/environment/cluster catalog
- policy result storage
- drift result storage
- cost warning storage
- audit event storage
- service catalog
- runbook catalog
- dashboard API
- remote runner orchestration
- approval workflow
- notifications

Non-goals:
- replacing Terraform
- storing cloud credentials directly
- automatic remediation by default
- unrestricted apply button
- plugin marketplace
- public SaaS multi-tenancy in first version

Core components:
1. API server
2. Database
3. CLI integration
4. Runner agent
5. Dashboard
6. Notification system
7. Policy engine integration
8. Audit event pipeline

Main resources:
- Organization
- Workspace
- Project
- Environment
- Cluster
- Stack
- App
- Service
- Runbook
- PlanRequest
- ApplyRequest
- PolicyResult
- DriftResult
- CostReport
- AuditEvent
- Runner
- Approval

Architecture options:
1. Local-only mode
2. Self-hosted control plane
3. Enterprise internal deployment
4. Future SaaS mode

Security model:
- authentication
- authorization
- service accounts
- runner tokens
- no cloud credentials in API by default
- audit log
- approval workflow
- least privilege runner

Execution model:
- API creates plan request.
- Runner checks out repo.
- Runner executes ClusterForge/Terraform.
- Runner uploads sanitized results.
- Approval required before apply.
- Apply runs only after explicit approval.

Docs must include:
- sequence diagram in Mermaid
- deployment diagram in Mermaid
- data flow
- trust boundaries
- threat model summary

Do not implement code in this prompt.
Update:
- ROADMAP_V0.5.md if it exists
- docs/roadmap.md
```
