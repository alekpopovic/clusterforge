# Prompt 178 — Service catalog API and dashboard

```text
Implement service catalog support in Control Plane and Dashboard.

Control Plane:
- tables/resources:
  - services
  - service_dependencies
  - service_environments

API:
- GET /api/v1/services
- POST /api/v1/services
- GET /api/v1/services/{id}
- GET /api/v1/services/{id}/dependencies
- GET /api/v1/services/{id}/environments

CLI:
- cf api push-service-catalog
- cf service sync

Dashboard:
- /services
- /services/[id]

Service page:
- owner
- lifecycle
- tier
- repositories
- environments
- dependencies
- runbooks
- apps
- policy status
- drift status

Import:
- from service-catalog.yaml
- from app manifests
- from Backstage catalog if available later

Tests:
- service import
- dependencies import
- API list/get
- dashboard build

Docs:
- docs/control-plane-service-catalog.md

Rules:
- Metadata only.
- No secrets.
- Keep schema flexible.
```
