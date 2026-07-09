## Prompt 63 — Upgrade and migration framework

```text
Design and implement a basic upgrade framework for ClusterForge CLI.

Goal:
Allow future ClusterForge versions to migrate clusterforge.yaml and generated project layouts safely.

Commands:
- cf upgrade check
- cf upgrade plan
- cf upgrade apply

Versioning:
- Add config field:
  clusterforge_version: "0.1.0"

Behavior:
cf upgrade check:
- Reads clusterforge.yaml.
- Compares config version with CLI supported version.
- Reports whether migration is needed.

cf upgrade plan:
- Shows proposed changes.
- Does not write files.

cf upgrade apply:
- Requires confirmation unless --yes.
- Backs up changed files to .cf/backups/<timestamp>/.
- Applies migrations.

First migrations:
1. Add missing clusterforge_version.
2. Normalize environment paths.
3. Add default policies if missing.
4. Add backend config skeleton if missing.

Rules:
- Never overwrite user changes without backup.
- Never modify Terraform state.
- Never run terraform apply.
- Keep migrations idempotent.
- Add tests for migration logic.

Docs:
- docs/upgrades.md

Run:
- gofmt
- go test ./...
```

---
