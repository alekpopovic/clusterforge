# v0.2 Release Gate

Date: 2026-07-09

## 1. Release Readiness Summary

Recommended status: **ready with warnings, blocked for final tag until lint and
release evidence are resolved**.

ClusterForge has broad repository validation coverage, usable CLI workflows,
documented safety gates, and experimental multi-cloud support. It is not ready
to tag as v0.2.0 without addressing the known lint failures and recording
smoke-test evidence honestly.

## 2. Supported Functionality

| Area | Status | Notes |
| --- | --- | --- |
| AWS EKS | Supported with warnings | Modules, examples, live roots, smoke runbook, and validation exist. Real smoke pass is not recorded. |
| AWS ECS | Supported with warnings | Cluster/service modules, examples, and validation exist. Real smoke pass is not recorded. |
| Existing Kubernetes | Partial | Workload examples and smoke runbook exist; first-class generated `existing-kubernetes` environment is planned for Prompt 083. |
| AKS | Experimental | Azure network and AKS MVP modules validate, but no real Azure smoke evidence exists. |
| GKE | Experimental | GCP network and GKE MVP modules validate, but no real GCP smoke evidence exists. |
| K3s/RKE2 | Experimental | User-data generation modules validate; no VM provisioning or conformance evidence yet. |
| App manifests | Supported | CLI app add/list/validate/render paths exist for Kubernetes-family and ECS targets. |
| Policy packs | Partial | Pack docs and CLI selectors exist; several checks remain advisory. |
| Template packs | Partial | Config, listing, validation, and env generation support exist; app renderer override remains limited. |

## 3. Test Status

| Check | Result | Evidence |
| --- | --- | --- |
| `make fmt-check` | Passed | Terraform fmt check and Go gofmt check passed. |
| `make lint` | Failed | TFLint found 38 warnings, mostly unused declarations in examples/placeholders and experimental modules. |
| `make test` | Passed | Go tests/build plus Terraform validation completed successfully. |
| `make validate` | Passed | 83 directories validated, 0 skipped; core Terraform native tests passed. |
| `make security` | Passed with skips | Checkov and Trivy were not installed, so no scanner ran. |
| `cd cli && go build -o cf .` | Passed | CLI binary built. |
| `./cf version` | Passed | Printed `version: dev`, `commit: unknown`, `date: unknown`, Go version. |
| `./cf doctor` | Failed in repo root | Expected hard failure because `clusterforge.yaml` does not exist in the repository root. Optional `tofu`, `kubectl`, and `helm` were missing. |
| Integration tests | Not run | Real integration tests are opt-in only. |
| Smoke tests | Not run | No real-cloud smoke test pass was claimed. |

## 4. Security Status

- Secrets handling: `.gitignore`, docs, smoke runbooks, and state helpers warn
  against committing credentials, kubeconfigs, tfstate, tfplan, and tfvars.
- Production apply protection: production apply requires a plan file and
  explicit confirmation.
- Destroy protection: production destroy is blocked by default.
- State protection: read-only state helper commands warn about sensitivity and
  require explicit output for state pull.
- Policy pack status: baseline/production/Kubernetes/AWS packs are documented,
  but some controls are advisory until OPA/Checkov rules are implemented.
- Supply chain: release workflow creates checksums; SBOM and signing are still
  documented follow-ups.

## 5. Documentation Status

| Area | Status |
| --- | --- |
| README | Present with quickstart, safety model, tutorials, and roadmap links. |
| CLI docs | Present and include operational helper commands. |
| Module docs | Required module READMEs exist; `MODULE_CATALOG.md` tracks stability. |
| Tutorials | Present for local Kubernetes, ECS, EKS, GitOps, and production safety. |
| Smoke tests | Present for AWS EKS, AWS ECS, existing Kubernetes, and cleanup. |
| Release docs | Present for release process, checklist, supply chain, and version support. |

## 6. Known Limitations

- Real cloud smoke tests have not been run or recorded for v0.2.
- TFLint currently fails when installed because warning-level findings are
  treated as lint failure.
- `cf doctor` fails in a plain repository root without `clusterforge.yaml`;
  this is useful for project health but noisy for release gate checks.
- AKS, GKE, K3s, and RKE2 remain experimental.
- Existing Kubernetes generation is not first-class yet.
- Several policy packs are advisory.
- Template pack app rendering is not a full override system yet.
- Checkov, Trivy, `tofu`, `kubectl`, and `helm` were unavailable in this local
  verification environment.

## 7. Required Fixes Before Release

Blockers before a final v0.2.0 tag:

- Decide whether TFLint warnings should fail `make lint`, or fix/suppress the
  38 current findings.
- Run and record at least one real smoke test or explicitly release with a
  documented no-smoke-test warning.
- Run security scanners in an environment with Checkov and/or Trivy installed,
  or document the skip in release notes.

Non-blocking warnings:

- Improve `cf doctor` release-check behavior for repository roots that are not
  initialized ClusterForge projects.
- Convert advisory policy pack items into enforceable checks.
- Add first-class existing Kubernetes generation.
- Add stronger generator golden tests and non-cloud e2e CLI tests.

Deferred tasks:

- Azure/GCP real-account smoke tests.
- K3s/RKE2 VM-based conformance checks.
- SBOM/signing implementation.
