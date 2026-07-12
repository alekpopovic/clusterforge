# Prompt 180 — Git provider integration RFC

```text
Create Git provider integration RFC.

Create:
- docs/rfcs/018-git-provider-integration.md

Goal:
Design integration with GitHub and GitLab for plan comments, PR checks, and repository metadata.

Use cases:
- plan on pull request
- policy check on pull request
- cost warning on pull request
- drift report issue
- release automation
- template pack sources
- service catalog import

Providers:
- GitHub
- GitLab

Authentication:
- GitHub App preferred for future
- GitHub token for MVP optional
- GitLab project token / OAuth later
- no tokens committed
- environment variables or secret store

MVP recommendation:
- CLI-generated CI templates first
- Control Plane webhook integration later
- PR comments optional

Security:
- no apply from PR automatically
- fork PR restrictions
- secret exposure prevention
- plan output redaction
- permission scoping

Do not implement code in this prompt.
Update roadmap.
```
