## Prompt 131 — Terraform Cloud / HCP Terraform integration

```text
Add optional Terraform Cloud / HCP Terraform integration support.

Goal:
Support users who want remote execution and remote state.

Do not make Terraform Cloud mandatory.

Config:
clusterforge.yaml:
  terraform_cloud:
    enabled: true
    organization: example-org
    project: clusterforge
    workspaces:
      dev:
        name: clusterforge-dev
      prod:
        name: clusterforge-prod

CLI:
- cf tfe workspace list
- cf tfe workspace generate <env>
- cf tfe backend render <env>

Generated backend:
terraform {
  cloud {
    organization = "example-org"

    workspaces {
      name = "clusterforge-dev"
    }
  }
}

Docs:
- docs/terraform-cloud.md

Cover:
- remote state
- remote execution
- variable sets
- sensitive variables
- workspace naming
- VCS-driven workflow
- CLI-driven workflow
- production approval model

Rules:
- Do not store Terraform Cloud tokens.
- Do not call Terraform Cloud API unless token/config is explicitly provided.
- Initial implementation may only generate config.
- Keep S3 backend support.

Tests:
- Config parsing.
- Backend generation.
- Missing organization fails.
- TFE disabled no-op.

Run:
- gofmt
- go test ./...
- terraform fmt -recursive
```


---
