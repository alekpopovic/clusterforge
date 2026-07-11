# RFC 011: Future ClusterForge platform API

Status: design only; no API server is implemented or committed by this RFC.

## Goals

Provide a backend for central inventory, environment/cluster status, plan
requests, policy results, audit events, service catalog, read-only fleet
operations and a future UI. Terraform/OpenTofu files and reviewed plans remain
the infrastructure source of truth.

## Non-goals

- Replacing Terraform/OpenTofu, providers, state backends or Git review.
- Storing cloud credentials directly in the API database.
- Automatic remediation or production mutation by default.
- Multi-tenant hosted SaaS in the first version.

## Resource model

Organization owns Workspaces and Projects. Projects contain Environments;
Environments contain Stacks and reference Clusters. Apps and Services describe
workloads/ownership. Plans are immutable request/result records with artifact
references. PolicyResult and AuditEvent records attach to operations/resources.
Runbooks are metadata references, not executable procedures.

Every resource needs a stable ID, organization boundary, version for optimistic
concurrency, timestamps, labels, lifecycle status and actor/audit references.

## Execution options

1. **Local CLI only:** API stores/imports metadata; execution remains local.
2. **GitOps/PR:** API opens or observes reviewed changes and CI plans.
3. **Remote runner:** short-lived isolated worker claims an approved job.
4. **HCP Terraform:** API records workspace/run metadata and delegates runs.

The first implementation should prefer local/GitOps metadata and read-only
status. Any remote runner requires a separate threat model, queue/lease design,
workspace isolation, artifact integrity, cancellation and failure recovery.

## Security and approvals

Authentication should use organization OIDC with short-lived sessions.
Authorization needs organization/workspace/project roles plus environment and
operation policy; production approvals must enforce separation of duties.
Every read of sensitive metadata and every requested/approved/executed action
must produce an append-oriented audit event.

The API stores secret references, never cloud access keys or Terraform variable
secrets. Runners should use workload identity, bounded credentials and isolated
work directories. Plan/state artifacts may contain secrets and require
encryption, narrow access, retention limits, redaction and integrity checks.

## Storage

PostgreSQL is the candidate metadata/relationship store. Object storage may
hold checksummed plans, logs, reports and SBOMs with short retention. Existing
remote backends continue to own Terraform state; the API does not store state
by default. Backups, migrations, regional recovery and deletion semantics need
design before production use.

## Candidate API and CLI

Versioned HTTP resources should expose inventory/status and asynchronous plan
requests with explicit approval states. Future commands may include `cf login`,
`cf api status`, `cf plan request` and `cf approval list`. CLI local mode must
remain functional when the API is absent.

## Concurrency and state

Use idempotency keys, immutable plan/config digests, optimistic record versions
and one active mutation lease per state workspace. The runner must acquire the
real backend lock; an API lease never substitutes for Terraform locking.
Approval becomes invalid when commit, variables, policy or plan digest changes.

## Risks

- Scope creep from inventory API into a full platform/SaaS.
- Credential, plan and artifact confidentiality.
- Concurrent/stale approvals and duplicate execution.
- Backend state-lock coordination and orphaned jobs.
- Enterprise SSO, residency, retention, audit and availability requirements.

## Proposed phases

1. Read-only local API prototype over inventory, service catalog and audit data.
2. Authenticated organization/workspace authorization and durable audit.
3. GitOps/CI plan request observation without credential custody.
4. Separately approved remote-runner or HCP integration, if justified.

Advancement requires API schema tests, authorization threat modeling, audit and
retention tests, failure injection, and an explicit production operations plan.
