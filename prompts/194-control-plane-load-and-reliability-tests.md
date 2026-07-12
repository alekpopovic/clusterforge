# Prompt 194 — Control Plane load and reliability tests

```text
Add basic load and reliability tests for Control Plane.

Goal:
Ensure API behaves reasonably under moderate local load.

Create:
- control-plane/tests/load/
- scripts/control-plane-load-test.sh
- docs/control-plane-load-testing.md

Scenarios:
- create 100 projects
- create 500 environments
- create 1000 policy results
- create 1000 audit events
- list with pagination
- runner heartbeat burst
- plan request creation burst

Metrics:
- response time summary
- error count
- database errors
- memory usage if available

Rules:
- Local only.
- No cloud resources.
- Not part of default CI unless lightweight.
- Mark as optional.

Final response:
- Include how to run load test.
```
