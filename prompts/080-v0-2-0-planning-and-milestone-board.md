## Prompt 80 — v0.2.0 planning and milestone board

```text
Create v0.2.0 planning documents.

Goal:
Convert all remaining work into a clear roadmap.

Create:
- ROADMAP_V0.2.md
- ROADMAP_V0.3.md
- BACKLOG.md

ROADMAP_V0.2.md:
Theme:
- production validation
- multi-cloud MVP
- better CLI workflow
- policy packs
- template packs

Must include:
- goals
- non-goals
- milestones
- acceptance criteria
- risks
- dependencies

ROADMAP_V0.3.md:
Theme:
- enterprise adoption
- plugin system
- advanced deployment strategies
- UI/API exploration

BACKLOG.md:
Group by:
- CLI
- Terraform modules
- Kubernetes
- ECS
- Nomad
- Docker
- Multi-cloud
- Security
- Docs
- CI/CD
- Testing
- Product

For every item:
- title
- priority
- status
- estimated complexity
- notes

Rules:
- Do not create vague tasks.
- Every roadmap item must be testable.
- Clearly separate must-have from nice-to-have.

Final response:
- Summarize recommended v0.2 scope.
```

---

## Preporučeni nastavak rada

Kreni ovako:

```text
51. Real MVP acceptance test
52. Terraform native tests for core modules
53. Plan-mode tests for AWS modules
54. Real cloud smoke test runbook
56. Version support matrix
57. Module release packaging
58. GitHub release automation
61. Drift detection
63. Upgrade/migration framework
80. v0.2.0 planning
```

Posle toga multi-cloud:

```text
66. AKS RFC
67. AKS MVP
68. GKE RFC
69. GKE MVP
70. K3s/RKE2
```

Zatim enterprise/product sloj:

```text
71. Policy packs
76. Plugin architecture RFC
77. Template packs
78. Docs website
79. Tutorials
```
