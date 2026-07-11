# AWS ECS incident

## Severity classification
SEV-1 for broad production outage/compromise; SEV-2 for unhealthy service or failed deployment; SEV-3 for limited task degradation.

## Symptoms
ECS service unhealthy, tasks stop, image pulls fail, ALB targets fail health checks, deployment circuit breaker activates.

## Initial checks
Read ECS service events, stopped-task reasons, task definition/image digest, ALB target health, CloudWatch logs and capacity/provider status.

## Containment
Pause pipeline deployment; reduce traffic through approved routing only; revoke compromised task roles/secrets.

## Diagnosis
Check CPU/memory, port mappings, health paths, security groups, subnets/NAT/endpoints, ECR permissions, execution/task roles and secret references.

## Remediation
Register/deploy a reviewed task definition or repair dependency/IAM/network configuration; verify steady state and application probes.

## Rollback
Redeploy the last known task definition. **High-risk:** scaling to zero, ALB rule changes or service replacement require approval.

## Communication notes
Report cluster/service, failed deployment ID, task/target counts, endpoints affected and mitigation status.

## Evidence collection
Preserve service events, stopped reasons, task definition revision, image digest, target health, logs and CloudTrail references.

## Postmortem checklist
Timeline; task/app/platform cause; detection/rollback gaps; load/capacity test; owners and deadlines.
