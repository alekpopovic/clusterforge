# Prompt 221 — GitHub webhook integration

```text
Implement GitHub webhook integration for Control Plane.

Goal:
Allow Control Plane to receive GitHub pull request and push events.

Webhook endpoint:
- POST /api/v1/webhooks/github

Security:
- validate GitHub webhook signature
- secret from env:
  CLUSTERFORGE_GITHUB_WEBHOOK_SECRET
- reject invalid signatures
- audit received events
- do not log payload secrets

Events:
- pull_request
- push
- check_suite optional later
- workflow_run optional later

Behavior:
- On pull_request opened/synchronize:
  - create policy check job
  - create plan job if configured
- On push to main/default branch:
  - sync inventory if configured
  - update project commit metadata

Config:
projects:
  repo_url:
  github:
    enabled: true
    plan_on_pr: true
    policy_on_pr: true
    allowed_branches:
      - main

Tests:
- valid signature accepted
- invalid signature rejected
- PR event creates job
- push event updates metadata
- disabled project ignored

Docs:
- docs/control-plane/github-webhooks.md

Rules:
- No automatic apply from webhook.
- Fork PRs must be treated conservatively.
- Do not expose secrets.
```
