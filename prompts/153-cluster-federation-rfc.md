## Prompt 153 — Cluster federation RFC

```text
Create cluster federation RFC.

Create:
- docs/rfcs/013-cluster-federation.md

Goal:
Evaluate whether ClusterForge should support cluster federation and multi-cluster application placement.

Cover:
1. Use cases:
   - active/active apps
   - region failover
   - tenant isolation
   - edge clusters
   - disaster recovery

2. Options:
   - GitOps multi-cluster
   - Argo CD ApplicationSet
   - Flux multi-cluster
   - Kubernetes Cluster API
   - service mesh multi-cluster
   - DNS-based failover

3. Non-goals:
   - automatic global scheduler in early versions
   - replacing cloud load balancers
   - hidden failover magic

4. Proposed initial support:
   - inventory
   - GitOps rendering
   - DNS failover docs
   - fleet health
   - no automatic cross-cluster scheduling

5. Risks:
   - complexity
   - network connectivity
   - data consistency
   - operational burden
   - secrets replication

Do not implement code.
Update roadmap.
```


---
