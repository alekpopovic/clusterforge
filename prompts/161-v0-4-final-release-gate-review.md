# Prompt 161 — v0.4 final release gate review

```text
Perform the final v0.4 release gate review for ClusterForge.

Goal:
Determine whether ClusterForge v0.4.0 is ready to be tagged and released.

Create:
- RELEASE_GATE_V0.4.md

Inspect:
- CLI
- Terraform modules
- provider compatibility
- plugin system
- template pack registry
- policy engine v2
- AWS multi-account support
- multi-region metadata
- upgrade planners
- Terraform Cloud support
- GitLab CI templates
- service catalog
- Backstage export
- compliance mapping
- air-gapped bundle support
- docs
- examples
- tests
- release automation

RELEASE_GATE_V0.4.md must include:

1. Release decision
   - ready
   - ready with warnings
   - blocked

2. Feature readiness
   For each major v0.4 feature:
   - status
   - tests
   - docs
   - known limitations
   - release risk

3. CLI readiness
   Verify:
   - cf version
   - cf doctor
   - cf project init
   - cf env create
   - cf generate
   - cf app commands
   - cf policy commands
   - cf plugin commands
   - cf template commands
   - cf inventory commands
   - cf compliance commands
   - cf bundle commands
   - cf migrate analyze

4. Terraform readiness
   - stable modules
   - beta modules
   - experimental modules
   - placeholder modules
   - module catalog accuracy
   - module conformance check

5. Security readiness
   - no secrets committed
   - state files ignored
   - audit log safe
   - policy engine works
   - production apply protection
   - destroy protection
   - supply chain controls
   - air-gapped bundle excludes secrets

6. Test readiness
   - Go unit tests
   - CLI e2e tests
   - golden tests
   - Terraform tests
   - module conformance
   - policy tests
   - smoke test evidence

7. Docs readiness
   - README
   - quickstart
   - CLI docs
   - architecture docs
   - security docs
   - compliance docs
   - migration docs
   - release docs

8. Blockers
   - critical
   - important
   - deferred

Run:
- make fmt-check
- make lint
- make test
- make validate
- make security
- make check-modules
- cd cli && go test ./...
- cd cli && go build -o cf .
- ./cf version
- ./cf doctor

Rules:
- Do not add new features.
- Fix only release-blocking issues.
- Do not claim cloud smoke tests passed unless evidence exists.
- No credentials.
- No apply.

Final response:
- State v0.4 release readiness.
- List blockers.
- List commands run.
- List changed files.
```
