# orchestrators/ecs/cluster

## Purpose

Creates an AWS ECS cluster suitable for Fargate services. The module manages
the ECS cluster, Container Insights setting, attached capacity providers, and
the default capacity provider strategy.

Provider configuration belongs in the root module. This module declares the
AWS provider requirement but does not configure the provider.

Use `modules/workloads/ecs/service` to create ECS task definitions and
services on top of this cluster.

## Root Usage

```hcl
provider "aws" {
  region = "us-east-1"
}

module "ecs_cluster" {
  source = "../../../modules/orchestrators/ecs/cluster"

  name        = "clusterforge-dev-ecs"
  environment = "dev"

  tags = {
    Project = "clusterforge"
  }
}
```

## Capacity Providers

By default the module attaches `FARGATE` and `FARGATE_SPOT`, and uses FARGATE
as the default strategy. Override `default_capacity_provider_strategy` in a
root module when you want a different default.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
