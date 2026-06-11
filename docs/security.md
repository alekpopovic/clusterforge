# Security

ClusterForge is designed to make infrastructure changes explicit and
reviewable. It does not remove the need for normal cloud, platform, and CI/CD
security controls.

## Secrets

Never commit:

- Cloud credentials
- kubeconfig files
- Private keys
- Terraform state files
- Real secret values in `tfvars`

Use references to external secret systems instead:

- AWS Secrets Manager
- AWS SSM Parameter Store
- Vault
- Kubernetes Secrets managed outside Terraform
- External Secrets Operator
- Cloud-native secret services

Be aware that some Terraform resources store sensitive values in state. Avoid
those patterns unless there is no practical alternative.

## Production Safety

Production operations require deliberate review:

- Production apply requires an existing plan file.
- Production destroy is blocked by default.
- The CLI does not add `--auto-approve`.
- Risk summaries count creates, updates, deletes, and replacements.
- Delete actions in production require an explicit override.

Recommended flow:

```bash
cf plan prod --out .cf/plans/prod.tfplan --risk-summary
cf policy check prod --plan-file .cf/plans/prod.tfplan
cf apply prod --plan-file .cf/plans/prod.tfplan --confirm-prod
```

If the reviewed production plan intentionally deletes resources, the apply
command must include `--allow-destroy`:

```bash
cf apply prod --plan-file .cf/plans/prod.tfplan --confirm-prod --allow-destroy
```

Production destroy requires both explicit flags:

```bash
cf destroy prod --allow-destroy --confirm-prod
```

## CI Scans

GitHub Actions run:

- Terraform formatting and validation
- CLI tests and build
- Checkov IaC scan
- Trivy config scan

These scans do not require real cloud credentials.

## State

State belongs to root environments. Keep state separate by environment and
stack. Do not share state between production and non-production.

Remote state backends should use encryption, locking, and access controls.
