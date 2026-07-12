# Prompt 191 — Control Plane observability

```text
Add observability support for ClusterForge Control Plane.

Control Plane API:
- structured logs
- request ID
- metrics endpoint:
  GET /metrics
- Prometheus metrics:
  - HTTP requests total
  - request duration
  - API errors
  - plan requests by status
  - apply requests by status
  - runner heartbeats
  - job duration
  - policy results by severity

Runner:
- metrics endpoint optional
- job duration metrics
- job success/failure counters

Dashboard:
- show basic runner/API health if API exposes it

Docs:
- docs/control-plane-observability.md

Kubernetes:
- ServiceMonitor optional in Helm chart if Prometheus Operator exists
- values:
  serviceMonitor.enabled

Tests:
- metrics endpoint exists
- metrics include expected counters where practical

Rules:
- No sensitive labels in metrics.
- Do not expose metrics publicly by default.
```
