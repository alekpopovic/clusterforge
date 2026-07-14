# ClusterForge v0.3.0 release gate review

Review date: 2026-07-11

Reviewed commit: `e580c65` (`main`)

Decision: **BLOCKED — do not release v0.3.0 yet**

## 1. Executive summary

ClusterForge is not ready for a v0.3.0 release. The Go unit, CLI end-to-end,
and golden generator test packages pass, the CLI builds, formatting passes,
Gitleaks found no leaks in the scanned repository history, and the module
conformance command completes with warnings. Documentation and feature breadth
are substantial.

The release gate is nevertheless blocked because the repository still declares
version `0.1.0` and CLI support `0.1.x`, the default Terraform validation and
lint targets fail on generated golden roots with invalid relative module paths,
the security target reports 128 Checkov failures, all recorded real smoke tests
remain `Not run`, and multiple modules do not declare a stability status. No
new features or release-blocking code fixes were made during this review.

## 2. Supported targets

| Target | Current classification | Release evidence |
| --- | --- | --- |
| AWS EKS | Supported in `VERSION_MATRIX.md`; primary modules are beta in the catalog | Validation/plan coverage exists, but real EKS smoke test is not run; security findings remain |
| AWS ECS | Supported in `VERSION_MATRIX.md`; primary modules are beta | Validation coverage exists, but real ECS and ECS+ALB smoke tests are not run |
| Existing Kubernetes | Supported | CLI templates/docs/tests exist; real existing-cluster smoke test is not run |
| Local Kind/K3d | Development target | Docs, CLI local commands, tests, and examples exist; no release evidence proves a full current lifecycle run |
| AKS | Experimental | Network/AKS modules exist; catalog records format-only testing and Checkov reports production-hardening gaps |
| GKE | Experimental | Network/GKE modules exist; catalog records format-only testing and Checkov reports production-hardening gaps |
| K3s/RKE2 | Experimental | User-data modules and examples exist; ClusterForge does not provision their hosts |
| Nomad | Experimental | Cluster, workload, and platform modules exist; production validation is incomplete |
| Docker/Swarm | Experimental | Engine/container/swarm modules and examples exist; production lifecycle commitment is not established |

Only targets with both a supported classification and current evidence should
be advertised as release-ready. Experimental targets must remain explicitly
experimental in v0.3 documentation.

## 3. CLI readiness

| Capability | Status | Evidence / limitation |
| --- | --- | --- |
| `project init` | Pass with warning | Covered by CLI e2e tests; no clean-install tutorial execution recorded in this review |
| `env create` | Pass with warning | Covered by unit/e2e and golden tests; generated complete roots currently break repository validation due to module paths |
| `generate` | Pass with warning | Golden tests pass; generated fixture validation fails |
| `app add/list/validate/render` | Pass | Unit/e2e and renderer golden packages pass |
| `doctor` | Conditional | Binary runs and validates installed tools; from repository root it exits 1 because `clusterforge.yaml` is absent, which is expected for a non-project directory |
| `drift check` | Beta | Read-only implementation/docs exist; no real backend/cloud evidence recorded |
| `fleet` commands | Beta | Unit tests pass and commands are documented as read-only; no heterogeneous real-fleet evidence |
| `policy` commands | Beta | Unit/e2e tests pass; policy coverage is not equivalent to full security enforcement |
| `module check` | Pass with warnings | Command exits 0, but 18 modules report missing declared status |
| `upgrade` commands | Experimental/beta | Commands and framework exist; no real migration/rollback evidence in the release gate |

The direct CLI build succeeded with Go 1.22.5. `cf version` returned `dev`,
`commit: unknown`, and `date: unknown` because the prompt-required direct build
did not supply release linker flags. The release workflow must prove tagged
version metadata before publishing.

## 4. Terraform module readiness

The repository contains approximately 70 module `main.tf` files and 44 example
roots. The catalog currently gives explicit classifications to only a subset:

- **Stable:** core naming, tags, and labels. These have Terraform native tests.
- **Beta:** AWS network, tfstate backend, DNS, IRSA, EKS, ECS cluster, Kubernetes
  app, and ECS service in the catalog. Coverage varies from native/plan tests to
  validation only.
- **Experimental:** Azure/AKS, GCP/GKE, K3s, RKE2, and ECS CodeDeploy blue/green
  entries in the catalog. Nomad and Docker are experimental in the version
  matrix.
- **Implicit beta/experimental:** the catalog says unlisted platform/workload
  modules must be treated as beta or experimental until reviewed. This is not
  precise enough for a release artifact.
- **Placeholder modules:** none were conclusively identified as an intentional
  empty placeholder by the conformance command; an explicit catalog audit is
  still required before release.
- **Missing examples:** the module checker verifies required module files, not
  a one-to-one example for every module. The catalog is incomplete, so absence
  cannot currently be certified reliably.
- **Missing tests:** only seven `*.tftest.hcl` files exist for roughly 70
  modules. Most modules rely on validation or format checks rather than native
  behavioral tests.

`make check-modules` exited 0 with warnings. The AWS network, ECS cluster, EKS,
K3s, RKE2, Kubernetes bootstrap/cert-manager/external-dns/ingress-nginx/
metrics-server, several Docker/ECS/Kubernetes/Nomad workload modules, and other
listed paths lack a declared module status according to the checker.

## 5. Security readiness

| Area | Status | Evidence / gap |
| --- | --- | --- |
| State safety | Partial | Ignore rules, backend guidance/modules, encryption and locking guidance exist; backend posture is operator-controlled |
| Secrets handling | Pass with warning | Gitleaks scanned 109 commits and about 1.56 MB with no leaks; external secret references are documented; scanner success cannot prove absence of every secret |
| Production apply safety | Implemented in CLI | Existing reviewed plan file and explicit production confirmation are required; direct Terraform use bypasses CLI controls |
| Destroy protection | Implemented in CLI | Production destroy is blocked by default and requires explicit flags; direct Terraform remains outside this protection |
| Image policy | Partial | Image-security workflow and policy guidance exist; enforcement depends on user configuration |
| Policy packs | Partial | Policy modules and CLI checks exist; coverage is advisory/optional in several paths |
| Secret scanning | Pass for this run | Gitleaks found no leaks; release CI evidence is still required on the final tagged commit |
| Audit logs | Implemented locally | CLI audit logging exists; retention, access, and SIEM export remain operator concerns |

`make security` failed. Checkov reported 481 passed and 128 failed Terraform
checks, including findings affecting public/network exposure, EKS/AKS/GKE
hardening, encryption, IAM scope, logging, database/cache resilience, S3,
container security contexts, probes, resources, and immutable images. Some
findings may be deliberate examples or require documented suppressions, but the
gate cannot accept them without triage. Checkov also could not resolve module
paths in several golden fixtures. Trivy did not run because `security.sh` stops
after the failing Checkov command.

## 6. Test readiness

| Test class | Result |
| --- | --- |
| Go unit tests | **Pass:** all packages completed successfully |
| CLI e2e tests | **Pass:** `cli/e2e` completed successfully |
| Golden generator tests | **Pass:** generator and renderer test packages completed successfully |
| Terraform native tests | **Not reached by default gate:** validation fails first; only seven test files exist, with core coverage documented |
| Module conformance | **Pass with warnings:** exit 0, multiple missing stability declarations |
| Smoke tests | **Not run:** every row in `SMOKE_TEST_MATRIX.md` remains `Not run` |
| Skipped validation | App-only golden fragments skip because they lack `versions.tf`; cloud/provider roots may require initialized providers or credentials as documented by validation scripts |

The principal deterministic test blocker is `make validate`: Terraform cannot
read `../../../modules` from `cli/testdata/golden/aws-ecs-simple` (and the same
layout affects other complete golden roots). `make lint` reports TFLint failures
for the AWS ECS, AWS EKS, and backend golden fixtures. Because `make test`
depends on validation, it exits non-zero after its Go test/build phase passes.

## 7. Documentation readiness

| Documentation area | Status |
| --- | --- |
| Root README | Present; architecture, quickstart, safety, development, and validation are covered |
| CLI docs | Present, with command-specific supporting documents |
| Architecture docs | Present |
| Module docs | Every inspected module contract has a README, but catalog/stability coverage is incomplete |
| Tutorials | Getting-started and onboarding examples exist; release acceptance still requires a verified install/quickstart/cleanup execution |
| Production guides | AWS EKS, operations, backends, security, promotion, and related guides exist; findings show hardening review is incomplete |
| Security docs | Security guide, threat model, checklist, policy and scanning docs exist and avoid compliance overclaims |
| Backup/DR docs | Backup/restore and multiple DR runbooks exist; no recorded real restore exercise proves them |
| Smoke-test docs | EKS, ECS, existing Kubernetes, and cleanup runbooks exist; evidence matrix is entirely not run |

## 8. Release blockers

### Critical

1. `VERSION` is `0.1.0`, `VERSION_MATRIX.md` declares CLI `0.1.x`, and direct
   build metadata is `dev/unknown`; release identity is not v0.3.0.
2. Required `make lint`, `make test`, and `make validate` gates fail due to
   broken module paths in complete generated golden roots.
3. `make security` fails with 128 untriaged Checkov findings; the scan does not
   proceed to Trivy.
4. No real smoke test or cleanup evidence exists for any supported target.

### Important

1. Stability is undeclared for multiple modules and most modules lack native
   Terraform tests.
2. The module catalog is incomplete and cannot provide a definitive inventory
   of missing examples/tests or registry readiness.
3. Restore, upgrade, drift, and fleet capabilities lack current real-environment
   evidence.
4. The final tagged release build, checksums, SBOM, version metadata, and
   installation verification have not been exercised for v0.3.0.

### Deferred (must remain explicitly experimental or out of scope)

- Production-grade AKS/GKE parity.
- Production Nomad and Docker/Swarm guarantees.
- Automatic remediation, write-capable fleet orchestration, plugin marketplace,
  hosted SaaS, and advanced service mesh defaults.

## 9. Recommended release decision

**Do not release v0.3.0 yet.** Fix the deterministic lint/validation problem,
triage security findings with fixes or narrowly justified suppressions, execute
at least one supported-target smoke test including cleanup, complete stability
classification, and align version/release metadata. Then rerun this gate from
the proposed release commit and retain command outputs as release evidence.

## Commands run

The following prompt-required commands were run on 2026-07-11:

| Command | Result |
| --- | --- |
| `make fmt-check` | Pass |
| `make lint` | Fail: TFLint failures in three complete golden fixtures |
| `make test` | Fail overall: Go tests/build pass, dependent Terraform validation fails |
| `make validate` | Fail: unreadable relative module directories in golden fixture |
| `make security` | Fail: Gitleaks pass, Checkov 481 pass/128 fail, Trivy not reached |
| `make check-modules` | Pass with warnings about missing module status |
| `cd cli && go test ./...` | Pass |
| `cd cli && go build -o cf .` | Pass |
| `cd cli && ./cf version` | Pass; reports development/unknown metadata |
| `cd cli && ./cf doctor` | Conditional fail: tool checks pass, root lacks `clusterforge.yaml` |

Tool versions used included Go 1.22.5, Terraform 1.15.8, OpenTofu 1.8.2,
TFLint 0.63.1, Checkov 3.3.8, Trivy 0.72.0, kubectl 1.30.3, Helm 3.15.3,
pre-commit 3.7.1, and Gitleaks 8.30.1.

## Files changed by this review

- `RELEASE_GATE_V0.3.md` — added this evidence-based release decision.

The locally built `cli/cf` is a disposable artifact and is not part of the
release review change set.
