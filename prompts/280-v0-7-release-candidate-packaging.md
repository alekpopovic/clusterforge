# Prompt 280 — v0.7 release candidate packaging

```text
Prepare ClusterForge v0.7 release candidate.

Goal:
Package the repository for v0.7.0 RC with SaaS-ready enterprise governance foundations.

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
   - create RELEASE_NOTES_V0.7.md
   - include:
     - tenant isolation
     - scoped data access
     - quotas and rate limits
     - usage metering
     - onboarding/offboarding
     - KMS-backed artifact encryption
     - secret reference broker
     - Kubernetes Job runner executor
     - ephemeral runner pools
     - job fairness
     - advanced approvals
     - policy exceptions
     - risk acceptance
     - compliance evidence
     - evidence bundles
     - immutable audit log
     - SCIM provisioning
     - data residency model
     - cluster blueprints
     - blueprint registry
     - fleet add-on manager
     - GitOps status ingestion
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
   - create two organizations
   - verify cross-tenant access denied
   - create workspace/project/environment
   - configure RBAC
   - configure quota
   - sync sample project
   - upload artifact with encryption enabled
   - verify artifact download with RBAC
   - start runner in fake mode
   - create plan request
   - verify usage event recorded
   - create approval request
   - verify advanced approval requirements
   - create policy exception
   - create evidence bundle
   - verify audit hash chain
   - verify dashboard build

5. Security review:
   - tenant isolation tests pass
   - no raw tokens stored
   - no cloud credentials stored
   - no Terraform state stored
   - raw plan upload disabled by default
   - artifacts encrypted when configured
   - audit hash chain verifies
   - SCIM tokens hashed
   - quota/rate limit behavior documented
   - secret reference broker stores references only

6. Docs review:
   Verify:
   - docs/control-plane/saas-readiness.md
   - docs/control-plane/tenant-isolation-testing.md
   - docs/control-plane/scoped-data-access.md
   - docs/control-plane/quotas-rate-limits.md
   - docs/control-plane/usage-metering.md
   - docs/control-plane/onboarding.md
   - docs/control-plane/offboarding.md
   - docs/control-plane/artifact-encryption-kms.md
   - docs/control-plane/secret-reference-broker.md
   - docs/control-plane/kubernetes-job-executor.md
   - docs/control-plane/advanced-approvals.md
   - docs/control-plane/policy-exceptions.md
   - docs/control-plane/risk-acceptance.md
   - docs/control-plane/evidence-collection.md
   - docs/control-plane/evidence-bundles.md
   - docs/control-plane/immutable-audit.md
   - docs/control-plane/scim.md
   - docs/control-plane/data-residency.md
   - docs/cluster-blueprints.md
   - docs/blueprint-registry.md
   - docs/fleet-addons.md
   - docs/gitops-status.md

7. Create:
   - RELEASE_CANDIDATE_V0.7.md

RELEASE_CANDIDATE_V0.7.md must include:
- release decision
- included features
- excluded features
- breaking changes
- migration notes
- test results
- smoke test status
- security status
- tenant isolation status
- compliance evidence status
- known limitations
- upgrade notes
- blockers
- deferred items

Rules:
- Do not claim public SaaS readiness.
- Do not claim compliance certification.
- Do not hide skipped tests.
- No credentials.
- No real cloud apply.
- Fix only release-blocking issues.

Final response:
- State v0.7 RC status.
- List blockers.
- List commands run.
- List changed files.
```
