# Prompt 200 — v0.5 release candidate packaging

```text
Prepare ClusterForge v0.5 release candidate.

Goal:
Package the repository for v0.5.0 RC with Control Plane MVP.

Tasks:

1. Versioning:
   - update VERSION
   - update CLI version
   - update control-plane version
   - update runner version
   - update dashboard package version if applicable
   - update CHANGELOG.md

2. Release notes:
   - create RELEASE_NOTES_V0.5.md
   - include:
     - new Control Plane
     - runner MVP
     - dashboard MVP
     - approval workflow
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
   - docker build checks if Docker available

4. Control Plane smoke test:
   - start local database
   - start API server
   - run CLI login
   - sync sample project
   - start runner in fake executor mode
   - create plan request
   - verify plan completes
   - create apply request
   - verify approval required
   - verify audit events exist

5. Docs:
   Verify:
   - docs/control-plane/*
   - docs/remote-runner.md
   - docs/approval-workflow.md
   - docs/dashboard.md
   - docs/control-plane-security.md
   - docs/control-plane-deployment.md

6. Create:
   - RELEASE_CANDIDATE_V0.5.md

RELEASE_CANDIDATE_V0.5.md must include:
- release decision
- included features
- excluded features
- test results
- smoke test status
- security status
- known limitations
- upgrade notes
- blockers
- deferred items

Rules:
- Do not claim production readiness beyond what is tested.
- Do not hide skipped tests.
- No credentials.
- No real cloud apply.
- Fix only release-blocking issues.

Final response:
- State v0.5 RC status.
- List blockers.
- List commands run.
- List changed files.
```
