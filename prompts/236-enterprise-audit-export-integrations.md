# Prompt 236 — Enterprise audit export integrations

```text
Add enterprise audit export integration design and MVP.

Goal:
Allow audit events to be exported to external systems.

Supported MVP:
- JSONL file export
- webhook sink

Future:
- Splunk
- Datadog
- OpenSearch
- CloudWatch Logs

Config:
audit_export:
  enabled: true
  sinks:
    - name: security-webhook
      type: webhook
      url_env: CLUSTERFORGE_AUDIT_WEBHOOK
      batch_size: 100
      interval: 1m

API:
- GET /api/v1/audit-exports
- POST /api/v1/audit-exports/test

CLI:
- cf api audit export --format jsonl
- cf audit sink test <name>

Behavior:
- batch export
- retry lightly
- track delivery status
- never export secrets
- redact metadata

Tests:
- webhook payload
- redaction
- retry behavior
- failed delivery recorded

Docs:
- docs/control-plane/audit-export.md

Rules:
- No real external service calls in tests.
- Do not enable by default.
```
