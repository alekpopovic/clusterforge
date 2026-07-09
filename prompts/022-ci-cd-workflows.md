## Prompt 22 — CI/CD workflows

```text
Create GitHub Actions workflows for ClusterForge.

Files:
.github/workflows/terraform-validate.yml
.github/workflows/cli-test.yml
.github/workflows/security-scan.yml

terraform-validate.yml:
- Trigger on pull_request and push to main.
- Checkout.
- Setup Terraform.
- Run terraform fmt -check -recursive.
- Validate example root modules where possible.
- Do not require cloud credentials.
- Skip modules that require provider credentials unless validation can run without them.
- Make script resilient and explicit.

cli-test.yml:
- Trigger on pull_request and push to main.
- Setup Go.
- Run:
  cd cli
  go mod download
  gofmt check
  go test ./...
  go build -o cf .

security-scan.yml:
- Trigger on pull_request and push to main.
- Add Checkov if practical.
- Add Trivy config scan if practical.
- Make failures explicit but avoid requiring unavailable credentials.

Also update:
- scripts/lint.sh
- scripts/validate.sh
- scripts/test-cli.sh

Scripts:
- Use bash strict mode.
- Print clear sections.
- Exit non-zero on failure.
- Work from repo root.

README:
- Add CI status section placeholders.
- Add local validation commands.

Do not add secrets or real cloud account information.
```

---
