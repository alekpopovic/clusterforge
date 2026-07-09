## Prompt 51 — Real MVP acceptance test

```text
Perform a real MVP acceptance review for ClusterForge.

Goal:
Determine whether ClusterForge is truly ready for v0.1.0 usage.

Inspect the whole repository and create:
- ACCEPTANCE_REPORT.md

The report must answer:

1. Can a new user clone the repo and understand it?
2. Can a new user build the CLI?
3. Can a new user run:
   - cf project init
   - cf env create
   - cf generate
   - cf app add
   - cf app validate
   - cf app render
   - cf doctor

4. Can Terraform formatting pass?
5. Can Terraform validation pass where expected?
6. Can Go tests pass?
7. Are safety policies enforced?
8. Are secrets avoided in examples?
9. Are production warnings clear?
10. Are placeholder modules clearly marked?

Run:
- make fmt-check
- make lint
- make test
- make validate
- cd cli && go build -o cf .
- ./cf version
- ./cf doctor if possible

Rules:
- Do not add new features.
- Fix only small blocking issues.
- Do not rewrite architecture.
- Do not add cloud credentials.
- Be honest about what was not tested.

ACCEPTANCE_REPORT.md must include:
- Passed checks
- Failed checks
- Skipped checks
- Known limitations
- Required fixes before public release
- Recommended v0.1.0 status:
  - ready
  - ready with warnings
  - not ready

Final response:
- State release readiness.
- List blockers.
- List commands run.
```

---
