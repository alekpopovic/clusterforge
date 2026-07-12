# ClusterForge v0.5 backlog

Theme: verified platform support, explicit module maturity, provider
compatibility, and provenance-aware releases.

Status values: `done`, `in-progress`, `planned`, `blocked`, `deferred`.
Priorities: P0 release blocker, P1 must-have, P2 useful if low risk.

## Credential-free release scope

These items must be implementable and testable without cloud credentials.

| Item | Priority | Complexity | Status | Acceptance criteria | Notes |
|---|---|---:|---|---|---|
| Publish module maturity metadata | P0 | medium | planned | Every catalogued module has a documented maturity level, supported providers/platforms, limitations, and evidence link; catalog checks fail on missing metadata | A maturity label is not a production-readiness claim |
| Build the provider compatibility matrix | P0 | high | planned | Supported Terraform/OpenTofu and provider combinations initialize and validate from reviewed lock files in CI; unsupported combinations are documented | No cloud apply is implied |
| Stabilize CLI machine-readable contracts | P0 | medium | planned | Versioned JSON schemas and golden tests cover drift, state, cost, inventory, catalog, compliance, and policy output; incompatible changes fail CI | Additive evolution within a schema version |
| Complete non-interactive CLI coverage | P1 | medium | planned | CI covers missing input, exit codes, JSON output, wizard defaults, production plan-file enforcement, and non-prompting behavior | Destructive safeguards remain mandatory |
| Harden generated-artifact secret exclusions | P0 | medium | planned | Fixtures prove generated roots, exports, bundles, logs, and release archives exclude credentials, kubeconfigs, private keys, state, and plan files | Use synthetic identifiers only |
| Add policy-pack OPA validation | P1 | medium | planned | Selected advisory controls have deterministic Rego tests, stable IDs, documented limitations, and opt-in enforcement guidance | New controls default to audit/advisory |
| Add release dry-run automation | P0 | high | planned | A non-publishing workflow reproducibly builds archives, checksums, SBOMs, and provenance metadata from a clean tag-shaped ref | Dry-run must never publish or sign with production identity |
| Add supported-platform install smoke | P0 | medium | planned | Linux, macOS, and Windows artifacts run version, help, and a credential-free generated-project smoke test | Platform claims are limited to tested artifacts |
| Refresh API, CLI, and configuration references | P1 | medium | planned | Documented commands and schemas match generated help and tested fixtures; stale proposed commands are removed | Include upgrade and rollback notes |
| Audit experimental and unsupported claims | P0 | low | planned | AKS, GKE, K3s, RKE2, Nomad, Docker, edge, federation, service mesh, and Windows claims consistently match recorded evidence | Scaffolding and validation are not runtime evidence |

## Credential-gated evidence

These items require operator-supplied credentials or disposable infrastructure.
They must run only through explicitly approved workflows, must not expose secrets,
and must always include cleanup evidence. Their absence limits support claims; it
must not be reported as a passing test.

| Item | Priority | Complexity | Status | Acceptance criteria | Credential and safety boundary |
|---|---|---:|---|---|---|
| AWS EKS lifecycle smoke | P0 | high | planned | A pinned example completes plan, apply, workload health, upgrade/rollback exercise, destroy, and cost/resource cleanup review; date and tester are recorded | Requires a disposable AWS account, scoped credentials, budget guardrails, and explicit approval |
| AWS ECS lifecycle smoke | P0 | high | planned | Cluster and Fargate service complete plan, apply, health, deployment rollback, destroy, and cleanup review with recorded versions | Requires a disposable AWS account and scoped credentials; CodeDeploy blue/green remains separate evidence until exercised |
| AKS production-hardening evidence | P1 | high | planned | Private networking, identity, node-pool, upgrade, workload, destroy, and cleanup checklist passes before any beta claim | Requires a disposable Azure subscription and scoped service principal |
| GKE production-hardening evidence | P1 | high | planned | Private cluster, workload identity, node-pool, upgrade, workload, destroy, and cleanup checklist passes before any beta claim | Requires a disposable GCP project and scoped service account |
| K3s and RKE2 lifecycle validation | P2 | high | planned | Installation, workload, disruption, upgrade, rollback/rebuild, and teardown runbooks pass on declared VM/OS profiles | Requires disposable VMs; no production edge claim |
| Admission audit-to-enforce exercise | P1 | medium | planned | Gatekeeper and Kyverno examples prove audit findings, narrow remediation, staged enforcement, rollback, and cleanup | Requires a disposable Kubernetes cluster; enforcement is never enabled on an existing cluster by default |
| Backup and restore exercise | P1 | high | planned | A disposable workload backup and restore meets documented recovery checks and records tool/provider versions | Requires disposable cluster and object storage; no real application data |

## Release evidence and governance

| Item | Priority | Complexity | Status | Acceptance criteria | Notes |
|---|---|---:|---|---|---|
| Dependency and scanner finding gate | P0 | medium | planned | No unresolved critical shipped/runtime finding; exceptions identify owner, scope, expiry, and evidence | Blanket suppressions are prohibited |
| Evidence index and support matrix | P0 | medium | planned | Each support claim links to current local, CI, or credential-gated evidence with tool versions, date, and known limitations | Missing/expired evidence downgrades the claim |
| v0.5 release go/no-go | P0 | medium | planned | Clean tree/tag, full credential-free gate, release dry-run, evidence index, known limitations, and approval are recorded | Skipped checks are not passes |

## Explicitly deferred

Hosted SaaS, automatic remediation, automatic regional failover, global
scheduling, cross-cluster data or secret replication, public plugin marketplace,
compliance certification, cluster federation, general edge lifecycle automation,
and production Windows workload support require separate RFC and release
decisions. They are not part of v0.5.
