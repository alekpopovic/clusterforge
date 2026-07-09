## Prompt 95 — Disaster recovery runbooks

```text
Create disaster recovery runbooks for ClusterForge-managed platforms.

Create:
- docs/dr/
  aws-eks-dr.md
  aws-ecs-dr.md
  existing-kubernetes-dr.md
  state-recovery.md
  cluster-restore.md
  app-restore.md
  dns-failover.md

Each runbook must include:
- Scope
- Assumptions
- Prerequisites
- Recovery steps
- Validation steps
- Rollback steps
- Data loss risks
- Estimated downtime categories
- Required access
- Common failure modes

Specific topics:
1. Lost Terraform state
2. Corrupted state
3. Deleted Kubernetes namespace
4. Broken ingress
5. EKS cluster recreation
6. ECS service rollback
7. DNS misconfiguration
8. Secret manager outage
9. Backup restore from Velero
10. Region-level disaster planning

Rules:
- Do not imply DR is automatic.
- Do not include fake RTO/RPO guarantees.
- State that restore procedures must be tested.
- Avoid real account IDs/domains.

Update:
- docs/security.md
- docs/operations.md if it exists
- README with DR documentation link
```

---
