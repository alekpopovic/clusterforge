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
cf apply prod --plan-file .cf/plans/prod.tfplan --confirm-prod
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

## Development

Run all CLI checks:

```bash
./scripts/test-cli.sh
```

When adding CLI behavior:

- Put command wiring in `cli/cmd/`.
- Put business logic in `cli/internal/`.
- Add unit tests for parsers, renderers, policy checks, and command builders.
- Avoid destructive defaults.
- Do not hide generated Terraform logic.
