---
title: Failed deployment
category: deployment
severity: high
tags: [terraform, kubernetes, ecs, argocd, incident]
---

# Failed deployment

## Severity classification
SEV-1 if deployment causes broad production outage/data risk; SEV-2 for major degraded release; SEV-3 for blocked/non-production deployment.

## Symptoms
Terraform apply fails midway, Kubernetes rollout stalls, Argo CD sync fails, ECS deployment unhealthy, or new image cannot pull/start.

## Initial checks
Freeze further releases; identify commit, plan/run, image digest, environment and first failing event without rerunning blindly.

## Containment
Stop automatic promotion/sync where approved; keep healthy old replicas/targets serving; revoke suspect artifact/credential access.

## Diagnosis
Compare desired/live state; inspect provider error, rollout/events, Argo diff, ECS events, image registry and dependencies.

## Remediation
Correct the smallest configuration/artifact/dependency issue, generate a fresh plan/diff and use normal reviewed deployment controls.

## Rollback
Redeploy the last known immutable artifact/config. **Destructive:** Terraform resource replacement, database rollback or forced GitOps deletion requires explicit approval and backup evidence.

## Communication notes
Report release ID, blast radius, customer symptoms, containment, next decision point and rollback readiness.

## Evidence collection
Preserve logs, plan/diff, task/pod events, image digest/SBOM, commit, approvals and timing.

## Postmortem checklist
Change/failure timeline; test and canary gaps; rollback time; ownership; corrective tests and deadlines.
