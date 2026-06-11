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

- Purpose
- Status
- Usage example
- Inputs and outputs summary
- Provider notes
- Safety or lifecycle notes

For placeholder modules, be explicit that no real resources are created yet.

## Development Checklist

Before committing a module change:

```bash
terraform fmt -recursive
./scripts/lint.sh
./scripts/validate.sh
```

Add focused tests or examples when the module behavior is non-trivial.
