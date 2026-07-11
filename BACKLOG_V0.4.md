# ClusterForge v0.4 backlog

Status values: `done`, `in-progress`, `planned`, `blocked`, `deferred`.
Priorities: P0 release blocker, P1 must-have, P2 useful if low risk.

## CLI

| Item | Priority | Complexity | Status | Acceptance criteria | Notes |
|---|---|---:|---|---|---|
| Stabilize plugin MVP | P0 | high | in-progress | Disabled-by-default discovery, trust, timeout/error and security tests pass | Plugins execute local code; no marketplace |
| Freeze enterprise JSON schemas | P1 | medium | planned | Versioned golden inventory/catalog/compliance outputs pass | Additive changes only after freeze |
| Validate non-interactive workflows | P1 | medium | planned | CI commands fail on missing input and never prompt | Include wizard defaults fixture |

## Terraform modules

| Item | Priority | Complexity | Status | Acceptance criteria | Notes |
|---|---|---:|---|---|---|
| AWS multi-account safeguards | P0 | high | in-progress | Shared-prod-account validation and provider-alias examples pass | No credential custody |
| Multi-region metadata | P1 | medium | in-progress | Regions render explicitly with deterministic inventory | No failover resources implied |
| Provider compatibility validation | P0 | high | planned | Supported module/example matrix validates with reviewed locks | Cloud smoke remains credential-gated |

## Policy

| Item | Priority | Complexity | Status | Acceptance criteria | Notes |
|---|---|---:|---|---|---|
| Policy engine v2 schema/packs | P0 | high | in-progress | Baseline/production/override tests and stable IDs pass | New controls audit first |
| Template pack registry hardening | P0 | high | in-progress | Pin/checksum/path/cache tests pass | No unpinned remote execution |
| Admission pack rollout evidence | P2 | high | planned | Audit-to-enforce example is validated on disposable cluster | Enforcement remains opt-in |

## Security

| Item | Priority | Complexity | Status | Acceptance criteria | Notes |
|---|---|---:|---|---|---|
| Plugin threat model | P0 | high | planned | Trust boundary, discovery paths and malicious fixture coverage reviewed | Marketplace deferred |
| Generated artifact secret review | P0 | medium | planned | Scans and explicit state/kubeconfig/credential exclusion tests pass | Includes bundles/exports |
| Dependency critical finding gate | P0 | medium | blocked | No unresolved critical release-scope finding | Must triage current Dependabot report |

## Compliance

| Item | Priority | Complexity | Status | Acceptance criteria | Notes |
|---|---|---:|---|---|---|
| Review mapping pack accuracy | P1 | medium | in-progress | Every row has status, evidence, limitation and disclaimer | No certification claim |
| Stable compliance CLI reports | P1 | low | in-progress | Unknown pack fails; Markdown/JSON goldens pass | Implementation aid only |

## Fleet

| Item | Priority | Complexity | Status | Acceptance criteria | Notes |
|---|---|---:|---|---|---|
| Fleet health semantics | P1 | medium | planned | Unknown/unreachable/unhealthy are distinct and read-only | Never triggers remediation |
| Asset/dashboard export stability | P1 | medium | in-progress | Schema version/redaction/deterministic ordering tests pass | No cloud calls by default |
| Multi-cluster GitOps inventory | P2 | medium | in-progress | Two-cluster and filtered rendering tests pass | No registration/scheduling |

## Docs

| Item | Priority | Complexity | Status | Acceptance criteria | Notes |
|---|---|---:|---|---|---|
| Enterprise workflow documentation | P0 | high | planned | All v0.4 workflows linked, scoped and limitation-reviewed | Include rollback/cleanup |
| API/CLI/config reference refresh | P1 | medium | planned | Examples match actual help/schema and pass copy checks | Remove stale proposed commands |
| Experimental status audit | P1 | low | planned | Edge/federation/Windows/SaaS claims are consistently labelled | Docs are not support evidence |

## Testing

| Item | Priority | Complexity | Status | Acceptance criteria | Notes |
|---|---|---:|---|---|---|
| Full Go/e2e/golden gate | P0 | high | planned | Go test/vet, e2e and reviewed goldens pass cleanly | No unexplained skips |
| Terraform conformance matrix | P0 | high | planned | fmt and supported module/example validation pass | Record tool/provider versions |
| Enterprise workflow fixtures | P1 | medium | planned | Multi-account/TFC/catalog/compliance/offline fixtures pass | Synthetic identifiers only |

## Release

| Item | Priority | Complexity | Status | Acceptance criteria | Notes |
|---|---|---:|---|---|---|
| RC packaging and SBOM | P0 | high | planned | Archives, SHA256SUMS and SBOM reproduce from clean tag | Signing/provenance stated honestly |
| Supported-platform install smoke | P0 | medium | planned | Linux/macOS/Windows artifacts run version/help/basic offline smoke | Failures block platform claim |
| Release evidence and go/no-go | P0 | medium | planned | Evidence index complete; clean tree/tag; approvals recorded | Skipped checks are not pass |

## Explicitly deferred

Full SaaS, automatic remediation/failover/global scheduling, public plugin
marketplace, compliance certification, direct SIEM pushes, cluster federation,
edge lifecycle automation, and production Windows support are v0.5+ or separate
RFC decisions.
