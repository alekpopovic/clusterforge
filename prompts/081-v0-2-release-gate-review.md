## Prompt 81 — v0.2 release gate review

```text
Perform a v0.2 release gate review for ClusterForge.

Goal:
Determine whether the repository is ready for a v0.2.0 release after the post-MVP work.

Create:
- RELEASE_GATE_V0.2.md

Inspect:
- CLI commands
- Terraform modules
- docs
- examples
- CI workflows
- tests
- release automation
- smoke test docs
- security policies
- template packs
- multi-cloud support

RELEASE_GATE_V0.2.md must include:

1. Release readiness summary
   - ready
   - ready with warnings
   - blocked

2. Supported functionality
   - AWS EKS
   - AWS ECS
   - existing Kubernetes
   - AKS if implemented
   - GKE if implemented
   - K3s/RKE2 if implemented
   - app manifests
   - policy packs
   - template packs

3. Test status
   - Go unit tests
   - Terraform fmt
   - Terraform validate
   - Terraform native tests
   - integration tests
   - smoke tests

4. Security status
   - secrets handling
   - production apply protection
   - destroy protection
   - state protection
   - policy pack status

5. Documentation status
   - README
   - CLI docs
   - module docs
   - tutorials
   - smoke tests
   - release docs

6. Known limitations
   - untested providers
   - placeholder modules
   - cloud-specific caveats
   - features not production-ready

7. Required fixes before release
   - blockers
   - non-blocking warnings
   - deferred tasks

Run:
- make fmt-check
- make lint
- make test
- make validate
- make security if available
- cd cli && go build -o cf .
- ./cf version
- ./cf doctor

Rules:
- Do not add major new features.
- Fix only small release-blocking issues.
- Do not claim cloud smoke tests passed unless evidence exists.
- Be explicit about skipped checks.

Final response:
- State v0.2 readiness.
- List blockers.
- List commands run.
- List files changed.
```

---
