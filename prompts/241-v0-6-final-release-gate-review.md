# Prompt 241 — v0.6 final release gate review

```text
Perform the final v0.6 release gate review for ClusterForge.

Goal:
Determine whether ClusterForge v0.6.0 is ready to be tagged and released.

Create:
- RELEASE_GATE_V0.6.md

Inspect:
- CLI
- Control Plane API
- Control Plane database
- RBAC v2
- OIDC/SSO
- user/group/service account management
- token rotation
- artifact storage
- artifact retention
- job queue hardening
- runner pools
- runner sandboxing
- environment locks
- freeze windows
- incident mode
- break-glass workflow
- change history
- rollback planner
- GitHub/GitLab webhooks
- preview environments
- HA deployment
- database migrations
- retention policies
- disaster recovery drill docs
- docs
- tests
- security controls
- Helm charts
- Docker images

RELEASE_GATE_V0.6.md must include:

1. Release decision
   - ready
   - ready with warnings
   - blocked

2. Feature readiness table
   For each v0.6 feature include:
   - status
   - tests
   - docs
   - known limitations
   - release risk

3. Security readiness
   - RBAC deny-by-default
   - OIDC security
   - token hashing
   - token rotation
   - runner scoping
   - artifact encryption
   - no Terraform state in API
   - no cloud credentials in API
   - audit redaction
   - break-glass auditability

4. Operational readiness
   - HA docs
   - backup docs
   - DR drill docs
   - migration policy
   - retention policy
   - observability dashboards
   - runner deployment
   - production deployment values

5. Test readiness
   - Go unit tests
   - CLI integration tests
   - Control Plane API tests
   - runner tests
   - dashboard build
   - Helm lint
   - Docker build
   - security tests
   - E2E workflow tests
   - skipped tests and why

6. Documentation readiness
   - README
   - docs/control-plane/*
   - docs/security/*
   - docs/rbac
   - docs/oidc
   - docs/artifacts
   - docs/jobs
   - docs/runner
   - docs/incident-mode
   - docs/break-glass
   - docs/rollback
   - docs/ha
   - docs/data-retention
   - docs/dr

7. Release blockers
   - critical
   - important
   - deferred

Run:
- make fmt-check
- make lint
- make test
- make validate
- make security
- make check-modules
- cd cli && go test ./...
- cd control-plane && go test ./...
- cd runner && go test ./...
- cd dashboard && npm run build if dashboard exists
- helm lint charts/clusterforge-control-plane if helm exists
- docker build checks if Docker is available

Rules:
- Do not add new features.
- Fix only release-blocking issues.
- Do not claim production readiness beyond what is tested.
- Do not claim compliance certification.
- Do not hide skipped tests.
- No credentials.
- No real cloud apply.

Final response:
- State v0.6 release readiness.
- List blockers.
- List commands run.
- List files changed.
```
