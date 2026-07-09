## Prompt 118 — GitHub issue and PR templates

```text
Add GitHub community and maintainer templates.

Create:
- .github/ISSUE_TEMPLATE/bug_report.yml
- .github/ISSUE_TEMPLATE/feature_request.yml
- .github/ISSUE_TEMPLATE/module_request.yml
- .github/ISSUE_TEMPLATE/security_hardening.yml
- .github/PULL_REQUEST_TEMPLATE.md
- .github/CODEOWNERS placeholder
- SECURITY.md if missing

Bug report template must ask for:
- ClusterForge version
- CLI version
- Terraform/OpenTofu version
- provider versions
- cloud/orchestrator
- command run
- expected behavior
- actual behavior
- logs with secrets redacted

PR template must include checklist:
- tests added/updated
- docs updated
- terraform fmt run
- go test run
- no secrets
- production safety considered
- module README updated

Rules:
- Keep templates practical.
- Do not require private information.
- Emphasize redacting secrets.

Update CONTRIBUTING.md with issue/PR process.
```

---
