# DNS incident

## Severity classification
SEV-1 for hijack or broad production misrouting; SEV-2 for major resolution/certificate impact; SEV-3 for partial/staging issue.

## Symptoms
DNS changed incorrectly, NXDOMAIN, wrong endpoint, resolution split, external-dns conflict or certificate DNS challenge failure.

## Initial checks
Query authoritative and multiple recursive resolvers; inspect zone delegation, records/TTL, external-dns ownership, change/audit history and target health.

## Containment
Pause DNS automation and protect the zone/account. If compromise is suspected, revoke sessions/roles and lock registrar changes.

## Diagnosis
Distinguish delegation, propagation/cache, record, health evaluation, load balancer, external-dns and certificate challenge issues.

## Remediation
Apply the smallest reviewed record/delegation correction and verify authoritative answers before traffic validation.

## Rollback
Restore the last known record set with reviewed TTL. **High-risk:** delegation, hosted-zone deletion and broad wildcard changes require explicit approval.

## Communication notes
Report affected names/record types, resolver regions, TTL/propagation expectations and traffic impact without private zone data.

## Evidence collection
Capture dig output, zone change IDs, audit logs, old/new record sets, ownership TXT records and timestamps.

## Postmortem checklist
Timeline; automation/human/provider cause; TTL/monitoring gaps; registrar/zone protections; owners/dates.
