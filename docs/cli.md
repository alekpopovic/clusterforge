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
cf completion bash
cf project init demo
cf project init
cf env create dev --cloud aws --orchestrator eks --region eu-central-1
cf env create
cf backend configure dev --backend local
cf backend show dev
cf app validate api
cf app validate --all
cf env list
cf generate dev --layout simple
cf generate dev --layout stacked
cf init dev
cf plan dev --stack network
cf plan dev --out .cf/plans/dev.tfplan --risk-summary
cf policy check dev --plan-file .cf/plans/dev.tfplan
cf apply prod --plan-file .cf/plans/prod.tfplan --confirm-prod
cf destroy prod --allow-destroy --confirm-prod
cf doctor
cf doctor --json
```

## Interactive Wizard Mode

ClusterForge prompts for missing required values when running interactively:

```bash
cf project init
cf env create
cf app add
```

Wizard prompts are intentionally limited to non-secret project metadata,
environment settings, and app shape. They never ask for credentials, token
values, private keys, or secret contents.

Each wizard flow prints a summary before writing files.

## Non-Interactive Mode

Use `--non-interactive` in scripts and CI so missing required values fail
instead of prompting:

```bash
cf --non-interactive project init demo
cf --non-interactive env create dev --cloud aws --orchestrator eks --region eu-central-1
cf --non-interactive app add api --image ghcr.io/company/api:1.0.0 --port 8080 --autoscaling
```

## Install From GitHub Releases

Linux and macOS releases include `install.sh`, platform binaries, individual
`.sha256` files, and `SHA256SUMS`:

```bash
curl -fsSL https://github.com/alekpopovic/clusterforge/releases/latest/download/install.sh | bash
```

The installer detects `linux|darwin` and `amd64|arm64`, downloads `cf` plus its
checksum, verifies it, and installs to `$HOME/.local/bin`. Pin production
automation instead of following `latest`:

```bash
curl -fsSL https://github.com/alekpopovic/clusterforge/releases/download/v0.4.0/install.sh \
  | VERSION=v0.4.0 INSTALL_DIR="$HOME/bin" bash
```

For higher assurance, download `install.sh` and `SHA256SUMS`, inspect the script,
verify the release through the organization's trusted process, then execute it.
Windows users should download `cf-windows-amd64.exe` and its checksum manually.

Environment variables:

- `VERSION`: `latest`, `v0.4.0`, or `0.4.0`.
- `INSTALL_DIR`: destination directory.
- `BINARY_NAME`: installed filename, default `cf`.
- `CLUSTERFORGE_REPO`: GitHub `owner/repo`, useful for approved mirrors/forks.
- `CLUSTERFORGE_DOWNLOAD_BASE_URL`: explicit release asset base URL for testing or
  an internal mirror.

## Install From Source

Use the install script to build `cf` and install it into `/usr/local/bin`:

```bash
./scripts/install-cli.sh
```

Install into a user-writable directory without sudo:

```bash
INSTALL_DIR="$HOME/.local/bin" ./scripts/install-cli.sh
```

The script builds `cli/cf`, injects version metadata from Git, installs the
binary, and prints `cf version`.

## Build Manually

For local development:

```bash
cd cli
go build -o cf .
```

For a metadata-injected build:

```bash
cd cli
go build -trimpath \
  -ldflags "-s -w -X github.com/textracta/clusterforge/cli/cmd.Version=$(git describe --tags --always --dirty) -X github.com/textracta/clusterforge/cli/cmd.Commit=$(git rev-parse --short HEAD) -X github.com/textracta/clusterforge/cli/cmd.Date=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
  -o cf .
```

`make build-cli` uses the same ldflags pattern.

## Version Command

```bash
cf version
```

The command prints:

- version
- commit
- build date
- Go runtime version

## Shell Completion

Generate completion scripts with:

```bash
cf completion bash
cf completion zsh
cf completion fish
cf completion powershell
```

Example bash installation:

```bash
cf completion bash > ~/.local/share/bash-completion/completions/cf
```

Example zsh installation:

```bash
cf completion zsh > "${fpath[1]}/_cf"
```

## Release Artifacts

Tagged pushes matching `v*` run the CLI release workflow. It cross-compiles:

- `linux amd64`
- `linux arm64`
- `darwin amd64`
- `darwin arm64`
- `windows amd64`

The workflow uploads binaries and SHA256 checksum files as workflow artifacts
and GitHub release artifacts. It does not publish packages to external package
registries.

## App Manifests

Create an app manifest:

```bash
cf app add api --image ghcr.io/company/api:1.0.0 --port 8080 --replicas 2 --autoscaling
```

Create one with prompts:

```bash
cf app add
```

Validate manifests before rendering:

```bash
cf app validate api
cf app validate --all
```

Render it into an environment:

```bash
cf app render api --env dev
```

The renderer writes Terraform module calls into `env.path/apps/<name>.tf`.

See [App Manifest]({{ '/app-manifest/' | relative_url }}) for the full schema
and validation rules.

## JSON Output

Selected commands support JSON output for CI and scripts:

```bash
cf env list --json
cf app list --json
cf doctor --json
cf policy check dev --json
cf plan dev --risk-summary --json
```

See [CLI JSON Output]({{ '/cli-json/' | relative_url }}) for stable response
schemas.

## Engine Selection

Terraform is the default engine. Use OpenTofu with:

```bash
cf --engine tofu plan dev
```

Engine binaries are configured in `clusterforge.yaml`.

## Environment Layouts

ClusterForge supports simple and stacked environment layouts:

```bash
cf generate dev --layout simple
cf generate dev --layout stacked
```

Simple layout writes one Terraform root at `env.path`. Stacked layout writes
separate roots under `network`, `cluster`, `platform`, and `apps`.

For stacked environments, stack-aware commands accept `--stack`:

```bash
cf plan dev
cf plan dev --stack network
cf apply dev --stack cluster
cf output dev --stack network --json
```

Without `--stack`, `plan`, `init`, and `apply` run stacks in dependency order:
`network`, `cluster`, `platform`, then `apps`. Destroy runs the reverse order.

See [Environments]({{ '/environments/' | relative_url }}) for layout details
and production recommendations.

## Backend Configuration

Configure backend generation before running `cf generate`:

```bash
cf backend configure dev --backend local
cf backend configure prod \
  --backend s3 \
  --bucket my-terraform-state-bucket \
  --region eu-central-1 \
  --dynamodb-table my-terraform-locks \
  --key-prefix clusterforge/prod
cf backend show prod
cf generate prod --force
```

The CLI never asks for or writes cloud credentials into backend configuration.
Production environments should use a remote backend with locking.

See [Backends]({{ '/backends/' | relative_url }}) for bootstrap guidance.

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

## Operational Helpers

Read-only and advisory workflows:

```bash
cf drift check dev
cf drift check dev --json
cf state list dev
cf state show dev module.network.aws_vpc.this
cf state pull dev --output /secure/path/state.json
cf cost scan dev --plan-file .cf/plans/dev.tfplan
cf policy list
cf policy check prod --pack production --plan-file .cf/plans/prod.tfplan
```

Local project workflow helpers:

```bash
cf upgrade check
cf upgrade plan
cf upgrade apply --yes
cf promote diff --from staging --to prod
cf template list
cf template validate
cf generate dev --template-pack company-standard
```

These commands do not auto-remediate infrastructure. State output and plan files
may contain sensitive values and should stay out of git.

## Doctor

Use `doctor` to check local tooling, project config, safety settings, and common
Git hygiene issues:

```bash
cf doctor
```

The command prints a pass/warn/fail table and exits non-zero only for hard
failures such as missing required binaries, invalid config, or missing
environment paths.

JSON output is useful for CI or support bundles:

```bash
cf doctor --json
```

`doctor` warns when optional tools such as `tofu`, `kubectl`, `helm`, or `go`
are missing. It also warns if production environments use a local backend, if
production safety policies are disabled, or if sensitive files such as
`tfstate`, `.env`, or kubeconfig files are tracked by Git.

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
