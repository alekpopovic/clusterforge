# Prompt 223 — PR/MR check status API

```text
Implement PR/MR check status integration abstraction.

Goal:
Support reporting ClusterForge plan/policy results back to GitHub/GitLab as statuses/checks.

Create package:
- control-plane/internal/vcs

Providers:
- GitHub
- GitLab

Capabilities:
- create/update check status
- create/update PR/MR comment
- link to Control Plane plan result
- mark status:
  - pending
  - success
  - failure
  - neutral

Control Plane:
- when plan/policy job starts, set pending status
- when succeeds/fails, update status
- if provider credentials not configured, skip with warning

Config:
vcs:
  github:
    token_env: GITHUB_TOKEN
  gitlab:
    token_env: GITLAB_TOKEN

Tests:
- provider interface tests
- GitHub mocked status update
- GitLab mocked status update
- missing token skips safely
- token not logged

Docs:
- docs/control-plane/vcs-status.md

Rules:
- No raw plan output.
- Credentials from env only.
- Fork PR handling documented.
```
