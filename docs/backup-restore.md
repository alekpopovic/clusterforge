# Backup And Restore

ClusterForge provides Velero building blocks, but backup installation is not a
restore guarantee. Restore procedures must be tested before a platform is
treated as recoverable.

## Scope

- `modules/platform/kubernetes/velero` installs Velero with Helm.
- `modules/cloud/aws/velero-backup-bucket` creates an encrypted S3 bucket for
  backup storage.
- Cloud IAM roles, bucket policies, and restore approvals remain environment
  responsibilities.

## Backup Runbook

1. Confirm Velero is installed:

   ```bash
   kubectl -n velero get deploy
   velero version
   ```

2. Create a test namespace and workload.
3. Run a backup:

   ```bash
   velero backup create smoke-test --include-namespaces smoke-test
   ```

4. Confirm the backup completed:

   ```bash
   velero backup describe smoke-test
   velero backup logs smoke-test
   ```

5. Confirm backup objects exist in the configured S3 bucket.

## Restore Runbook

1. Choose a tested backup.
2. Restore into a safe namespace first when possible:

   ```bash
   velero restore create smoke-test-restore --from-backup smoke-test
   ```

3. Validate Kubernetes objects, pods, services, ingress, and application-level
   health checks.
4. Confirm data-bearing workloads restored the expected data. Velero object
   restore does not automatically prove application consistency.

## Disaster Recovery Limitations

- Velero cannot recover data that was never backed up.
- Failed cloud IAM permissions can block backup or restore.
- Storage provider outages can block restore.
- Cluster-scoped resources and CRDs need careful ordering.
- Some stateful applications require database-native backup and restore in
  addition to Kubernetes object restore.

## Testing Restore

At minimum, test a namespace restore after every material backup configuration
change. For production, test restore into an isolated cluster or namespace on a
regular schedule and retain evidence of the result.
