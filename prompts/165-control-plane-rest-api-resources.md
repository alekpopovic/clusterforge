# Prompt 165 — Control Plane REST API resources

```text
Implement REST API endpoints for Control Plane inventory resources.

Resources:
- organizations
- workspaces
- projects
- environments
- clusters
- apps
- services
- runbooks

Endpoints:
- GET /api/v1/organizations
- POST /api/v1/organizations
- GET /api/v1/organizations/{id}

- GET /api/v1/workspaces
- POST /api/v1/workspaces
- GET /api/v1/workspaces/{id}

- GET /api/v1/projects
- POST /api/v1/projects
- GET /api/v1/projects/{id}

- GET /api/v1/environments
- POST /api/v1/environments
- GET /api/v1/environments/{id}

- GET /api/v1/clusters
- POST /api/v1/clusters
- GET /api/v1/clusters/{id}

- GET /api/v1/apps
- POST /api/v1/apps
- GET /api/v1/apps/{id}

Requirements:
- JSON request/response
- input validation
- pagination:
  - limit
  - offset
- filtering:
  - project_id
  - environment_id
  - cloud
  - orchestrator
  - status
- consistent error format
- request ID in response headers

Create:
- OpenAPI draft:
  control-plane/openapi.yaml

Tests:
- create/list/get for each resource
- validation errors
- pagination
- filtering
- not found responses

Rules:
- No auth complexity yet if auth mode none is configured.
- Do not expose internal DB errors.
- Do not accept secret values in metadata without redaction.

Run:
- cd control-plane && gofmt -w .
- cd control-plane && go test ./...
```
