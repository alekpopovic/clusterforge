# Prompt 239 — v0.6 roadmap and release plan

```text
Create v0.6.0 roadmap and release plan.

Create:
- ROADMAP_V0.6.md
- RELEASE_PLAN_V0.6.md
- BACKLOG_V0.6.md

Theme:
ClusterForge v0.6 makes the Control Plane production-grade for internal enterprise use.

Goals:
- multi-tenancy design and implementation
- RBAC v2
- OIDC/SSO
- user/group/service account management
- token rotation
- artifact storage
- job queue hardening
- runner pools and labels
- runner sandboxing
- environment locks
- freeze windows
- incident mode
- change history
- rollback planner MVP
- GitHub/GitLab webhook integration
- preview environments MVP
- HA deployment docs
- migration hardening
- retention policies
- DR drill docs

Non-goals:
- public SaaS
- billing
- plugin marketplace
- automatic remediation
- fully automatic rollback
- global multi-cluster scheduler
- compliance certification

Milestones:
1. v0.5 release gate
2. multi-tenancy and RBAC
3. authentication and token lifecycle
4. artifacts and job queue
5. runner hardening
6. locks, incidents and approvals
7. VCS integrations
8. operational hardening
9. v0.6 release candidate

Acceptance criteria:
- v0.5 gate reviewed
- RBAC tests pass
- OIDC mock tests pass
- artifact tests pass
- runner pool tests pass
- job queue tests pass
- locks/freeze policy tests pass
- webhook tests pass
- docs complete
- no secrets stored
- release gate passes

BACKLOG_V0.6.md:
Group tasks by:
- API
- CLI
- runner
- dashboard
- security
- RBAC
- artifacts
- jobs
- VCS
- operations
- docs
- tests
- release

Each item:
- priority
- complexity
- owner placeholder
- status
- acceptance criteria
- notes

Final response:
- Summarize recommended v0.6 scope.
- List must-have.
- List deferred.
```
