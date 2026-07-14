# ClusterForge v0.3 roadmap

Theme: production hardening, local development, fleet operations, and security.

## Goals

- Make the local Kubernetes target useful for credential-free development.
- Support adopting an existing Kubernetes cluster without owning its lifecycle.
- Strengthen EKS production options while keeping provider configuration in roots.
- Provide backup/restore modules and honest, testable recovery runbooks.
- Establish tenant and security baselines for Kubernetes workloads.
- Add read-only fleet inventory and inspection operations.
- Strengthen CLI, generator, module, and platform conformance tests.
- Improve installation, quickstart, cleanup, development, and operations tutorials.

## Non-goals

- A full hosted SaaS platform.
- Automatic remediation or unattended production mutation.
- A plugin marketplace.
- Advanced service mesh enabled by default.
- Production-grade parity across every supported cloud.

## Milestones

| Milestone | Outcome | Exit signal |
| --- | --- | --- |
| M1: Testing and release gate | Reliable non-cloud regression suite and explicit release evidence | CLI e2e, golden generator, core Terraform, conformance, and release-gate checks pass |
| M2: Local/existing Kubernetes | Fast local target and safe attachment to operator-owned clusters | Documented create/use/cleanup flows work without ClusterForge owning an existing cluster |
| M3: AWS production hardening | Safer EKS, registry, encryption, networking, identity, and data-service building blocks | Plan tests and security guidance cover supported production options |
| M4: Backup and security baseline | Recoverability and tenant/workload protections are composable | Velero/runbooks plus namespace, quota, policy, identity, and admission examples are validated |
| M5: Fleet operations | Read-only multi-cluster inventory and inspection | Fleet commands have deterministic output, tests, and no mutation path |
| M6: Docs and examples | Users can install, start, inspect, and clean up safely | Docs cover install, quickstart, examples, security limitations, and cleanup |

## Acceptance criteria

- CLI end-to-end tests pass.
- Golden generator tests pass without unexplained fixture drift.
- Core module Terraform tests pass.
- At least one real smoke test is documented with date, version, operator, and cleanup result.
- Secret scanning finds no known exposure in release scope or history reviewed for the release.
- Documentation covers installation, quickstart, and cleanup.
- Failed, skipped, credential-gated, and experimental checks are stated explicitly.

## Risks

| Risk | Response |
| --- | --- |
| Cloud tests cost money | Use credential-free tests first, budgets and short-lived fixtures for scheduled smoke tests |
| Provider compatibility changes | Maintain the version matrix and compatibility CI; review lock changes |
| Kubernetes chart drift | Pin tested chart versions and run platform conformance tests |
| IAM complexity | Prefer focused roles, workload identity, plan review, and documented trust boundaries |
| Too many features before stabilization | Treat release-gate must-haves as blockers and defer optional parity/experiments |

Detailed sequencing and release evidence live in `RELEASE_PLAN_V0.3.md`; the
work-item inventory lives in `BACKLOG_V0.3.md`.
