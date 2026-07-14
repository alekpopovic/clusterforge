# Prompt 243 — Tenant isolation enforcement tests

```text
Implement tenant isolation enforcement tests.

Goal:
Prove that users, runners, service accounts, artifacts, audit events, jobs, and API resources cannot cross tenant boundaries.

Test scope:
- organizations
- workspaces
- projects
- environments
- clusters
- apps
- artifacts
- audit events
- policy results
- drift results
- cost reports
- runners
- jobs
- plan requests
- apply requests
- approvals
- service catalog
- runbooks

Create:
- control-plane/tests/tenant_isolation/
- docs/control-plane/tenant-isolation-testing.md

Test scenarios:

1. Organization isolation
   - user in org A cannot read org B projects
   - user in org A cannot list org B artifacts
   - user in org A cannot read org B audit events
   - user in org A cannot approve org B apply request

2. Workspace isolation
   - workspace viewer cannot read unrelated workspace resources
   - workspace admin cannot mutate another workspace

3. Project isolation
   - project operator can request plan only for assigned project
   - project operator cannot access another project artifacts

4. Environment isolation
   - environment viewer can read environment status
   - environment viewer cannot mutate locks or approvals

5. Runner isolation
   - runner in pool A cannot claim jobs from pool B
   - dev runner cannot claim prod jobs
   - runner token cannot access normal user APIs

6. Artifact isolation
   - artifact download requires scoped permission
   - artifact deletion requires scoped permission
   - signed URL generation requires scoped permission

7. Audit isolation
   - auditor can read scoped audit events
   - auditor cannot mutate resources

Requirements:
- Use fixture data with at least two organizations.
- Use table-driven tests.
- Every API endpoint must have at least one cross-tenant denial test.
- Add helper to assert 403 for cross-tenant access.
- Add coverage report section in docs.

Rules:
- No real cloud resources.
- No external network.
- No secrets.
- Tests must be deterministic.

Run:
- cd control-plane && go test ./...
```
