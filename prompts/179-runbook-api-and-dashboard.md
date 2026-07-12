# Prompt 179 — Runbook API and dashboard

```text
Implement runbook catalog in Control Plane and Dashboard.

Sources:
- docs/incident-response
- docs/dr
- project-local runbooks
- service catalog references

Control Plane:
- runbooks table:
  - id
  - title
  - category
  - severity
  - path
  - tags_json
  - content_markdown optional
  - created_at
  - updated_at

API:
- GET /api/v1/runbooks
- GET /api/v1/runbooks/{id}
- POST /api/v1/runbooks/import

CLI:
- cf api push-runbooks
- cf runbook sync

Dashboard:
- /runbooks
- /runbooks/[id]

Features:
- search by title/tag/category
- link runbooks to services/environments
- render Markdown safely

Rules:
- Do not execute runbook commands.
- Sanitize rendered Markdown.
- No secrets in runbooks.
- Do not expose local filesystem paths unexpectedly.

Tests:
- import runbooks
- search runbooks
- API get runbook
- dashboard build

Docs:
- docs/control-plane-runbooks.md
```
