# workloads/nomad/service

## Purpose

Deploys a service-style Nomad job using the Terraform Nomad provider and the
Docker task driver.

Provider configuration belongs in the root module. This module declares the
Nomad provider requirement but does not configure the provider.

Use this module when a workload should run directly on Nomad instead of ECS,
Kubernetes, or Docker Swarm.

## Basic Service

```hcl
module "api" {
  source = "../../../modules/workloads/nomad/service"

  name        = "api"
  datacenters = ["dc1"]
  image       = "nginx:1.27"
  task_count  = 2
  cpu         = 500
  memory      = 256
}
```

## Ports And Service Registration

```hcl
module "web" {
  source = "../../../modules/workloads/nomad/service"

  name        = "web"
  datacenters = ["dc1"]
  image       = "hashicorp/http-echo:1.0"
  args        = ["-listen", ":8080", "-text", "hello from nomad"]

  ports = [
    {
      label = "http"
      to    = 8080
    }
  ]

  service = {
    enabled    = true
    name       = "web"
    port_label = "http"
    tags       = ["traefik.enable=true"]
  }
}
```

Service registration is useful for discovery and ingress, but a Nomad cluster
usually needs Consul or another service discovery integration for it to become
useful outside the job itself.

## Inputs

| Name | Description | Type | Default |
| --- | --- | --- | --- |
| `name` | Nomad job, group, and task name. | `string` | required |
| `datacenters` | Nomad datacenters where the job can run. | `list(string)` | required |
| `namespace` | Nomad namespace for the job. | `string` | `"default"` |
| `type` | Nomad job type. | `string` | `"service"` |
| `image` | Docker image reference for the task. | `string` | required |
| `task_count` | Number of task group instances. Terraform reserves `count`, so the module exposes `task_count` instead. | `number` | `1` |
| `command` | Optional command override. | `string` | `""` |
| `args` | Optional command arguments. | `list(string)` | `[]` |
| `env` | Plain environment variables. Do not put secrets here. | `map(string)` | `{}` |
| `ports` | Network ports exposed by the task group. | `list(object)` | `[]` |
| `cpu` | CPU shares allocated to the task. | `number` | `500` |
| `memory` | Memory allocated to the task in MiB. | `number` | `256` |
| `service` | Optional Nomad service registration. | `object` | disabled |

## Outputs

| Name | Description |
| --- | --- |
| `job_id` | Nomad job resource ID. |
| `job_name` | Nomad job name. |
