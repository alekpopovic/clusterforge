## Prompt 21 — Policy i risk summary

```text
Implement basic policy and risk summary features in ClusterForge CLI.

Commands:
- cf policy check <env>
- cf plan <env> --risk-summary
- cf apply <env> --plan-file <file>

Goal:
Before apply, especially in prod, show a simple risk summary and enforce safety rules.

Policy rules:
1. Production apply requires --plan-file.
2. Production destroy is blocked unless --allow-destroy is provided.
3. apply must not use --auto-approve by default.
4. If plan JSON contains delete actions in prod, block apply unless --allow-destroy is passed.
5. If plan JSON contains replacement actions in prod, print high risk warning.
6. If plan JSON cannot be parsed, fail closed for prod and warn for non-prod.

Terraform plan JSON:
- Use terraform show -json <planfile> to obtain JSON.
- Parse enough JSON to count:
  - creates
  - updates
  - deletes
  - replacements
  - no-ops
- List top 20 changed resource addresses.
- Print summary table.

Implementation:
- Create cli/internal/policy package.
- Create cli/internal/terraform/planjson package if useful.
- Add tests using fixture JSON files.
- Do not require Terraform binary for pure parser tests.

Risk levels:
- LOW: only creates/updates in non-prod
- MEDIUM: creates/updates in prod, or replacement in non-prod
- HIGH: delete/replacement in prod
- BLOCKED: forbidden operation by policy

Output example:
Plan summary for prod:
  Create: 3
  Update: 1
  Delete: 0
  Replace: 0
Risk: MEDIUM
Policy: apply allowed only with plan file

Run:
- gofmt
- go test ./...
```

---
