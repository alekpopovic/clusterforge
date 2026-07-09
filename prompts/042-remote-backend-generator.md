## Prompt 42 — Remote backend generator

```text
Add remote backend generation support.

Goal:
Generate safe Terraform backend configuration templates for environments and stacks.

Supported backends:
1. local
2. s3
3. azurerm placeholder
4. gcs placeholder

For first real implementation:
- local
- s3

Update clusterforge.yaml:
backends:
  dev:
    type: local
  prod:
    type: s3
    bucket: my-terraform-state-bucket
    region: eu-central-1
    dynamodb_table: my-terraform-locks
    key_prefix: clusterforge/prod

CLI:
- cf backend configure <env>
- cf backend show <env>
- cf generate <env> should use backend config.

S3 backend generation:
backend.tf should contain:
terraform {
  backend "s3" {
    bucket         = "..."
    key            = "..."
    region         = "..."
    dynamodb_table = "..."
    encrypt        = true
  }
}

Rules:
- Do not create the backend bucket automatically in the same root that uses it.
- Add separate bootstrap example for backend resources:
  examples/aws-tfstate-backend
- Do not put credentials in backend config.
- For prod, warn if local backend is used.
- Support --backend local|s3.

Docs:
- docs/backends.md
- Explain why remote state matters.
- Explain state contains sensitive data.
- Explain backend bootstrap flow.

Tests:
- Generate local backend.
- Generate s3 backend.
- Missing s3 bucket fails validation.
- Prod local backend warning.

Run:
- gofmt
- go test ./...
```

---
