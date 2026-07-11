# Terraform state incident

## Severity classification
SEV-1 for leaked/corrupted production state or accidental destroy; SEV-2 for lock/access failure or failed apply midway; SEV-3 for non-production issue.

## Symptoms
Apply failed midway, state lock stuck, resources diverge, state missing/corrupt, or accidental resource destroy.

## Initial checks
Stop all writers; record backend/workspace, run read-only state list/pull to a protected location, inspect cloud resources/audit logs and the saved plan.

## Containment
Disable pipelines and restrict backend writes. For exposure, treat every state secret as compromised and follow the secret runbook.

## Diagnosis
Compare state, config, plan and real inventory; identify the last successful state version and partially completed provider operations.

## Remediation
Prefer refresh-only assessment and a new reviewed plan. Restore a backend version only after peer review and protected backup.

## Rollback
Infrastructure rollback is a new reviewed plan, not blind Git revert. **Destructive:** state rm/mv/import, force-unlock, backend restore and resource recreation require explicit approval and backups.

## Communication notes
Report impacted workspace/resources, confidentiality/integrity/availability impact and writer freeze status without attaching state.

## Evidence collection
Retain encrypted state versions, checksums, plan/run IDs, provider logs, audit events, commit SHA and resource inventory.

## Postmortem checklist
Timeline; backend/provider/operator cause; rotation/recovery; locking/versioning gaps; restore test; owners/dates.
