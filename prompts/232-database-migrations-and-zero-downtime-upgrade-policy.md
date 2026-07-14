# Prompt 232 — Database migrations and zero-downtime upgrade policy

```text
Harden database migrations and upgrade policy.

Goal:
Make Control Plane upgrades safer.

Tasks:
1. Document migration policy:
   - backward-compatible migrations preferred
   - expand/contract pattern
   - no destructive migrations without backup
   - migration rollback limitations

2. Add migration metadata table:
   - version
   - applied_at
   - checksum
   - description

3. CLI/server commands:
   - clusterforge-server migrate status
   - clusterforge-server migrate up
   - clusterforge-server migrate down only for dev or disabled if unsafe
   - clusterforge-server migrate validate

4. Startup behavior:
   - server refuses to start if migrations are missing unless auto_migrate enabled
   - auto_migrate default false for production

5. Tests:
   - migration status
   - migration checksum
   - missing migration detected
   - destructive down disabled in production

Docs:
- docs/control-plane/migrations.md
- docs/control-plane/upgrades.md

Rules:
- No silent destructive migrations.
- Backup before production migrations.
```
