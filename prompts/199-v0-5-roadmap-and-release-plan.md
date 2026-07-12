# Prompt 199 — v0.5 roadmap and release plan

```text
Create v0.5.0 roadmap and release plan.

Create:
- ROADMAP_V0.5.md
- RELEASE_PLAN_V0.5.md
- BACKLOG_V0.5.md

Theme:
ClusterForge v0.5 introduces the optional self-hosted Control Plane.

Goals:
- Control Plane API server MVP
- database schema
- CLI API integration
- runner MVP
- plan request workflow
- approval workflow
- dashboard MVP
- audit trail
- notification MVP
- Kubernetes deployment chart
- Docker images
- security hardening
- E2E tests

Non-goals:
- public SaaS
- unrestricted apply UI
- plugin marketplace
- automatic remediation
- storing cloud credentials in API
- replacing Terraform Cloud
- full multi-tenant isolation

Milestones:
1. API and DB
2. CLI sync
3. Runner
4. Plan and approval workflow
5. Dashboard
6. Deployment
7. Security and tests
8. Release candidate

Acceptance criteria:
- local control plane can start
- CLI can sync project
- runner can process fake validate/plan job
- approval workflow blocks apply
- dashboard shows inventory
- audit events stored
- tests pass
- docs complete
- no secret storage

BACKLOG_V0.5.md:
Group tasks by:
- API
- CLI
- runner
- dashboard
- security
- deployment
- docs
- tests
- release

For each item:
- priority
- complexity
- status
- acceptance criteria

Final response:
- Summarize recommended v0.5 scope.
```
