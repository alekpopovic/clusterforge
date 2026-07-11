## Prompt 134 — Nomad production MVP

```text
Implement Nomad production MVP support.

Goal:
Make Nomad support more than a single job wrapper.

Create or improve modules:
- modules/orchestrators/nomad/cluster
- modules/platform/nomad/consul
- modules/platform/nomad/ingress
- modules/workloads/nomad/service
- modules/workloads/nomad/batch

Nomad cluster module:
- For MVP, avoid provisioning servers directly if too provider-specific.
- Support configuration templates and bootstrap docs.
- Optionally support cloud-init user_data generation.
- Output install/config notes.

Consul module:
- Install/configure Consul integration if practical.
- Otherwise provide docs and example config.

Nomad workload modules:
- service job
- batch job
- Docker driver
- resources
- env
- ports
- service registration

CLI:
- cf env create dev --cloud existing --orchestrator nomad
- cf generate dev

Docs:
- docs/nomad.md
- docs/nomad-production.md

Examples:
- examples/nomad-service
- examples/nomad-batch
- examples/existing-nomad

Rules:
- Do not pretend full Nomad cluster automation is done if it is only documented.
- Keep Nomad support clearly labeled beta or experimental.
- Provider configured in root.

Run:
- terraform fmt -recursive
- gofmt
- go test ./...
```


---
