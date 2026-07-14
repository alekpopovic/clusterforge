# Prompt 206 — Token rotation and runner credential lifecycle

```text
Implement token rotation and runner credential lifecycle.

Goal:
Make runner and service account credentials operationally safe.

Features:
1. token expiration
2. token rotation
3. token revocation
4. last used tracking
5. token scope display
6. runner registration lifecycle

Control Plane:
- service account token expiration
- runner token expiration
- token revoked_at
- token last_used_at
- token name/description

API:
- POST /api/v1/service-accounts/{id}/tokens/{token_id}/rotate
- POST /api/v1/service-accounts/{id}/tokens/{token_id}/revoke
- GET /api/v1/tokens/expiring
- POST /api/v1/runners/{id}/rotate-token
- POST /api/v1/runners/{id}/revoke-token

CLI:
- cf token list
- cf token rotate <id>
- cf token revoke <id>
- cf runner token rotate <runner>
- cf runner token revoke <runner>
- cf token expiring

Notifications:
- optional warning for tokens expiring within configured window

Audit:
- token created
- token rotated
- token revoked
- token used after revocation attempt
- runner token rotated

Tests:
- expired token rejected
- revoked token rejected
- last_used_at updated
- rotation creates new token and revokes old one if requested
- token value displayed once
- audit events created

Docs:
- docs/control-plane/token-rotation.md
- docs/control-plane/runner-credentials.md

Rules:
- Never store raw tokens.
- Never log tokens.
- Token rotation must be explicit.
```
