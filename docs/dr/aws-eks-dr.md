# AWS EKS Disaster Recovery

## Scope

Recovery guidance for ClusterForge-managed AWS EKS clusters, platform add-ons,
and workloads.

## Assumptions

- Terraform state is remote, encrypted, and access-controlled.
- Cluster manifests and generated Terraform are in version control.
- Velero or another backup tool has been configured and tested.
- DR is not automatic; operators execute and validate each step.

## Prerequisites

- AWS access for VPC, IAM, EKS, EC2, KMS, CloudWatch, S3, and Route53 when DNS is involved.
- Terraform/OpenTofu and the `cf` CLI.
- `kubectl`, `helm`, and Velero CLI when restoring Kubernetes resources.
- Access to backup buckets and secret managers.

## Recovery Steps

1. Freeze non-essential changes and identify the failure mode.
2. Check Terraform state health with `terraform state list` and backend access.
3. Recreate missing infrastructure with a reviewed plan file, not ad hoc console edits.
4. Recreate or repair the EKS cluster from the current ClusterForge root.
5. Restore platform add-ons, then namespaces, then workloads.
6. Restore from Velero only after confirming the target cluster and CRDs are ready.
7. Reconcile secrets from the approved secret manager.
8. Validate ingress and DNS before sending production traffic.

## Validation Steps

- `aws eks describe-cluster` returns active cluster status.
- Managed node groups are ready.
- `kubectl get nodes` shows ready nodes.
- CoreDNS, ingress, external-dns, cert-manager, and observability add-ons are healthy.
- Application health checks pass.
- Backup restore evidence is attached to the incident record.

## Rollback Steps

- Stop traffic with DNS or load balancer controls.
- Revert the last Terraform change with a reviewed plan.
- Restore the previous known-good application version.
- If restore introduced bad data, stop writes and use application-specific restore procedures.

## Data Loss Risks

- Velero may not capture application-consistent database state.
- Deleted backup objects cannot be recovered after retention expires.
- Secret manager outage may block workload startup.
- Region-level events may make backup storage unavailable if not replicated.

## Estimated Downtime Categories

- Control plane repair: minutes to hours.
- Cluster recreation: hours.
- Stateful application restore: hours to days depending on data size.
- Region-level disaster: depends on pre-tested multi-region design.

## Required Access

- AWS break-glass role.
- Terraform backend access.
- Kubernetes cluster admin.
- Backup bucket and KMS decrypt access.
- Secret manager read access.

## Common Failure Modes

- Lost or corrupted Terraform state.
- Deleted Kubernetes namespace.
- Broken ingress controller or certificate chain.
- EKS node group failure.
- Velero restore blocked by missing CRDs.
- DNS misconfiguration.
- Secret manager outage.
