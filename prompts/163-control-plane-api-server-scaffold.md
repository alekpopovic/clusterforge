# Prompt 163 — Control Plane API server scaffold

```text
Implement the initial ClusterForge Control Plane API server scaffold.

Location:
- control-plane/

Language:
- Go

Recommended structure:
control-plane/
  go.mod
  main.go
  cmd/
    server.go
    migrate.go
  internal/
    api/
    auth/
    config/
    db/
    models/
    services/
    middleware/
    logging/
    health/
  migrations/
  README.md

HTTP framework:
- Use standard net/http or chi.
- Prefer simple and maintainable code.
- Do not introduce heavy framework unless needed.

API endpoints:
- GET /healthz
- GET /readyz
- GET /version
- GET /api/v1/info

Config:
- config file path via --config
- environment variables supported
- server address
- database URL
- log level
- auth mode:
  - none for local dev
  - token for MVP

Create:
- control-plane/config.example.yaml
- control-plane/README.md

Requirements:
- graceful shutdown
- structured logging
- request ID middleware
- panic recovery middleware
- JSON error responses
- version metadata via ldflags

Tests:
- health endpoint
- version endpoint
- config loading
- JSON error response

Rules:
- No cloud credentials.
- No Terraform execution yet.
- No dashboard yet.
- Keep API scaffold small.

Run:
- cd control-plane && gofmt -w .
- cd control-plane && go test ./...
- cd control-plane && go build -o clusterforge-server .
```
