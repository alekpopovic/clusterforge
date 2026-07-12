# Prompt 166 — Control Plane authentication MVP

```text
Implement authentication MVP for ClusterForge Control Plane.

Auth modes:
1. none
   - local development only
2. static_token
   - bearer token from config or environment
3. service_token
   - runner/service account tokens

Config:
auth:
  mode: static_token
  static_tokens:
    - name: local-admin
      token_env: CLUSTERFORGE_ADMIN_TOKEN
      role: admin

Roles:
- admin
- operator
- viewer
- runner

Permissions:
- viewer:
  - read inventory
  - read reports
- operator:
  - create plan requests
  - read audit events
- admin:
  - all actions
- runner:
  - claim jobs
  - upload job results
  - heartbeat

Implement:
- auth middleware
- role middleware
- current actor context
- token redaction in logs
- 401/403 JSON responses

Endpoints:
- GET /api/v1/me

Tests:
- unauthenticated request blocked in token mode
- valid token accepted
- invalid token rejected
- viewer cannot create restricted resources
- runner can access runner endpoints only

Docs:
- docs/control-plane-auth.md

Rules:
- Do not implement OAuth yet.
- Do not store tokens in database in plaintext unless service token feature requires hashed tokens.
- Prefer environment variables for tokens.
- No default admin token in production config.

Run:
- cd control-plane && go test ./...
```
