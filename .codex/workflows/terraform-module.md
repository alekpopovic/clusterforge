# Terraform Module Workflow

1. Confirm module scope and provider ownership.
2. Create or update `main.tf`, `variables.tf`, `outputs.tf`, `versions.tf`,
   and `README.md`.
3. Add examples when the module is user-facing.
4. Keep variables typed and documented.
5. Mark sensitive outputs.
6. Run formatting and validation.

Commands:

```bash
terraform fmt -recursive
make validate
```
