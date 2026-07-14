# Prompt 217 — Rollback workflow design

```text
Create rollback workflow RFC.

Create:
- docs/rfcs/022-rollback-workflow.md
- docs/control-plane/rollback.md

Goal:
Define how ClusterForge should support safe rollback without pretending rollback is always trivial.

Cover:
1. Terraform rollback reality
   - Terraform does not have universal automatic rollback
   - rollback often means applying a previous desired state
   - state and external data may not be reversible

2. Kubernetes app rollback
   - GitOps revert
   - image tag rollback
   - Deployment rollout undo
   - Argo Rollouts rollback

3. ECS rollback
   - previous task definition
   - CodeDeploy rollback
   - ALB target group rollback

4. Infrastructure rollback
   - revert Git commit
   - plan carefully
   - backup/restore for data resources
   - manual intervention for destructive changes

5. Proposed ClusterForge support
   - record plan/apply history
   - record app image versions
   - generate rollback plan from previous manifest
   - link runbooks
   - require approval
   - never auto-rollback infra by default

6. Future CLI:
   - cf rollback plan
   - cf rollback app
   - cf rollback history
   - cf rollback runbook

Do not implement rollback execution yet.
Update roadmap.
```
