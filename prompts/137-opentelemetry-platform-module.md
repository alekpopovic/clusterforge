## Prompt 137 — OpenTelemetry platform module

```text
Add OpenTelemetry Collector platform module.

Create module:
- modules/platform/kubernetes/opentelemetry-collector

Purpose:
Install OpenTelemetry Collector via Helm for traces, metrics, and logs pipelines.

Inputs:
- namespace default "observability"
- create_namespace default true
- chart_version default ""
- mode default "deployment"
- values list(string)
- presets object optional
- service_account_annotations map(string)
- labels map(string)

Outputs:
- namespace
- release_name

Docs:
- docs/opentelemetry.md

Cover:
- traces
- metrics
- logs
- collector deployment modes
- exporting to vendor backends
- secrets handling for API keys
- resource impact
- production tuning

Example:
- examples/kubernetes-opentelemetry

Bootstrap:
- Add enable_opentelemetry_collector optional.

Rules:
- Do not include vendor credentials.
- Do not enable by default.
- Keep values minimal.

Run:
- terraform fmt -recursive
```


---
