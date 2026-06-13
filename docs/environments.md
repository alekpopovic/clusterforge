---
title: Environments
permalink: /environments/
---

# Environments

ClusterForge environments describe where Terraform/OpenTofu roots live and
which cloud/orchestrator adapter they target.

## Simple Layout

Simple layout keeps one Terraform root per environment:

```text
live/dev/aws-eks/
  versions.tf
  backend.tf
  providers.tf
  main.tf
  variables.tf
  outputs.tf
```

Use it for learning, small systems, examples, and early prototyping.

```yaml
environments:
  dev:
    cloud: aws
    region: eu-central-1
    orchestrator: eks
    path: live/dev/aws-eks
    layout: simple
```

Generate it with:

```bash
cf generate dev --layout simple
```

## Stacked Layout

Stacked layout splits one environment into explicit Terraform roots:

```text
live/dev/aws-eks/
  network/
  cluster/
  platform/
  apps/
```

ClusterForge uses this dependency order:

1. `network`
2. `cluster`
3. `platform`
4. `apps`

Use stacked layout for production-style environments, larger teams, and
systems where network, cluster, platform add-ons, and applications should have
separate state files and review boundaries.

```yaml
environments:
  dev:
    cloud: aws
    region: eu-central-1
    orchestrator: eks
    path: live/dev/aws-eks
    layout: stacked
    stacks:
      network:
        path: live/dev/aws-eks/network
      cluster:
        path: live/dev/aws-eks/cluster
      platform:
        path: live/dev/aws-eks/platform
      apps:
        path: live/dev/aws-eks/apps
```

Generate it with:

```bash
cf generate dev --layout stacked
```

## Stack Commands

For stacked environments, commands default to all stacks in dependency order:

```bash
cf init dev
cf plan dev
cf apply dev
```

Run a single stack with `--stack`:

```bash
cf plan dev --stack network
cf apply dev --stack cluster
cf output dev --stack network --json
```

Destroy runs all stacks in reverse order when no stack is specified.

## Passing Outputs Between Stacks

Keep stack boundaries explicit. Downstream stacks should read upstream outputs
through `terraform_remote_state` or another reviewed handoff mechanism.

For local examples, a local backend can be used carefully:

```hcl
data "terraform_remote_state" "network" {
  backend = "local"

  config = {
    path = "../network/terraform.tfstate"
  }
}
```

For production, use a remote backend with locking. Avoid copying IDs manually
between `tfvars` files unless the values are intentionally static.

## Recommendation

Start with simple layout for experiments. Move to stacked layout before
production so state, plans, approvals, and blast radius are easier to reason
about.
