# Operations

ClusterForge operations are deliberate, reviewable workflows around generated
Terraform/OpenTofu. The CLI helps with planning and safety checks, but it does
not make production recovery automatic.

## Core References

- [State](state.md)
- [Backup and restore](backup-restore.md)
- [Security](security.md)
- [DR runbooks](dr/state-recovery.md)

## Recovery Principles

- Stop unsafe writes before restoring stateful systems.
- Prefer reviewed Terraform plans over console edits.
- Record manual emergency changes and reconcile them back into Terraform.
- Test restore procedures before relying on them.
- Do not publish RTO/RPO guarantees unless they were agreed and measured.
