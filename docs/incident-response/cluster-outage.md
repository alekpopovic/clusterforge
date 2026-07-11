# Cluster outage

## Severity classification
SEV-1 for complete production cluster outage or control-plane compromise; SEV-2 for major capacity/AZ degradation; SEV-3 for non-production.

## Symptoms
API unavailable, scheduler cannot place workloads, most nodes NotReady, ingress down, or core DNS/network/storage failure.

## Initial checks
Confirm external monitoring, cloud/provider status, API reachability, node/control-plane health, recent changes and regional dependencies.

## Containment
Declare change freeze, protect credentials/state/evidence and redirect traffic only through an approved tested DR procedure.

## Diagnosis
Separate control plane, node capacity, network/CNI/DNS, identity, storage, quotas, certificates and upstream cloud failure.

## Remediation
Use provider-supported recovery and reviewed infrastructure plans; restore core services in dependency order, then workloads and traffic.

## Rollback
Return traffic to the last verified healthy target when available. **Destructive:** cluster recreation, state restore, DNS failover and data restore require incident-command approval.

## Communication notes
Maintain scheduled updates with region/cluster scope, customer impact, recovery stage, data risk and next checkpoint.

## Evidence collection
Preserve provider/audit logs, events, metrics, plans, DNS changes, backup/restore IDs, versions and cleanup proof.

## Postmortem checklist
Timeline; failure domains; RTO/RPO actuals; capacity/DR gaps; restore evidence; action owners/dates.
