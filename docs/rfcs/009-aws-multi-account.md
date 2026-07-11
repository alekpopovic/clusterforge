# RFC 009: AWS multi-account model

Status: implemented metadata/provider MVP.

ClusterForge maps an environment to a named `aws_accounts` entry containing an
account ID, default region, optional local profile, and optional deployment role
ARN. Generated AWS provider roots set region, optional profile, optional
`assume_role`, and stable default tags.

The model deliberately excludes credentials, AWS Organizations mutation,
account vending, and automatic role creation. Trust policies and GitHub Actions
OIDC subjects remain organization-owned and must be reviewed independently.
