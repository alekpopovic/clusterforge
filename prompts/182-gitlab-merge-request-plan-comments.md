# Prompt 182 — GitLab merge request plan comments

```text
Implement GitLab merge request comment support.

CLI command:
- cf gitlab comment-plan
- cf gitlab comment-policy
- cf gitlab comment-cost

Inputs:
- --project-id
- --mr number
- --token-env GITLAB_TOKEN
- --plan-summary file
- --policy-results file
- --cost-report file
- --gitlab-url default https://gitlab.com

Behavior:
- Create or update a bot comment/note.
- Include sanitized plan summary.
- Include policy results.
- Include cost warnings.
- Do not include raw plan output.

Tests:
- GitLab API client mocked
- comment rendering
- existing note update
- token redaction

Docs:
- docs/gitlab-mr-comments.md
- example GitLab CI template

Rules:
- No automatic apply.
- Token only from env.
- No secrets in comments.
```
