# Prompt 245 — API rate limiting and tenant quotas

```text
Implement API rate limiting and tenant quotas.

Goal:
Protect the Control Plane from accidental overload and future abuse.

Rate limits:
- per IP
- per actor
- per organization
- per token/service account
- per endpoint class

Endpoint classes:
- read
- write
- auth
- job_polling
- artifact_download
- artifact_upload
- webhook

Config:
rate_limits:
  enabled: true
  defaults:
    read_per_minute: 600
    write_per_minute: 120
    auth_per_minute: 30
    artifact_download_per_minute: 60
    job_polling_per_minute: 300

Quotas:
- max projects per organization
- max environments per organization
- max runners per organization
- max artifacts storage bytes
- max active jobs
- max preview environments
- max API tokens
- max service accounts

API:
- GET /api/v1/quotas
- GET /api/v1/rate-limits
- GET /api/v1/usage/current

CLI:
- cf quota show
- cf quota check
- cf rate-limit show

Behavior:
- return 429 for rate limit exceeded
- include retry-after header
- quota violations return clear JSON error
- audit quota-related denials if useful
- metrics for rate limit/quota denials

Tests:
- rate limit by token
- rate limit by IP
- quota exceeded for artifacts
- quota exceeded for runners
- 429 includes retry-after
- admin can view quotas

Docs:
- docs/control-plane/quotas-rate-limits.md

Rules:
- Keep defaults generous for self-hosted.
- Allow disabling in local dev.
- Do not break runner polling with too-low defaults.
```
