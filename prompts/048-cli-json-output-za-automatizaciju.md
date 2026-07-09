## Prompt 48 — CLI JSON output za automatizaciju

```text
Add JSON output support to selected ClusterForge CLI commands.

Commands:
- cf env list --json
- cf app list --json
- cf doctor --json
- cf policy check <env> --json
- cf plan <env> --risk-summary --json

Requirements:
- Human output remains default.
- JSON output must be stable and documented.
- JSON must not include secrets.
- Errors should still be useful.
- Add internal/ui package support for human vs JSON output.

JSON schema examples:
env list:
{
  "environments": [
    {
      "name": "dev",
      "cloud": "aws",
      "orchestrator": "eks",
      "region": "eu-central-1",
      "path": "live/dev/aws-eks"
    }
  ]
}

doctor:
{
  "status": "warn",
  "checks": [
    {
      "name": "terraform binary",
      "status": "pass",
      "message": "terraform found"
    }
  ]
}

Tests:
- Validate JSON output can be unmarshaled.
- Ensure no secret fields are printed.
- Human output still works.

Docs:
- docs/cli-json.md

Run:
- gofmt
- go test ./...
```

---
