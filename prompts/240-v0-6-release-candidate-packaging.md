# Prompt 240 — v0.6 release candidate packaging

```text
Prepare ClusterForge v0.6 release candidate.

Goal:
Package the repository for v0.6.0 RC with production-grade Control Plane hardening.

Tasks:

1. Versioning:
   - update VERSION
   - update CLI version
   - update control-plane version
   - update runner version
   - update dashboard package version if applicable
   - update Helm chart version
   - update CHANGELOG.md

2. Release notes:
   - create RELEASE_NOTES_V0.6.md
   - include:
     - multi-tenancy
     - RBAC v2
     - OIDC/SSO
     - token lifecycle
     - artifact storage
     - job queue hardening
     - runner pools
     - runner sandboxing
     - environment locks
     - incident mode
     - change history
     - rollback planner
     - VCS webhooks
     - preview environments
     - HA deployment docs
     - retention policies
     - known limitations
     - migration notes
     - security warnings

3. Validation:
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

4. Control Plane smoke test:
   - start local database
   - start API server
   - run migrations
   - create organization/workspace/project
   - configure RBAC
   - login through token auth
   - sync sample project
   - start runner in fake executor mode
   - create plan request
   - verify plan completes
   - create apply request
   - verify approval required
   - verify environment lock blocks apply
   - verify audit events exist
   - verify artifact upload/download
   - verify dashboard build

5. Security review:
   - no raw tokens stored
   - no secrets in artifacts
   - no Terraform state stored
   - no cloud credentials stored
   - raw plan upload disabled by default
   - audit events redacted
   - RBAC deny-by-default verified

6. Docs review:
   Verify:
   - docs/control-plane/*
   - docs/control-plane/rbac.md
   - docs/control-plane/oidc.md
   - docs/control-plane/artifacts.md
   - docs/control-plane/jobs.md
   - docs/control-plane/runner-pools.md
   - docs/control-plane/environment-locks.md
   - docs/control-plane/incident-mode.md
   - docs/control-plane/rollback.md
   - docs/control-plane/high-availability.md
   - docs/control-plane/data-retention.md

7. Create:
   - RELEASE_CANDIDATE_V0.6.md

RELEASE_CANDIDATE_V0.6.md must include:
- release decision
- included features
- excluded features
- breaking changes
- migration notes
- test results
- smoke test status
- security status
- known limitations
- upgrade notes
- blockers
- deferred items

Rules:
- Do not claim SaaS readiness.
- Do not claim compliance certification.
- Do not hide skipped tests.
- No credentials.
- No real cloud apply.
- Fix only release-blocking issues.

Final response:
- State v0.6 RC status.
- List blockers.
- List commands run.
- List changed files.
```
