---
title: CLI
permalink: /cli/
---

# CLI

The ClusterForge CLI is a wrapper and generator. It helps create project files,
environment roots, app manifests, and safer Terraform/OpenTofu workflows. It is
not a replacement for reviewing Terraform.

## Structure

```text
cli/
  main.go
  cmd/                  Cobra commands
  internal/
    app/                App manifests and rendering
    config/             clusterforge.yaml schema and loader
    generator/          Environment file generation
    policy/             Safety checks and risk evaluation
    terraform/          Terraform/OpenTofu runner
    ui/                 Output helpers
  templates/            Terraform root templates
```

## Common Commands

```bash
cf version
cf project init demo
cf env create dev --cloud aws --orchestrator eks --region eu-central-1
cf env list
cf generate dev
cf init dev
cf plan dev --out .cf/plans/dev.tfplan --risk-summary
cf policy check dev --plan-file .cf/plans/dev.tfplan
cf apply prod --plan-file .cf/plans/prod.tfplan --confirm-prod
cf destroy prod --allow-destroy --confirm-prod
cf doctor
```

## App Manifests

Create an app manifest:

```bash
cf app add api --image ghcr.io/company/api:1.0.0 --port 8080 --replicas 2
```

Render it into an environment:

```bash
cf app render api --env dev
```

The renderer writes Terraform module calls into `env.path/apps/<name>.tf`.

## Engine Selection

Terraform is the default engine. Use OpenTofu with:

```bash
cf --engine tofu plan dev
```

Engine binaries are configured in `clusterforge.yaml`.

## Policy And Risk Summaries

Use risk summaries before applying reviewed plans:

```bash
cf plan prod --out .cf/plans/prod.tfplan --risk-summary
cf policy check prod --plan-file .cf/plans/prod.tfplan
```

Production apply requires a plan file and explicit confirmation:

```bash
cf apply prod --plan-file .cf/plans/prod.tfplan --confirm-prod
```

If a production plan contains delete actions, apply is blocked unless the
operator also passes `--allow-destroy`.

## Development

Run all CLI checks:

```bash
make test-cli
```

Build the CLI:

```bash
make build-cli
```

Run the default local CI checks without cloud credentials:

```bash
make ci
```

When adding CLI behavior:

- Put command wiring in `cli/cmd/`.
- Put business logic in `cli/internal/`.
- Add unit tests for parsers, renderers, policy checks, and command builders.
- Avoid destructive defaults.
- Do not hide generated Terraform logic.
