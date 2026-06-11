# docker-swarm-service

Example root module for `modules/workloads/docker/swarm-service`.

This example deploys an `nginx` Docker Swarm service with one published port.

## Provider

The Docker provider uses `var.docker_host`, defaulting to:

```text
unix:///var/run/docker.sock
```

Run against a Docker Swarm manager:

```bash
terraform init
terraform validate
terraform plan
```

The Docker daemon must be reachable from the machine running Terraform. The
target Docker Engine must be initialized as a Swarm manager for
`docker_service` resources.
