# Prompt 177 — Dashboard operations pages

```text
Implement dashboard operations pages.

Pages:
- Policy results
- Drift results
- Cost reports
- Audit events
- Runner status
- Plan requests
- Apply requests
- Approvals

Policy page:
- severity filters
- status filters
- policy pack filters
- environment filters

Drift page:
- environment
- stack
- drift status
- last checked
- changed resources count if available

Cost page:
- warnings
- expensive resource categories
- environment filters

Audit page:
- actor
- action
- resource
- timestamp
- filters

Runner page:
- runner name
- status
- last seen
- allowed job types
- active job

Approvals page:
- pending approvals
- approve/reject action only if API supports auth/permission
- show strong warnings for prod apply

Rules:
- No apply button.
- Approval action must clearly show plan/apply details.
- Do not show raw secrets.
- Redact sensitive metadata.

Tests/build:
- run dashboard build
```
