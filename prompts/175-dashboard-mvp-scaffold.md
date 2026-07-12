# Prompt 175 — Dashboard MVP scaffold

```text
Implement ClusterForge dashboard MVP scaffold.

Location:
- dashboard/

Preferred stack:
- Next.js with TypeScript
or
- simple React/Vite app if simpler

Goal:
Create a read-only dashboard for Control Plane data.

Pages:
- /
  overview
- /projects
- /projects/[id]
- /environments
- /clusters
- /apps
- /policy
- /drift
- /cost
- /audit
- /runners
- /approvals

Requirements:
- API base URL configurable
- no secrets in frontend
- read-only for first version except approval actions may be added later
- simple layout
- status badges
- loading/error states
- basic responsive design

Create:
- dashboard/README.md
- dashboard/.env.example

API integration:
- fetch from Control Plane REST API
- typed API client if practical

Tests:
- basic build
- component tests if setup exists
- API client tests if practical

Docs:
- docs/dashboard.md

Rules:
- Do not add apply button.
- Do not store tokens insecurely beyond local dev limitations.
- Keep dashboard MVP simple.

Run:
- cd dashboard && npm install if package manager already chosen
- cd dashboard && npm run build
```
