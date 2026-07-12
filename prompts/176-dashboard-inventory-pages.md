# Prompt 176 — Dashboard inventory pages

```text
Implement dashboard inventory pages.

Pages:
- Projects
- Project detail
- Environments
- Environment detail
- Clusters
- Cluster detail
- Apps
- App detail

Each page should show:
- list table
- status
- cloud
- orchestrator
- region
- owner if available
- last updated
- links to related resources

Project detail:
- environments
- clusters
- apps
- recent policy results
- recent drift results

Environment detail:
- stacks
- cluster
- apps
- policy status
- drift status
- cost warnings
- recent audit events

Cluster detail:
- cloud
- orchestrator
- version if known
- status
- platform add-ons if known
- runbooks

Requirements:
- filtering
- search
- empty states
- error states

Rules:
- Read-only.
- No cloud API calls from frontend.
- No secret values.
- Handle missing data gracefully.

Run:
- dashboard tests/build
```
