# Prompt State

Track numbered prompt execution here. Update this file whenever a prompt is
completed or intentionally skipped.

## Current Position

| Field | Value |
| --- | --- |
| Last executed prompt | `088-platform-conformance-tests-for-kubernetes-add-ons` |
| Next prompt to execute | `089-eks-production-hardening-options` |
| Prompt directory | `prompts/` |
| Last updated | 2026-07-09 |

## Rules

- `Last executed prompt` is the highest prompt that was actually implemented,
  verified, committed, and pushed.
- `Next prompt to execute` is the next numbered prompt that should be run by
  default.
- Splitting prompt files does not count as executing those prompts.
- If prompts are executed out of order, record the exception in `Notes`.
- Do not mark a prompt complete if real cloud, smoke, or integration evidence
  was required but not collected.
- After each prompt execution, update this file before running `git add`.
- Commit the prompt result and this state file together.
- The commit message for a prompt must be the prompt title without the prompt
  number.

## Current Prompt Template

Use this block in `Notes` when executing a prompt:

```text
Prompt: <NNN-slug>
Title: <prompt title without number>
Result: <completed | skipped | blocked>
Validation: <commands run or skip reason>
Evidence: <paths, reports, or not applicable>
Commit: <filled after commit if useful>
```

## Notes

- Prompts `000` through `080` have repository artifacts from prior execution.
- Prompts `081` through `120` have been split into files but have not been
  executed as implementation prompts yet.

```text
Prompt: 081-v0-2-release-gate-review
Title: v0.2 release gate review
Result: completed
Validation: make fmt-check passed; make lint failed on TFLint warnings; make test passed; make validate passed; make security skipped missing scanners; cli build/version passed; cf doctor failed without clusterforge.yaml
Evidence: RELEASE_GATE_V0.2.md
Commit: pending
```

```text
Prompt: 082-local-kubernetes-development-target-with-kind-or-k3d
Title: Local Kubernetes development target with Kind or K3d
Result: completed
Validation: terraform fmt -recursive examples/local-kind-app examples/local-k3d-app passed; gofmt passed; cd cli && go test ./... passed
Evidence: docs/local-development.md; examples/local-kind-app; examples/local-k3d-app; cli/cmd/local.go
Commit: pending
```

```text
Prompt: 083-existing-kubernetes-environment-support
Title: Existing Kubernetes environment support
Result: completed
Validation: terraform fmt -recursive examples/existing-kubernetes-basic-app examples/existing-kubernetes-platform-bootstrap passed; gofmt passed; cd cli && go test ./... passed
Evidence: docs/existing-kubernetes.md; cli/templates/env/existing-kubernetes; examples/existing-kubernetes-basic-app; examples/existing-kubernetes-platform-bootstrap
Commit: pending
```

```text
Prompt: 084-provider-compatibility-matrix-ci
Title: Provider compatibility matrix CI
Result: completed
Validation: git diff --check passed; YAML reviewed
Evidence: docs/provider-compatibility.md; .github/workflows/provider-compatibility.yml
Commit: pending
```

```text
Prompt: 085-golden-tests-for-cli-generators
Title: Golden tests for CLI generators
Result: completed
Validation: CLUSTERFORGE_UPDATE_GOLDEN=true go test ./internal/generator ./internal/app passed; cd cli && go test ./... passed; rg secret/path scan over cli/testdata/golden passed
Evidence: cli/testdata/golden; cli/internal/generator/golden_test.go; cli/internal/app/render_golden_test.go; docs/testing-generators.md
Commit: pending
```

```text
Prompt: 086-cli-end-to-end-non-cloud-tests
Title: CLI end-to-end non-cloud tests
Result: completed
Validation: gofmt passed; cd cli && go test ./... passed
Evidence: cli/e2e/project_init_test.go; cli/e2e/env_create_test.go; cli/e2e/generate_test.go; cli/e2e/app_flow_test.go; cli/e2e/policy_test.go
Commit: pending
```

```text
Prompt: 087-module-conformance-checker
Title: Module conformance checker
Result: completed
Validation: gofmt passed; cd cli && go test ./... passed; make check-modules passed with warning status; cf module check --json verified
Evidence: cli/internal/modulecheck; cli/cmd/module.go; .github/workflows/module-conformance.yml; docs/module-conformance.md
Commit: pending
```

```text
Prompt: 088-platform-conformance-tests-for-kubernetes-add-ons
Title: Platform conformance tests for Kubernetes add-ons
Result: completed
Validation: scripts/check-platform-modules.sh passed with bootstrap pass-through warnings; terraform fmt -recursive passed; git diff --check passed; cd cli && go test ./... passed
Evidence: scripts/check-platform-modules.sh; docs/platform-module-conventions.md
Commit: pending
```
