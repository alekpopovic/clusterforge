---
title: Security
permalink: /security/
---

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

## Recommended Kubernetes Secret Strategy

For Kubernetes workloads, prefer this flow:

1. Store secret values in a cloud secret manager such as AWS Secrets Manager or
   SSM Parameter Store.
2. Install External Secrets Operator with
   `modules/platform/kubernetes/external-secrets`.
3. Configure a `ClusterSecretStore` or `SecretStore` that references the cloud
   secret manager and uses IRSA or another workload identity mechanism.
4. Let External Secrets Operator sync values into Kubernetes Secrets.
5. Configure workload modules with `secret_env` references to Kubernetes Secret
   names and keys.

Terraform should manage references, IAM roles, service accounts, and External
Secrets manifests. It should not manage the secret values themselves. Remember
that Terraform state can contain rendered manifests and references, so state
still needs encryption and tight access controls.

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

See [security-scanning.md](security-scanning.md) for local scanner
configuration, thresholds, exclusions, and limitations.

## State

State belongs to root environments. Keep state separate by environment and
stack. Do not share state between production and non-production.

Remote state backends should use encryption, locking, and access controls.
