## Prompt 109 — Multi-cluster inventory model

```text
Add multi-cluster inventory model.

Goal:
Allow ClusterForge to understand multiple environments and clusters across clouds.

Update clusterforge.yaml schema:
clusters:
  dev-eks:
    environment: dev
    cloud: aws
    orchestrator: eks
    region: eu-central-1
    path: live/dev/aws-eks
    status: experimental
  prod-eks:
    environment: prod
    cloud: aws
    orchestrator: eks
    region: eu-central-1
    path: live/prod/aws-eks
    status: production

CLI commands:
- cf cluster list
- cf cluster show <name>
- cf cluster doctor <name>
- cf cluster kubeconfig <name> if supported
- cf cluster outputs <name>

Behavior:
- Existing environments remain supported.
- clusters is optional initially.
- env list and cluster list should be consistent.
- JSON output supported.

Docs:
- docs/multi-cluster.md

Tests:
- Load config with clusters.
- List clusters.
- Show cluster.
- Missing cluster fails clearly.

Rules:
- Do not require live cloud access for inventory.
- Do not store kubeconfig content.
- Keep backward compatibility.

Run:
- gofmt
- go test ./...
```

---
