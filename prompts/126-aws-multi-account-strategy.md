## Prompt 126 — AWS multi-account strategy

```text
Design and implement AWS multi-account support.

Goal:
Allow ClusterForge projects to target separate AWS accounts for dev, staging, prod, shared services, and security.

Create docs:
- docs/aws-multi-account.md
- docs/rfcs/009-aws-multi-account.md

Config:
clusterforge.yaml:
  aws_accounts:
    dev:
      account_id: "111111111111"
      region: eu-central-1
      profile: dev
      role_arn: ""
    prod:
      account_id: "222222222222"
      region: eu-central-1
      profile: prod
      role_arn: "arn:aws:iam::222222222222:role/ClusterForgeDeployRole"

  environments:
    prod:
      cloud: aws
      account: prod
      region: eu-central-1
      orchestrator: eks
      path: live/prod/aws-eks

CLI:
- cf account list
- cf account show <name>
- cf account doctor <name>
- cf env doctor <env> should check account config

Generator:
- AWS provider templates should support:
  - profile
  - assume_role
  - region
  - default_tags

Security:
- Warn if prod and dev use same account unless explicitly allowed.
- Warn if root account appears to be used.
- Do not store AWS credentials.

Tests:
- Config loading.
- Provider template generation with profile.
- Provider template generation with assume_role.
- Prod account warning.

Docs:
- Explain AWS Organizations pattern.
- Explain deployment role pattern.
- Explain GitHub Actions OIDC role assumption.
- Explain state bucket per account vs centralized state account.

Run:
- gofmt
- go test ./...
- terraform fmt -recursive
```


---
