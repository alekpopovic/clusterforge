# Cluster Restore

## Scope

Generic cluster restore sequence for Kubernetes clusters managed or targeted by
ClusterForge.

## Assumptions

- Infrastructure and workload code are in version control.
- Backups exist and have been tested.
- Restore is coordinated through an incident process.

## Prerequisites

- Terraform/OpenTofu, `cf`, `kubectl`, `helm`, and backup tooling.
- Access to cloud credentials, state backend, backup storage, and secrets.
- Approved recovery target cluster or environment.

## Recovery Steps

1. Recreate or repair network, IAM, and cluster infrastructure.
2. Install CRDs and platform controllers before restoring dependent resources.
3. Restore namespaces in dependency order.
4. Restore workloads and verify secrets are present.
5. Restore ingress, DNS, and certificates.
6. Resume traffic only after smoke tests pass.

## Validation Steps

- Nodes and system pods are ready.
- Platform add-ons are healthy.
- Restored namespaces contain expected resources.
- Application smoke tests pass.
- Metrics and logs confirm normal operation.

## Rollback Steps

- Stop traffic to the restored cluster.
- Delete partially restored namespaces if they are unsafe.
- Return traffic to the previous healthy cluster when available.

## Data Loss Risks

- Application data may require database-native restore.
- Backups may be stale.
- Restored resources may reference missing cloud resources.

## Estimated Downtime Categories

- Platform-only restore: minutes to hours.
- Full cluster restore: hours.
- Stateful restore: hours to days.

## Required Access

- Cloud operator role.
- Kubernetes admin.
- Backup storage and KMS access.
- Secret manager access.

## Common Failure Modes

- CRDs restored after custom resources.
- Broken ingress.
- Missing storage classes.
- Incompatible Kubernetes versions.
- Secret manager outage.
