# Prompt 270 — Data residency model

```text
Create data residency model for ClusterForge.

Goal:
Define how self-hosted and future managed deployments can control where metadata and artifacts are stored.

Create:
- docs/rfcs/034-data-residency.md
- docs/control-plane/data-residency.md

Cover:

1. Data categories
   - organization metadata
   - project metadata
   - environment metadata
   - policy results
   - drift results
   - cost reports
   - artifacts
   - audit events
   - usage events
   - tokens
   - runner metadata

2. Residency controls
   - deployment region
   - database region
   - artifact bucket region
   - runner region
   - backup region
   - logs/metrics region

3. Organization-level settings
   - preferred_region
   - allowed_regions
   - artifact_region
   - backup_region

4. Enforcement
   - API refuses artifact storage outside allowed region
   - runner pool must match allowed region for regulated environments
   - backups respect residency settings

5. Non-goals
   - legal advice
   - automatic compliance certification
   - public SaaS guarantees

6. Future commands
   - cf residency show
   - cf residency validate
   - cf residency report

Do not implement code unless simple validation is already available.
Update roadmap.
```
