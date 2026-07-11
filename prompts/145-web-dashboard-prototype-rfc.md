## Prompt 145 — Web dashboard prototype RFC

```text
Create an RFC and lightweight prototype plan for a ClusterForge web dashboard.

Create:
- docs/rfcs/012-web-dashboard.md
- docs/dashboard.md

Dashboard goals:
- environment inventory
- cluster status
- app catalog
- drift results
- policy results
- cost warnings
- audit events
- runbook links
- release status

Non-goals for first prototype:
- direct apply button
- storing cloud credentials
- replacing Git workflow
- full SaaS multi-tenancy

Possible stack:
- Next.js or simple static UI
- API later
- local JSON import first

Prototype mode:
- CLI exports:
  cf dashboard export --output dashboard-data.json
- Static dashboard reads dashboard-data.json

Docs:
- Explain dashboard data model.
- Explain security considerations.
- Explain why no apply button in early version.

Do not implement full dashboard yet.
Only RFC and data model proposal.

Final response:
- Summarize recommended dashboard MVP.
```


---
