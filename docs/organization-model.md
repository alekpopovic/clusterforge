# Organization and workspace model

ClusterForge keeps organization structure as local configuration metadata. It
does not require a SaaS control plane or change Terraform backend ownership.

- A **project** is one ClusterForge configuration and generated source tree.
- A **workspace** groups environments for navigation and future API/dashboard use.
- An **environment** is a deployable Terraform/OpenTofu root or stack layout.
- A **cluster** records an orchestrator target associated with an environment.
- A **team** records owners and Kubernetes namespace responsibility.

```yaml
organization:
  name: example-org
  owner: platform-team
  contact: platform@example.com
workspaces:
  platform:
    description: Core platform environments
    environments: [dev, staging, prod]
teams:
  payments:
    owners: [payments@example.com]
    namespaces: [payments-dev, payments-prod]
```

Inspect metadata with `cf workspace list`, `cf workspace show platform`,
`cf workspace doctor platform`, `cf team list`, and `cf team show payments`.

Small teams can omit all three sections and keep a single project with a few
environments. Larger organizations should use workspaces for lifecycle or
platform boundaries and teams for ownership, while keeping credentials,
authorization, and secrets in their existing identity systems. This model is
descriptive and does not itself grant access to clouds, state, or namespaces.
