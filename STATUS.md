# ClusterForge Status

Last audited: 2026-06-11

## 1. Repository Summary

ClusterForge is in early but usable development. The repository has the intended
top-level structure, a working Go CLI foundation, several implemented Terraform
modules, practical examples, validation scripts, CI workflows, and project
documentation.

Main implemented areas:

- Core Terraform metadata modules: `modules/core/naming`, `modules/core/tags`,
  and `modules/core/labels`.
- AWS foundation/orchestrator path: `modules/cloud/aws/network`,
  `modules/orchestrators/kubernetes/eks`, and
  `modules/orchestrators/ecs/cluster`.
- Kubernetes platform Helm wrappers and bootstrap composition.
- Kubernetes app and CronJob workload modules.
- ECS Fargate service workload module.
- Nomad service workload module.
- Docker container and Swarm service workload modules.
- Go/Cobra CLI with config, environment generation, app manifests, Terraform
  runner, policy checks, and risk summaries.
- Examples for AWS network, EKS, ECS, Kubernetes app, Nomad, Docker Swarm, and
  core modules.

Main missing or placeholder-only areas:

- AWS DNS, IAM, and storage foundation modules are placeholders.
- Generic Kubernetes, Nomad cluster, Docker Engine, and Docker Swarm
  orchestrator modules are placeholders.
- ECS platform modules, Nomad platform modules, ECS scheduled task,
  Kubernetes Helm app, and Nomad batch modules are placeholders.
- Security policy directories exist, but custom Checkov and Conftest policies
  are not implemented yet.
- Helm chart versions are not pinned in the wrapper modules.
- Some initialized Terraform artifacts exist in the working tree as ignored
  files. They are not tracked, but the workspace should be cleaned periodically.

Risky assumptions:

- Terraform validation proves syntax and provider schema compatibility, not
  successful cloud deployment.
- EKS IAM/OIDC and add-on behavior still need apply-time testing in a real AWS
  account.
- ECS service defaults are useful for Fargate but need broader production
  review around IAM, load balancers, logging, and deployment controls.
- The CLI safety model relies on environment names `prod` or `production` for
  production-specific policy checks.

## 2. Terraform Module Status

Validation status reflects this audit run of `./scripts/validate.sh`, which
ran `terraform init -backend=false`, `terraform validate`, and available
`terraform test` suites.

| Module path | Status | Providers used | Has README | Has example | Validation status | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| `modules/core/naming` | implemented | none | yes | yes, `examples/core-naming` | pass, `terraform test` pass | Generates deterministic names and DNS/label-safe variants. |
| `modules/core/tags` | implemented | none | yes | yes, `examples/core-metadata` | pass, `terraform test` pass | Produces cloud tag maps. |
| `modules/core/labels` | implemented | none | yes | yes, `examples/core-metadata` | pass, `terraform test` pass | Produces Kubernetes-compatible labels. |
| `modules/cloud/aws/network` | implemented | `hashicorp/aws` | yes | yes, `examples/aws-network`, EKS/ECS examples | pass | Creates VPC, public/private subnets, routing, IGW, NAT, and Kubernetes subnet tags. |
| `modules/cloud/aws/iam` | placeholder | none | yes | no | pass | Placeholder only; no IAM resources yet. |
| `modules/cloud/aws/dns` | placeholder | none | yes | no | pass | Placeholder only; no Route 53 resources yet. |
| `modules/cloud/aws/storage` | placeholder | none | yes | no | pass | Placeholder only; no S3/EBS/EFS resources yet. |
| `modules/orchestrators/kubernetes/eks` | implemented | `hashicorp/aws` | yes | yes, `examples/aws-eks-minimal`, `live/dev/aws-eks` | pass | Creates EKS cluster, IAM roles, managed node groups, optional add-ons, and EBS CSI IRSA. |
| `modules/orchestrators/kubernetes/generic` | placeholder | none | yes | no | pass | Placeholder for future generic Kubernetes integration. |
| `modules/orchestrators/ecs/cluster` | implemented | `hashicorp/aws` | yes | yes, `examples/ecs-cluster-minimal`, `examples/ecs-fargate-app` | pass | Creates ECS cluster and Fargate capacity provider attachment. |
| `modules/orchestrators/nomad/cluster` | placeholder | none | yes | no | pass | Placeholder only; no Nomad server/client resources yet. |
| `modules/orchestrators/docker/engine` | placeholder | none | yes | no | pass | Placeholder only; no Docker host provisioning yet. |
| `modules/orchestrators/docker/swarm-service` | placeholder | none | yes | no | pass | Name overlaps with workload module; clarify intended responsibility. |
| `modules/platform/kubernetes/bootstrap` | implemented | child modules use Helm/Kubernetes | yes | no dedicated example | pass | Composes Helm wrapper modules conditionally. |
| `modules/platform/kubernetes/ingress-nginx` | implemented | `hashicorp/helm`, `hashicorp/kubernetes` | yes | via bootstrap only | pass | Helm release wrapper with optional namespace. |
| `modules/platform/kubernetes/cert-manager` | implemented | `hashicorp/helm`, `hashicorp/kubernetes` | yes | via bootstrap only | pass | Helm release wrapper with optional namespace. |
| `modules/platform/kubernetes/external-dns` | implemented | `hashicorp/helm`, `hashicorp/kubernetes` | yes | via bootstrap only | pass | Helm release wrapper; cloud IAM integration not included. |
| `modules/platform/kubernetes/metrics-server` | implemented | `hashicorp/helm`, `hashicorp/kubernetes` | yes | via bootstrap only | pass | Helm release wrapper. |
| `modules/platform/kubernetes/prometheus-stack` | implemented | `hashicorp/helm`, `hashicorp/kubernetes` | yes | via bootstrap only | pass | Helm release wrapper; production values need hardening. |
| `modules/platform/kubernetes/loki` | implemented | `hashicorp/helm`, `hashicorp/kubernetes` | yes | via bootstrap only | pass | Helm release wrapper; storage values need production design. |
| `modules/platform/kubernetes/argocd` | implemented | `hashicorp/helm`, `hashicorp/kubernetes` | yes | via bootstrap only | pass | Helm release wrapper; GitOps bootstrap policy needs design. |
| `modules/platform/ecs/alb` | placeholder | none | yes | no | pass | Placeholder only; ALB/listener/target group module not implemented. |
| `modules/platform/ecs/cloudwatch` | placeholder | none | yes | no | pass | Placeholder only; logging/alarms module not implemented. |
| `modules/platform/nomad/consul` | placeholder | none | yes | no | pass | Placeholder only; Consul integration not implemented. |
| `modules/platform/nomad/ingress` | placeholder | none | yes | no | pass | Placeholder only; ingress/proxy integration not implemented. |
| `modules/workloads/kubernetes/app` | implemented | `hashicorp/kubernetes` | yes | yes, `examples/kubernetes-basic-app` | pass | Deploys namespace, deployment, service, ingress, and HPA. |
| `modules/workloads/kubernetes/cronjob` | implemented | `hashicorp/kubernetes` | yes | no dedicated example | pass | Deploys namespace and CronJob; should get a practical example. |
| `modules/workloads/kubernetes/helm-app` | placeholder | none | yes | no | pass | Placeholder only; generic Helm workload wrapper not implemented. |
| `modules/workloads/ecs/service` | implemented | `hashicorp/aws` | yes | yes, `examples/ecs-fargate-app` | pass | Deploys ECS task definition, service, roles, logs, optional LB and autoscaling. |
| `modules/workloads/ecs/scheduled-task` | placeholder | none | yes | no | pass | Placeholder only; EventBridge/ECS scheduled task not implemented. |
| `modules/workloads/nomad/service` | implemented | `hashicorp/nomad` | yes | yes, `examples/nomad-service` | pass | Renders readable Nomad Docker service job spec. |
| `modules/workloads/nomad/batch` | placeholder | none | yes | no | pass | Placeholder only; batch job module not implemented. |
| `modules/workloads/docker/container` | implemented | `kreuzwerker/docker` | yes | no dedicated example | pass | Deploys Docker image and container. |
| `modules/workloads/docker/swarm-service` | implemented | `kreuzwerker/docker` | yes | yes, `examples/docker-swarm-service` | pass | Deploys Docker image and Swarm service. |

## 3. CLI Status

The CLI test suite passed with `GOCACHE` and `GOPATH` pointed at `/tmp`.

| Command | Status | Tests | Notes |
| --- | --- | --- | --- |
| `cf version` | implemented | indirect only | Prints static CLI version. |
| `cf project init <name>` | implemented | config/project behavior partially covered through config tests | Creates `clusterforge.yaml`, `apps/`, `live/`, and `.cf/`; uses `--force` for overwrite. |
| `cf env create <name>` | implemented | config tests cover environment save/load | Adds config environment and creates matching live path. |
| `cf env list` | implemented | indirect only | Lists environments from config. |
| `cf generate <env>` | implemented | generator unit tests | Supports first-version `aws+eks` and `aws+ecs`; has `--force`, `--dry-run`, and target overrides. |
| `cf init <env>` | implemented | runner argument tests | Runs selected Terraform/OpenTofu binary in env path. |
| `cf plan <env>` | implemented | runner and policy parser tests | Supports `--out` and `--risk-summary`; risk summary requires a plan output file. |
| `cf apply <env>` | implemented | policy tests and runner tests | Production apply requires `--plan-file` and `--confirm-prod`; prod delete plans require `--allow-destroy`. |
| `cf destroy <env>` | implemented | policy tests and runner tests | Production destroy blocked unless `--allow-destroy` and `--confirm-prod` are provided. |
| `cf doctor` | implemented | indirect only | Checks for Terraform/OpenTofu binaries. |
| `cf app add <name>` | implemented | app unit tests | Creates app manifests; no overwrite without `--force`. |
| `cf app list` | implemented | app unit tests | Lists manifests under `apps/`. |
| `cf app render <name> --env <env>` | implemented | app render tests | Renders Kubernetes-family and ECS module calls. |
| `cf app remove <name>` | implemented | indirect only | Removes an app manifest. |
| `cf policy check <env>` | implemented | policy and plan JSON tests | Summarizes plan risk when `--plan-file` is provided. |

CLI gaps:

- Command packages have useful short help, but limited long-form examples.
- Root command tests are limited; most tests target internal packages.
- Policy production detection is name-based and should later support explicit
  environment classification in `clusterforge.yaml`.

## 4. Examples Status

| Example path | Status | Can run `terraform validate`? | Requires cloud credentials? | Notes |
| --- | --- | --- | --- | --- |
| `examples/core-naming` | implemented | yes, pass | no | Demonstrates AWS, Kubernetes, and platform names. |
| `examples/core-metadata` | implemented | yes, pass | no | Demonstrates tags and labels modules. |
| `examples/aws-network` | implemented | yes, pass | yes for plan/apply | Safe VPC example with no backend and no real credentials. |
| `examples/aws-eks-minimal` | implemented | yes, pass | yes for plan/apply | Composes tags, network, and EKS. |
| `examples/ecs-cluster-minimal` | implemented | yes, pass | yes for plan/apply | Minimal ECS cluster composition. |
| `examples/ecs-fargate-app` | implemented | yes, pass | yes for plan/apply | Composes network, ECS cluster, security group, and service. |
| `examples/kubernetes-basic-app` | implemented | yes, pass | requires existing Kubernetes endpoint for plan/apply | Uses Kubernetes workload app module. |
| `examples/kubernetes-with-ingress` | placeholder | yes, pass | likely requires Kubernetes endpoint once implemented | Placeholder root; should become a real ingress example. |
| `examples/nomad-service` | implemented | yes, pass | requires Nomad endpoint for plan/apply | Uses Nomad service workload module. |
| `examples/docker-swarm-service` | implemented | yes, pass | requires Docker Swarm manager for plan/apply | Uses Docker Swarm workload module. |

## 5. Security Status

Secrets exposure risk:

- No committed credentials, kubeconfigs, private keys, `tfstate`, or plan files
  were found by this audit.
- Example secrets are references to existing Kubernetes secrets, SSM parameters,
  or Secrets Manager ARNs rather than plaintext values.
- `.tfvars` files are ignored, while safe `.tfvars.example` files are present.

tfstate protection:

- `.gitignore` excludes `.terraform/`, `*.tfstate`, `*.tfstate.*`, and
  `*.tfplan`.
- Backends are intentionally example/commented or omitted; production backends
  still need explicit S3/DynamoDB or equivalent guidance per environment.

Production safety:

- CLI production apply requires a plan file and explicit `--confirm-prod`.
- CLI production destroy is blocked by default and requires `--allow-destroy`
  plus `--confirm-prod`.
- Plan JSON delete actions in prod are blocked unless `--allow-destroy` is set.

Destructive command protections:

- The Terraform runner never adds `--auto-approve` by default.
- CLI policy tests cover production apply, destroy, delete actions,
  replacements, and parse failures.

Missing guardrails:

- Policy directories are placeholders; no custom Rego or Checkov policy rules
  are implemented.
- Production detection should become explicit metadata, not only environment
  name matching.
- Generated/live roots should eventually include stronger backend and lock-file
  guidance.
- Helm chart versions should be pinned before production use.

## 6. CI Status

Workflows present:

- `.github/workflows/terraform-validate.yml`
- `.github/workflows/cli-test.yml`
- `.github/workflows/security-scan.yml`

What each workflow does:

- Terraform Validate checks out the repo, installs Terraform, runs
  `./scripts/lint.sh`, and runs `./scripts/validate.sh` with fake AWS
  credentials for provider validation.
- CLI Test checks out the repo, installs Go 1.22, and runs
  `./scripts/test-cli.sh`.
- Security Scan runs Checkov Terraform scanning and Trivy config scanning.

What might fail:

- Security scans may fail on intentionally incomplete placeholder or example
  infrastructure unless suppressions or project-specific policies are added.
- Terraform validation may require network access to download providers on a
  fresh runner.
- Fake AWS credentials are suitable for validation, but not for plan/apply.

What needs improvement:

- Add custom Conftest/Checkov policies instead of relying only on generic
  scanner defaults.
- Add workflow syntax/lint checks if the repo grows more CI logic.
- Consider caching Terraform providers in CI.
- Add a matrix for Terraform and OpenTofu once OpenTofu is actively validated.

## 7. Immediate Next Tasks

Critical:

- Add explicit production classification to environment config and policy
  checks instead of relying only on the environment name.
- Pin Helm chart versions in platform modules and generated examples.
- Decide whether root `.terraform.lock.hcl` files should be committed; current
  `.gitignore` says no, but several ignored lock files exist in the workspace.
- Apply-test AWS EKS and ECS examples in a real sandbox account before claiming
  production readiness.

Important:

- Implement `modules/platform/ecs/alb` so ECS service examples do not need
  ad-hoc load balancer resources.
- Implement `modules/workloads/ecs/scheduled-task`.
- Add a practical `examples/kubernetes-with-ingress` implementation.
- Add custom Conftest and Checkov policies for secrets, public exposure, and
  production safety.
- Expand CLI command tests around Cobra command behavior and help output.

Nice to have:

- Add a dedicated Kubernetes CronJob example.
- Add examples for platform bootstrap add-ons.
- Add OpenTofu validation to scripts and CI.
- Add generated documentation tables for module inputs and outputs.
- Add a cleanup script for ignored local Terraform initialization artifacts.

## 8. Recommended Next Milestone

Next milestone: harden the first real production path, `aws+eks`, while keeping
the rest of the repository readable and honest about placeholders.

Recommended task order:

1. Add explicit environment classification such as `tier: dev|staging|prod` or
   `protected: true` to `clusterforge.yaml`, config validation, and policy
   checks.
2. Pin default Helm chart versions in all Kubernetes platform wrapper modules
   and document upgrade behavior.
3. Add a real `examples/kubernetes-with-ingress` composition using
   `modules/workloads/kubernetes/app`.
4. Implement a focused `modules/platform/ecs/alb` module.
5. Update `examples/ecs-fargate-app` to use the ECS ALB module.
6. Implement `modules/workloads/ecs/scheduled-task`.
7. Add custom Conftest policies for blocked plaintext secrets and production
   destroy/apply guardrails.
8. Add custom Checkov suppressions or policies for known-safe examples.
9. Apply-test `examples/aws-network`, `examples/aws-eks-minimal`, and
   `examples/ecs-fargate-app` in a sandbox AWS account.
10. Add OpenTofu validation support to CI after Terraform validation remains
    stable.

## Audit Commands Run

Passed:

- `terraform fmt -check -recursive`
- `env GOCACHE=/tmp/clusterforge-go-cache GOPATH=/tmp/clusterforge-go go test ./...`
  from `cli/`
- `./scripts/validate.sh`

Initial sandbox-limited failure:

- `go test ./...` from `cli/` failed because the default Go build cache under
  `/home/popac/.cache/go-build` is read-only in this environment. It passed
  after setting `GOCACHE` and `GOPATH` to `/tmp`.
