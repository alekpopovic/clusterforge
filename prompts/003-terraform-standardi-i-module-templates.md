## Prompt 3 — Terraform standardi i module templates

```text
Create a consistent Terraform module template across all empty module directories.

For each Terraform module that currently contains only placeholders, update the files with a consistent style:

versions.tf:
- Include terraform required_version >= 1.6.0.
- Include required_providers only when the module directly uses provider resources.
- For placeholder modules that do not yet use providers, include only required_version.

variables.tf:
- Include common variables where appropriate:
  - name
  - environment
  - labels or tags depending on module type
- Add descriptions and types.
- Add validation for name/environment where useful.

main.tf:
- Add locals for common naming conventions where appropriate.
- Do not add fake resources.
- If module is not implemented, add clear TODO comments.

outputs.tf:
- Add useful outputs only if there is something real to output.
- For placeholder modules, do not create fake outputs.

README.md:
- Add module title.
- Add purpose.
- Add status: placeholder or implemented.
- Add expected future resources.
- Add usage skeleton.

Add a root-level docs/conventions.md file explaining:
- Module structure
- Naming conventions
- Environment separation
- Provider placement
- State separation
- Secret handling
- Production safety rules

Run terraform fmt -recursive if available.
```

---
