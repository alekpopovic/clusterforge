# Prompt 235 — Performance profiling and optimization

```text
Add performance profiling and optimization plan.

Goal:
Understand and improve Control Plane performance before larger deployments.

Create:
- docs/control-plane/performance.md
- scripts/control-plane-profile.sh

Add optional pprof support:
Config:
debug:
  pprof_enabled: false
  pprof_addr: "127.0.0.1:6060"

Rules:
- pprof disabled by default
- bind localhost by default
- document security risks

Benchmarks:
- API list endpoints
- policy result ingestion
- audit event ingestion
- artifact metadata queries
- job claim loop

Tests/benchmarks:
- Go benchmark tests where practical
- load test script updates

Docs:
- how to run pprof locally
- common bottlenecks
- database indexes
- pagination requirements

Run:
- go test ./...
- go test -bench=. where practical

Rules:
- No premature rewrites.
- Focus on indexes, pagination, and query efficiency.
```
