# Prompt 276 — Kubernetes fleet add-on manager

```text
Implement Kubernetes fleet add-on manager.

Goal:
Track and compare platform add-ons across multiple clusters.

Data sources:
- clusterforge.yaml
- generated Terraform
- Helm releases if cluster access available
- Control Plane inventory

Tracked add-ons:
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
- opentelemetry-collector

CLI:
- cf fleet addons list
- cf fleet addons diff
- cf fleet addons report --format markdown|json
- cf fleet addons drift

Control Plane:
- add-on inventory table optional:
  - cluster_id
  - name
  - desired_version
  - observed_version
  - status
  - source

Dashboard:
- fleet add-ons page

Behavior:
- show desired vs observed
- show missing add-ons
- show version skew
- show CRD-sensitive upgrades
- no automatic upgrade

Tests:
- compare two clusters
- missing add-on detected
- version skew detected
- JSON output

Docs:
- docs/fleet-addons.md

Rules:
- Read-only.
- No Helm upgrade.
- No cluster mutation.
```
