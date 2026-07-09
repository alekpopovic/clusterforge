## Prompt 44 — Observability stack: Prometheus, Loki, Grafana values

```text
Improve Kubernetes observability platform modules.

Target modules:
- modules/platform/kubernetes/prometheus-stack
- modules/platform/kubernetes/loki

Goal:
Provide practical but safe defaults for observability.

Tasks:
1. prometheus-stack module:
   - Ensure Helm release works.
   - Inputs:
     - namespace default "monitoring"
     - chart_version
     - values
     - enable_grafana_ingress bool default false
     - grafana_host string default ""
     - storage_enabled bool default false
     - storage_class_name string default ""
   - If ingress enabled, require grafana_host.
   - Do not hardcode admin password.
   - Document how to set secrets safely.

2. loki module:
   - Ensure Helm release works.
   - Inputs:
     - namespace default "logging"
     - chart_version
     - values
     - storage_enabled bool default false
     - storage_class_name string default ""

3. Add module:
   - modules/platform/kubernetes/promtail
   Or document using Grafana Alloy if preferred by current chart direction.
   Keep it simple.

4. Update bootstrap module:
   - enable_prometheus_stack
   - enable_loki
   - enable_promtail or enable_log_agent

5. Create example:
   - examples/kubernetes-observability

Docs:
- docs/observability.md
- Explain metrics vs logs.
- Explain storage implications.
- Explain production values should be tuned.

Rules:
- Do not include real passwords.
- Do not expose Grafana publicly by default.
- Do not enable persistent storage by default unless clearly documented.

Run:
- terraform fmt -recursive
```

---
