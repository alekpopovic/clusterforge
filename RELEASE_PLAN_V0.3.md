# ClusterForge v0.3.0 release plan

## Recommended scope

v0.3.0 consolidates already introduced local/existing Kubernetes, EKS
hardening, backup, security-baseline, fleet, and testing capabilities into a
reviewable release. Stabilization and evidence take precedence over adding a
new cloud or mutation workflow.

## Release gates

### Must have

- All roadmap acceptance criteria have recorded evidence.
- `make fmt-check`, CLI tests/e2e tests, generator golden tests, module
  conformance, and supported credential-free Terraform tests pass.
- Local Kubernetes and existing-cluster workflows document ownership boundaries
  and cleanup.
- EKS production options, workload identity, backup/restore, tenancy, admission,
  and network controls have limitations and examples documented.
- Fleet operations remain read-only and cover machine-readable output.
- Security scans report no known committed secret; release artifacts include
  checksums and an SBOM.
- Installation, quickstart, development, upgrade notes, and cleanup paths are
  accurate for the release tag.

### Should have

- A disposable real EKS or existing-Kubernetes smoke test with recorded cost,
  versions, cleanup, and redacted evidence.
- Provider and Kubernetes compatibility jobs cover the supported matrix.
- Recovery runbooks have a tabletop review and at least one restore exercise.

### Deferred from v0.3

- Full production AKS/GKE parity and broad real-cloud matrices.
- Automatic drift or security remediation.
- Executable third-party plugins or a plugin marketplace.
- Advanced service mesh defaults and complex progressive delivery guarantees.
- Hosted UI/API/SaaS and write-capable fleet orchestration.

## Milestone sequence

1. **M1 — Testing and release gate:** freeze fixture formats, fix deterministic
   failures, and establish release evidence.
2. **M2 — Local/existing Kubernetes:** verify local lifecycle and non-ownership
   of attached clusters.
3. **M3 — AWS production hardening:** review EKS, KMS, VPC endpoints, ECR,
   identity, DNS, and data-service plans.
4. **M4 — Backup and security baseline:** validate backup configuration,
   recovery docs, tenancy, quotas, policy engines, and network boundaries.
5. **M5 — Fleet operations:** verify read-only semantics, audit events, and JSON
   output across mixed inventory.
6. **M6 — Docs and examples:** run every documented install/quickstart/cleanup
   command and close release documentation gaps.

## Validation record

For each gate, record command, commit, tool/provider versions, date, operator,
result, and a reason for every skip. Cloud evidence must also record region,
approximate cost, resource inventory, and cleanup verification without exposing
credentials, account IDs, state, or kubeconfigs.

Recommended final checks:

```bash
make fmt-check
make lint
make test
make validate
make security
cd cli && go build -o cf . && go test ./...
```

The release owner updates `VERSION`, `CHANGELOG.md`, `VERSION_MATRIX.md`, smoke
test evidence, and known limitations before tagging. A green workflow without
the required evidence is not by itself a release approval.

## Rollback and follow-up

Because Terraform module interfaces and generated roots affect stateful
infrastructure, rollback is release-specific: document incompatible input or
address changes, preserve prior binaries and checksums, and publish migration
guidance. Never recommend reverting applied infrastructure without reviewing a
fresh plan. File deferred and failed gate items with owners before release.
