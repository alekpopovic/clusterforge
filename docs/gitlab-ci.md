# GitLab CI templates

Reusable examples live under `ci/gitlab/`. Include the files you need from the
same repository or copy and pin them in an internal CI template project:

```yaml
include:
  - local: ci/gitlab/cli-test.gitlab-ci.yml
  - local: ci/gitlab/terraform-plan.gitlab-ci.yml
  - local: ci/gitlab/terraform-apply-manual.gitlab-ci.yml
```

The plan template includes AWS EKS, AWS ECS, and existing-Kubernetes examples.
Adapt `CF_ENV` to the actual generated root layout. Plan files and JSON are
short-lived artifacts; restrict artifact access because plans can contain
sensitive values.

Production apply is manual, limited to the default branch, and consumes an
existing plan artifact. No automatic production apply or destroy job is
provided. Configure protected environments and required approvers in GitLab.

For cloud authentication, add GitLab `id_tokens` and exchange the OIDC token
for a narrowly scoped cloud role in the consuming project. The templates contain
no project IDs, account IDs, credentials, tokens, or real secrets.
