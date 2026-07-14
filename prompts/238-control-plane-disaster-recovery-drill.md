# Prompt 238 — Control Plane disaster recovery drill

```text
Create disaster recovery drill workflow for Control Plane.

Goal:
Make DR testable, not just documented.

Create:
- docs/control-plane/dr-drill.md
- examples/control-plane-dr-drill/

DR drill steps:
1. backup database
2. deploy fresh Control Plane
3. restore database
4. verify migrations
5. verify users/RBAC
6. verify projects/environments
7. verify runners reconnect or re-register
8. verify artifacts availability
9. verify dashboard
10. verify audit history
11. run test plan request

CLI:
- cf api dr check
- cf api dr report

Evidence file:
control-plane-dr-evidence.yaml

Tests:
- evidence parsing
- missing restore test warns
- DR report JSON

Rules:
- Do not perform destructive restore automatically.
- Do not include credentials.
- Mark manual steps clearly.
```
