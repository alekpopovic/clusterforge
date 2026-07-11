# ClusterForge v0.4 roadmap

Theme: enterprise extensibility, fleet visibility, policy/compliance, and
operational maturity.

## Goals

- Stabilize the plugin system MVP with explicit trust and CI behavior.
- Stabilize the versioned template pack registry and override validation.
- Ship policy engine v2 with deterministic packs, severity and override behavior.
- Support an explicit AWS multi-account configuration model and safeguards.
- Represent multi-region inventory metadata without automatic failover.
- Provide read-only platform upgrade planners and Kubernetes upgrade planners.
- Support reviewed Terraform Cloud organization/project/workspace configuration.
- Provide non-secret GitLab CI templates alongside GitHub Actions guidance.
- Stabilize the service catalog manifest and CLI validation/export.
- Generate Backstage catalog metadata without credentials or API mutation.
- Publish non-certifying compliance control mappings and reports.
- Produce checksummed, manifest-only air-gapped bundles without downloading
  artifacts.

## Non-goals

- A hosted or full SaaS control plane.
- Automatic remediation or unattended infrastructure mutation.
- A global application scheduler or cross-cluster data placement system.
- A public plugin marketplace or execution of untrusted plugins.
- Compliance certification, attestation, or legal advice.
- Automatic regional failover, traffic switching, database promotion, or secret
  replication.

## Milestones

| Milestone | Outcome | Exit signal |
|---|---|---|
| M1: v0.3 release gate and stabilization | Known v0.3 behavior is a reliable baseline | v0.3 gate evidence is archived; regressions and critical dependency findings are resolved or explicitly block v0.4 |
| M2: plugin/template/policy extensibility | Versioned extensions have predictable discovery, trust, validation and failure behavior | Plugin, template registry and policy engine suites pass with malicious/invalid fixtures |
| M3: enterprise config model | Accounts, regions, workspaces, teams and Terraform Cloud compose without hidden credentials | Config round-trip/validation tests and multi-account safeguards pass |
| M4: upgrade and operations tooling | Operators can inspect health, drift, cost, audit, backup evidence and upgrade plans read-only | CLI reports are deterministic, redacted and documented; mutation remains separately confirmed |
| M5: catalog/compliance/inventory | Service, Backstage, asset, fleet and compliance exports share stable schemas | Golden JSON/Markdown tests pass and limitations/disclaimers are visible |
| M6: release hardening | v0.4 is reproducible, documented and supportable | Full release gate, packaging, checksums, SBOM, install smoke and clean-tree verification pass |

## Acceptance criteria

- Plugin system unit/security tests pass and plugins remain disabled by default.
- Template registry resolution, pinning, checksum and traversal tests pass.
- Policy engine v2 tests pass for baseline, production and override behavior.
- Golden generator tests pass with every fixture change reviewed.
- CLI end-to-end tests pass, including read-only enterprise workflows.
- Documentation covers configuration, trust, upgrade, fleet, catalog, compliance,
  offline and rollback workflows with honest limitations.
- Generated configs, exports, bundles, fixtures and logs contain no secret values,
  state, kubeconfigs or credentials.
- The v0.4 release gate passes from a clean checkout using documented tool versions;
  credential-gated cloud tests have current, attributable evidence or are stated as
  unverified.

## Risks

| Risk | Response |
|---|---|
| Feature creep | Freeze scope at M2; additions require removing equal release risk/work or deferring to v0.5 |
| Plugin security | Disable by default, require explicit trust, constrain discovery, redact I/O and document plugins as local code execution |
| Policy false positives | Default new controls to advisory/audit, provide narrow overrides and test representative fixtures |
| Provider compatibility | Pin safe constraints/locks, run compatibility matrices and avoid claiming untested cloud support |
| Enterprise complexity | Prefer explicit typed configuration, stable exports and provider-native concepts over a universal hidden abstraction |

## Must-have release scope

MVP plugins, template registry, policy engine v2, enterprise config validation,
read-only upgrade/operations reports, fleet/catalog inventory, compliance mapping,
offline manifest bundles, complete docs, and a passing release gate are blockers.

## Deferred

Hosted services, automatic remediation/failover/scheduling, marketplace
distribution, certification claims, direct SIEM/cloud mutations, federation,
edge lifecycle automation, and production Windows workload claims remain outside
v0.4. Experimental RFCs and scaffolding may exist but are not release guarantees.

Detailed sequencing is in `RELEASE_PLAN_V0.4.md`; tracked work is in
`BACKLOG_V0.4.md`.
