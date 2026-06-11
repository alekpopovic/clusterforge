# workloads/docker/swarm-service

## Purpose

Deploys a Docker Swarm service using the Terraform Docker provider.

Provider configuration belongs in the root module. This module declares the
`kreuzwerker/docker` provider requirement but does not configure the provider.

This target is useful for simple self-hosted clusters and small edge
deployments. It is not recommended as the main production target when
Kubernetes, ECS, or Nomad are available.

## Example

```hcl
module "web" {
  source = "../../../modules/workloads/docker/swarm-service"

  name     = "web"
  image    = "nginx:1.27"
  replicas = 2

  ports = [
    {
      target_port    = 80
      published_port = 8080
    }
  ]

  labels = {
    "clusterforge.io/workload" = "web"
  }
}
```

## Example With Environment And Network

```hcl
module "worker" {
  source = "../../../modules/workloads/docker/swarm-service"

  name     = "worker"
  image    = "busybox:1.36"
  replicas = 1
  networks = ["app-net"]

  env = {
    APP_ENV = "dev"
  }
}
```

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
