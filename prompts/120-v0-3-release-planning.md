## Prompt 120 — v0.3 release planning

```text
Create v0.3.0 release plan.

Create:
- ROADMAP_V0.3.md
- RELEASE_PLAN_V0.3.md
- BACKLOG_V0.3.md

Theme:
ClusterForge v0.3 should focus on production hardening, local development, fleet operations and security.

ROADMAP_V0.3.md must include:
1. Goals
   - local Kubernetes target
   - existing Kubernetes target
   - stronger EKS production options
   - backup/restore support
   - tenant/security baseline
   - fleet read-only operations
   - stronger tests
   - better docs and tutorials

2. Non-goals
   - full SaaS platform
   - automatic remediation
   - plugin marketplace
   - advanced service mesh by default
   - production-grade all-cloud parity

3. Milestones:
   - M1: Testing and release gate
   - M2: Local/existing Kubernetes
   - M3: AWS production hardening
   - M4: Backup and security baseline
   - M5: Fleet operations
   - M6: Docs and examples

4. Acceptance criteria:
   - CLI e2e tests pass
   - golden generator tests pass
   - core module terraform tests pass
   - at least one real smoke test documented
   - no known secret exposure
   - docs cover install, quickstart and cleanup

5. Risks:
   - cloud tests cost money
   - provider compatibility
   - Kubernetes chart drift
   - IAM complexity
   - too many features before stabilization

BACKLOG_V0.3.md:
Group tasks by:
- CLI
- Terraform modules
- Kubernetes
- AWS
- Multi-cloud
- Security
- Testing
- Docs
- Release

For each item include:
- priority
- complexity
- owner placeholder
- status
- acceptance criteria

Final response:
- Summarize recommended v0.3 scope.
- List must-have items.
- List deferred items.
```

---

# Preporučeni redosled posle Prompt 80

Najpraktičniji redosled nije striktno 81–120 redom. Bolji balans je:

```text
81  v0.2 release gate review
85  golden tests for CLI generators
86  CLI end-to-end non-cloud tests
87  module conformance checker
82  local Kubernetes target
83  existing Kubernetes target
89  EKS production hardening
92  ECR registry module
94  Velero backup module
95  DR runbooks
103 Kubernetes tenant model
105 Kyverno policy module
109 multi-cluster inventory
110 fleet operations CLI
119 security threat model
120 v0.3 release planning
```

Ovaj redosled prvo stabilizuje projekat, zatim dodaje lokalni/dev UX, zatim produkcioni AWS/Kubernetes hardening, a na kraju fleet/security sloj.
