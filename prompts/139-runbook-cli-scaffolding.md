## Prompt 139 — Runbook CLI scaffolding

```text
Add runbook discovery support to ClusterForge CLI.

Goal:
Let users discover operational runbooks from the CLI.

Commands:
- cf runbook list
- cf runbook show <name>
- cf runbook search <query>
- cf runbook open <name>

Behavior:
- Read Markdown runbooks from docs/incident-response and docs/dr.
- Do not execute runbook steps.
- Search by title, filename, tags if frontmatter exists.
- Print path and summary.
- open command can print file path or open browser/editor only if safe and explicit.

Add optional frontmatter to runbooks:
---
title: Failed Terraform Apply
category: terraform
severity: high
tags:
  - terraform
  - apply
  - incident
---

Tests:
- list runbooks
- search runbooks
- show runbook
- missing runbook fails clearly

Docs:
- docs/runbooks.md

Rules:
- Read-only.
- No destructive action.
- No background processes.

Run:
- gofmt
- go test ./...
```


---
