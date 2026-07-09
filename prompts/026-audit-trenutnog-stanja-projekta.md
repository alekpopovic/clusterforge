## Prompt 26 — Audit trenutnog stanja projekta

```text
Perform a full audit of the current ClusterForge repository.

Goal:
Create a clear status report of what exists, what is implemented, what is placeholder-only, what is broken, and what should be done next.

Inspect:
- modules/
- live/
- examples/
- cli/
- policies/
- scripts/
- .github/workflows/
- docs/
- README.md
- AGENTS.md

Create or update:
- STATUS.md

STATUS.md must include:

1. Repository summary
   - Current project state
   - Main implemented areas
   - Main missing areas

2. Terraform module status table
   Columns:
   - Module path
   - Status: implemented / partial / placeholder / missing
   - Providers used
   - Has README
   - Has example
   - Validation status
   - Notes

3. CLI status table
   Columns:
   - Command
   - Status
   - Tests
   - Notes

4. Examples status table
   Columns:
   - Example path
   - Status
   - Can run terraform validate?
   - Requires cloud credentials?
   - Notes

5. Security status
   - Secrets exposure risk
   - tfstate protection
   - production safety
   - destructive command protections
   - missing guardrails

6. CI status
   - Workflows present
   - What each workflow does
   - What might fail
   - What needs improvement

7. Immediate next tasks
   Group into:
   - Critical
   - Important
   - Nice to have

8. Recommended next milestone
   Define the next 5–10 tasks in exact order.

Rules:
- Do not make large code changes.
- Only create/update STATUS.md.
- You may fix tiny documentation typos only if necessary.
- Do not implement new modules in this prompt.
- Run lightweight checks if possible:
  - terraform fmt -check -recursive
  - go test ./... inside cli if cli exists
  - scripts/validate.sh if safe

Final response:
- Summarize audit findings.
- List commands run and whether they passed.
- Mention top 5 next recommended tasks.
```

---
