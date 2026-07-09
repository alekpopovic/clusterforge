## Prompt 113 — CLI audit log

```text
Add local audit log support to ClusterForge CLI.

Goal:
Record important CLI operations for troubleshooting and accountability.

Config:
clusterforge.yaml:
audit:
  enabled: true
  path: .cf/audit.log

Commands to log:
- project init
- env create
- generate
- app add/render/remove
- plan
- apply
- destroy
- drift check
- policy check
- upgrade apply

Log fields:
- timestamp
- user if available
- command
- args with sensitive values redacted
- working directory
- environment
- stack
- result
- duration
- CLI version

Rules:
- Do not log secrets.
- Redact values for flags containing:
  password
  token
  secret
  key
- Audit log is local file only.
- Do not send telemetry anywhere.
- Add .gitignore entry for .cf/audit.log unless project decides otherwise.

CLI:
- cf audit show
- cf audit tail
- cf audit clear with confirmation

Tests:
- Log entry written.
- Sensitive flag redacted.
- Audit disabled works.
- Clear requires confirmation or --yes.

Docs:
- docs/audit-log.md

Run:
- gofmt
- go test ./...
```

---
