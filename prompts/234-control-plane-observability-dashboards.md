# Prompt 234 — Control Plane observability dashboards

```text
Add observability dashboards for ClusterForge Control Plane.

Create:
- dashboards/grafana/
  clusterforge-control-plane.json
  clusterforge-runner.json

Dashboards:
1. API dashboard
   - request rate
   - error rate
   - latency
   - active users if available
   - DB errors
   - job queue depth

2. Runner dashboard
   - runner heartbeats
   - job duration
   - job failures
   - jobs by type
   - queue wait time

3. Workflow dashboard
   - plan requests
   - apply requests
   - approvals pending
   - policy blocked
   - drift detected

Docs:
- docs/control-plane/observability-dashboards.md

Helm chart:
- optional ConfigMap dashboard provisioning if enabled:
  grafanaDashboards.enabled

Rules:
- Dashboards must not include secrets.
- Metrics names must match implementation.
- If metrics are not implemented, mark dashboard panels TODO clearly.
```
