## Prompt 30 — TFLint, Checkov i Trivy config scanning

```text
Add static analysis and security scanning configuration.

Goal:
Improve Terraform quality and security checks without requiring cloud credentials.

Create or update:
- .tflint.hcl
- .checkov.yml
- trivy.yaml or documented Trivy config if appropriate
- scripts/security.sh
- scripts/lint.sh
- .github/workflows/security-scan.yml
- docs/security-scanning.md

scripts/security.sh must:
- Use bash strict mode.
- Run checkov if installed.
- Run trivy config if installed.
- Print clear warnings if tools are missing.
- Exit non-zero when tools are installed and findings exceed configured threshold.
- Avoid scanning .terraform directories, tfstate, vendor/cache directories.

scripts/lint.sh must:
- Run terraform fmt -check -recursive.
- Run tflint recursively where applicable if installed.
- Run gofmt check for CLI if present.
- Run go vet for CLI if practical.

TFLint:
- Configure base Terraform rules.
- Add AWS ruleset if practical.
- Do not require AWS credentials.

Checkov:
- Skip or document rules that are intentionally not applicable to placeholder modules.
- Do not hide real critical issues.

GitHub Actions:
- security-scan.yml should run on pull_request and push to main.
- It should not require cloud credentials.
- It should show useful output.

Update README:
- Add local commands:
  make lint
  make security

Rules:
- Do not suppress findings broadly.
- Any skipped rule must have a comment explaining why.
- Do not add fake resources just to satisfy scanners.

Final response:
- List tools configured.
- List files changed.
- List commands run.
- Mention any known scanner limitations.
```

---
