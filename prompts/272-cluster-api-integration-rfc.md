# Prompt 272 — Cluster API integration RFC

```text
Create Cluster API integration RFC.

Goal:
Evaluate whether ClusterForge should support Kubernetes Cluster API as a cluster lifecycle backend.

Create:
- docs/rfcs/035-cluster-api-integration.md
- docs/cluster-api.md

Cover:

1. Goals
   - Kubernetes-native cluster lifecycle
   - self-service cluster templates
   - multi-cloud cluster creation
   - GitOps-friendly cluster management
   - standardized cluster blueprints

2. Non-goals
   - replacing Terraform modules immediately
   - supporting every Cluster API provider
   - making Cluster API mandatory

3. Target providers
   - AWS provider
   - Azure provider
   - GCP provider
   - Docker/local provider for testing
   - others later

4. Architecture options
   - Terraform provisions management cluster
   - Cluster API manages workload clusters
   - ClusterForge renders Cluster API manifests
   - GitOps applies manifests
   - Control Plane tracks inventory/status

5. Terraform vs Cluster API boundary
   - Terraform for foundational cloud resources
   - Cluster API for cluster lifecycle
   - GitOps for reconciliation
   - Control Plane for visibility/governance

6. CLI future
   - cf cluster-template list
   - cf cluster create --template eks-small
   - cf cluster delete
   - cf cluster status

7. Risks
   - management cluster dependency
   - provider maturity differences
   - IAM complexity
   - debugging complexity
   - lifecycle ownership conflicts

Do not implement code.
Update roadmap.
```
