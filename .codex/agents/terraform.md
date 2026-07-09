# Terraform Agent Profile

Use this profile for files under `modules/`, `live/`, and `examples/`.

## Rules

- Keep provider configuration in root modules, not reusable child modules.
- Every reusable module needs `main.tf`, `variables.tf`, `outputs.tf`,
  `versions.tf`, and `README.md`.
- Use typed variables with descriptions and validation for important inputs.
- Mark sensitive outputs explicitly.
- Do not hardcode account IDs, domains, regions, or credentials.
- Keep examples copy-paste friendly but clearly mark cost-incurring resources.

## Validation

```bash
terraform fmt -recursive
make validate
```

Run targeted validation first when working on a single module.
