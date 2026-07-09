## Prompt 47 — CLI doctor command hardening

```text
Improve cf doctor command.

Goal:
Diagnose local setup and project health.

Checks:
1. Required binaries:
   - terraform
   - tofu optional
   - git
   - kubectl optional
   - helm optional
   - go only when developing CLI

2. Terraform:
   - version
   - can run terraform version
   - warn if version below minimum

3. Project:
   - clusterforge.yaml exists
   - project.name exists
   - environments configured
   - environment paths exist

4. CLI:
   - version info available
   - config can be loaded

5. Safety:
   - prod environments use remote backend or warn
   - prod destroy block enabled
   - require_plan_file_for_apply enabled for prod

6. Git:
   - warn if not inside Git repository
   - warn if tfstate files are tracked
   - warn if .env or kubeconfig files are tracked

Output:
- Clear pass/warn/fail table.
- Return non-zero only for hard failures.
- Add --json output option.

Tests:
- Mock checks where possible.
- Test config validation path.

Docs:
- Update docs/cli.md with doctor examples.

Run:
- gofmt
- go test ./...
```

---
