## Prompt 112 — Scheduled drift check workflow templates

```text
Add scheduled drift check workflow templates.

Goal:
Provide safe examples for teams to run drift detection in CI.

Create:
- .github/workflows/examples/drift-check-aws-eks.yml
- .github/workflows/examples/drift-check-aws-ecs.yml
- docs/ci-drift-checks.md

Workflow examples:
- scheduled cron
- manual workflow_dispatch
- configure cloud credentials through GitHub OIDC
- run cf drift check
- upload plan summary artifact
- do not apply
- fail or warn based on config

Docs:
- Explain required secrets/permissions.
- Explain GitHub OIDC.
- Explain why drift check should not auto-apply.
- Explain how to tune schedule.
- Explain handling failures.

Rules:
- Put workflows under examples or disabled path so they do not run by default.
- Do not include real account IDs.
- Do not include credentials.
- No apply.

Final response:
- List workflow templates.
- Explain how users enable them.
```

---
