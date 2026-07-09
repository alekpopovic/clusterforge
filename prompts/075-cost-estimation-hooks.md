## Prompt 75 — Cost estimation hooks

```text
Add cost-awareness hooks to ClusterForge.

Goal:
Warn users about potentially expensive resources before apply.

Create:
- docs/cost-awareness.md
- cli/internal/cost/
- cf cost estimate <env>
- cf cost scan <env>

Implementation:
- Do not build a full cost engine.
- Parse Terraform plan JSON.
- Detect expensive resource categories:
  - NAT Gateway
  - EKS cluster
  - managed node groups
  - Load Balancers
  - RDS if added later
  - persistent volumes
  - CloudWatch log retention
- Print advisory warnings.

Optional:
- Integrate Infracost if installed.
- If infracost is missing, print install/documentation hint.
- Do not require Infracost by default.

Output:
- Human summary.
- JSON output.

Rules:
- Do not claim exact cost unless using real pricing data.
- Mark built-in scanner as heuristic.
- No apply.
- No cloud credentials required for heuristic mode.

Tests:
- Plan JSON with NAT Gateway triggers warning.
- Plan JSON with EKS triggers warning.
- Empty plan returns no warnings.

Run:
- gofmt
- go test ./...
```

---
