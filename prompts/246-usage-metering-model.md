# Prompt 246 — Usage metering model

```text
Create usage metering model for ClusterForge.

Goal:
Track operational usage for capacity planning, enterprise reporting, and possible future billing, without implementing billing.

Create:
- docs/rfcs/026-usage-metering.md
- docs/control-plane/usage-metering.md

Metrics to meter:
- organizations
- workspaces
- projects
- environments
- clusters
- apps
- runners
- jobs by type
- plan requests
- apply requests
- policy checks
- drift checks
- cost scans
- artifact storage bytes
- artifact downloads
- audit event volume
- preview environments
- API requests
- active users

Metering principles:
- no secret values
- no raw Terraform state
- no raw plan content
- tenant-scoped
- exportable
- retention controlled
- disabled or reduced mode for privacy-sensitive installs

Data model:
- usage_events
- usage_rollups_daily
- usage_rollups_monthly

Dimensions:
- organization_id
- workspace_id
- project_id
- environment_id
- event_type
- quantity
- metadata_json sanitized

Do not implement code in this prompt.

Include:
- data retention policy
- privacy considerations
- self-hosted reporting
- future SaaS billing note
- examples of usage reports

Update:
- ROADMAP_V0.7.md if it exists
```
