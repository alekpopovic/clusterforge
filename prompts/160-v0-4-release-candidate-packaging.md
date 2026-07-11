## Prompt 160 — v0.4 release candidate packaging

```text
Prepare ClusterForge v0.4 release candidate.

Goal:
Package the repository for a v0.4.0 release candidate.

Tasks:

1. Update version files:
   - VERSION
   - cli version constants
   - CHANGELOG.md

2. Update release docs:
   - RELEASE_PLAN_V0.4.md
   - RELEASE_NOTES_V0.4.md
   - docs/release-checklist.md

3. Run release validation:
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

4. Verify docs:
   - README quickstart
   - docs/cli.md
   - docs/plugins.md
   - docs/template-pack-registry.md
   - docs/policy-engine.md
   - docs/aws-multi-account.md
   - docs/compliance
   - docs/air-gapped.md

5. Verify examples:
   - local Kubernetes example
   - existing Kubernetes example
   - AWS EKS example
   - AWS ECS example
   - policy pack example
   - template pack example
   - service catalog example

6. Create:
   - RELEASE_CANDIDATE_V0.4.md

RELEASE_CANDIDATE_V0.4.md must include:
- included features
- excluded features
- breaking changes
- migration notes
- known limitations
- test results
- release blockers
- smoke test status
- security notes
- recommended release decision

Rules:
- Do not claim production cloud tests passed unless evidence exists.
- Do not hide skipped tests.
- Do not add major new features.
- Fix only release-blocking issues.
- No credentials.
- No apply.

Final response:
- State release candidate status.
- List blockers.
- List commands run.
- List changed files.
```

---

## Preporučeni redosled za batch 121–160

Ne moraš sve redom. Najbolji redosled je:

```text
121  v0.3 release gate
122  plugin system MVP
123  template pack registry
124  policy engine v2
126  AWS multi-account
128  Kubernetes upgrade planner
129  platform add-on upgrade planner
130  execution profiles
131  Terraform Cloud support
140  FinOps v2
141  inventory export
142  Backstage integration
143  service catalog
148  secret rotation workflow
150  compliance mapping
155  air-gapped bundle manifest
157  migration analyzer
159  v0.4 roadmap
160  v0.4 release candidate
```

Najpraktičniji sledeći prompt odmah posle 120 je:

```text
Prompt 121 — v0.3 release gate review
```

On prvo kaže da li je v0.3 stvarno spreman pre nego što uđeš u v0.4 enterprise funkcije.
