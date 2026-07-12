# ClusterForge v0.4.0-rc.1 assessment

Assessment date: 2026-07-12  
Candidate version: `0.4.0-rc.1`  
Status: **NO-GO / release candidate not ready for final v0.4.0**

## Included features

- Plugin MVP, template pack registry and policy engine v2.
- AWS account/region metadata, Terraform Cloud configuration and execution
  profiles.
- Kubernetes/platform upgrade planners, fleet health, asset/dashboard inventory,
  audit and cost reporting.
- Service catalog, Backstage export, compliance mappings and read-only operational
  workflows for secrets, backup evidence and migration assessment.
- Multi-cluster Argo CD rendering and checksummed manifest-only offline bundles.
- v2 scaffolding wizard and expanded enterprise/operations documentation.

## Excluded features

Hosted SaaS, automatic remediation, global scheduling, automatic failover/data or
secret replication, plugin marketplace, compliance certification, direct SIEM
push, cluster federation, edge lifecycle automation and production Windows
workload support are not included. RFC/scaffolding content is not a support claim.

## Breaking and behavior changes

- `cf init` without an environment now launches project scaffolding; `cf init
  <env>` retains Terraform/OpenTofu initialization behavior.
- New app manifests include Linux/amd64 platform intent; Windows values are
  experimental and Kubernetes rendering remains Linux-oriented.
- Enterprise config fields and versioned report/export schemas are additive, but
  consumers must tolerate new fields and check schema versions.
- Admission and policy extensions remain audit/advisory by default; opting into
  enforcement can reject workloads and requires a staged rollout.

## Migration notes

1. Back up configuration and state through the existing protected process.
2. Run `cf upgrade check` and `cf upgrade plan`; review generated/config diffs.
3. Keep plugins disabled until each executable/source is reviewed and pinned.
4. Re-run policy and generator golden checks; resolve or narrowly document new
   advisory findings before enforcing them.
5. Use Terraform `moved`/`import` and saved plans for adoption; never copy/edit
   state or run a production apply without the existing plan gate.

## Known limitations

- No production-cloud apply/smoke test was run for this candidate.
- Static migration, inventory, secret reference and offline dependency discovery
  can miss dynamic/transitive behavior.
- `cf doctor` requires project context; the prescribed invocation from `cli/`
  fails because that directory intentionally has no `clusterforge.yaml`.
- Module checks return success with warnings for modules lacking explicit status
  declarations.
- Full repository Terraform validation is long/provider-dependent and was not
  completed in this run; golden snapshots are intentionally validated by Go
  golden tests rather than as roots in their storage directory.

## Test results

| Command/check | Result | Evidence/notes |
|---|---|---|
| `make fmt-check` | PASS | Terraform and Go formatting clean |
| `make lint` | FAIL | TFLint cannot evaluate local module paths in three golden fixture directories; a no-module-call diagnostic also exposes 36 existing warnings |
| `make test` | INCOMPLETE | Go unit/e2e/golden and CLI build pass; nested Terraform validation did not finish |
| `make validate` | INCOMPLETE | Golden snapshots skipped with explicit reason; several AWS examples validated before duplicate long-running invocations were terminated |
| `make security` | FAIL | Gitleaks: 150 commits/~1.94 MB, no leaks; Checkov: 481 passed, 128 failed |
| `make check-modules` | PASS WITH WARNINGS | Exit 0, `status: warn`; multiple modules lack status metadata |
| `cd cli && go test ./...` | PASS | All CLI unit/e2e packages passed |
| `cd cli && go build -o cf .` | PASS | Local Linux binary built |
| `./cf version` | PASS | Reports `0.4.0-rc.1`; commit/date unknown without release ldflags |
| `./cf doctor` from `cli/` | FAIL | Project config missing; all preceding binary/tool/repository safety checks passed |
| doctor in temporary default scaffold | PASS WITH WARNING | Project/config/safety checks pass; expected warning because temp directory is not Git |
| required docs/examples presence | PASS | Requested docs, local/existing K8s, AWS EKS/ECS, policy/plugin and service catalog paths exist |

No result above implies production readiness or semantic validation of every
Terraform example.

## Release blockers

1. Triage/fix or explicitly baseline the TFLint gate so `make lint` passes for the
   intended source roots without hiding genuine warnings.
2. Triage 128 Checkov failures. Fix release-scope high-risk findings and document
   narrowly justified suppressions; do not blanket-ignore them.
3. Investigate GitHub Dependabot alert 2 (critical) and resolve it or prove with
   recorded dependency analysis that it is outside shipped/runtime scope.
4. Complete `make validate` and `make test` once, serially, from a clean checkout;
   preserve validated/skipped directory evidence.
5. Add explicit module stability/status metadata or make the warning disposition
   part of the supported module catalog.
6. Run the release workflow/archives, checksums, SBOM and supported-platform
   install smoke from the candidate commit.

## Smoke test status

- Local CLI build/version: passed.
- Default scaffold plus doctor in an isolated temporary directory: passed with a
  non-Git warning.
- Local/existing Kubernetes example apply: not run.
- AWS EKS/ECS plan/apply and cleanup: not run; no credentials used.
- Azure/GCP/edge/Windows production tests: not run and not claimed.
- Cross-platform packaged artifact install: not run.

## Security notes

Gitleaks found no leaks in the scanned history. This does not clear the Checkov or
Dependabot blockers. No credentials were supplied, no cloud APIs or Terraform
apply were invoked, and no state/kubeconfig was generated or committed. The local
`cli/cf` smoke binary is ignored build output and is not a release artifact.

## Recommended release decision

Do not tag or publish final `v0.4.0`. Keep `rc.1` as an internal assessment commit,
resolve the blockers, then create `rc.2` and rerun the full gate serially from a
clean checkout. Production-cloud checks may remain credential-gated, but their
absence must remain visible and corresponding support claims limited.
