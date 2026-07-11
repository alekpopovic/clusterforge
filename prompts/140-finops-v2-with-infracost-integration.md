## Prompt 140 — FinOps v2 with Infracost integration

```text
Improve FinOps support with optional Infracost integration.

Goal:
Provide practical cost visibility while keeping built-in cost scanner heuristic.

CLI:
- cf cost estimate <env>
- cf cost estimate <env> --stack <stack>
- cf cost estimate <env> --infracost
- cf cost diff <env>
- cf cost scan <env>
- cf cost report <env> --json

Behavior:
- Built-in heuristic mode works without credentials.
- Infracost mode runs infracost if installed and configured.
- If Infracost is missing, print clear setup instructions.
- Do not require Infracost by default.

Built-in warnings:
- NAT gateways
- load balancers
- EKS clusters
- managed node groups
- RDS
- ElastiCache
- CloudWatch retention
- persistent volumes
- multi-AZ databases
- cross-region resources

Docs:
- docs/finops.md

Include:
- cost review in PRs
- tagging strategy
- budgets
- cost allocation tags
- cleanup of test environments
- NAT gateway alternatives
- right-sizing node pools
- log retention

Tests:
- Plan fixture with NAT Gateway.
- Plan fixture with RDS.
- JSON output.
- Infracost missing path.

Rules:
- Do not claim exact prices in heuristic mode.
- Do not fail apply based on cost unless policy explicitly configured.
- No cloud credentials required for heuristic mode.

Run:
- gofmt
- go test ./...
```


---
