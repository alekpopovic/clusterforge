---
title: Module Conventions
permalink: /module-conventions/
---

# Module Conventions

Reusable Terraform modules live under `modules/`. Each module should do one
thing and expose a small, typed interface.

## Required Files

Every module must include:

- `main.tf`
- `variables.tf`
- `outputs.tf`
- `versions.tf`
- `README.md`

Examples are useful when the module is meant to be used directly or has
non-obvious inputs.

## Variables

Use typed variables and descriptions for every input.

Add validation when invalid input could create confusing plans or unsafe
infrastructure. Common validations include:

- Non-empty names and environment values
- Allowed enum values
- CIDR syntax
- Matching list lengths
- Min/max bounds

Do not hardcode environment-specific values in reusable modules.

## Outputs

Outputs should expose useful real values:

- IDs
- ARNs
- Names
- Endpoints
- Subnet lists

Do not create fake outputs in placeholder modules.

## Providers

Reusable modules should declare required providers only when they use provider
resources. Do not configure providers in child modules unless there is no
reasonable alternative.

Provider blocks belong in `live/` roots and runnable examples.

## Naming And Metadata

Use locals for common names, tags, and labels. Prefer the core metadata
modules:

- `modules/core/naming`
- `modules/core/tags`
- `modules/core/labels`

Keep naming logic deterministic and readable.

## README Requirements

Each module README should include:

- Title
- Purpose
- Status
- Usage example
- Generated Terraform documentation
- Notes
- TODOs when the module is partial or placeholder-only

For placeholder modules, be explicit that no real resources are created yet.

Generated documentation must be wrapped in terraform-docs markers:

```markdown
<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
```

The generated section is managed by `terraform-docs` and should include
requirements, providers, inputs, and outputs. Keep hand-written context,
examples, lifecycle notes, and TODOs outside the generated block so useful
documentation is not lost.

Run module documentation generation from the repository root:

```bash
make docs
```

or:

```bash
./scripts/docs.sh
```

`scripts/docs.sh` only scans Terraform module directories under `modules/`.
It does not generate documentation for `live/` roots or examples.

Every implemented module should have at least one practical usage example in
its README. Prefer a minimal copy-paste friendly example that shows provider
configuration in the root module and the module call separately when that
distinction matters.

## Development Checklist

Before committing a module change:

```bash
terraform fmt -recursive
./scripts/lint.sh
./scripts/validate.sh
```

Add focused tests or examples when the module behavior is non-trivial.
