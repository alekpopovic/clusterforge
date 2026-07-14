# Prompt 201 — v0.5 final release gate review

```text
Perform the final v0.5 release gate review for ClusterForge.

Goal:
Determine whether ClusterForge v0.5.0 with Control Plane MVP is ready to be tagged and released.

Create:
- RELEASE_GATE_V0.5.md

Inspect:
- CLI
- Control Plane API
- Control Plane database
- Runner
- Dashboard
- approval workflow
- plan workflow
- apply safeguards
- audit trail
- notification system
- Helm chart
- Docker images
- docs
- tests
- CI
- security controls

RELEASE_GATE_V0.5.md must include:

1. Release decision
   - ready
   - ready with warnings
   - blocked

2. Control Plane readiness
   For each component:
   - API server
   - database migrations
   - CLI sync
   - runner
   - plan requests
   - apply requests
   - approvals
   - audit trail
   - notifications
   - dashboard
   - Helm chart
   - Docker images

   Include:
   - status
   - tests
   - docs
   - known limitations
   - release risk

3. Security readiness
   - authentication
   - authorization
   - runner tokens
   - audit logging
   - secret redaction
   - artifact handling
   - plan/apply safeguards
   - no cloud credentials stored in API
   - no Terraform state stored in API

4. Operational readiness
   - deployment docs
   - database docs
   - backup/restore docs
   - observability docs
   - health checks
   - metrics
   - logs

5. Test readiness
   - Go unit tests
   - CLI integration tests
   - Control Plane API tests
   - Runner tests
   - E2E tests
   - Dashboard build
   - Helm lint
   - Docker build
   - security checks

6. Known limitations
   - MVP limitations
   - production warnings
   - deferred v0.6 work
   - missing cloud smoke tests
   - missing HA features

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
- helm lint charts/clusterforge-control-plane if helm is available
- docker build checks if Docker is available

Rules:
- Do not add new features.
- Fix only release-blocking issues.
- Do not claim production readiness beyond what is actually tested.
- Do not hide skipped checks.
- No credentials.
- No real cloud apply.

Final response:
- State v0.5 release readiness.
- List blockers.
- List commands run.
- List files changed.
```
