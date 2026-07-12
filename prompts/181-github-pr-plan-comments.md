# Prompt 181 — GitHub PR plan comments

```text
Implement GitHub PR plan comment support.

Goal:
Allow ClusterForge to post sanitized Terraform plan summaries to GitHub pull requests.

Modes:
1. CLI mode in GitHub Actions
2. Control Plane mode later

CLI command:
- cf github comment-plan
- cf github comment-policy
- cf github comment-cost

Inputs:
- --repo owner/name
- --pr number
- --token-env GITHUB_TOKEN
- --plan-summary file
- --policy-results file
- --cost-report file

Behavior:
- Create or update a bot comment.
- Include:
  - environment
  - stack
  - create/update/delete/replace counts
  - risk level
  - policy status
  - cost warnings
  - link to artifacts if provided
- Redact sensitive data.
- Do not post raw plan output.

Tests:
- GitHub API client mocked
- comment body rendering
- existing comment update
- token not logged

Docs:
- docs/github-pr-comments.md
- example workflow

Rules:
- No automatic apply.
- No raw secrets.
- Token only from env.
```
