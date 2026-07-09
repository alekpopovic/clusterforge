# Roadmap v0.2

Theme: production validation, multi-cloud MVP, better CLI workflow, policy
packs, and template packs.

## Goals

- record real AWS smoke-test evidence
- stabilize drift, state, upgrade, cost, promotion, and template commands
- validate AKS and GKE MVP modules in real accounts
- formalize policy packs and release workflow

## Non-Goals

- SaaS control plane
- plugin marketplace
- automatic cloud remediation
- production imports without manual review

## Milestones

| Milestone | Must-have acceptance criteria |
| --- | --- |
| Production validation | `SMOKE_TEST_MATRIX.md` has at least one redacted EKS and ECS run. |
| Multi-cloud MVP | AKS and GKE examples validate and have smoke runbooks. |
| CLI workflows | Drift/state/upgrade/cost/template commands have unit tests and docs. |
| Policy packs | Baseline and production packs have enforceable checks or explicit advisory labels. |
| Release readiness | Release workflow produces binaries and checksums from a test tag. |

## Risks

- provider schema churn
- real cloud cost
- incomplete enterprise policy coverage
- pre-1.0 module interface changes

## Dependencies

- Terraform/OpenTofu provider availability
- maintainers with disposable AWS/Azure/GCP accounts
- GitHub release permissions
