# ClusterForge MVP Acceptance Report

Date: 2026-06-13

## Recommendation

Recommended v0.1.0 status: **ready with warnings**.

ClusterForge is acceptable as an initial MVP for local evaluation, project
generation, Terraform review, and non-production experimentation. It is not
production-ready, and it should not be marketed as production-hardened.

The strongest acceptance signal is that a fresh user workflow works end to end:
project initialization, environment creation, Terraform generation, app
manifest creation, manifest validation, app rendering, and doctor diagnostics.

The main warnings are:

- No real cloud apply was tested.
- Security scanner findings from the previous release-candidate pass still need
  triage before recommending production use.
- TFLint is installed locally but its plugin handshake fails in this
  environment, so `make lint` passes with Terraform static lint degraded.
- Terraform validation depends heavily on provider plugin availability. In one
  full `make test` run, 66 directories validated and 1 provider-heavy module was
  skipped after timeout. A later standalone `make validate` run validated 19
  directories and skipped 48 due provider/plugin availability.

## Acceptance Questions

### 1. Can a new user clone the repo and understand it?

Yes.

Evidence:

- `README.md` explains the project pitch, architecture, quickstart, CLI usage,
  module development, validation, and roadmap.
- Required docs are present:
  - `docs/architecture.md`
  - `docs/cli.md`
  - `docs/app-manifest.md`
  - `docs/environments.md`
  - `docs/backends.md`
  - `docs/security.md`
  - `docs/gitops.md`
  - `docs/roadmap.md`
- `CONTRIBUTING.md`, `STATUS.md`, `FINAL_MVP_REPORT.md`, `CHANGELOG.md`, and
  `docs/release-checklist.md` add useful context for contributors.

### 2. Can a new user build the CLI?

Yes.

Command run:

```bash
cd cli && go build -o cf .
```

Result: passed.

Follow-up command:

```bash
cd cli && ./cf version
```

Result: passed.

Observed output:

```text
version: dev
commit: unknown
date: unknown
go: go1.26.1
```

Release builds should inject version, commit, and date via ldflags.

### 3. Can a new user run the core CLI workflow?

Yes.

Smoke test directory:

```text
/tmp/clusterforge-acceptance-thc9Gl
```

Commands run:

```bash
cf project init demo --non-interactive
cf env create dev --cloud aws --orchestrator eks --region eu-central-1 --non-interactive
cf generate dev
cf app add api --image ghcr.io/example/api:1.0.0 --port 8080 --replicas 2 --host api.dev.example.com --non-interactive
cf app validate api
cf app render api --env dev
cf doctor
```

Results:

- `cf project init`: passed.
- `cf env create`: passed.
- `cf generate`: passed.
- `cf app add`: passed.
- `cf app validate`: passed.
- `cf app render`: passed.
- `cf doctor`: passed in the generated project with one warning because the
  temporary directory is not a Git repository.

Generated app render path:

```text
live/dev/aws-eks/apps/api.tf
```

Repository-level doctor behavior:

```bash
cd cli && ./cf doctor
```

Result: failed as expected because the repository root under `cli/` is not a
ClusterForge project and has no `clusterforge.yaml`. This is acceptable for
`doctor`; it correctly reports the missing project config as a hard failure.

### 4. Can Terraform formatting pass?

Yes.

Command run:

```bash
make fmt-check
```

Result: passed.

### 5. Can Terraform validation pass where expected?

Yes, with provider/plugin caveats.

Command run:

```bash
make validate
```

Result: passed.

Summary from the final standalone run:

- Validated directories: 19
- Skipped directories: 48
- Core module `terraform test`: passed for `core/labels`, `core/naming`, and
  `core/tags`

Skip reason:

```text
terraform init could not complete in this local environment, usually due to provider download or plugin availability
```

Command run as part of `make test`:

```bash
make test
```

Result: passed.

Summary from that full run:

- Validated directories: 66
- Skipped directories: 1
- Skipped module: `modules/cloud/aws/tfstate-backend`
- Skip reason: Terraform init exceeded 45 seconds, likely waiting on provider
  installation or local plugin startup.

Acceptance interpretation:

Terraform syntax and module structure are in good shape. Provider-dependent
validation is still environment-sensitive and should be improved before relying
on it as a strict release gate.

### 6. Can Go tests pass?

Yes.

`make test` runs `scripts/test-cli.sh`, which runs Go module download, gofmt
check, `go test ./...`, and a CLI build check.

Result: passed.

### 7. Are safety policies enforced?

Mostly yes for the CLI policy surface.

Evidence:

- `cli/internal/policy/checks.go` enforces:
  - Production apply requires a plan file.
  - Production apply requires `--confirm-prod`.
  - Production delete actions require `--allow-destroy`.
  - Production destroy is blocked by default.
  - Production plan JSON parse failure fails closed.
- `cli/internal/policy/checks_test.go` covers production apply, destroy,
  delete, replacement, and parse-error behavior.
- `docs/security.md` and `docs/cli.md` document production apply/destroy
  requirements.

Not tested:

- A real production apply was not run.
- A real destructive Terraform plan was not applied.

### 8. Are secrets avoided in examples?

Yes, based on repository scan.

Scan command:

```bash
rg -n "AKIA|BEGIN (RSA |OPENSSH |EC |)PRIVATE KEY|aws_secret_access_key|password\\s*=\\s*\\\"|token\\s*=\\s*\\\"|kubeconfig|tfstate" . -g '!**/.terraform/**' -g '!**/.git/**'
```

Result:

- No committed AWS access keys, private keys, hardcoded provider tokens, or
  plain password assignments were found.
- Matches were documentation references, variable names, safe placeholders, and
  tests that verify detection of tracked `tfstate` and kubeconfig files.

### 9. Are production warnings clear?

Yes.

Production warnings appear in:

- `README.md`
- `docs/security.md`
- `docs/cli.md`
- `docs/backends.md`
- `docs/environments.md`
- live environment READMEs
- generated environment README templates
- module READMEs for Karpenter, observability, Docker, tfstate backend, and
  Helm app usage

The warnings are clear that:

- Production apply requires a reviewed plan file.
- Production destroy is blocked by default.
- Secrets must not be placed in `tfvars`.
- Production state should use remote backends with locking.
- Helm chart versions and values should be pinned/tuned before production.

### 10. Are placeholder modules clearly marked?

Yes.

Examples of clearly marked placeholder modules:

- `modules/cloud/aws/iam`
- `modules/cloud/aws/storage`
- `modules/orchestrators/kubernetes/generic`
- `modules/orchestrators/nomad/cluster`
- `modules/orchestrators/docker/engine`
- `modules/orchestrators/docker/swarm-service`
- `modules/platform/ecs/cloudwatch`
- `modules/platform/nomad/consul`
- `modules/platform/nomad/ingress`
- `modules/workloads/ecs/scheduled-task`
- `modules/workloads/nomad/batch`

These modules have README status sections and TODO comments in `main.tf`. They
create no fake resources.

## Passed Checks

- Repository has a coherent top-level README and required docs.
- All 45 Terraform module directories have the expected file contract:
  `main.tf`, `variables.tf`, `outputs.tf`, `versions.tf`, and `README.md`.
- CLI builds with `go build`.
- CLI version command works.
- Fresh-project CLI workflow passes:
  - project init
  - env create
  - generate
  - app add
  - app validate
  - app render
  - doctor
- `make fmt-check` passes.
- `make lint` exits successfully.
- `make test` passes.
- `make validate` passes.
- Core Terraform tests pass.
- Safety policy unit tests pass through `make test`.
- No real credentials were found in examples or docs.
- Placeholder modules are clearly marked.

## Failed Checks

No requested acceptance command ended in a final non-zero status during the
clean final run, except repository-level `cd cli && ./cf doctor`, which fails
correctly because `cli/` is not a generated ClusterForge project and does not
contain `clusterforge.yaml`.

Degraded checks:

- TFLint could not initialize plugins in this local environment. `make lint`
  reports the warning and continues with Go formatting and `go vet`.
- Provider-dependent Terraform validation is inconsistent depending on local
  plugin/cache state.

## Skipped Checks

- Real cloud `terraform plan` and `terraform apply` were not run.
- Kubernetes/Helm providers were not tested against a real cluster.
- Nomad and Docker providers were not tested against real runtimes.
- Production apply/destroy behavior was tested through unit tests and CLI policy
  code, not against real infrastructure.
- `make validate` skipped many provider-dependent roots in the standalone run
  because provider init could not complete in this local environment.

## Known Limitations

- v0.1.0 is an MVP, not a production platform guarantee.
- Security scan findings from `FINAL_MVP_REPORT.md` remain unresolved:
  - EKS public endpoint and public CIDR defaults.
  - Missing EKS secrets encryption configuration.
  - Public subnet public IP behavior.
  - ALB HTTP/public exposure warnings.
  - tfstate backend SSE-S3 vs customer-managed KMS key warning.
- Provider plugin initialization can be flaky in restricted or partially cached
  environments.
- Several modules are intentionally placeholder-only.
- No cloud credentials were used, so AWS, Kubernetes, Nomad, and Docker runtime
  behavior is not proven.
- Generated environments are readable and useful, but they still require user
  review before use.

## Required Fixes Before Public Production Recommendation

These are not blockers for a clearly labeled MVP release, but they are blockers
for production recommendation:

1. Triage and either fix or narrowly document security scanner findings.
2. Add EKS KMS secrets encryption support and safer public endpoint defaults.
3. Harden ECS ALB defaults or make public HTTP behavior more explicit.
4. Add optional KMS key support for the tfstate backend module.
5. Make provider-dependent validation more deterministic in CI and local
   development.
6. Run real AWS EKS and ECS integration tests in a sandbox account.
7. Add a release smoke script for the fresh-project CLI workflow.

## Command Results

| Command | Result | Notes |
| --- | --- | --- |
| `make fmt-check` | Pass | Terraform and Go formatting are clean. |
| `make lint` | Pass with warning | TFLint plugin initialization fails locally; Go formatting and `go vet` pass. |
| `make test` | Pass | CLI tests pass; Terraform validation validates 66 dirs and skips `modules/cloud/aws/tfstate-backend` after timeout. |
| `make validate` | Pass with skips | Validates 19 dirs and skips 48 provider-dependent roots in the final standalone run. |
| `cd cli && go build -o cf .` | Pass | Built CLI successfully. |
| `cd cli && ./cf version` | Pass | Prints dev build metadata and Go version. |
| `cd cli && ./cf doctor` | Expected fail | Fails because `cli/` is not a ClusterForge project and lacks `clusterforge.yaml`. |
| `cf doctor` in generated smoke project | Pass with warning | Warns that the temp project is not a Git repository. |

## Final Acceptance Decision

ClusterForge v0.1.0 is **ready with warnings** for MVP usage:

- good for repository evaluation,
- good for CLI workflow trials,
- good for readable Terraform generation,
- good for non-production experimentation,
- not ready for production infrastructure without further hardening and real
  cloud validation.
