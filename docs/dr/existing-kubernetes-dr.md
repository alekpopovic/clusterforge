# Existing Kubernetes Disaster Recovery

## Scope

Recovery guidance for Kubernetes clusters not created by ClusterForge but used
as ClusterForge workload or platform targets.

## Assumptions

- Cluster lifecycle is owned outside ClusterForge.
- ClusterForge manages selected platform add-ons and workloads.
- Restore procedures are tested against this specific cluster type.

## Prerequisites

- Cluster admin access.
- Access to the external cluster provider or owner.
- `kubectl`, `helm`, Terraform/OpenTofu, `cf`, and backup tooling.
- Backup storage and secret manager access.

## Recovery Steps

1. Confirm whether the cluster itself or only ClusterForge-managed resources are affected.
2. Coordinate cluster-level recovery with the owning team.
3. Reapply ClusterForge-managed platform modules after the cluster is healthy.
4. Restore deleted namespaces or resources from Velero or another backup source.
5. Reconcile external secrets and ingress/DNS integrations.

## Validation Steps

- Cluster API is reachable.
- Required namespaces exist.
- Platform controllers are healthy.
- Workloads are ready.
- Ingress and DNS resolve to expected endpoints.

## Rollback Steps

- Revert the last ClusterForge platform or workload change.
- Restore the previous namespace backup.
- Ask the cluster owner to roll back cluster-level changes.

## Data Loss Risks

- ClusterForge may not control the underlying backup system.
- Namespace restore may miss external databases.
- Secret manager outage may prevent full workload recovery.

## Estimated Downtime Categories

- Namespace restore: minutes to hours.
- Cluster-owner repair: unknown until provider confirms impact.
- Region-level or provider-level disaster: depends on external DR plan.

## Required Access

- Kubernetes admin or delegated namespace admin.
- Backup storage access.
- Terraform backend access for ClusterForge roots.
- External provider support channel.

## Common Failure Modes

- Kubeconfig or context points to the wrong cluster.
- CRDs missing before restore.
- Ingress class mismatch.
- Backup tool not installed or not tested.
- Secret references point to unavailable external systems.
