# Application Restore

## Scope

Application-level recovery for workloads rendered or managed through
ClusterForge.

## Assumptions

- Image references are versioned tags or digests.
- Application data restore is documented by the application owner.
- ClusterForge does not provide automatic application rollback.

## Prerequisites

- Access to app manifests and generated Terraform.
- Registry access.
- Kubernetes or ECS operator access.
- Application backup and secret access.

## Recovery Steps

1. Identify the last known-good image and configuration.
2. Re-render or reapply the workload module if Terraform-managed.
3. Roll back image tags or digests deliberately.
4. Restore application data with app-specific tools.
5. Recreate secrets from the approved secret manager.
6. Validate health checks and user flows.

## Validation Steps

- Deployment or service reaches ready state.
- Logs show successful startup.
- App endpoint passes smoke tests.
- Data integrity checks pass.

## Rollback Steps

- Revert to previous manifest or task definition.
- Stop traffic to failed version.
- Restore previous data snapshot when needed.

## Data Loss Risks

- Backups may not include recent writes.
- Queue or stream consumers may replay messages.
- Schema migrations may not be reversible.

## Estimated Downtime Categories

- Stateless rollback: minutes.
- Stateful restore: hours.
- Cross-region recovery: depends on application design.

## Required Access

- Cluster or ECS service access.
- Registry read access.
- Secret manager read access.
- Application database restore access.

## Common Failure Modes

- `latest` tag points to unexpected image.
- Missing secret keys.
- Database migration mismatch.
- Broken ingress or target group.
- External dependency outage.
