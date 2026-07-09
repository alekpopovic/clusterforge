## Prompt 17 — CLI config loader

```text
Implement cli/internal/config package.

Files:
- schema.go
- loader.go
- writer.go
- defaults.go
- loader_test.go

Requirements:
- Define Go structs for clusterforge.yaml.
- Load YAML from disk.
- Validate required fields.
- Apply defaults:
  - project.default_engine = terraform
  - engines.terraform.binary = terraform
  - engines.opentofu.binary = tofu
  - defaults.cloud = aws
  - defaults.region = eu-central-1
  - defaults.orchestrator = eks
- Support saving config back to disk.
- Preserve enough structure for future use; exact comments do not need to be preserved.

Validation:
- project.name required
- default engine must exist in engines
- environment names must not be empty
- environment path must not be empty
- cloud must be one of aws, azure, gcp, hetzner, local for now
- orchestrator must be one of eks, ecs, kubernetes, nomad, docker, swarm, k3s, rke2, aks, gke for now

Tests:
- Load valid config.
- Fail on missing project name.
- Apply defaults.
- Add environment and save.
- Validate unknown orchestrator fails.

Use gopkg.in/yaml.v3 unless another YAML package is already used.

Run:
- gofmt
- go test ./...
```

---
