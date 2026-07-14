# Prompt 222 — GitLab webhook integration

```text
Implement GitLab webhook integration for Control Plane.

Webhook endpoint:
- POST /api/v1/webhooks/gitlab

Security:
- validate GitLab token from header
- token from env:
  CLUSTERFORGE_GITLAB_WEBHOOK_TOKEN
- reject invalid token
- audit received events

Events:
- merge request events
- push events
- pipeline events optional later

Behavior:
- On merge request opened/update:
  - create policy check job
  - create plan job if configured
- On push to default branch:
  - sync inventory if configured

Config:
projects:
  gitlab:
    enabled: true
    plan_on_mr: true
    policy_on_mr: true
    allowed_branches:
      - main

CLI:
- cf gitlab webhook test optional

Tests:
- valid token accepted
- invalid token rejected
- MR event creates job
- disabled config ignored

Docs:
- docs/control-plane/gitlab-webhooks.md

Rules:
- No automatic apply.
- No raw plan output in MR comments.
```
