## Prompt 114 — Pre-commit hooks

```text
Add pre-commit hook support.

Goal:
Help contributors catch issues before pushing.

Create:
- .pre-commit-config.yaml
- docs/pre-commit.md

Hooks:
- trailing whitespace
- end of file fixer
- YAML check
- Terraform fmt
- Go fmt
- secret scan if tool available
- shellcheck if scripts exist

Scripts:
- scripts/install-hooks.sh

Rules:
- Hooks should be helpful but not overly slow.
- Do not require cloud credentials.
- Secret scanning should avoid false positives where possible.
- Document how to skip hooks in emergencies.

Update:
- README development section
- CONTRIBUTING.md

Run:
- pre-commit run --all-files if pre-commit is available
- otherwise verify config syntax manually where possible

Final response:
- List hooks added.
- Mention whether pre-commit was available.
```

---
