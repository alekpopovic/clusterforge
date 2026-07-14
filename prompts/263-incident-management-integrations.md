# Prompt 263 — Incident management integrations

```text
Add incident management integration support.

Goal:
Send or create incident notifications in external systems through generic webhook first.

Supported MVP:
- generic webhook
- Slack/Teams notification reuse
- Pager-style integration documented through webhook, not provider-specific API unless already implemented

Events:
- incident.started
- incident.resolved
- break_glass.requested
- break_glass.used
- apply.failed
- drift.detected.prod
- runner.offline.prod
- policy.blocked.prod

Config:
incident_integrations:
  enabled: true
  sinks:
    - name: incident-webhook
      type: webhook
      url_env: CLUSTERFORGE_INCIDENT_WEBHOOK
      events:
        - incident.started
        - apply.failed

API:
- GET /api/v1/incident-integrations
- POST /api/v1/incident-integrations/test

CLI:
- cf incident integration test <name>

Tests:
- webhook payload rendered
- URL redacted
- disabled integration skipped
- failure recorded
- incident event triggers integration

Docs:
- docs/control-plane/incident-integrations.md

Rules:
- Do not store webhook URL plaintext if avoidable.
- Do not call real external services in tests.
- Do not make incidents auto-resolve.
```
