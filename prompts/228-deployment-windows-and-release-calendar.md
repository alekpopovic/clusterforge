# Prompt 228 — Deployment windows and release calendar

```text
Implement deployment windows and release calendar support.

Goal:
Allow teams to define allowed deployment windows per environment.

Config:
deployment_windows:
  prod:
    timezone: Europe/Belgrade
    allowed:
      - days: ["Mon", "Tue", "Wed", "Thu"]
        start: "09:00"
        end: "17:00"
    blocked_dates:
      - "2026-12-31"

Control Plane:
- deployment_windows table
- API for CRUD windows
- policy integration

CLI:
- cf deploy-window list prod
- cf deploy-window create prod
- cf deploy-window check prod
- cf calendar export prod --format ics

Behavior:
- apply blocked outside allowed window unless override
- freeze windows still override deployment windows
- export calendar optional

Tests:
- allowed window passes
- outside window blocks
- blocked date blocks
- timezone handling
- override requires permission

Docs:
- docs/control-plane/deployment-windows.md

Rules:
- Use explicit timezones.
- Do not rely on local machine timezone.
- Override audited.
```
