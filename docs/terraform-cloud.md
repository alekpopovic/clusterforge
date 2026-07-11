# Optional HCP Terraform integration

ClusterForge can generate a Terraform `cloud` block for remote state and
remote execution. The integration is optional and makes no API calls in this
MVP. It never stores or requests a Terraform Cloud token.

Configure organization, project metadata, and an explicit workspace name per
environment. Use `cf tfe workspace list`, `cf tfe backend render dev`, and
`cf tfe workspace generate dev`. Generation replaces that environment's
`backend.tf`; review the diff before initialization or state migration.

HCP Terraform variable sets can hold shared non-secret values, while sensitive
variables must be marked sensitive and managed outside Git. Teams may use a
VCS-driven workflow or CLI-driven runs, but production applies still require
reviewed plans, policy gates, and manual approval. Use stable workspace naming
and separate production permissions. Existing S3, Azure Storage, GCS, and local
backend support remains available.
