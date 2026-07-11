## Prompt 125 — Organization and workspace model

```text
Add an organization/workspace model to ClusterForge configuration.

Goal:
Support larger teams managing multiple projects, environments, and clusters.

Config additions:
clusterforge.yaml:
  organization:
    name: example-org
    owner: platform-team
    contact: platform@example.com

  workspaces:
    platform:
      description: Core platform environments
      environments:
        - dev
        - staging
        - prod

    apps:
      description: Application workloads
      environments:
        - dev
        - staging
        - prod

  teams:
    platform:
      owners:
        - platform@example.com
      namespaces:
        - platform-system

    payments:
      owners:
        - payments@example.com
      namespaces:
        - payments-dev
        - payments-prod

CLI:
- cf workspace list
- cf workspace show <name>
- cf workspace doctor <name>
- cf team list
- cf team show <name>

Behavior:
- Existing single-project config remains valid.
- Workspaces are metadata in this phase.
- No remote backend or SaaS required.
- Future dashboard/API can use this model.

Docs:
- docs/organization-model.md
- Explain project vs workspace vs environment vs cluster vs team.
- Explain recommended layouts for small teams and large organizations.

Tests:
- Load config with organization.
- Load config without organization.
- List workspaces.
- List teams.
- Validation errors for duplicate workspace names.

Run:
- gofmt
- go test ./...
```


---
