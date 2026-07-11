## Prompt 129 — Platform add-on upgrade planner

```text
Implement platform add-on upgrade planner.

Goal:
Track and plan upgrades for Helm-installed platform components.

CLI:
- cf platform versions <env>
- cf platform upgrade plan <env>
- cf platform upgrade check <env>

Components:
- ingress-nginx
- cert-manager
- external-dns
- external-secrets
- metrics-server
- prometheus-stack
- loki
- argocd
- kyverno
- gatekeeper
- velero
- argo-rollouts

Config:
clusterforge.yaml:
  platform_versions:
    ingress-nginx: "4.10.0"
    cert-manager: "v1.14.0"
    argocd: "6.7.0"

Behavior:
- Read configured chart versions from generated Terraform or config.
- Compare to desired versions in config.
- Print upgrade plan.
- Detect CRD-related components.
- Warn that CRD upgrades require review.
- Do not upgrade automatically.

Docs:
- docs/platform-upgrades.md

Tests:
- Version diff detection.
- CRD warning for cert-manager/argocd/kyverno.
- Unknown component warning.
- JSON output.

Rules:
- No internet lookup.
- No automatic Helm upgrade.
- No apply.

Run:
- gofmt
- go test ./...
```


---
