## Prompt 121 — v0.3 release gate review

```text
Perform a v0.3 release gate review for ClusterForge.

Goal:
Determine whether ClusterForge is ready for a v0.3.0 release.

Create:
- RELEASE_GATE_V0.3.md

Inspect:
- CLI
- Terraform modules
- Kubernetes platform modules
- AWS production modules
- local Kubernetes support
- existing Kubernetes support
- fleet operations
- security baseline
- policy modules
- backup/DR docs
- tests
- CI
- release workflows
- docs
- examples

RELEASE_GATE_V0.3.md must include:

1. Executive summary
   - ready
   - ready with warnings
   - blocked

2. Supported targets
   - AWS EKS
   - AWS ECS
   - existing Kubernetes
   - local Kind/K3d
   - AKS if implemented
   - GKE if implemented
   - K3s/RKE2 if implemented
   - Nomad if implemented
   - Docker if implemented

3. CLI readiness
   - project init
   - env create
   - generate
   - app add/list/validate/render
   - doctor
   - drift check
   - fleet commands
   - policy commands
   - module check
   - upgrade commands

4. Terraform module readiness
   - stable modules
   - beta modules
   - experimental modules
   - placeholder modules
   - modules missing examples
   - modules missing tests

5. Security readiness
   - state safety
   - secrets handling
   - production apply safety
   - destroy protection
   - image policy
   - policy packs
   - secret scanning
   - audit logs

6. Test readiness
   - Go unit tests
   - CLI e2e tests
   - golden generator tests
   - Terraform native tests
   - module conformance checks
   - smoke test evidence
   - skipped tests and why

7. Documentation readiness
   - README
   - CLI docs
   - architecture docs
   - module docs
   - tutorials
   - production guides
   - security docs
   - DR docs
   - smoke tests

8. Release blockers
   - critical
   - important
   - deferred

9. Recommended release decision
   - release now
   - release after fixes
   - do not release yet

Run:
- make fmt-check
- make lint
- make test
- make validate
- make security if available
- make check-modules if available
- cd cli && go test ./...
- cd cli && go build -o cf .
- ./cf version
- ./cf doctor

Rules:
- Do not add new features.
- Fix only tiny release-blocking issues.
- Do not claim cloud apply or smoke tests passed unless there is evidence.
- Be explicit about skipped checks.

Final response:
- State whether v0.3.0 is ready.
- List blockers.
- List commands run.
- List files changed.
```


---
