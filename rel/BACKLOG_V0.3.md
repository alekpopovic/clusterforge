# ClusterForge v0.3 backlog

Owner is deliberately a placeholder until maintainers assign work. Status uses
`done`, `in progress`, `planned`, or `deferred`; each item still requires
release-gate evidence even when implementation already exists.

## CLI

| Item | Priority | Complexity | Owner | Status | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| Stabilize CLI e2e suite | P0 | M | TBD | in progress | Non-cloud project, environment, app, and policy flows pass with isolated temp directories |
| Stabilize golden generator tests | P0 | M | TBD | in progress | Supported templates match reviewed fixtures and drift produces an actionable diff |
| Harden doctor and JSON diagnostics | P1 | S | TBD | done | Missing tools and unsafe configuration return stable exit codes and valid JSON |
| Keep fleet commands read-only | P0 | M | TBD | done | Tests prove inventory/status paths do not invoke Terraform mutation commands |
| Write-capable fleet orchestration | P3 | L | TBD | deferred | Requires a separate safety RFC and explicit production authorization model |

## Terraform modules

| Item | Priority | Complexity | Owner | Status | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| Core native Terraform tests | P0 | M | TBD | done | Naming and labels modules pass supported Terraform tests |
| Module conformance gate | P0 | M | TBD | done | Every reusable module has required files, descriptions, and provider-boundary checks |
| Module release packaging review | P1 | M | TBD | planned | Published module set has versions, checksums, catalog entries, and upgrade notes |
| Validate state-sensitive interface changes | P0 | M | TBD | planned | Plans and migration notes cover address or type changes before release |

## Kubernetes

| Item | Priority | Complexity | Owner | Status | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| Local Kubernetes target | P0 | M | TBD | done | Documented local create/use/cleanup flow passes without cloud credentials |
| Existing Kubernetes target | P0 | M | TBD | done | Bootstrap uses supplied provider configuration and never owns cluster lifecycle |
| Tenant and resource baseline | P0 | M | TBD | done | Namespace, RBAC, quotas, limits, and default-deny options have examples and tests |
| Admission policy baseline | P1 | M | TBD | done | Kyverno and Gatekeeper alternatives are optional, documented, and conformance-tested |
| Platform add-on conformance | P0 | L | TBD | in progress | Supported chart/module combinations render and plan against declared Kubernetes versions |
| Advanced service mesh default | P3 | L | TBD | deferred | RFC and operational evidence are required before any default integration |

## AWS

| Item | Priority | Complexity | Owner | Status | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| EKS production options | P0 | L | TBD | done | Private access, encryption, logging, managed nodes, OIDC/IRSA, and add-on roles have plan tests |
| Registry, KMS, and VPC endpoint modules | P1 | M | TBD | done | Focused modules pass fmt/validate/plan checks and document IAM/network implications |
| Backup and Route53 integration | P0 | L | TBD | in progress | Velero and DNS01/external-dns examples use scoped identity and include restore/cleanup steps |
| AWS real smoke test | P1 | L | TBD | planned | Disposable EKS or ECS run records versions, cost, evidence, and verified cleanup |

## Multi-cloud

| Item | Priority | Complexity | Owner | Status | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| Preserve AKS/GKE experimental modules | P2 | M | TBD | planned | Credential-free validation passes and experimental status is explicit |
| Production AKS/GKE parity | P3 | XL | TBD | deferred | Separate roadmap includes identity, networking, upgrades, backup, and real-cloud evidence |
| Self-hosted K3s/RKE2 documentation | P2 | M | TBD | planned | Ownership, bootstrap, upgrade, and cleanup limitations are explicit |

## Security

| Item | Priority | Complexity | Owner | Status | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| Threat model and security checklist | P0 | M | TBD | done | Assets, threats, implemented controls, gaps, and review triggers are documented honestly |
| Secret scanning release gate | P0 | M | TBD | done | CI and optional local scanning run; findings and exclusions are reviewed |
| Supply-chain release evidence | P0 | M | TBD | in progress | Every CLI release includes checksums and SBOM; signing gap is documented |
| Workload identity adoption | P1 | L | TBD | in progress | Supported AWS and Kubernetes examples avoid static workload credentials |
| Automatic security remediation | P3 | XL | TBD | deferred | Explicitly outside v0.3; would require policy, authorization, and rollback design |

## Testing

| Item | Priority | Complexity | Owner | Status | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| Credential-free release test suite | P0 | L | TBD | in progress | Required CLI, golden, Terraform, and conformance suites pass on release commit |
| Provider compatibility CI | P0 | M | TBD | done | Supported constraint edges are exercised and failures identify provider/tool versions |
| Ephemeral integration harness | P1 | L | TBD | planned | Creates isolated fixtures, captures redacted evidence, and always attempts cleanup |
| Real smoke test evidence | P0 | L | TBD | planned | At least one supported target has dated, versioned, cleaned-up evidence |

## Docs

| Item | Priority | Complexity | Owner | Status | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| Install and development setup | P0 | S | TBD | done | Native, devcontainer, and optional asdf paths are clear and do not require credentials |
| Quickstart and cleanup tutorial | P0 | M | TBD | planned | Commands run from a clean checkout and include ownership-safe cleanup |
| Security and recovery documentation | P0 | M | TBD | in progress | Threat model, checklist, backup, restore, and DR limitations are cross-linked |
| Screenshot-heavy product tutorial | P3 | M | TBD | deferred | Text-first reproducible guides take priority for v0.3 |

## Release

| Item | Priority | Complexity | Owner | Status | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| v0.3 release-gate report | P0 | M | TBD | planned | Every must-have gate has evidence, failure, or explicit approved exception |
| Version, changelog, and matrix update | P0 | S | TBD | planned | Files agree on v0.3 support, experimental targets, and known limitations |
| Release artifacts and verification | P0 | M | TBD | planned | Expected platform binaries, checksums, and SBOM publish and are independently verified |
| Migration and rollback notes | P0 | M | TBD | planned | Interface changes list safe migration steps; rollback never assumes blind infrastructure revert |
