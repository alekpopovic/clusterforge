## Prompt 39 — CLI interactive wizard

```text
Add optional interactive wizard mode to ClusterForge CLI.

Goal:
Make it easier to create projects, environments, and apps.

Commands:
- cf project init
  When name is not provided, prompt for name.

- cf env create
  When required flags are missing, prompt for:
    - environment name
    - cloud
    - orchestrator
    - region
    - path confirmation

- cf app add
  When required flags are missing, prompt for:
    - app name
    - type
    - image
    - port
    - replicas
    - ingress host
    - autoscaling yes/no

Use a Go prompt library only if already acceptable for this project.
Otherwise implement simple stdin prompts.

Requirements:
- Add --non-interactive flag.
- In --non-interactive mode, missing required values must return clear errors.
- Interactive mode must never ask for secret values.
- Interactive mode must show generated summary before writing files.
- Existing flags should still work.

Tests:
- Test non-interactive validation.
- Keep prompt logic separable to avoid hard-to-test code.

Docs:
- Update docs/cli.md with wizard examples.

Rules:
- Do not make interactive prompts mandatory.
- CI must run in non-interactive mode.
- Do not block automation.

Run:
- gofmt
- go test ./...
```

---
