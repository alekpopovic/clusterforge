## Prompt 132 — GitLab CI templates

```text
Add GitLab CI support templates.

Goal:
Support teams using GitLab instead of GitHub Actions.

Create:
- ci/gitlab/
  terraform-plan.gitlab-ci.yml
  terraform-apply-manual.gitlab-ci.yml
  drift-check.gitlab-ci.yml
  cli-test.gitlab-ci.yml
  security-scan.gitlab-ci.yml

Docs:
- docs/gitlab-ci.md

Templates must include:
- stages:
  - validate
  - plan
  - policy
  - apply
- manual production apply
- no automatic prod apply
- artifacts for plan output
- optional OIDC/cloud auth placeholders
- no secrets in templates

Include examples:
- AWS EKS plan
- AWS ECS plan
- existing Kubernetes plan

Rules:
- Do not assume real GitLab project IDs.
- Do not include credentials.
- Production apply must be manual.
- Destroy must be absent or explicitly manual with warnings.

Update README:
- Add GitLab CI support section.

Final response:
- List templates created.
- Explain how users include them.
```


---
