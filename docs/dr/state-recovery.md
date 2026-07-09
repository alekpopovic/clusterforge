# Terraform State Recovery

## Scope

Recovery guidance for lost, corrupted, or inaccessible Terraform/OpenTofu state
used by ClusterForge roots.

## Assumptions

- State is remote for production.
- State backend access is audited.
- Recovery is manual and must be reviewed.

## Prerequisites

- Backend bucket/table access.
- Terraform/OpenTofu and `cf`.
- Current repository checkout.
- Cloud read access for imported resources.

## Recovery Steps

1. Stop all applies for the affected environment.
2. Identify whether state is lost, corrupted, locked, or only inaccessible.
3. Restore the backend object from versioning or backup when available.
4. If state cannot be restored, rebuild it with `terraform import` resource by resource.
5. Run `terraform plan -refresh-only` and review drift before any apply.
6. Commit any import documentation or root fixes.

## Validation Steps

- `terraform state list` succeeds.
- `terraform plan` shows only expected drift.
- Backend locking works.
- A second operator can read the restored state.

## Rollback Steps

- Restore the previous state object version.
- Keep a copy of corrupted state for audit, but do not commit it.
- Revert root changes that were made only for recovery.

## Data Loss Risks

- Incorrect imports can bind Terraform to the wrong resource.
- Corrupted state edits can cause destructive plans.
- Missing state for data-bearing resources can lead to accidental recreation.

## Estimated Downtime Categories

- Backend access issue: minutes to hours.
- Versioned state restore: minutes.
- Full import rebuild: hours to days.

## Required Access

- State backend read/write.
- Lock table access.
- Cloud read permissions.
- Repository write access for recovery notes.

## Common Failure Modes

- Lost Terraform state.
- Corrupted state.
- Stale lock records.
- Backend KMS permission failure.
- Human import error.
