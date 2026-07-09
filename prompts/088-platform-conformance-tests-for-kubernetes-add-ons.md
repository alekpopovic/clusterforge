## Prompt 88 — Platform conformance tests for Kubernetes add-ons

```text
Add conformance checks for Kubernetes platform modules.

Goal:
Validate that Helm-based platform modules are structured consistently.

Target modules:
- ingress-nginx
- cert-manager
- external-dns
- external-secrets
- metrics-server
- prometheus-stack
- loki
- argocd
- karpenter if implemented

Checks:
- Helm repository input exists or is documented.
- chart name is correct in module.
- chart_version input exists.
- values input exists.
- namespace input exists.
- create_namespace input exists where appropriate.
- release_name output exists.
- namespace output exists.
- README has example.
- bootstrap module wires inputs correctly.

Create:
- scripts/check-platform-modules.sh
- docs/platform-module-conventions.md

Optional:
- Add CLI module check rules for platform modules.

Rules:
- Do not pin chart versions globally unless this is already project policy.
- Do not install charts during tests.
- No Kubernetes cluster required.

Run:
- scripts/check-platform-modules.sh
- terraform fmt -recursive
```

---
