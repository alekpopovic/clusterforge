## Prompt 157 — Migration analyzer for existing Terraform repos

```text
Design and implement a read-only migration analyzer for existing Terraform repositories.

Goal:
Help teams evaluate how hard it would be to adopt ClusterForge.

CLI:
- cf migrate analyze --path <terraform-repo>
- cf migrate analyze --path <terraform-repo> --json
- cf migrate report --path <terraform-repo>

Analyzer should detect:
- providers used
- root modules
- module sources
- AWS VPC resources
- EKS resources
- ECS resources
- Kubernetes resources
- Helm releases
- backend configuration
- tfvars files
- possible secrets
- state files present
- environment layout
- module structure

Output:
- summary
- detected architecture
- ClusterForge equivalent modules
- risks
- suggested migration steps
- import/adoption notes

Rules:
- Read-only.
- Do not modify repo.
- Do not run terraform apply.
- Do not read secret values beyond detection/redaction.
- Warn about tfstate sensitivity.
- No cloud API calls.

Docs:
- docs/migration-analyzer.md

Tests:
- Analyze fixture Terraform repo.
- Detect EKS.
- Detect ECS.
- Detect backend.
- Redact secrets.

Run:
- gofmt
- go test ./...
```


---
