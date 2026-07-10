# Multi-cluster inventory

ClusterForge can describe several clusters across environments and clouds without
contacting cloud APIs. Inventory is metadata; Terraform remains the source of
truth for infrastructure.

```yaml
environments:
  dev:
    cloud: aws
    region: eu-central-1
    orchestrator: eks
    path: live/dev/aws-eks
  prod:
    cloud: aws
    region: eu-central-1
    orchestrator: eks
    path: live/prod/aws-eks

clusters:
  dev-eks:
    environment: dev
    cloud: aws
    orchestrator: eks
    region: eu-central-1
    path: live/dev/aws-eks
    status: experimental
    kubeconfig_path: ~/.kube/config
  prod-eks:
    environment: prod
    cloud: aws
    orchestrator: eks
    region: eu-central-1
    path: live/prod/aws-eks
    status: production
```

`clusters` is optional. When absent, `cf cluster list` presents each existing
environment as a legacy inventory record, keeping `env list` and cluster views
consistent. Explicit clusters must reference a configured environment. Supported
statuses are `experimental`, `development`, `staging`, `production`, and
`deprecated`.

Commands:

```bash
cf cluster list
cf cluster list --json
cf cluster show prod-eks --json
cf cluster doctor prod-eks
cf cluster outputs prod-eks --json
cf cluster kubeconfig dev-eks
```

`cluster doctor` checks local metadata and paths only; it does not require live
cloud access. `cluster outputs` runs the configured Terraform/OpenTofu binary in
the inventory path and may therefore require backend access.

`kubeconfig_path` is only a path reference. ClusterForge never stores, prints, or
generates kubeconfig content through these inventory commands. Keep kubeconfig
files outside the repository and protect them as credentials.
