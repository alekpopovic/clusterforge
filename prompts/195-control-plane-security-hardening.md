# Prompt 195 — Control Plane security hardening

```text
Harden Control Plane security.

Tasks:
1. HTTP security:
   - request size limits
   - timeouts
   - CORS config
   - secure headers where applicable
   - panic recovery
   - JSON error sanitization

2. Auth:
   - token hash storage for service tokens if persisted
   - token redaction in logs
   - role enforcement
   - deny by default

3. Audit:
   - auth failures audited if practical
   - approval actions audited
   - runner actions audited

4. Input validation:
   - metadata size limits
   - allowed enum values
   - no arbitrary path traversal
   - no unsafe file reads

5. Runner security:
   - allowed job types
   - workspace cleanup
   - no arbitrary command execution
   - token redaction

6. Dashboard:
   - no secrets in frontend
   - safe Markdown rendering
   - no unsafe HTML from runbooks

Create:
- docs/control-plane-security.md
- SECURITY.md update if needed

Tests:
- unauthorized access blocked
- role restrictions
- request too large rejected
- sensitive log redaction
- path traversal blocked

Run:
- go test ./...
```
