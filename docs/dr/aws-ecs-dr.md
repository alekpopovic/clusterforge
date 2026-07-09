# AWS ECS Disaster Recovery

## Scope

Recovery guidance for ClusterForge-managed AWS ECS clusters, services, load
balancers, task definitions, and supporting AWS resources.

## Assumptions

- Terraform state and ClusterForge configuration are available.
- Images are stored in a registry with immutable version tags or digests.
- DR is manual unless a separate tested automation exists.

## Prerequisites

- AWS access for ECS, IAM, EC2, ELB, CloudWatch, ECR, Route53, KMS, and S3.
- Terraform/OpenTofu and `cf`.
- Access to ECR repositories and application secret stores.

## Recovery Steps

1. Stop unsafe deploys and identify the last known-good task definition.
2. Check ECS service events, target group health, and CloudWatch logs.
3. Repair infrastructure through Terraform with a reviewed plan.
4. Roll back the ECS service to a known-good image or task definition.
5. Recreate missing services or clusters from ClusterForge roots.
6. Restore DNS or load balancer routing after health checks pass.

## Validation Steps

- ECS cluster is active.
- Services reach desired count.
- Target groups are healthy.
- Application endpoints pass smoke tests.
- Logs show normal startup and no secret resolution failures.

## Rollback Steps

- Revert to the previous task definition.
- Shift traffic back to the previous target group or listener rule.
- Revert Terraform changes with a reviewed plan.

## Data Loss Risks

- ECS rollback does not restore application databases.
- Queue consumers may replay or drop work depending on application behavior.
- Missing secrets can cause repeated task launch failures.

## Estimated Downtime Categories

- Service rollback: minutes.
- Cluster or load balancer repair: minutes to hours.
- Region-level recovery: depends on pre-built capacity and tested DNS failover.

## Required Access

- AWS operator role.
- Terraform backend access.
- ECR read access.
- Secret manager read access.

## Common Failure Modes

- Bad image pushed to mutable tag.
- Missing task execution role permissions.
- Broken target group health checks.
- DNS points at the wrong load balancer.
- Secret manager or KMS outage.
