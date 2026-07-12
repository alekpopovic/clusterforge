# Prompt 193 — Control Plane E2E tests

```text
Add end-to-end tests for Control Plane workflow.

Goal:
Test API + CLI + runner together without cloud credentials.

Test environment:
- temporary database or SQLite
- control-plane test server
- fake project repository
- fake runner executor

E2E flows:
1. API health
2. CLI login
3. CLI sync inventory
4. Create plan request
5. Runner claims job
6. Fake plan succeeds
7. Create apply request
8. Approval required
9. Approval accepted
10. Runner refuses apply unless allowed
11. Audit events created

Tests:
- control-plane/e2e/
- cli API integration tests if needed
- runner integration tests

Rules:
- No real Terraform required unless mocked/fake executor mode.
- No cloud credentials.
- No network outside localhost.
- Tests must clean up temp files.

Run:
- go test ./... across control-plane, cli, runner
```
