## Prompt 127 — Multi-region environment strategy

```text
Add multi-region environment strategy.

Goal:
Support region-aware environment generation and documentation.

Config:
clusterforge.yaml:
  regions:
    primary: eu-central-1
    secondary: eu-west-1

  environments:
    prod-eu-central:
      cloud: aws
      region: eu-central-1
      orchestrator: eks
      path: live/prod/eu-central-1/aws-eks

    prod-eu-west:
      cloud: aws
      region: eu-west-1
      orchestrator: eks
      path: live/prod/eu-west-1/aws-eks

CLI:
- cf region list
- cf region show <name>
- cf env create prod-eu-central --region eu-central-1
- cf fleet status --region eu-central-1

Docs:
- docs/multi-region.md

Topics:
- active/passive
- active/active
- DNS failover
- data replication
- Terraform state separation
- backup/restore
- RTO/RPO planning
- cost implications

Do not implement automatic failover.

Tests:
- Config with multiple regions.
- Fleet filtering by region.
- Env generation uses correct region.
- Duplicate region aliases fail validation.

Rules:
- No fake DR guarantees.
- No automatic cross-region data replication.
- Keep region model metadata-first.

Run:
- gofmt
- go test ./...
```


---
