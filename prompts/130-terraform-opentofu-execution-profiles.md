## Prompt 130 — Terraform/OpenTofu execution profiles

```text
Add execution profiles for Terraform/OpenTofu operations.

Goal:
Let teams define safe defaults for local, CI, and production execution.

Config:
clusterforge.yaml:
  execution_profiles:
    local:
      engine: terraform
      parallelism: 10
      refresh: true
      lock_timeout: 5m

    ci:
      engine: terraform
      parallelism: 5
      refresh: true
      lock_timeout: 10m
      input: false

    prod:
      engine: terraform
      parallelism: 3
      refresh: true
      lock_timeout: 20m
      input: false
      require_plan_file: true

CLI:
- cf plan dev --profile local
- cf plan prod --profile prod
- cf apply prod --profile prod --plan-file ...
- cf profile list
- cf profile show <name>

Terraform runner:
- Support:
  - -parallelism
  - -refresh
  - -lock-timeout
  - -input
- Do not add auto-approve by default.

Tests:
- Profile parsing.
- Plan args generated correctly.
- Prod profile enforces plan file.
- Unknown profile fails.

Docs:
- docs/execution-profiles.md

Run:
- gofmt
- go test ./...
```


---
