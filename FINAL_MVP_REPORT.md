# ClusterForge v0.1.0 MVP Report

Date: 2026-06-13

## Release Candidate Status

ClusterForge has enough repository structure, modules, CLI surface, examples,
documentation, and validation automation to be considered a v0.1.0 MVP release
candidate.

The release should not be tagged yet. Two blockers remain:

- `make security` reports real HIGH/CRITICAL findings that need triage or
  documented, narrow exceptions.
- A full `make test` run can wedge locally when already-initialized Terraform
  provider plugins are reused, even though CLI tests pass and bounded
  validation completes with explicit skips.

No real cloud apply was tested.

## Included in v0.1.0

- Four-layer Terraform framework layout:
  - Foundation
  - Orchestrator
  - Platform
  - Workload
- Core modules:
  - `modules/core/naming`
  - `modules/core/tags`
  - `modules/core/labels`
- AWS foundation modules:
  - VPC/networking
  - Route53 DNS
  - Terraform state backend resources
  - IRSA role helpers
  - Karpenter IRSA foundations
- Orchestrator modules:
  - AWS EKS with managed node groups, add-ons, OIDC, IRSA foundations, and EBS
    CSI integration support
  - AWS ECS cluster
  - Placeholder or lightweight foundations for generic Kubernetes, Nomad, and
    Docker orchestrator paths
- Platform modules:
  - Kubernetes Helm add-ons and bootstrap composition
  - Argo CD app-of-apps bootstrap
  - External Secrets Operator and AWS SecretStore manifests
  - cert-manager issuer resources
  - Karpenter Helm install
  - Observability modules
  - Pod Security and NetworkPolicy baseline modules
  - ECS Application Load Balancer module
- Workload modules:
  - Kubernetes app, worker, CronJob, and Helm app
  - ECS Fargate service
  - Nomad service
  - Docker container and Swarm service
- Go CLI:
  - Project/environment initialization
  - Simple and stacked environment generation
  - AWS EKS and AWS ECS generation
  - App manifest add/list/validate/render
  - Backend configuration generation
  - Terraform/OpenTofu runner
  - Policy and risk summary support
  - Doctor command
  - JSON output for selected commands
  - Shell completion
  - Build metadata support
- Documentation:
  - Architecture
  - CLI
  - App manifests
  - Environments
  - Backends
  - Security
  - GitOps
  - Roadmap
  - Validation
  - Observability
  - Autoscaling
  - Release checklist
- Developer automation:
  - `Makefile`
  - Formatting, linting, validation, CLI test, docs, security, and build scripts
  - GitHub Actions workflows for validation, CLI tests, security scans, docs,
    and release artifacts

## Not Included

- Verified production AWS EKS or ECS apply.
- Verified Kubernetes add-on installation against a real cluster.
- Verified GitOps bootstrap against a real Argo CD installation.
- AKS, GKE, K3s, and RKE2 production implementations.
- Complete production IAM review for every AWS module.
- Production chart version recommendations for every Helm-based platform
  module.
- Automated backend bucket bootstrapping from the same root that consumes the
  backend. This is intentionally avoided.

## Known Limitations

- Terraform state can contain sensitive values. Users must use remote state,
  encryption, access controls, and external secret stores.
- Some validation paths require provider plugins that are not always available
  in local restricted environments.
- Security scanners flag intentionally public infrastructure examples, such as
  public subnets and public ALBs. These need narrow triage before release.
- EKS public endpoint defaults and missing secrets encryption are not production
  hardened enough for a release tag.
- The Helm app module now targets the Helm provider v3 schema.
- The CLI is an MVP helper. Terraform/OpenTofu remains the source of truth.

## Commands Run

| Command | Status | Notes |
| --- | --- | --- |
| `make fmt` | Passed | Ran `terraform fmt -recursive` and `gofmt`. |
| `make lint` | Passed with warning | Fixed TFLint config path handling and Go 1.22 test compatibility. TFLint plugin initialization is unhealthy in this local environment, so lint warns and skips TFLint static analysis; `gofmt` and `go vet` pass. |
| `make test` | Failed/interrupted | CLI test phase passed. Full command later wedged at Terraform validation for `modules/cloud/aws/tfstate-backend` when initialized provider plugins were reused. |
| `timeout 180s make validate` | Passed with skips | 19 directories validated, 48 skipped with explicit provider/plugin availability reasons, and core `terraform test` checks passed. |
| `GOCACHE=/tmp/clusterforge-go-cache GOPATH=/tmp/clusterforge-go ./scripts/test-cli.sh` | Passed | CLI unit tests and CLI build check passed independently. |
| `bash -n scripts/validate.sh scripts/lint.sh` | Passed | Shell syntax check passed for touched scripts. |
| `make security` | Failed | Checkov is not installed. Trivy ran with embedded checks after network fallback and reported HIGH/CRITICAL findings. |
| `cd cli && go build -o cf .` | Failed | Default Go build cache under `/home/popac/.cache` is read-only in this sandbox. |
| `cd cli && GOCACHE=/tmp/clusterforge-go-cache GOPATH=/tmp/clusterforge-go go build -o cf .` | Passed | Built `cli/cf`. |
| `./cli/cf version` | Passed | Printed dev build metadata and Go version. |

## Security Scan Findings To Triage

Trivy reported findings including:

- Public subnet public IP assignment in `modules/cloud/aws/network`.
- Terraform state bucket using SSE-S3 instead of a customer-managed KMS key in
  `modules/cloud/aws/tfstate-backend`.
- EKS cluster secrets encryption not configured.
- EKS public endpoint access and `0.0.0.0/0` public access CIDR defaults.
- Public ALB exposure, HTTP listener usage, unrestricted ALB egress, and missing
  invalid-header dropping in `modules/platform/ecs/alb`.

Some findings are expected for examples, but they should be handled with either
secure defaults, new inputs, or narrow documented scanner exceptions.

## Recommended Next Version Goals

1. Harden EKS defaults:
   - Add optional cluster KMS encryption.
   - Make public endpoint CIDRs safer in examples.
   - Document private endpoint production patterns.
2. Harden ECS ALB:
   - Enable invalid-header dropping.
   - Add internal ALB option.
   - Improve HTTPS examples with certificate placeholders.
3. Improve tfstate backend:
   - Add optional KMS key ARN input.
   - Document SSE-S3 vs SSE-KMS tradeoffs.
4. Stabilize validation:
   - Avoid provider plugin reuse hangs in `make test`.
   - Consider validating in clean temporary plugin/cache directories.
5. Triage Trivy/Checkov rules:
   - Fix real risks.
   - Add narrow exceptions only for intentional public examples.
6. Run real AWS EKS and ECS integration tests in a sandbox account.
7. Pin recommended Helm chart versions for production examples.
8. Add release smoke tests for `cf project init`, `cf env create`,
   `cf generate`, and `cf app render` in a temporary directory.

## Verdict

v0.1.0 is not ready to tag yet.

The MVP content is assembled, but the release candidate is blocked by security
scan findings and the full local `make test` stability issue. Once those are
triaged, the project is close to an honest initial release.
