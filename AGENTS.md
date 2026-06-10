# ClusterForge Agent Notes

ClusterForge is a Terraform/OpenTofu framework for container orchestrators.

When changing this repository:

- Keep modules narrow and readable.
- Do not hide provider configuration in child modules.
- Add or update module README files with any input/output changes.
- Prefer validation in variables over clever local logic.
- Do not commit secrets, real kubeconfigs, credentials, or plain text tfvars.
- Run `./scripts/lint.sh` and `./scripts/validate.sh` when Terraform changes.
