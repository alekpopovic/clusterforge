## Prompt 41 — Environment manifest i multi-stack layout

```text
Introduce environment stack layout support.

Goal:
Allow each environment to be split into separate Terraform stacks:
- network
- cluster
- platform
- apps

Current simple layout may be:
live/dev/aws-eks/

New supported layout:
live/dev/aws-eks/
  network/
  cluster/
  platform/
  apps/

Update clusterforge.yaml schema:
environments:
  dev:
    cloud: aws
    region: eu-central-1
    orchestrator: eks
    path: live/dev/aws-eks
    layout: stacked
    stacks:
      network:
        path: live/dev/aws-eks/network
      cluster:
        path: live/dev/aws-eks/cluster
      platform:
        path: live/dev/aws-eks/platform
      apps:
        path: live/dev/aws-eks/apps

CLI updates:
- cf generate dev --layout simple
- cf generate dev --layout stacked
- cf plan dev
  Defaults to planning all stacks in dependency order.
- cf plan dev --stack network
- cf apply dev --stack cluster
- cf output dev --stack network

Dependency order:
1. network
2. cluster
3. platform
4. apps

Important:
- Use remote state data sources or documented outputs between stacks.
- For local examples, use terraform_remote_state with local backend only if safe.
- Make stack boundaries explicit.

Rules:
- Preserve existing simple layout support.
- Do not break existing environments.
- Do not force stacked mode.
- Generated files must remain readable.

Tests:
- Generate simple layout.
- Generate stacked layout.
- Plan command resolves stack path.
- Unknown stack returns useful error.

Docs:
- docs/environments.md
- Explain simple vs stacked layout.
- Recommend stacked layout for production.

Run:
- gofmt
- go test ./...
```

---
