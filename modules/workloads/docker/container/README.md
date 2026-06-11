# workloads/docker/container

## Purpose

Runs a single Docker container on a local or self-hosted Docker Engine using
the Terraform Docker provider.

Provider configuration belongs in the root module. This module declares the
`kreuzwerker/docker` provider requirement but does not configure the provider.

This target is useful for development, homelab, edge, or very small
self-hosted deployments. It is not recommended as the main production target
when Kubernetes, ECS, or Nomad are available.

## Example

```hcl
module "web" {
  source = "../../../modules/workloads/docker/container"

  name  = "web"
  image = "nginx:1.27"

  ports = [
    {
      internal = 80
      external = 8080
    }
  ]

  labels = {
    "clusterforge.io/workload" = "web"
  }
}
```

## Example With Env And Volume

```hcl
module "app" {
  source = "../../../modules/workloads/docker/container"

  name  = "app"
  image = "busybox:1.36"
  command = [
    "sh",
    "-c",
    "while true; do echo hello; sleep 30; done"
  ]

  env = {
    APP_ENV = "dev"
  }

  volumes = [
    {
      host_path      = "/srv/app"
      container_path = "/app"
      read_only      = true
    }
  ]
}
```

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
