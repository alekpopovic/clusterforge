# Prompt 249 — Organization onboarding workflow

```text
Implement organization onboarding workflow.

Goal:
Make it easy to bootstrap a new organization/workspace/project safely.

Control Plane:
- onboarding state per organization
- setup checklist

Checklist:
- organization created
- admin user configured
- workspace created
- first project created
- RBAC bindings configured
- runner registered
- artifact storage configured
- notification channel configured
- first inventory sync completed
- first policy check completed
- backup guidance acknowledged

API:
- GET /api/v1/onboarding
- POST /api/v1/onboarding/steps/{step}/complete
- POST /api/v1/onboarding/reset

CLI:
- cf org init
- cf onboarding status
- cf onboarding complete <step>

Dashboard:
- onboarding checklist page
- setup progress

Docs:
- docs/control-plane/onboarding.md

Tests:
- new organization has checklist
- completing step updates status
- onboarding status scoped by organization
- unauthorized user cannot complete admin steps

Rules:
- Do not require SaaS.
- Must work in self-hosted mode.
- No secrets in onboarding data.
```
