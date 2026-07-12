# ClusterForge v0.4 release plan

## Release policy

v0.4.0 is a stabilization release for enterprise workflows already in scope. A
feature is not complete when code merely exists: tests, redaction, documentation,
upgrade/rollback impact and clean-checkout evidence are required. Experimental or
credential-gated areas must be labelled rather than inferred as supported.

## Sequence and gates

### M1 — baseline

- Run and archive the v0.3 release gate from a clean checkout.
- Triage dependency/security findings; unresolved critical findings block release.
- Freeze generator/config compatibility expectations and document migrations.

Exit: baseline failures have owners, priority and release disposition.

### M2 — extensibility

- Threat-model plugin discovery/execution and verify disabled-by-default behavior.
- Test template source/version/checksum, cache and path safety.
- Freeze policy result schema, pack identifiers, severity and override behavior.

Exit: extension failures are deterministic and never silently weaken policy.

### M3 — enterprise configuration

- Round-trip organization, teams, workspaces, AWS accounts/regions and Terraform
  Cloud configuration.
- Exercise shared production account safeguards and multi-region metadata.
- Validate GitLab templates without embedding IDs, tokens or credentials.

Exit: minimal and representative enterprise configs validate and generate.

### M4 — operations

- Exercise health, drift, cost, audit export, backup evidence and upgrade plans.
- Verify default commands are read-only and production mutations retain gates.
- Test JSON/text schemas, exit behavior and redaction.

Exit: documented operator workflows produce attributable local evidence.

### M5 — visibility and mappings

- Validate fleet/asset/dashboard, service catalog and Backstage exports.
- Review compliance mappings for status, evidence, limitation and disclaimer.
- Verify offline bundle exclusions, checksum tamper detection and acquisition lists.

Exit: stable golden outputs contain no secret/state/kubeconfig data.

### M6 — release candidate

1. Freeze code and docs; update version/changelog/release notes.
2. Run formatting, Go tests/vet, CLI e2e and golden tests.
3. Run Terraform formatting, module/example validation and conformance scripts.
4. Run secret, dependency, license and security scans; review generated artifacts.
5. Build supported CLI archives, SHA256SUMS and SBOM from a clean tagged commit.
6. Install artifacts on supported OS/architecture targets and run offline help/basic
   command smoke tests.
7. Attach current cloud smoke evidence where required, including cleanup/cost.
8. Verify repository is clean, tag is signed/approved, and published checksums match.

Exit: all mandatory checks pass or the release is stopped. Skips are not success.

## Required test matrix

| Area | Required evidence |
|---|---|
| CLI | `go test ./...`, `go vet ./...`, e2e and command help smoke |
| Generator | reviewed golden suite for supported environment targets |
| Terraform | `terraform fmt -check -recursive`, module/example validate/conformance |
| Extensibility | plugin, template registry and policy engine security/invalid-input tests |
| Enterprise | config round-trip, AWS account safeguards, Terraform Cloud generation, GitLab lint |
| Visibility | stable inventory/catalog/Backstage/compliance JSON and Markdown fixtures |
| Offline | secret/state/kubeconfig exclusion and SHA256 tamper/unlisted-file rejection |
| Supply chain | secret scan, dependency review, SBOM, checksums and artifact install smoke |

## Documentation gate

Architecture, configuration reference, plugin trust, templates, policies,
multi-account/region, Terraform Cloud, GitLab, upgrade operations, catalog,
compliance and air-gapped workflows must be discoverable. Every workflow states
mutation behavior, credential boundary, limitations and cleanup/rollback where
applicable.

## Go/no-go review

Must-have backlog items are complete; no known secret exposure or unresolved
critical release-scope vulnerability exists; schemas and migrations are reviewed;
cloud support claims match evidence; artifacts reproduce and verify. Otherwise
create a new RC after remediation—do not waive a safety gate to meet a date.

## RC1 execution status (2026-07-12)

`v0.4.0-rc.1` is currently **no-go**. Formatting, Go tests/build, CLI version,
temporary-project doctor, Gitleaks and module contract execution succeeded.
Release blockers are recorded in `RELEASE_CANDIDATE_V0.4.md`: lint failures,
128 Checkov findings, an untriaged critical Dependabot alert, and absent
production-cloud smoke evidence. Full Terraform validation now passes with golden
snapshots explicitly delegated to the passing Go golden tests. M6 remains open.
