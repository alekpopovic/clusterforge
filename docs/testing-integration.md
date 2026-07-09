# Integration Testing

ClusterForge integration tests are optional real-environment checks. They are
never part of default CI and must be started explicitly:

```bash
CLUSTERFORGE_RUN_INTEGRATION_TESTS=true scripts/integration-test.sh aws-eks
CLUSTERFORGE_RUN_INTEGRATION_TESTS=true scripts/integration-test.sh aws-ecs
CLUSTERFORGE_RUN_INTEGRATION_TESTS=true scripts/integration-test.sh existing-kubernetes
```

The script prints a cost warning, creates a temporary generated project, uses a
unique environment name, and attempts cleanup through traps. AWS targets require
`AWS_REGION` and valid AWS credentials outside the repository. Existing
Kubernetes requires `KUBECONFIG`.

Do not run these tests against production accounts or clusters. Do not commit
tfvars, state, plan files, kubeconfigs, credentials, or unredacted evidence.
