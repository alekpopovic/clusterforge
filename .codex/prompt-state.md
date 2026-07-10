# Prompt State

Track numbered prompt execution here. Update this file whenever a prompt is
completed or intentionally skipped.

## Current Position

| Field | Value |
| --- | --- |
| Last executed prompt | `114-pre-commit-hooks` |
| Next prompt to execute | `115-secret-scanning-baseline` |
| Prompt directory | `prompts/` |
| Last updated | 2026-07-10 |

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

```text
Prompt: 089-eks-production-hardening-options
Title: EKS production hardening options
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for modules/orchestrators/kubernetes/eks and examples/aws-eks-production-hardened; git diff --check passed
Evidence: modules/orchestrators/kubernetes/eks; docs/aws-eks-production.md; examples/aws-eks-production-hardened
Commit: pending
```

```text
Prompt: 090-aws-kms-reusable-module
Title: AWS KMS reusable module
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for modules/cloud/aws/kms-key and modules/cloud/aws/tfstate-backend; make check-modules passed with warning status; git diff --check passed
Evidence: modules/cloud/aws/kms-key; modules/cloud/aws/tfstate-backend; docs/aws-eks-production.md
Commit: pending
```

```text
Prompt: 091-aws-vpc-endpoints-module
Title: AWS VPC endpoints module
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for modules/cloud/aws/vpc-endpoints and examples/aws-vpc-endpoints-private-eks; git diff --check passed
Evidence: modules/cloud/aws/vpc-endpoints; examples/aws-vpc-endpoints-private-eks
Commit: pending
```

```text
Prompt: 092-ecr-registry-module
Title: ECR registry module
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for modules/cloud/aws/ecr and examples/aws-ecr-repositories; git diff --check passed
Evidence: modules/cloud/aws/ecr; examples/aws-ecr-repositories
Commit: pending
```

```text
Prompt: 093-container-image-security-workflow
Title: Container image security workflow
Result: completed
Validation: gofmt passed; cd cli && go test ./... passed; git diff --check passed
Evidence: docs/image-security.md; .github/workflows/image-security-example.yml; cli/internal/app image policy warnings
Commit: pending
```

```text
Prompt: 094-velero-backup-module
Title: Velero backup module
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for modules/cloud/aws/velero-backup-bucket, modules/platform/kubernetes/velero, and examples/aws-eks-velero; make check-modules passed with warning status; git diff --check passed
Evidence: modules/platform/kubernetes/velero; modules/cloud/aws/velero-backup-bucket; docs/backup-restore.md; examples/aws-eks-velero
Commit: pending
```

```text
Prompt: 095-disaster-recovery-runbooks
Title: Disaster recovery runbooks
Result: completed
Validation: git diff --check passed; DR section coverage scan passed for all runbooks; no real account IDs or fake RTO/RPO guarantees added
Evidence: docs/dr; docs/operations.md; docs/security.md; README.md
Commit: pending
```

```text
Prompt: 096-external-dns-production-hardening
Title: External DNS production hardening
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for modules/cloud/aws/external-dns-irsa, modules/platform/kubernetes/external-dns, and examples/aws-eks-external-dns; make check-modules passed with warning status; git diff --check passed
Evidence: modules/cloud/aws/external-dns-irsa; modules/platform/kubernetes/external-dns; examples/aws-eks-external-dns
Commit: pending
```

```text
Prompt: 097-cert-manager-route53-dns01-iam-module
Title: Cert-manager Route53 DNS01 IAM module
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for modules/cloud/aws/cert-manager-route53-irsa, modules/platform/kubernetes/cert-manager, modules/platform/kubernetes/cert-manager-issuer, and examples/aws-eks-cert-manager-dns01; make check-modules passed with warning status; git diff --check passed
Evidence: modules/cloud/aws/cert-manager-route53-irsa; docs/tls-cert-manager.md; examples/aws-eks-cert-manager-dns01
Commit: pending
```

```text
Prompt: 098-aws-rds-postgresql-module
Title: AWS RDS PostgreSQL module
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for modules/cloud/aws/rds-postgres and examples/aws-rds-postgres; make check-modules passed with warning status; git diff --check passed
Evidence: modules/cloud/aws/rds-postgres; examples/aws-rds-postgres
Commit: pending
```

```text
Prompt: 099-aws-elasticache-redis-module
Title: AWS ElastiCache Redis module
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for modules/cloud/aws/elasticache-redis and examples/aws-elasticache-redis; make check-modules passed with warning status; git diff --check passed
Evidence: modules/cloud/aws/elasticache-redis; examples/aws-elasticache-redis
Commit: pending
```

```text
Prompt: 100-aws-messaging-modules-sqs-and-sns
Title: AWS messaging modules: SQS and SNS
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for modules/cloud/aws/sqs, modules/cloud/aws/sns, examples/aws-sqs-worker, and examples/aws-sns-topic; make check-modules passed with warning status; git diff --check passed
Evidence: modules/cloud/aws/sqs; modules/cloud/aws/sns; examples/aws-sqs-worker; examples/aws-sns-topic
Commit: pending
```

```text
Prompt: 101-workload-cloud-identity-abstraction
Title: Workload cloud identity abstraction
Result: completed
Validation: gofmt passed with a temporary Go toolchain; cd cli && go test ./... passed; terraform fmt passed; terraform validate passed for ECS service, Kubernetes worker, and Kubernetes cronjob modules; git diff --check passed
Evidence: docs/rfcs/007-workload-identity.md; CLI manifest rendering tests; ECS task role policy support; Kubernetes annotated service accounts
Commit: pending
```

```text
Prompt: 102-service-binding-pattern-for-apps
Title: Service binding pattern for apps
Result: completed
Validation: gofmt passed with a temporary Go toolchain; cd cli && go test ./... passed; git diff --check passed
Evidence: cli/internal/bindings; app manifest dependency rendering tests; docs/service-bindings.md
Commit: pending
```

```text
Prompt: 103-kubernetes-tenant-model
Title: Kubernetes tenant model
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for the tenant module and example; make check-modules completed with only pre-existing warning status; git diff --check passed
Evidence: modules/platform/kubernetes/tenant; examples/kubernetes-tenant; docs/kubernetes-tenancy.md
Commit: pending
```

```text
Prompt: 104-resourcequota-and-limitrange-baseline-modules
Title: ResourceQuota and LimitRange baseline modules
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for both modules and the example; make check-modules passed for both new modules with only pre-existing repository warnings; git diff --check passed
Evidence: modules/platform/kubernetes/resource-quota; modules/platform/kubernetes/limit-range; examples/kubernetes-resource-governance; docs/kubernetes-resource-governance.md
Commit: pending
```

```text
Prompt: 105-kyverno-policy-module
Title: Kyverno policy module
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for the Kyverno module and example; make check-modules passed for the new module with only pre-existing repository warnings; git diff --check passed; live Helm install and policy admission tests skipped because no cluster was provided
Evidence: modules/platform/kubernetes/kyverno; examples/kubernetes-kyverno-baseline; docs/kubernetes-policy-kyverno.md
Commit: pending
```

```text
Prompt: 106-opa-gatekeeper-alternative-module
Title: OPA Gatekeeper alternative module
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for the Gatekeeper module and example; make check-modules passed for the new module with only pre-existing repository warnings; git diff --check passed; live admission tests skipped because no cluster was provided
Evidence: modules/platform/kubernetes/gatekeeper; examples/kubernetes-gatekeeper-baseline; docs/kubernetes-policy-gatekeeper.md
Commit: pending
```

```text
Prompt: 107-progressive-delivery-with-argo-rollouts
Title: Progressive delivery with Argo Rollouts
Result: completed
Validation: terraform fmt -recursive passed; terraform validate passed for controller module, rollout workload module, and example; make check-modules passed for both new modules with only pre-existing repository warnings; git diff --check passed; live rollout tests skipped because no cluster was provided
Evidence: modules/platform/kubernetes/argo-rollouts; modules/workloads/kubernetes/rollout-app; examples/kubernetes-argo-rollouts-canary; docs/progressive-delivery.md
Commit: pending
```

```text
Prompt: 108-service-mesh-rfc
Title: Service mesh RFC
Result: completed
Validation: git diff --check passed; docs build skipped because mkdocs-material is unavailable
Evidence: docs/rfcs/008-service-mesh.md; docs/roadmap.md
Commit: pending
```

```text
Prompt: 109-multi-cluster-inventory-model
Title: Multi-cluster inventory model
Result: completed
Validation: gofmt passed with a temporary Go toolchain; cd cli && go test ./... passed; git diff --check passed
Evidence: cli/internal/config cluster schema; cli/internal/inventory; cli/cmd/cluster.go; docs/multi-cluster.md
Commit: pending
```

```text
Prompt: 110-fleet-operations-cli
Title: Fleet operations CLI
Result: completed
Validation: gofmt passed with a temporary Go toolchain; cd cli && go test ./... passed; git diff --check passed; real fleet drift was not run because no initialized fleet/backend credentials were provided
Evidence: cli/internal/fleet; cli/cmd/fleet.go; fleet aggregation and filter tests; docs/fleet-operations.md
Commit: pending
```

```text
Prompt: 111-environment-graph-visualization
Title: Environment graph visualization
Result: completed
Validation: gofmt passed with a temporary Go toolchain; cd cli && go test ./... passed; git diff --check passed; Terraform graph execution was not run because no initialized target root was requested
Evidence: cli/internal/graph; cli/cmd/graph.go; Terraform runner graph support; docs/graphs.md
Commit: pending
```

```text
Prompt: 112-scheduled-drift-check-workflow-templates
Title: Scheduled drift check workflow templates
Result: completed
Validation: git diff --check passed; both workflow examples parsed successfully as YAML with Ruby; actionlint was unavailable; workflows were not executed because they are disabled examples and no AWS OIDC role was provided
Evidence: .github/workflows/examples/drift-check-aws-eks.yml; .github/workflows/examples/drift-check-aws-ecs.yml; docs/ci-drift-checks.md
Commit: pending
```

```text
Prompt: 113-cli-audit-log
Title: CLI audit log
Result: completed
Validation: gofmt passed with a temporary Go toolchain; cd cli && go test ./... passed; git diff --check passed
Evidence: cli/internal/audit; CLI audit middleware and commands; audit redaction/disabled/confirmation tests; docs/audit-log.md; .gitignore
Commit: pending
```

```text
Prompt: 114-pre-commit-hooks
Title: Pre-commit hooks
Result: completed
Validation: .pre-commit-config.yaml parsed successfully with Ruby; shell scripts passed bash -n; optional wrappers reported gitleaks and shellcheck unavailable; git diff --check passed; pre-commit run --all-files skipped because pre-commit is unavailable
Evidence: .pre-commit-config.yaml; scripts/install-hooks.sh; scripts/pre-commit-secret-scan.sh; scripts/pre-commit-shellcheck.sh; docs/pre-commit.md; README.md; CONTRIBUTING.md
Commit: pending
```
