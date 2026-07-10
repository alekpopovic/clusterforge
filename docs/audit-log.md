# Local CLI audit log

New ClusterForge projects enable a local JSON-lines audit trail:

```yaml
audit:
  enabled: true
  path: .cf/audit.log
```

Existing projects remain disabled until this block is added. The audit log is a
local troubleshooting and accountability aid; ClusterForge sends no telemetry or
log data to any remote service.

Recorded operations include project initialization, environment creation,
generation, app add/render/remove, plan, apply, destroy, drift and policy checks,
and upgrade apply. Each entry includes UTC timestamp, local user, command,
redacted arguments, working directory, environment/stack when available, result,
duration, and CLI version.

Flags whose names contain `password`, `token`, `secret`, or `key` have their
values replaced with `[REDACTED]`. Do not treat redaction as a license to pass
secrets on command lines: process lists, shells, and other tools may still expose
them. The audit logger deliberately does not record command output, Terraform
plans, state, environment variables, or error text.

```bash
cf audit show
cf audit show --json
cf audit tail --lines 50
cf audit clear
cf audit clear --yes --non-interactive
```

Clearing requires interactive confirmation or `--yes`. Audit management commands
do not audit themselves, so clearing does not recreate the file. `.cf/audit.log`
is explicitly ignored by Git and created with owner-only file permissions.

Local users with filesystem access can modify or delete this log. It is not an
immutable compliance record, cryptographically signed trail, or centralized SIEM
integration. Organizations needing those properties should collect reviewed logs
outside ClusterForge without adding secrets to repository configuration.
