# DNS Failover

## Scope

Recovery guidance for DNS misconfiguration and traffic failover for
ClusterForge-managed platforms.

## Assumptions

- DNS records are managed through reviewed Terraform or an approved DNS process.
- Low TTLs are configured before planned failover.
- DR does not provide fake RTO or RPO guarantees.

## Prerequisites

- Route53 or DNS provider access.
- Load balancer or ingress endpoint information.
- Application smoke tests.
- Communication channel for customer-impacting changes.

## Recovery Steps

1. Confirm the current DNS record, TTL, and target.
2. Validate the recovery endpoint before changing DNS.
3. Update records through Terraform when time allows.
4. For urgent incidents, make an approved manual change and record it for later reconciliation.
5. Monitor propagation and application health.

## Validation Steps

- `dig` or equivalent resolves to the expected target.
- TLS certificate matches the hostname.
- Application health checks pass from multiple networks.
- Terraform state is reconciled after manual changes.

## Rollback Steps

- Restore the previous DNS target.
- Revert Terraform DNS changes with a reviewed plan.
- Keep traffic on the old endpoint until the new endpoint is healthy.

## Data Loss Risks

- DNS failover can split traffic between old and new systems.
- Stateful apps may accept writes in multiple places if not fenced.
- Cached DNS can extend impact beyond the TTL.

## Estimated Downtime Categories

- Correcting a typo: minutes plus cache delay.
- Regional failover: depends on warm standby readiness.
- Certificate or ingress repair: minutes to hours.

## Required Access

- DNS provider write access.
- Terraform backend access.
- Load balancer or ingress read access.
- Certificate manager access.

## Common Failure Modes

- DNS points at the wrong load balancer.
- Broken ingress.
- Missing or expired certificate.
- External-dns overwrites manual change.
- Region-level disaster planning was not tested.
