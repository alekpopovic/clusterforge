# Smoke Test Matrix

ClusterForge smoke tests are manual real-environment checks. They are not part
of default CI because they can create billable cloud resources and require real
credentials or cluster access.

Do not mark a row as passed unless the smoke test was actually run. Do not add
real account IDs, credentials, kubeconfigs, plan files, state files, private
keys, or unredacted endpoint details.

## Matrix

| Provider | Orchestrator | Status | Last tested date | Tester | Terraform/OpenTofu version | Kubernetes version | Result | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| AWS | EKS | Not run | `<YYYY-MM-DD or not run>` | `<tester>` | `<version>` | `<version>` | `<not run, passed, failed, blocked>` | Evidence: `<path or issue link>` |
| AWS | ECS | Not run | `<YYYY-MM-DD or not run>` | `<tester>` | `<version>` | `N/A` | `<not run, passed, failed, blocked>` | Evidence: `<path or issue link>` |
| Existing cluster | Kubernetes | Not run | `<YYYY-MM-DD or not run>` | `<tester>` | `<version>` | `<version>` | `<not run, passed, failed, blocked>` | Evidence: `<path or issue link>` |
| AWS | EKS with Route53 | Not run | `<YYYY-MM-DD or not run>` | `<tester>` | `<version>` | `<version>` | `<not run, passed, failed, blocked>` | Optional DNS path. Evidence: `<path or issue link>` |
| AWS | ECS with ALB | Not run | `<YYYY-MM-DD or not run>` | `<tester>` | `<version>` | `N/A` | `<not run, passed, failed, blocked>` | Optional workload path. Evidence: `<path or issue link>` |

## Evidence Requirements

Each completed run should link to redacted evidence stored outside the
repository or in an approved redacted location:

- tool versions
- redacted identity or cluster context confirmation
- saved plan summary
- apply result summary
- orchestrator health checks
- workload or demo app checks
- endpoint or ingress status, when available
- cleanup confirmation

## Runbooks

- [AWS EKS smoke test](docs/smoke-tests/aws-eks.md)
- [AWS ECS smoke test](docs/smoke-tests/aws-ecs.md)
- [Existing Kubernetes cluster smoke test](docs/smoke-tests/kubernetes-existing-cluster.md)
- [Cleanup checklist](docs/smoke-tests/cleanup.md)
