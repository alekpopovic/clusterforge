# Contributing to ClusterForge

Thanks for helping shape ClusterForge. This project aims to keep
Infrastructure-as-Code readable, reviewable, and production-minded.

Before contributing, read [`AGENTS.md`](AGENTS.md). It contains the repository
rules for humans and AI agents.

## Issues and pull requests

Use the structured GitHub form that best matches a bug, feature, module, or
security-hardening proposal. Suspected exploitable vulnerabilities must follow
[`SECURITY.md`](SECURITY.md) and must not be filed publicly. Redact credentials,
state values, kubeconfigs, account identifiers, private hostnames, and other
sensitive data from every issue and log.

Before opening a pull request, keep the change focused, add tests and docs where
behavior changes, run the relevant Make targets, and complete the pull request
checklist. Record checks that could not run instead of claiming they passed.
Maintainers review generated Terraform readability, production safety,
backward compatibility, and module documentation in addition to code behavior.

## Development Setup

The recommended reproducible setup is the optional devcontainer. Install
Docker, Visual Studio Code, and the Dev Containers extension, then choose
**Dev Containers: Reopen in Container**. See
[`docs/development-environment.md`](docs/development-environment.md). Local
development remains supported, and optional asdf pins are documented in
[`docs/tool-versions.md`](docs/tool-versions.md).

Install:

- Terraform or OpenTofu
- Go
- Git
- Optional: TFLint, Checkov, Trivy, terraform-docs

Build the CLI:

```bash
cd cli
go build -o cf .
cd ..
```

List local developer commands:

```bash
make help
```

Install the repository's optional pre-commit hook:

```bash
./scripts/install-hooks.sh
pre-commit run --all-files
```

The hooks format Terraform and Go, validate YAML, and use locally installed
Gitleaks/ShellCheck when available. See [`docs/pre-commit.md`](docs/pre-commit.md).

## Testing and Validation

Run the default local checks:

```bash
make fmt-check
make test
make validate
```

For CLI-only changes:

```bash
make test-cli
```

For security scans when tools are installed:

```bash
make security
```

For generated module documentation:

```bash
make docs
```

Do not run `terraform apply` or destructive commands as part of default
validation.

## Terraform Module Rules

Terraform modules live under `modules/`. Every module must include:

- `main.tf`
- `variables.tf`
- `outputs.tf`
- `versions.tf`
- `README.md`

Module expectations:

- Keep each module focused on one responsibility.
- Use typed variables with descriptions.
- Add validation blocks for important inputs.
- Add output descriptions.
- Keep provider configuration in root environments, not child modules.
- Do not hardcode environment-specific values.
- Do not add fake resources to satisfy scanners or examples.
- Avoid storing secrets in Terraform state unless there is no practical
  alternative and the README clearly documents the risk.
- Include a practical README usage example.

## CLI Rules

The CLI lives under `cli/` and is written in Go with Cobra.

Guidelines:

- Add commands under `cli/cmd`.
- Put reusable logic under `cli/internal`.
- Keep generated Terraform readable.
- Do not hide provider configuration or infrastructure logic.
- Add tests for config loading, generation, policy checks, app manifests, and
  command behavior where practical.
- Destructive commands must require explicit flags or confirmation.
- Production apply must require an existing plan file.
- Production destroy must be blocked by default.

## Documentation Rules

Documentation should be practical and direct.

- Update `README.md` when user-facing workflows change.
- Update the relevant file under `docs/` for architecture, CLI, security,
  environment layout, backend, GitOps, or validation changes.
- Keep examples copy-paste friendly and free of real account IDs, credentials,
  private keys, and secrets.
- Module READMEs should include purpose, status, usage, generated Terraform
  docs, notes, and TODOs when applicable.

## Security Rules

Never commit:

- Credentials
- `.env` files
- Kubeconfig files
- Private keys
- Terraform state files
- Terraform plan files
- Real secret values in `tfvars`

Prefer references to external secret stores. For Kubernetes workloads, the
recommended pattern is:

1. Store real values in a cloud secret manager.
2. Use External Secrets Operator to sync them into Kubernetes Secrets.
3. Reference Kubernetes Secret names and keys from workload modules.

## Release Contributions

Before a release candidate:

```bash
make fmt
make lint
make test
make validate
make security
cd cli && go build -o cf .
```

Record exact pass/fail results in the release report. Do not claim a cloud
apply, upgrade, or migration was tested unless it actually was.
