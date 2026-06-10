# ClusterForge Agent Guide

This file defines coding and infrastructure rules for contributors and AI
agents working in this repository.

## 1. Project Overview

- ClusterForge is a Terraform/OpenTofu framework for container orchestrators.
- ClusterForge supports Kubernetes, ECS, Nomad, and Docker targets.
- The CLI is a wrapper and generator, not a replacement for Terraform.
- Terraform/OpenTofu configuration must stay readable and reviewable.

## 2. Repository Conventions

- Terraform modules live under `modules/`.
- Real environment compositions live under `live/`.
- CLI source code lives under `cli/`.
- Examples live under `examples/`.
- Policies live under `policies/`.
- Scripts live under `scripts/`.

## 3. Terraform Module Rules

- Each module must have `main.tf`, `variables.tf`, `outputs.tf`,
  `versions.tf`, and `README.md`.
- Use typed variables.
- Use descriptions on all variables and outputs.
- Use validation blocks for important inputs.
- Do not hardcode environment-specific values.
- Do not configure providers inside reusable child modules unless unavoidable.
- Use `locals` for naming and common labels/tags.
- Do not store secrets in state unless absolutely unavoidable.
- Avoid excessive dynamic blocks unless they improve readability.
- Keep modules focused on one responsibility.

## 4. Provider Rules

- Provider configuration belongs in `live/` environment roots.
- Every root configuration must declare required providers.
- Pin provider versions using safe constraints.
- Prefer explicit provider aliases when multiple clusters or cloud accounts are
  involved.

## 5. CLI Rules

- The CLI is written in Go.
- Use Cobra for command structure.
- Use clear package boundaries under `cli/internal`.
- CLI commands must not perform destructive actions without confirmation.
- Production apply must require an existing plan file.
- Destroy in production must be blocked by default.
- The CLI must generate readable Terraform files.

## 6. Testing and Validation

- Run `terraform fmt -recursive`.
- Run `terraform validate` for examples where possible.
- Run `gofmt` and `go test` for CLI changes.
- Add shell scripts for repeatable validation.
- Add GitHub Actions for CI.

## 7. Documentation Rules

- Every module needs a `README.md` with a usage example.
- The root `README.md` must explain architecture, quickstart, and roadmap.
- Examples must be copy-paste friendly.

## 8. Security Rules

- Never commit credentials, kubeconfig files, private keys, or tfstate files.
- Do not generate real secrets into `tfvars`.
- Use references to external secret stores where possible.
- Production operations require explicit confirmation.

## 9. Codex Workflow Rules

- After every prompt that changes repository files, Codex must run `git status`,
  stage the completed changes with `git add`, create a commit, and push it.
- Do not create empty commits when no repository files changed.
- Do not commit or push secrets, credentials, kubeconfig files, private keys,
  tfstate files, or other sensitive artifacts.
