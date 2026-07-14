# Prompt 204 — OIDC / SSO authentication

```text
Add OIDC authentication support to ClusterForge Control Plane.

Goal:
Allow enterprise users to authenticate through an external identity provider.

Supported:
- generic OIDC
- GitHub OIDC or OAuth can be documented but not required
- Google Workspace / Okta / Azure AD via generic OIDC

Config:
auth:
  mode: oidc
  oidc:
    issuer_url: https://issuer.example.com
    client_id: clusterforge
    client_secret_env: CLUSTERFORGE_OIDC_CLIENT_SECRET
    redirect_url: https://clusterforge.example.com/auth/callback
    scopes:
      - openid
      - profile
      - email
      - groups
    groups_claim: groups
    email_claim: email

Endpoints:
- GET /auth/login
- GET /auth/callback
- POST /auth/logout
- GET /api/v1/me

Session:
- secure cookie
- configurable session lifetime
- CSRF protection for browser flows
- API token flow remains available for CLI/service accounts

CLI:
- cf login --browser
- cf login --token
- cf logout
- cf whoami

RBAC integration:
- map OIDC groups to ClusterForge groups
- allow default role mapping from config
- do not auto-admin unknown users

Docs:
- docs/control-plane/oidc.md
- docs/control-plane/sso.md

Tests:
- mock OIDC provider
- login callback creates/updates user
- group mapping works
- invalid issuer rejected
- missing email rejected
- session cookie security flags

Rules:
- Do not store client secret in config directly.
- Use env var reference.
- Keep static token auth for local/dev mode.
- Do not require OIDC for local development.

Run:
- cd control-plane && go test ./...
- cd cli && go test ./...
```
