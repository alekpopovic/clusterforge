# Prompt 226 — Preview TTL controller job

```text
Implement preview environment TTL cleanup support.

Goal:
Automatically identify and clean expired preview environments.

Control Plane:
- preview_environments table:
  - id
  - project_id
  - environment_id
  - app
  - pr_number
  - namespace
  - status
  - created_by
  - created_at
  - expires_at
  - deleted_at

Runner job:
- preview_cleanup

CLI:
- cf preview cleanup --dry-run
- cf preview cleanup --execute

Behavior:
- dry-run default
- execute requires confirmation or runner policy
- never cleanup non-preview resources
- only cleanup resources with preview labels
- audit cleanup

Tests:
- expired preview detected
- non-expired preview ignored
- cleanup dry-run does not delete
- missing labels block deletion
- audit created

Docs:
- docs/preview-environments.md

Rules:
- Be conservative.
- No production cleanup unless explicitly allowed.
```
