# Prompt 174 — Notification system MVP

```text
Implement Control Plane notification system MVP.

Goal:
Notify teams about important plan/apply/policy/drift events.

Notification channels:
1. webhook
2. Slack webhook
3. Microsoft Teams webhook

Config:
notifications:
  enabled: true
  channels:
    - name: platform-slack
      type: slack
      url_env: CLUSTERFORGE_SLACK_WEBHOOK
      events:
        - plan.failed
        - apply.pending_approval
        - apply.succeeded
        - apply.failed
        - drift.detected
        - policy.blocked

Events:
- plan.requested
- plan.succeeded
- plan.failed
- apply.pending_approval
- apply.approved
- apply.rejected
- apply.started
- apply.succeeded
- apply.failed
- drift.detected
- policy.blocked
- runner.offline

Requirements:
- URLs from env vars only.
- Do not store webhook secrets in database plaintext.
- Redact sensitive data.
- Retry lightly for transient failure.
- Log notification delivery status.

API:
- GET /api/v1/notifications/events
- GET /api/v1/notifications/deliveries

CLI:
- cf notification test <channel>
- cf notification list

Tests:
- webhook payload generated
- Slack payload generated
- Teams payload generated
- secret URL not logged
- disabled channel skipped

Docs:
- docs/notifications.md

Rules:
- Do not send notifications in tests to real endpoints.
- No credentials in config examples.
```
