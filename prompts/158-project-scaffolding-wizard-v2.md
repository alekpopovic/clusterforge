## Prompt 158 — Project scaffolding wizard v2

```text
Improve interactive project scaffolding wizard.

Goal:
Make the CLI easier for new users while preserving non-interactive automation.

Commands:
- cf init
- cf wizard
- cf project init --wizard
- cf env create --wizard
- cf app add --wizard

Wizard flow:
1. Project name
2. Team/owner
3. Target:
   - AWS EKS
   - AWS ECS
   - existing Kubernetes
   - local kind
   - AKS/GKE if implemented
4. Environment:
   - dev/staging/prod
5. Backend:
   - local
   - s3
   - terraform cloud
6. Platform add-ons:
   - ingress
   - cert-manager
   - external-secrets
   - argocd
   - monitoring
7. App:
   - optional demo app
8. Safety policies:
   - production plan required
   - destroy blocked
   - image latest policy

Output:
- Show summary before writing.
- Ask confirmation.
- Support --non-interactive.
- Support --defaults for quick demo.

Rules:
- Never ask for secret values.
- Never run apply.
- Generated files must be readable.
- CI must use non-interactive mode.

Tests:
- Non-interactive wizard validation.
- Defaults flow.
- Missing required inputs.
- Summary generation.

Docs:
- docs/wizard.md

Run:
- gofmt
- go test ./...
```


---
