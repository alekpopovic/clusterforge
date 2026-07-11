# AWS multi-account operations

Use separate AWS accounts for development, staging, production, shared
services, and security/log archives. Environments reference account metadata;
credentials are never stored in `clusterforge.yaml`.

Prefer AWS Organizations with short-lived deployment roles. Local operators
may select a named profile, while GitHub Actions should exchange its OIDC token
for a narrowly scoped role. Avoid long-lived access keys and root principals.

State may live in each workload account for isolation, or in a centralized
state account with tightly scoped cross-account roles. In both cases use
encryption, locking, versioning, audit logging, and separate state keys.

ClusterForge warns when production shares an account with another environment
unless `allow_shared_prod_aws_account: true` is explicitly set. Run
`cf account doctor prod` and `cf env doctor prod` before generation.
