## Prompt 148 — Secret rotation workflow

```text
Add secret rotation workflow documentation and CLI scaffolding.

Goal:
Help teams rotate secrets used by ClusterForge-managed workloads without putting secret values into Terraform.

Docs:
- docs/secrets-rotation.md

Cover:
1. AWS Secrets Manager rotation
2. SSM Parameter Store update
3. External Secrets Operator sync
4. Kubernetes secret refresh
5. Pod restart strategy
6. ECS task redeploy strategy
7. Database password rotation with RDS-managed password
8. Rollback risks
9. Audit trail

CLI:
- cf secrets check <env>
- cf secrets references <env>
- cf secrets references --app api
- cf secrets rotation-plan <env>

Behavior:
- Discover secret references from:
  - app manifests
  - Terraform files
  - External Secrets manifests
  - ECS service definitions
- Do not read secret values.
- Do not rotate secrets directly in MVP.
- Print references and recommended rotation steps.

Tests:
- Find Kubernetes secret_env references.
- Find ECS secrets value_from references.
- Ensure values are not printed.
- Rotation plan output.

Rules:
- Never ask user to type secret values into CLI.
- Never write secret values.
- No cloud API mutation.

Run:
- gofmt
- go test ./...
```


---
