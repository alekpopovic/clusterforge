## Prompt 147 — Audit event export and SIEM integration

```text
Add audit event export support.

Goal:
Allow local ClusterForge audit logs to be exported to external systems manually.

CLI:
- cf audit export --format jsonl|json|csv
- cf audit export --since 24h
- cf audit export --output audit.jsonl
- cf audit redact --input audit.log --output audit-redacted.jsonl

Docs:
- docs/audit-export.md

Integration docs:
- Splunk generic HTTP collector concept
- Datadog logs concept
- Elasticsearch/OpenSearch concept
- CloudWatch logs concept
- SIEM import via JSONL

Do not implement direct SIEM API push unless explicitly configured.

Rules:
- Redact sensitive args.
- Do not export secrets.
- Read-only.
- Keep audit log local by default.
- Do not add telemetry.

Tests:
- Export JSONL.
- Export CSV.
- Since filter.
- Redaction.

Run:
- gofmt
- go test ./...
```


---
