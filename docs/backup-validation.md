# Backup validation and restore testing

A successful backup job is not proof that data can be restored. ClusterForge
provides read-only local checks and a plan for recording restore-test evidence;
it never starts a restore, deletes a backup, or calls a cloud API.

```bash
cf backup check prod
cf backup plan prod
cf backup report prod
```

Checks look for detectable Velero and AWS backup-bucket Terraform configuration,
the expected Velero namespace, the restore runbook, and `backup-evidence.yaml`.
Static detection can miss indirect modules and cannot prove that a remote backup
is complete, encrypted, retained, or restorable.

## Evidence file

Keep non-secret test facts at the project root:

```yaml
backup_tests:
  prod:
    last_backup_test: "2026-07-01"
    last_restore_test: "2026-07-02"
    result: passed
    notes: "Restored namespace demo-restore"
```

Do not include backup contents, credentials, personal data, signed URLs, or
kubeconfigs. Reference protected CI/audit artifacts when stronger evidence is
needed. `cf backup report prod` emits JSON suitable for reviewed evidence
collection, but is not by itself proof of recovery.

## Restore-test environment and schedule

Restore into a dedicated, isolated non-production cluster or namespace with
restricted network access and synthetic or appropriately protected data. Prevent
connections to production queues, webhooks, email, payments, DNS, and databases.
Never run an automated restore test directly in production.

Test on a risk-based schedule and after material storage, encryption, Kubernetes,
Velero, application schema, or retention changes. A common starting point is a
monthly backup integrity review and quarterly restore exercise, tightened to the
documented RPO/RTO and regulatory needs. Measure recovery time and data age,
verify application behavior, clean up the isolated target through a reviewed
procedure, and retain the source backup.
