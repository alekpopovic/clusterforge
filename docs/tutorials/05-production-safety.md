# Tutorial 05: Production Safety

## Goal

Learn the production safety gates without applying real production changes.

## Prerequisites

- ClusterForge `cf`
- `terraform` or `tofu`

## Commands

```bash
cf policy list
cf policy check prod --pack production
cf plan prod --out .cf/plans/prod.tfplan --risk-summary
cf apply prod --plan-file .cf/plans/prod.tfplan --confirm-prod
```

Do not run the apply command against a real production environment unless the
plan has been reviewed and approved.

## Generated Files

- saved plan file under `.cf/plans/`

## What Terraform Will Create

Whatever the reviewed production plan states. ClusterForge does not hide the
Terraform plan.

## Validate

```bash
cf doctor
cf drift check prod
```

## Cleanup

Production cleanup must use an environment-specific runbook. `cf destroy prod`
is blocked by default unless explicit flags and approvals are used.

## Troubleshooting

Check backend locking, plan-file paths, policy pack warnings, and whether the
environment is named `prod` or `production`.
