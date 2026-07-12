# Prompt 186 — Scheduled cost reports in Control Plane

```text
Implement scheduled cost report support.

Goal:
Allow environments to have periodic cost warning reports.

Control Plane:
- cost schedules
- cost reports
- cost warnings

CLI:
- cf cost schedule list
- cf cost schedule create <env>
- cf cost report list
- cf cost report show <id>

Runner:
- support cost_scan job
- run heuristic cost scanner
- optionally run Infracost if configured
- upload sanitized report

Dashboard:
- cost reports page
- environment cost warnings

Rules:
- Heuristic mode must not claim exact prices.
- Infracost optional.
- No cloud credentials required for heuristic mode.
- No apply.

Tests:
- scheduled cost job
- cost report upload
- dashboard build
- JSON output

Docs:
- docs/control-plane-cost.md
```
