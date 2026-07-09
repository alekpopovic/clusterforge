# Codex Directory

This directory contains Codex-specific project guidance. It does not replace
`AGENTS.md`; it organizes practical context for recurring work.

## Layout

```text
.codex/
  agents/      Focused role profiles for common task types.
  workflows/   Repeatable implementation and validation flows.
  checklists/  Short pre-commit and release checks.
  context.md   High-signal project context for Codex sessions.
```

## How To Use

- For Terraform module work, read `agents/terraform.md`.
- For CLI changes, read `agents/cli.md`.
- For documentation tasks, read `agents/docs.md`.
- For security or release work, read the matching agent profile and checklist.
- For prompt execution, use the numbered files under `prompts/`.

Codex should keep changes readable, scoped, and reviewable. Generated Terraform
must remain understandable to humans.
