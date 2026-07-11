# Audit export and SIEM import

ClusterForge audit logs remain local by default. Operators can manually create a
redacted export without contacting a cloud service or enabling telemetry.

```bash
cf audit export --format jsonl --output audit.jsonl
cf audit export --format json --since 24h --output audit.json
cf audit export --format csv --output audit.csv
cf audit redact --input audit.log --output audit-redacted.jsonl
```

`--since` accepts Go duration syntax such as `30m`, `24h`, or `168h`. Export and
redaction apply the current sensitive-argument rules again and create owner-only
files (`0600`). Review exports before transferring them: user names, repository
paths, environment names, and command metadata can still be operationally
sensitive.

## SIEM integration concepts

JSONL is the recommended interchange format: each line is one complete JSON
event. Transfer the reviewed file using your organization's approved collector:

- Splunk: configure a generic HTTP Event Collector source and have an external
  agent submit each JSONL event. Keep the HEC token outside ClusterForge files.
- Datadog Logs: use the Datadog Agent or an approved log forwarder to tail the
  exported file; configure the source and service tags in the agent.
- Elasticsearch or OpenSearch: use Filebeat, Fluent Bit, or another bulk-capable
  shipper and map `timestamp` as a date.
- CloudWatch Logs: use the CloudWatch Agent or an organization-managed forwarding
  workflow to ingest the reviewed JSONL file.
- Other SIEMs: import JSONL and use `timestamp`, `command`, `environment`, and
  `result` as the primary search fields.

ClusterForge does not push directly to any of these systems. Collector endpoints,
API keys, tokens, and credentials must remain in the collector's secret store,
not in `clusterforge.yaml`, command arguments, or the exported file.
