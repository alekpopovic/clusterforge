# Incident response

Use these runbooks for ClusterForge-managed platforms. They are decision aids, not availability or recovery guarantees. Never paste credentials, state, kubeconfigs, tokens, customer data, or private endpoints into tickets.

Severity: **SEV-1** active security breach or broad production outage/data loss; **SEV-2** major degraded production service; **SEV-3** limited/non-production impact; **SEV-4** advisory/no active impact.

Runbooks: Kubernetes, AWS EKS, AWS ECS, Terraform state, secret leak, DNS, failed deployment, and cluster outage. Assign incident commander, operations lead, communications lead and scribe. Prefer read-only evidence collection before mutation.

## Initial checks

Confirm scope, start time, recent deploys/drift, environment identity, monitoring, audit logs and on-call ownership. Freeze unrelated changes.

## Containment

Revoke compromised access, pause pipelines or traffic changes as appropriate, without destroying evidence.

## Diagnosis

Build a timestamped hypothesis log and distinguish control-plane, workload, network, identity, DNS, state and dependency failure.

## Remediation and rollback

Use reviewed plans and provider-specific procedures. Mark destroy, state surgery, credential rotation and DNS cutover as destructive/high-risk actions requiring approval.

## Communication and evidence

Record severity, customer impact, decisions, command outputs with secrets redacted, plan IDs, commit SHAs, cloud/Kubernetes audit references and cleanup status.

## Postmortem checklist

Timeline; root/contributing causes; detection gaps; containment quality; recovery evidence; action owners/dates; runbook/test changes; secret rotation and artifact retention.
