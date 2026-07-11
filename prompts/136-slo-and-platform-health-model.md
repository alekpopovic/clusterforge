## Prompt 136 — SLO and platform health model

```text
Add SLO and platform health model.

Goal:
Define how ClusterForge environments express health, reliability, and operational expectations.

Create:
- docs/slo-model.md
- docs/platform-health.md

Config:
clusterforge.yaml:
  health:
    environments:
      prod:
        slo:
          availability_target: "99.9"
          latency_target_ms: 300
          error_rate_target: "1%"
        checks:
          kubernetes_nodes: true
          platform_addons: true
          ingress: true
          app_health: true

CLI:
- cf health check <env>
- cf health check <env> --json
- cf health report <env>

Initial checks:
- config exists
- environment path exists
- terraform state accessible if possible
- kubectl nodes if kubeconfig available
- namespace status if kubeconfig available
- platform add-on releases if helm available
- workload manifests present
- no apply or mutation

Docs:
- Explain SLO vs health check.
- Explain limitations.
- Explain manual production validation.

Tests:
- Config health check.
- Missing env path fails.
- JSON output.
- Live cluster checks skipped cleanly when unavailable.

Run:
- gofmt
- go test ./...
```


---
