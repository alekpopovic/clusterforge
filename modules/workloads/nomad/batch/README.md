# Nomad batch workload

Submits a Docker-driver batch job through the root-configured Nomad provider.
Supports image, arguments, non-secret environment variables, allocation count,
CPU and memory. Do not pass secret values in `env`.

```hcl
module "batch" {
  source = "../../modules/workloads/nomad/batch"
  name   = "backup"
  image  = "alpine:3.20"
  args   = ["echo", "reviewed-job"]
}
```
