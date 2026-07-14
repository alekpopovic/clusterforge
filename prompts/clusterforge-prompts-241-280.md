# ClusterForge prompts 241–280

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


---

# Prompt 242 — SaaS-readiness architecture RFC

```text
Create the SaaS-readiness architecture RFC for ClusterForge.

Goal:
Assess what would be required to evolve the self-hosted Control Plane into a SaaS-capable architecture later, without implementing SaaS now.

Create:
- docs/rfcs/025-saas-readiness.md
- docs/control-plane/saas-readiness.md

Cover:

1. Goals
   - strong tenant isolation
   - scalable API
   - scalable runners
   - usage metering
   - organization onboarding
   - audit evidence export
   - regional deployment
   - secure artifact storage
   - customer data deletion
   - data residency model

2. Non-goals
   - public SaaS launch
   - billing implementation
   - marketplace
   - storing customer cloud credentials
   - running untrusted arbitrary code
   - automatic remediation by default

3. Tenant model
   - organization
   - workspace
   - project
   - environment
   - runner pool
   - artifact namespace
   - audit namespace

4. Trust boundaries
   - user browser
   - CLI
   - API
   - database
   - artifact storage
   - runner
   - Git provider
   - cloud provider
   - Kubernetes clusters

5. Data classification
   - public docs
   - non-sensitive metadata
   - sensitive metadata
   - audit events
   - plan summaries
   - raw plan files
   - logs
   - tokens
   - secrets, which must not be stored

6. SaaS blockers
   - multi-tenant authorization bugs
   - runner isolation
   - artifact sensitivity
   - customer-managed credentials
   - data deletion
   - regional storage
   - operational monitoring
   - incident response
   - compliance evidence

7. Recommended path
   Phase 1:
   - self-hosted enterprise
   Phase 2:
   - single-tenant managed
   Phase 3:
   - limited multi-tenant private SaaS
   Phase 4:
   - broader SaaS only after security review

8. Required future features
   - tenant isolation test suite
   - per-tenant quotas
   - usage metering
   - SCIM
   - SSO hardening
   - artifact encryption
   - immutable audit log
   - data residency
   - customer deletion workflow
   - abuse prevention
   - support tooling

Include:
- Mermaid deployment diagram
- Mermaid trust boundary diagram
- risk matrix
- migration considerations

Do not implement code in this prompt.
Update:
- ROADMAP_V0.7.md if it exists
- docs/roadmap.md
```


---

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


---

# Prompt 244 — Tenant-scoped data access layer

```text
Harden the Control Plane data access layer for tenant scoping.

Goal:
Make tenant scoping difficult to forget and easy to test.

Tasks:
1. Introduce a scoped query context:
   - organization_id
   - workspace_id optional
   - project_id optional
   - environment_id optional
   - actor_id
   - permissions

2. Update repositories/services to require scoped context for:
   - projects
   - environments
   - clusters
   - apps
   - artifacts
   - audit events
   - jobs
   - plan requests
   - apply requests
   - approvals
   - policy results
   - drift results
   - cost reports
   - service catalog
   - runbooks

3. Add guardrails:
   - repository methods must not expose unscoped list methods except admin/system methods
   - system methods must be clearly named and audited
   - tests should fail if unscoped repository methods are used in API handlers

4. Add middleware:
   - resolves actor scope
   - attaches scope to request context
   - denies missing organization scope for tenant APIs

5. Add static/conformance check:
   - search for repository calls without scope
   - warn or fail in CI

Docs:
- docs/control-plane/scoped-data-access.md

Tests:
- unscoped access denied
- scoped list returns only allowed data
- admin/system access audited
- API handlers require scope

Rules:
- Preserve local development mode.
- Do not weaken RBAC.
- Do not rely on frontend filtering for security.
- Deny by default.

Run:
- cd control-plane && gofmt -w .
- cd control-plane && go test ./...
```


---

# Prompt 245 — API rate limiting and tenant quotas

```text
Implement API rate limiting and tenant quotas.

Goal:
Protect the Control Plane from accidental overload and future abuse.

Rate limits:
- per IP
- per actor
- per organization
- per token/service account
- per endpoint class

Endpoint classes:
- read
- write
- auth
- job_polling
- artifact_download
- artifact_upload
- webhook

Config:
rate_limits:
  enabled: true
  defaults:
    read_per_minute: 600
    write_per_minute: 120
    auth_per_minute: 30
    artifact_download_per_minute: 60
    job_polling_per_minute: 300

Quotas:
- max projects per organization
- max environments per organization
- max runners per organization
- max artifacts storage bytes
- max active jobs
- max preview environments
- max API tokens
- max service accounts

API:
- GET /api/v1/quotas
- GET /api/v1/rate-limits
- GET /api/v1/usage/current

CLI:
- cf quota show
- cf quota check
- cf rate-limit show

Behavior:
- return 429 for rate limit exceeded
- include retry-after header
- quota violations return clear JSON error
- audit quota-related denials if useful
- metrics for rate limit/quota denials

Tests:
- rate limit by token
- rate limit by IP
- quota exceeded for artifacts
- quota exceeded for runners
- 429 includes retry-after
- admin can view quotas

Docs:
- docs/control-plane/quotas-rate-limits.md

Rules:
- Keep defaults generous for self-hosted.
- Allow disabling in local dev.
- Do not break runner polling with too-low defaults.
```


---

# Prompt 246 — Usage metering model

```text
Create usage metering model for ClusterForge.

Goal:
Track operational usage for capacity planning, enterprise reporting, and possible future billing, without implementing billing.

Create:
- docs/rfcs/026-usage-metering.md
- docs/control-plane/usage-metering.md

Metrics to meter:
- organizations
- workspaces
- projects
- environments
- clusters
- apps
- runners
- jobs by type
- plan requests
- apply requests
- policy checks
- drift checks
- cost scans
- artifact storage bytes
- artifact downloads
- audit event volume
- preview environments
- API requests
- active users

Metering principles:
- no secret values
- no raw Terraform state
- no raw plan content
- tenant-scoped
- exportable
- retention controlled
- disabled or reduced mode for privacy-sensitive installs

Data model:
- usage_events
- usage_rollups_daily
- usage_rollups_monthly

Dimensions:
- organization_id
- workspace_id
- project_id
- environment_id
- event_type
- quantity
- metadata_json sanitized

Do not implement code in this prompt.

Include:
- data retention policy
- privacy considerations
- self-hosted reporting
- future SaaS billing note
- examples of usage reports

Update:
- ROADMAP_V0.7.md if it exists
```


---

# Prompt 247 — Usage metering implementation MVP

```text
Implement usage metering MVP.

Goal:
Record and report non-sensitive usage events.

Database:
- usage_events
  - id
  - organization_id
  - workspace_id nullable
  - project_id nullable
  - environment_id nullable
  - event_type
  - quantity
  - metadata_json
  - created_at

- usage_rollups_daily
  - id
  - organization_id
  - date
  - event_type
  - quantity
  - metadata_json
  - created_at

Events:
- api.request
- job.created
- job.completed
- plan.requested
- plan.completed
- apply.requested
- apply.completed
- policy.checked
- drift.checked
- cost.scanned
- artifact.uploaded
- artifact.downloaded
- preview.created
- preview.deleted
- runner.heartbeat

API:
- GET /api/v1/usage/events
- GET /api/v1/usage/summary
- POST /api/v1/usage/rollup/run

CLI:
- cf usage summary
- cf usage export --format json|csv

Dashboard:
- optional usage summary page

Requirements:
- tenant-scoped
- RBAC protected
- no secrets
- metadata sanitized
- retention controlled
- rollup job idempotent

Tests:
- usage event recorded for plan request
- usage event recorded for artifact upload
- summary aggregates by event type
- cross-tenant usage denied
- metadata redaction works

Docs:
- docs/control-plane/usage-metering.md

Run:
- cd control-plane && go test ./...
- cd cli && go test ./...
```


---

# Prompt 248 — Billing integration RFC

```text
Create billing integration RFC for future SaaS or managed offering.

Goal:
Design billing boundaries without implementing payments.

Create:
- docs/rfcs/027-billing-integration.md
- docs/control-plane/billing.md

Cover:

1. Non-goals for current project
   - no payment processor implementation
   - no invoicing
   - no subscription enforcement
   - no public SaaS launch

2. Possible future billing dimensions
   - active projects
   - active environments
   - active runners
   - job volume
   - artifact storage
   - audit retention
   - preview environments
   - enterprise features

3. Usage metering dependency
   - usage_events
   - daily rollups
   - monthly rollups

4. Billing account model
   - organization billing account
   - plan tier
   - quota mapping
   - trial mode
   - suspended mode

5. Safety
   - never block emergency access unexpectedly
   - support grace periods
   - read-only mode instead of destructive disablement
   - clear admin notifications

6. Data privacy
   - no secrets in usage data
   - minimal metadata
   - exportable reports
   - retention settings

7. Future APIs
   - GET /api/v1/billing/account
   - GET /api/v1/billing/usage
   - GET /api/v1/billing/invoices
   - POST /api/v1/billing/portal

Do not implement code.
Update roadmap and backlog.
```


---

# Prompt 249 — Organization onboarding workflow

```text
Implement organization onboarding workflow.

Goal:
Make it easy to bootstrap a new organization/workspace/project safely.

Control Plane:
- onboarding state per organization
- setup checklist

Checklist:
- organization created
- admin user configured
- workspace created
- first project created
- RBAC bindings configured
- runner registered
- artifact storage configured
- notification channel configured
- first inventory sync completed
- first policy check completed
- backup guidance acknowledged

API:
- GET /api/v1/onboarding
- POST /api/v1/onboarding/steps/{step}/complete
- POST /api/v1/onboarding/reset

CLI:
- cf org init
- cf onboarding status
- cf onboarding complete <step>

Dashboard:
- onboarding checklist page
- setup progress

Docs:
- docs/control-plane/onboarding.md

Tests:
- new organization has checklist
- completing step updates status
- onboarding status scoped by organization
- unauthorized user cannot complete admin steps

Rules:
- Do not require SaaS.
- Must work in self-hosted mode.
- No secrets in onboarding data.
```


---

# Prompt 250 — Organization offboarding and data deletion workflow

```text
Design and implement organization offboarding workflow.

Goal:
Allow safe deactivation and deletion of organization data.

Create docs:
- docs/control-plane/offboarding.md
- docs/rfcs/028-data-deletion.md

Control Plane states:
- active
- suspended
- deletion_requested
- deletion_in_progress
- deleted

API:
- POST /api/v1/organizations/{id}/suspend
- POST /api/v1/organizations/{id}/deletion-request
- POST /api/v1/organizations/{id}/deletion-confirm
- GET /api/v1/organizations/{id}/deletion-status

CLI:
- cf org suspend <org>
- cf org deletion request <org>
- cf org deletion confirm <org>
- cf org deletion status <org>

Deletion scope:
- projects
- environments
- clusters metadata
- apps metadata
- jobs
- artifacts
- policy results
- drift results
- cost reports
- usage events
- service catalog
- runbooks imported into DB

Retention exceptions:
- audit events may be retained or anonymized based on policy
- legal hold blocks deletion
- backups have separate retention policy

Safety:
- two-step confirmation
- admin-only
- audit event created
- export before deletion option
- dry-run deletion report
- no cloud resource deletion by default

Tests:
- suspend blocks writes
- deletion dry-run reports resources
- deletion request requires admin
- legal hold blocks deletion
- deleted org inaccessible
- audit event created

Rules:
- Do not delete customer cloud infrastructure.
- Do not delete Git repositories.
- Do not delete Terraform state.
- This deletes Control Plane metadata only unless explicitly extended.
```


---

# Prompt 251 — Customer-managed encryption keys RFC

```text
Create customer-managed encryption keys RFC.

Goal:
Design how ClusterForge could support organization-specific encryption keys for sensitive metadata and artifacts.

Create:
- docs/rfcs/029-customer-managed-keys.md
- docs/control-plane/customer-managed-keys.md

Cover:

1. Goals
   - per-organization encryption boundary
   - artifact encryption
   - token metadata encryption where useful
   - audit metadata protection
   - future SaaS readiness
   - key rotation

2. Non-goals
   - storing cloud provider root credentials
   - implementing every KMS provider immediately
   - encrypting Terraform state inside Control Plane, because state should not be stored there

3. Key providers
   - local development key
   - AWS KMS
   - Azure Key Vault future
   - GCP KMS future
   - HashiCorp Vault future

4. Data to encrypt
   - artifact payloads
   - sensitive metadata fields
   - webhook URLs if stored
   - notification secrets should use env references where possible
   - tokens are hashed, not encrypted

5. Data not to store
   - cloud credentials
   - secret values
   - kubeconfigs
   - Terraform state

6. Key hierarchy
   - platform master key
   - organization data key
   - artifact data key
   - envelope encryption

7. Rotation
   - create new key version
   - re-encrypt new artifacts
   - background re-encryption for old artifacts
   - key disabled/deleted behavior

8. Failure modes
   - key unavailable
   - access denied
   - key deleted
   - rotation partially complete

Do not implement code in this prompt.
Update roadmap.
```


---

# Prompt 252 — KMS-backed artifact encryption implementation

```text
Implement KMS-backed artifact encryption for AWS KMS as the first cloud KMS provider.

Goal:
Encrypt artifact payloads using envelope encryption backed by AWS KMS.

Config:
artifacts:
  encryption:
    provider: aws_kms
    aws_kms:
      key_id_env: CLUSTERFORGE_ARTIFACT_KMS_KEY_ID
      region: eu-central-1

Behavior:
- generate data key using KMS
- encrypt artifact payload locally with data key
- store encrypted data key with artifact metadata
- never store plaintext data key
- decrypt only for authorized download
- audit decrypt/download
- checksum encrypted and plaintext where appropriate

Database:
- artifacts table adds:
  - encryption_provider
  - encrypted_data_key
  - encryption_key_id
  - encryption_context_json

Security:
- encryption context includes organization_id and artifact_id
- RBAC required before decrypt
- downloads audited
- raw plan upload remains disabled by default

Tests:
- use mock KMS interface
- encrypt/decrypt round trip
- wrong encryption context fails
- missing key config fails
- artifact content not stored plaintext
- unauthorized download denied before decrypt

Docs:
- docs/control-plane/artifact-encryption-kms.md

Terraform:
- optional example using modules/cloud/aws/kms-key

Rules:
- Do not require AWS KMS for local dev.
- Keep local encryption provider supported.
- Do not log data keys.
```


---

# Prompt 253 — Vault integration RFC

```text
Create HashiCorp Vault integration RFC.

Goal:
Define how ClusterForge can integrate with Vault without storing secrets itself.

Create:
- docs/rfcs/030-vault-integration.md
- docs/control-plane/vault-integration.md

Use cases:
- runner obtains short-lived cloud credentials
- Control Plane stores no cloud credentials
- dynamic database credentials for Control Plane
- artifact encryption key access
- secret reference inventory
- workload secret references
- token signing or key wrapping in future

Non-goals:
- replacing External Secrets Operator
- copying secret values into ClusterForge
- making Vault mandatory
- secret editing UI

Design:
1. Vault auth methods
   - Kubernetes auth
   - AppRole
   - OIDC/JWT
   - cloud IAM auth

2. Runner integration
   - runner authenticates to Vault
   - runner retrieves short-lived credentials
   - credentials exist only in job workspace environment
   - redaction enforced

3. Control Plane integration
   - only stores Vault paths/references
   - never reads secret values unless explicitly needed for its own runtime
   - env references preferred

4. Policy
   - enforce no plaintext credentials
   - require secret references for prod

5. Future CLI:
   - cf vault check
   - cf secrets references
   - cf runner vault test

Do not implement code.
Update roadmap and security docs.
```


---

# Prompt 254 — Secret reference broker MVP

```text
Implement secret reference broker MVP.

Goal:
Centralize secret references and validation without reading secret values.

Control Plane:
- secret_references table:
  - id
  - organization_id
  - project_id
  - environment_id nullable
  - app_id nullable
  - provider
  - reference
  - key_name nullable
  - usage_type
  - source_path
  - metadata_json
  - created_at
  - updated_at

Providers:
- kubernetes_secret
- aws_secrets_manager
- aws_ssm_parameter
- external_secrets
- vault
- ecs_secret

API:
- GET /api/v1/secret-references
- POST /api/v1/secret-references/import
- GET /api/v1/secret-references/{id}

CLI:
- cf secrets inventory
- cf secrets sync
- cf secrets validate-references
- cf secrets report --format markdown|json

Validation:
- reference format validation only
- no secret value reads
- optional provider existence check only when explicitly enabled

Dashboard:
- secret references page
- app secret references
- environment secret references

Tests:
- import from app manifests
- import ECS value_from
- import ExternalSecret manifest
- Vault path reference
- no values stored
- cross-tenant access denied

Docs:
- docs/control-plane/secret-reference-broker.md

Rules:
- Never read secret values.
- Never display secret values.
- This is inventory and validation only.
```


---

# Prompt 255 — Kubernetes Job runner executor

```text
Implement Kubernetes Job executor for ClusterForge runner.

Goal:
Allow runner jobs to execute as isolated Kubernetes Jobs instead of long-running process execution.

Architecture:
- runner controller receives job from Control Plane
- creates Kubernetes Job per ClusterForge job
- Job runs a runner-worker container
- worker executes validate/plan/policy/drift/cost job
- worker uploads result/artifacts
- Kubernetes Job is deleted or retained based on policy

Config:
runner:
  executor: kubernetes_job
  kubernetes_job:
    namespace: clusterforge-runners
    service_account_name: clusterforge-runner-worker
    image: ghcr.io/example/clusterforge-runner-worker:latest
    cleanup_finished_jobs: true
    ttl_seconds_after_finished: 3600
    resources:
      requests:
        cpu: 250m
        memory: 512Mi
      limits:
        cpu: 1000m
        memory: 2Gi

Features:
- per-job workspace emptyDir
- config mounted from Secret/ConfigMap
- token from Secret
- resource limits
- node selector/tolerations
- labels for tenant/job/environment
- no privileged containers
- run as non-root

Tests:
- job manifest rendering
- labels included
- secrets referenced by name only
- unsupported apply blocked unless enabled
- cleanup policy

Docs:
- docs/control-plane/kubernetes-job-executor.md

Rules:
- Do not include real tokens.
- Do not grant cluster-admin.
- Apply disabled by default.
```


---

# Prompt 256 — Ephemeral runner pools

```text
Implement ephemeral runner pools design and MVP.

Goal:
Support temporary runners for isolated job execution.

Use cases:
- high-risk prod plan
- short-lived preview environment jobs
- per-tenant isolation
- air-gapped job execution
- burst capacity

Control Plane:
- runner_pool type:
  - persistent
  - ephemeral
- runner lifecycle:
  - requested
  - starting
  - active
  - draining
  - terminated
  - failed

API:
- POST /api/v1/runner-pools/{id}/scale-up
- POST /api/v1/runner-pools/{id}/drain
- GET /api/v1/runner-pools/{id}/capacity

Runner:
- supports drain mode
- stops claiming new jobs
- finishes active jobs
- exits when idle if ephemeral

CLI:
- cf runner pool scale-up <pool>
- cf runner pool drain <pool>
- cf runner drain <runner>

Docs:
- docs/control-plane/ephemeral-runners.md

Tests:
- draining runner stops claiming jobs
- active job can finish
- ephemeral runner exits when idle
- pool capacity reported

Rules:
- Do not auto-scale cloud resources yet unless explicit.
- No apply on untrusted ephemeral runners.
- Token expiration should be short for ephemeral runners.
```


---

# Prompt 257 — Runner autoscaling RFC

```text
Create runner autoscaling RFC.

Goal:
Design how ClusterForge could scale runners based on queued jobs.

Create:
- docs/rfcs/031-runner-autoscaling.md
- docs/control-plane/runner-autoscaling.md

Autoscaling signals:
- queued jobs
- queue wait time
- jobs by type
- environment priority
- runner pool labels
- max concurrent jobs
- failure rate
- cost constraints

Execution environments:
1. Kubernetes Deployment scale
2. Kubernetes Job-per-runner
3. VM auto scaling group
4. CI runner integration
5. manual scaling

Control Plane responsibilities:
- expose queue metrics
- define desired runner capacity
- enforce pool limits
- avoid dev runner processing prod jobs

Non-goals:
- building a full Kubernetes autoscaler
- cloud-specific autoscaling implementation in first phase
- automatic apply scaling without safety controls

Future commands:
- cf runner autoscaling status
- cf runner autoscaling configure
- cf runner autoscaling simulate

Do not implement code.
Update roadmap.
```


---

# Prompt 258 — Job fairness and tenant quotas

```text
Implement job fairness and tenant-level job quotas.

Goal:
Prevent one organization/project from starving others.

Features:
- per-organization active job limit
- per-project active job limit
- per-environment active job limit
- per-job-type limit
- priority classes:
  - low
  - normal
  - high
  - emergency
- fair scheduling across organizations
- FIFO within same priority/scope where practical

Config:
job_scheduling:
  max_active_jobs_per_org: 20
  max_active_jobs_per_project: 10
  max_active_jobs_per_environment: 3
  max_active_apply_jobs_per_environment: 1
  default_priority: normal

Control Plane:
- scheduler picks eligible job based on:
  - priority
  - quotas
  - runner pool
  - required labels
  - lease availability

API:
- GET /api/v1/job-scheduling/status
- GET /api/v1/job-scheduling/queues

CLI:
- cf job queue
- cf job priority set <job-id> --priority high

Tests:
- org active job limit enforced
- environment apply concurrency is 1
- high priority job scheduled first
- fair scheduling across orgs
- runner labels still enforced

Docs:
- docs/control-plane/job-scheduling.md

Rules:
- Emergency priority requires admin.
- Apply jobs remain conservative.
```


---

# Prompt 259 — Advanced approval policies

```text
Implement advanced approval policies.

Goal:
Allow organizations to define approval requirements based on risk, environment, resource changes, and policy results.

Policy examples:
approval_policies:
  prod:
    default_required_approvals: 1
    two_person_rule: true
    require_approval_for_destroy: true
    require_approval_for_replacements: true
    require_security_approval_for_public_ingress: true
    require_database_owner_approval_for_rds_changes: true
    require_platform_approval_for_cluster_changes: true

Approval dimensions:
- environment
- stack
- resource type
- risk level
- destroy operations
- replacement operations
- policy severity
- cost warning severity
- incident mode
- break-glass use

Roles:
- approver
- security_approver
- platform_approver
- database_approver
- incident_commander

Behavior:
- approval request includes required approval types
- approval cannot be completed until all required approvals exist
- self-approval blocked when two-person rule is enabled
- approval expires after configurable time

API:
- GET /api/v1/approval-policies
- POST /api/v1/approval-policies
- GET /api/v1/apply-requests/{id}/approval-requirements

CLI:
- cf approval requirements <apply-id>
- cf approval policy list
- cf approval policy validate

Tests:
- destroy requires extra approval
- database change requires database approver
- self-approval blocked
- expired approval invalid
- all requirements satisfied allows apply

Docs:
- docs/control-plane/advanced-approvals.md

Rules:
- Deny by default when approval requirements cannot be evaluated.
- Do not weaken existing prod approval behavior.
```


---

# Prompt 260 — Policy exceptions and waivers

```text
Implement policy exception and waiver workflow.

Goal:
Allow controlled temporary exceptions for policy findings.

Use cases:
- temporary use of public ingress
- temporary latest tag in dev
- known security exception
- migration period
- legacy module exception

Database:
- policy_exceptions
  - id
  - organization_id
  - project_id nullable
  - environment_id nullable
  - policy_id
  - scope_type
  - scope_id
  - reason
  - created_by
  - approved_by nullable
  - expires_at
  - status
  - created_at
  - updated_at

Statuses:
- requested
- approved
- rejected
- expired
- revoked

API:
- POST /api/v1/policy-exceptions
- GET /api/v1/policy-exceptions
- POST /api/v1/policy-exceptions/{id}/approve
- POST /api/v1/policy-exceptions/{id}/reject
- POST /api/v1/policy-exceptions/{id}/revoke

CLI:
- cf policy exception request --policy <id> --reason "..."
- cf policy exception list
- cf policy exception approve <id>
- cf policy exception revoke <id>

Behavior:
- exceptions are time-limited
- high severity exceptions require approval
- expired exceptions ignored
- policy results show waived status
- all exceptions audited
- dashboard displays active exceptions

Tests:
- exception waives policy
- expired exception does not waive
- high severity requires approval
- exception scoped to correct environment only
- audit events created

Docs:
- docs/control-plane/policy-exceptions.md

Rules:
- No permanent exceptions by default.
- Reason required.
- Approval required for prod/blocking policies.
```


---

# Prompt 261 — Risk acceptance workflow

```text
Implement risk acceptance workflow.

Goal:
Allow teams to explicitly accept operational/security risk with audit trail.

Difference from policy exception:
- policy exception waives a specific policy result
- risk acceptance records acknowledged risk for a defined scope and time

Database:
- risk_acceptances
  - id
  - organization_id
  - project_id nullable
  - environment_id nullable
  - title
  - description
  - risk_level
  - scope_type
  - scope_id
  - owner
  - created_by
  - approved_by
  - expires_at
  - status
  - mitigation_plan
  - metadata_json
  - created_at
  - updated_at

API:
- POST /api/v1/risk-acceptances
- GET /api/v1/risk-acceptances
- GET /api/v1/risk-acceptances/{id}
- POST /api/v1/risk-acceptances/{id}/approve
- POST /api/v1/risk-acceptances/{id}/revoke

CLI:
- cf risk accept
- cf risk list
- cf risk show <id>
- cf risk approve <id>
- cf risk revoke <id>
- cf risk report --format markdown|json

Dashboard:
- risk register page

Tests:
- create risk acceptance
- approval required for high risk
- expired risk shown as expired
- risk report includes active risks
- cross-tenant denied

Docs:
- docs/control-plane/risk-acceptance.md

Rules:
- Do not use risk acceptance to bypass destructive safety automatically.
- High risk requires approval.
- Expiration required by default.
```


---

# Prompt 262 — Change Advisory Board workflow

```text
Implement Change Advisory Board workflow support.

Goal:
Support formal change review for production environments without forcing it on all users.

Database:
- change_windows
- change_requests
- change_request_reviews

Change request fields:
- id
- organization_id
- project_id
- environment_id
- title
- description
- requested_by
- planned_start
- planned_end
- risk_level
- status
- related_plan_request_id
- related_apply_request_id
- metadata_json

Statuses:
- draft
- submitted
- approved
- rejected
- scheduled
- implemented
- canceled

API:
- POST /api/v1/change-requests
- GET /api/v1/change-requests
- GET /api/v1/change-requests/{id}
- POST /api/v1/change-requests/{id}/submit
- POST /api/v1/change-requests/{id}/approve
- POST /api/v1/change-requests/{id}/reject

CLI:
- cf change-request create
- cf change-request submit <id>
- cf change-request approve <id>
- cf change-request list
- cf change-request show <id>

Integration:
- prod apply can require approved change request
- apply request links to change request
- deployment windows checked
- approvals audited

Tests:
- prod apply blocked without change request when policy enabled
- approved change request allows apply request
- expired window blocks apply
- audit events created

Docs:
- docs/control-plane/change-advisory.md

Rules:
- Optional feature.
- Do not make CAB mandatory for all environments.
```


---

# Prompt 263 — Incident management integrations

```text
Add incident management integration support.

Goal:
Send or create incident notifications in external systems through generic webhook first.

Supported MVP:
- generic webhook
- Slack/Teams notification reuse
- Pager-style integration documented through webhook, not provider-specific API unless already implemented

Events:
- incident.started
- incident.resolved
- break_glass.requested
- break_glass.used
- apply.failed
- drift.detected.prod
- runner.offline.prod
- policy.blocked.prod

Config:
incident_integrations:
  enabled: true
  sinks:
    - name: incident-webhook
      type: webhook
      url_env: CLUSTERFORGE_INCIDENT_WEBHOOK
      events:
        - incident.started
        - apply.failed

API:
- GET /api/v1/incident-integrations
- POST /api/v1/incident-integrations/test

CLI:
- cf incident integration test <name>

Tests:
- webhook payload rendered
- URL redacted
- disabled integration skipped
- failure recorded
- incident event triggers integration

Docs:
- docs/control-plane/incident-integrations.md

Rules:
- Do not store webhook URL plaintext if avoidable.
- Do not call real external services in tests.
- Do not make incidents auto-resolve.
```


---

# Prompt 264 — Compliance evidence collection

```text
Implement compliance evidence collection.

Goal:
Collect non-sensitive evidence artifacts showing how ClusterForge controls are configured and operating.

Evidence types:
- policy check results
- approval records
- change records
- audit event exports
- RBAC bindings
- token rotation status
- artifact retention config
- backup evidence files
- DR drill evidence
- module conformance reports
- security scan summaries
- release gate reports
- smoke test matrix
- compliance mapping reports

Database:
- evidence_records
  - id
  - organization_id
  - project_id nullable
  - environment_id nullable
  - evidence_type
  - title
  - description
  - source_type
  - source_id nullable
  - artifact_id nullable
  - created_by
  - created_at
  - metadata_json

API:
- GET /api/v1/evidence
- POST /api/v1/evidence/collect
- GET /api/v1/evidence/{id}

CLI:
- cf evidence collect
- cf evidence list
- cf evidence show <id>
- cf evidence export --format markdown|json

Dashboard:
- evidence page

Tests:
- collect policy evidence
- collect approval evidence
- evidence export
- RBAC protects evidence
- no secrets in evidence metadata

Docs:
- docs/control-plane/evidence-collection.md

Rules:
- Do not claim compliance certification.
- Evidence is support material only.
- Redact sensitive data.
```


---

# Prompt 265 — Evidence bundle export

```text
Implement evidence bundle export.

Goal:
Allow teams to export a review-ready evidence package for auditors or internal reviews.

CLI:
- cf evidence bundle create
- cf evidence bundle create --environment prod
- cf evidence bundle create --framework soc2
- cf evidence bundle verify <bundle.zip>

Bundle contents:
- manifest.json
- README.md
- policy-results/
- approvals/
- change-records/
- audit-events/
- rbac/
- backup-evidence/
- dr-evidence/
- release-gates/
- module-conformance/
- compliance-mapping/
- checksums.txt

Control Plane:
- optional API:
  - POST /api/v1/evidence/bundles
  - GET /api/v1/evidence/bundles/{id}

Security:
- no secrets
- no raw Terraform state
- no raw plan files by default
- redaction pass before bundling
- checksums for files
- optional encryption documented

Tests:
- bundle created
- manifest includes files
- checksum verification
- redaction
- unknown framework fails clearly

Docs:
- docs/control-plane/evidence-bundles.md

Rules:
- Do not claim the bundle proves compliance.
- Bundle is evidence support only.
```


---

# Prompt 266 — Immutable audit log RFC

```text
Create immutable audit log RFC.

Goal:
Design tamper-evident audit storage for higher assurance environments.

Create:
- docs/rfcs/032-immutable-audit-log.md
- docs/control-plane/immutable-audit.md

Cover:
1. Goals
   - tamper evidence
   - append-only audit events
   - export to WORM-capable storage
   - cryptographic hash chain
   - periodic checkpoints
   - evidence bundle support

2. Non-goals
   - legal compliance guarantee
   - blockchain
   - replacing external SIEM
   - making audit events undeletable in dev

3. Design options
   - database append-only table
   - hash chain per organization
   - object storage export
   - external SIEM/webhook
   - WORM storage integration later

4. Data model
   - audit_event_hash
   - previous_hash
   - sequence_number
   - checkpoint records

5. Verification
   - cf audit verify
   - detect missing/modified events
   - export verification report

6. Risks
   - clock skew
   - partial writes
   - migration complexity
   - retention vs immutability conflict
   - legal hold

Do not implement code.
Update security roadmap.
```


---

# Prompt 267 — Audit hash chain implementation MVP

```text
Implement audit hash chain MVP.

Goal:
Make audit events tamper-evident within the database.

Database changes:
- audit_events:
  - sequence_number
  - previous_hash
  - event_hash
  - hash_algorithm

Behavior:
- event_hash computed from canonical event fields
- previous_hash links to prior event in same organization
- sequence_number increments per organization
- audit event updates/deletes remain disallowed
- migration initializes hash chain for existing events if possible

API:
- GET /api/v1/audit-events/verify
- GET /api/v1/audit-events/checkpoints

CLI:
- cf audit verify
- cf audit checkpoint create
- cf audit checkpoint list

Tests:
- hash chain valid
- modified event fails verification
- missing event fails verification
- per-organization chains separate
- checkpoint created

Docs:
- docs/control-plane/immutable-audit.md

Rules:
- This is tamper-evident, not tamper-proof.
- Do not claim legal WORM compliance.
- Do not allow audit event mutation.
```


---

# Prompt 268 — SCIM provisioning RFC

```text
Create SCIM provisioning RFC.

Goal:
Design automated user and group provisioning from enterprise identity providers.

Create:
- docs/rfcs/033-scim-provisioning.md
- docs/control-plane/scim.md

Cover:
1. Goals
   - user provisioning
   - user deprovisioning
   - group sync
   - group membership sync
   - RBAC group mapping
   - audit events

2. Non-goals
   - replacing OIDC login
   - storing identity provider credentials
   - supporting every SCIM edge case in MVP

3. Resources
   - Users
   - Groups
   - Group memberships

4. API endpoints
   - /scim/v2/Users
   - /scim/v2/Groups

5. Security
   - SCIM bearer token
   - token rotation
   - scoped to organization
   - audit all provisioning changes
   - deny login for deactivated users

6. Deprovisioning
   - deactivate user
   - remove group memberships
   - retain audit events
   - transfer ownership guidance

7. Testing
   - SCIM contract tests
   - create/update/delete user
   - group membership tests

Do not implement code in this prompt.
Update roadmap.
```


---

# Prompt 269 — SCIM provisioning MVP

```text
Implement SCIM provisioning MVP.

Goal:
Support basic enterprise user and group provisioning.

Endpoints:
- GET /scim/v2/ServiceProviderConfig
- GET /scim/v2/ResourceTypes
- GET /scim/v2/Schemas
- GET /scim/v2/Users
- POST /scim/v2/Users
- GET /scim/v2/Users/{id}
- PATCH /scim/v2/Users/{id}
- DELETE /scim/v2/Users/{id}
- GET /scim/v2/Groups
- POST /scim/v2/Groups
- GET /scim/v2/Groups/{id}
- PATCH /scim/v2/Groups/{id}
- DELETE /scim/v2/Groups/{id}

Features:
- create user
- deactivate user
- update user email/name
- create group
- update group memberships
- delete/deactivate group
- map SCIM groups to ClusterForge groups

Security:
- SCIM token per organization
- token stored hashed
- token expiration
- audit all changes
- no admin by default from SCIM unless mapped explicitly

Tests:
- create user
- update user
- deactivate user
- create group
- add member
- remove member
- invalid token rejected
- deactivated user cannot authenticate

Docs:
- docs/control-plane/scim.md

Rules:
- Keep implementation minimal and interoperable.
- Do not overclaim full SCIM compatibility if partial.
```


---

# Prompt 270 — Data residency model

```text
Create data residency model for ClusterForge.

Goal:
Define how self-hosted and future managed deployments can control where metadata and artifacts are stored.

Create:
- docs/rfcs/034-data-residency.md
- docs/control-plane/data-residency.md

Cover:

1. Data categories
   - organization metadata
   - project metadata
   - environment metadata
   - policy results
   - drift results
   - cost reports
   - artifacts
   - audit events
   - usage events
   - tokens
   - runner metadata

2. Residency controls
   - deployment region
   - database region
   - artifact bucket region
   - runner region
   - backup region
   - logs/metrics region

3. Organization-level settings
   - preferred_region
   - allowed_regions
   - artifact_region
   - backup_region

4. Enforcement
   - API refuses artifact storage outside allowed region
   - runner pool must match allowed region for regulated environments
   - backups respect residency settings

5. Non-goals
   - legal advice
   - automatic compliance certification
   - public SaaS guarantees

6. Future commands
   - cf residency show
   - cf residency validate
   - cf residency report

Do not implement code unless simple validation is already available.
Update roadmap.
```


---

# Prompt 271 — Regional Control Plane deployment patterns

```text
Create regional Control Plane deployment patterns.

Goal:
Document deployment topologies for single-region, multi-region, and disaster recovery.

Create:
- docs/control-plane/regional-deployment.md

Topologies:
1. Single-region self-hosted
   - simplest
   - one database
   - one artifact backend
   - one runner pool

2. Single-region HA
   - multiple API replicas
   - HA database
   - S3-compatible artifact storage
   - multiple runners

3. Multi-region active/passive
   - primary Control Plane
   - standby region
   - database backup/restore
   - artifact replication
   - runner pools per region

4. Multi-region active/active
   - not recommended for MVP
   - requires strong consistency decisions
   - tenant routing
   - data residency concerns

5. Per-region runner pools with central API
   - common enterprise pattern
   - low complexity
   - good separation

Include:
- diagrams
- pros/cons
- failure modes
- backup implications
- data residency implications
- recommended v0.7 pattern

No code changes required.
```


---

# Prompt 272 — Cluster API integration RFC

```text
Create Cluster API integration RFC.

Goal:
Evaluate whether ClusterForge should support Kubernetes Cluster API as a cluster lifecycle backend.

Create:
- docs/rfcs/035-cluster-api-integration.md
- docs/cluster-api.md

Cover:

1. Goals
   - Kubernetes-native cluster lifecycle
   - self-service cluster templates
   - multi-cloud cluster creation
   - GitOps-friendly cluster management
   - standardized cluster blueprints

2. Non-goals
   - replacing Terraform modules immediately
   - supporting every Cluster API provider
   - making Cluster API mandatory

3. Target providers
   - AWS provider
   - Azure provider
   - GCP provider
   - Docker/local provider for testing
   - others later

4. Architecture options
   - Terraform provisions management cluster
   - Cluster API manages workload clusters
   - ClusterForge renders Cluster API manifests
   - GitOps applies manifests
   - Control Plane tracks inventory/status

5. Terraform vs Cluster API boundary
   - Terraform for foundational cloud resources
   - Cluster API for cluster lifecycle
   - GitOps for reconciliation
   - Control Plane for visibility/governance

6. CLI future
   - cf cluster-template list
   - cf cluster create --template eks-small
   - cf cluster delete
   - cf cluster status

7. Risks
   - management cluster dependency
   - provider maturity differences
   - IAM complexity
   - debugging complexity
   - lifecycle ownership conflicts

Do not implement code.
Update roadmap.
```


---

# Prompt 273 — Cluster lifecycle blueprint model

```text
Implement cluster lifecycle blueprint model.

Goal:
Define reusable cluster templates independent of specific implementation backend.

Blueprint types:
- terraform_eks
- terraform_aks
- terraform_gke
- existing_kubernetes
- local_kind
- cluster_api future

Create:
- cluster-blueprints/
  aws-eks-small.yaml
  aws-eks-prod.yaml
  existing-kubernetes.yaml
  local-kind.yaml

Blueprint schema:
name: aws-eks-small
description: Small AWS EKS cluster for dev
backend: terraform
cloud: aws
orchestrator: eks
version: 0.1.0
inputs:
  region:
    type: string
    default: eu-central-1
  node_count:
    type: number
    default: 2
platform_addons:
  ingress_nginx: true
  cert_manager: true
  external_secrets: false
policies:
  production_ready: false

CLI:
- cf blueprint list
- cf blueprint show <name>
- cf blueprint validate <name>
- cf env create dev --blueprint aws-eks-small

Control Plane:
- optional blueprint catalog API:
  - GET /api/v1/blueprints

Tests:
- validate blueprint
- generate env from blueprint
- invalid input type fails
- blueprint list

Docs:
- docs/cluster-blueprints.md

Rules:
- Blueprints generate readable Terraform.
- Blueprints do not hide security warnings.
- Production blueprint must require remote backend and approvals.
```


---

# Prompt 274 — Blueprint registry support

```text
Add blueprint registry support.

Goal:
Allow organizations to maintain reusable environment and cluster blueprints outside the core repository.

Sources:
- local path
- Git source with ref
- archive file

Config:
blueprint_registries:
  - name: company-blueprints
    source: git::https://github.com/example/clusterforge-blueprints.git?ref=v0.1.0
    enabled: true

CLI:
- cf blueprint registry list
- cf blueprint registry fetch <name>
- cf blueprint registry update <name>
- cf blueprint registry validate <name>
- cf blueprint cache clear

Cache:
- .cf/cache/blueprints/<registry>/<version>

Validation:
- metadata.yaml exists
- blueprints are valid
- supported cloud/orchestrator declared
- no secrets in blueprints
- refs pinned
- warn on main/master branch refs

Security:
- no code execution
- templates only
- explicit fetch
- source shown to user

Tests:
- local registry
- Git source parsing
- cache path
- missing metadata fails
- unpinned ref warns
- blueprint validation

Docs:
- docs/blueprint-registry.md

Rules:
- Do not implement marketplace.
- Do not download remote code automatically.
```


---

# Prompt 275 — Crossplane integration RFC

```text
Create Crossplane integration RFC.

Goal:
Evaluate Crossplane as an optional backend for platform resource provisioning.

Create:
- docs/rfcs/036-crossplane-integration.md
- docs/crossplane.md

Cover:

1. Goals
   - Kubernetes-native cloud resource provisioning
   - platform APIs through compositions
   - GitOps-compatible infrastructure
   - developer self-service resources

2. Non-goals
   - replacing Terraform entirely
   - mandatory Crossplane dependency
   - managing every cloud resource through Crossplane

3. Use cases
   - app teams request database/cache/queue through claims
   - platform team owns compositions
   - GitOps reconciles claims
   - Control Plane tracks claims/status

4. Terraform boundary
   - Terraform creates cluster and installs Crossplane
   - Crossplane provisions app-level cloud resources
   - Terraform can still provision foundation resources

5. Proposed modules
   - modules/platform/kubernetes/crossplane
   - modules/platform/kubernetes/crossplane-provider-aws
   - modules/platform/kubernetes/crossplane-provider-azure
   - modules/platform/kubernetes/crossplane-provider-gcp

6. CLI future
   - cf resource claim create postgres
   - cf resource claim list
   - cf resource claim status

7. Risks
   - ownership conflicts with Terraform
   - provider credentials
   - debugging complexity
   - CRD lifecycle
   - state split across systems

Do not implement code.
Update roadmap.
```


---

# Prompt 276 — Kubernetes fleet add-on manager

```text
Implement Kubernetes fleet add-on manager.

Goal:
Track and compare platform add-ons across multiple clusters.

Data sources:
- clusterforge.yaml
- generated Terraform
- Helm releases if cluster access available
- Control Plane inventory

Tracked add-ons:
- ingress-nginx
- cert-manager
- external-dns
- external-secrets
- metrics-server
- prometheus-stack
- loki
- argocd
- kyverno
- gatekeeper
- velero
- argo-rollouts
- opentelemetry-collector

CLI:
- cf fleet addons list
- cf fleet addons diff
- cf fleet addons report --format markdown|json
- cf fleet addons drift

Control Plane:
- add-on inventory table optional:
  - cluster_id
  - name
  - desired_version
  - observed_version
  - status
  - source

Dashboard:
- fleet add-ons page

Behavior:
- show desired vs observed
- show missing add-ons
- show version skew
- show CRD-sensitive upgrades
- no automatic upgrade

Tests:
- compare two clusters
- missing add-on detected
- version skew detected
- JSON output

Docs:
- docs/fleet-addons.md

Rules:
- Read-only.
- No Helm upgrade.
- No cluster mutation.
```


---

# Prompt 277 — GitOps reconciliation status ingestion

```text
Implement GitOps reconciliation status ingestion.

Goal:
Allow Control Plane to display Argo CD or Flux sync/health status without becoming the GitOps controller.

Supported MVP:
- Argo CD Application status from Kubernetes API or exported JSON
- Flux Kustomization/HelmRelease status optional

CLI:
- cf gitops status --env prod
- cf gitops export-status --format json
- cf api push-gitops-status

Control Plane:
- gitops_app_status table:
  - id
  - project_id
  - environment_id
  - cluster_id
  - provider
  - app_name
  - sync_status
  - health_status
  - revision
  - message
  - observed_at
  - metadata_json

API:
- GET /api/v1/gitops/status
- POST /api/v1/gitops/status/import

Dashboard:
- GitOps status page
- environment detail shows GitOps status

Tests:
- import Argo CD app fixture
- import Flux fixture if implemented
- status visible by environment
- cross-tenant denied

Docs:
- docs/gitops-status.md

Rules:
- Read-only.
- Do not trigger sync.
- Do not store Git credentials.
- Do not replace Argo CD/Flux.
```


---

# Prompt 278 — Automated ticket creation RFC

```text
Create automated ticket creation RFC.

Goal:
Design integration for creating issues/tickets from drift, failed applies, policy blockers, and incidents.

Create:
- docs/rfcs/037-ticketing-integration.md
- docs/ticketing.md

Potential systems:
- GitHub Issues
- GitLab Issues
- Jira
- generic webhook

Use cases:
- drift detected in prod
- policy block in prod
- apply failed
- runner offline
- backup evidence missing
- token expiring
- compliance evidence missing
- incident started

Design:
- ticket templates
- deduplication key
- update existing ticket instead of spamming
- labels/severity
- owner mapping
- link to Control Plane record
- no secrets in ticket body

Config:
ticketing:
  enabled: true
  provider: github
  dedupe: true
  events:
    - drift.detected
    - apply.failed

Non-goals:
- full ITSM workflow
- automatic remediation
- storing provider tokens directly

Do not implement code.
Update roadmap.
```


---

# Prompt 279 — v0.7 roadmap and release plan

```text
Create v0.7.0 roadmap and release plan.

Create:
- ROADMAP_V0.7.md
- RELEASE_PLAN_V0.7.md
- BACKLOG_V0.7.md

Theme:
ClusterForge v0.7 focuses on SaaS-readiness, enterprise governance, tenant isolation, usage metering, runner scalability, compliance evidence, and blueprint-driven platform operations.

Goals:
- v0.6 final release gate
- SaaS-readiness architecture
- tenant isolation enforcement
- tenant-scoped data access layer
- API rate limits and quotas
- usage metering
- organization onboarding/offboarding
- customer-managed encryption design
- KMS-backed artifact encryption
- secret reference broker
- Kubernetes Job runner executor
- ephemeral runner pools
- job fairness
- advanced approval policies
- policy exceptions
- risk acceptance
- compliance evidence collection
- evidence bundles
- immutable audit log MVP
- SCIM provisioning MVP
- data residency model
- cluster blueprints
- blueprint registry
- fleet add-on manager
- GitOps status ingestion

Non-goals:
- public SaaS launch
- billing implementation
- plugin marketplace
- automatic rollback
- automatic remediation
- global scheduler
- legal compliance certification
- storing customer cloud credentials

Milestones:
1. v0.6 release validation
2. tenant isolation and quotas
3. encryption and artifact security
4. runner scalability
5. governance workflows
6. compliance evidence and audit immutability
7. identity provisioning and data residency
8. blueprints and fleet visibility
9. v0.7 release candidate

Acceptance criteria:
- tenant isolation tests pass
- RBAC scoped access tests pass
- quota tests pass
- usage metering tests pass
- artifact KMS mock tests pass
- runner executor tests pass
- advanced approval tests pass
- policy exception tests pass
- evidence bundle tests pass
- audit hash chain tests pass
- SCIM tests pass
- blueprint tests pass
- no secrets stored
- docs complete
- release gate passes

BACKLOG_V0.7.md:
Group by:
- API
- CLI
- runner
- dashboard
- security
- tenancy
- governance
- compliance
- identity
- artifacts
- blueprints
- GitOps
- fleet
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
- Summarize recommended v0.7 scope.
- List must-have items.
- List deferred items.
```


---

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


---

## Preporučeni redosled za batch 241–280

Najbolje nije raditi baš sve redom. Praktičan redosled:

```text
241  v0.6 final release gate
242  SaaS-readiness architecture RFC
243  tenant isolation enforcement tests
244  tenant-scoped data access layer
245  API rate limiting and tenant quotas
246  usage metering model
247  usage metering MVP
249  organization onboarding
250  organization offboarding
251  customer-managed encryption keys RFC
252  KMS-backed artifact encryption
254  secret reference broker MVP
255  Kubernetes Job runner executor
256  ephemeral runner pools
258  job fairness and tenant quotas
259  advanced approval policies
260  policy exceptions and waivers
261  risk acceptance workflow
264  compliance evidence collection
265  evidence bundle export
266  immutable audit log RFC
267  audit hash chain MVP
268  SCIM provisioning RFC
269  SCIM provisioning MVP
270  data residency model
273  cluster lifecycle blueprint model
274  blueprint registry support
276  Kubernetes fleet add-on manager
277  GitOps reconciliation status ingestion
279  v0.7 roadmap
280  v0.7 release candidate
```

Najbolji sledeći prompt odmah posle 240 je:

```text
Prompt 241 — v0.6 final release gate review
```
