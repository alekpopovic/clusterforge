# Secret rotation workflow

ClusterForge discovers references and prepares a rotation plan, but never reads,
prompts for, writes, or rotates secret values. Rotation remains an operator-owned
change in the external secret system.

```bash
cf secrets check prod
cf secrets references prod
cf secrets references prod --app api
cf secrets rotation-plan prod
```

Discovery covers app `secret_env` references, Terraform files, External Secrets
manifests, and ECS `value_from` references under the environment path. Results are
identifiers and file locations only. Review them because static discovery cannot
prove that every runtime consumer has been found.

## Provider workflows

1. For AWS Secrets Manager, use managed rotation where supported or create a new
   version through an approved operator workflow. Check staging labels and do not
   remove the previous version until consumers are healthy.
2. For SSM Parameter Store, update the SecureString with your approved tooling,
   preserving KMS and IAM controls. A changed parameter does not automatically
   restart a workload.
3. For External Secrets Operator, verify provider authentication, refresh policy,
   reconciliation status, and the resulting Kubernetes Secret metadata. Never put
   the remote value in the manifest.
4. Refresh Kubernetes consumers according to application behavior. Environment
   variables require new pods; mounted secret volumes may update asynchronously,
   but applications might still need a reload.
5. Restart pods with a controlled rollout, health gates, disruption budgets, and
   rollback awareness. A rollback to pods expecting an invalidated credential can
   deepen the outage.
6. For ECS, force a new service deployment or register/deploy the reviewed task
   definition so replacement tasks resolve the new value. Watch deployment circuit
   breakers and target health before stopping old tasks.
7. For an RDS-managed master password, use the RDS/Secrets Manager managed rotation
   path. Test application connection pooling and replicas; do not copy the master
   password into Terraform variables or outputs.

## Verification, rollback, and audit

Record the reference identifier, owner, approval, maintenance window, rotation
event/version identifier, affected workloads, rollout evidence, and revocation
time in the audit system without recording the value. Validate authentication and
application health before revoking the old version.

Secret rollback is not equivalent to application rollback: an old deployment may
reference a revoked credential, while restoring an old credential may reopen an
exposure. Prefer completing the forward rotation. If rollback is unavoidable,
require explicit security approval, tightly limit validity, and rotate again after
recovery.
