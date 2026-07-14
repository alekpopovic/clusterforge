# Prompt 231 — High availability deployment architecture

```text
Create high availability architecture for ClusterForge Control Plane.

Create:
- docs/control-plane/high-availability.md
- docs/rfcs/024-control-plane-ha.md

Cover:
1. API server replicas
2. database HA
3. artifact storage HA
4. runner pools
5. dashboard static hosting
6. load balancer/ingress
7. session handling
8. migrations
9. zero-downtime upgrades
10. backup/restore
11. disaster recovery

Kubernetes Helm chart updates:
- replicaCount > 1 support
- readiness/liveness probes
- PodDisruptionBudget
- topology spread constraints optional
- external PostgreSQL required for HA
- external artifact storage recommended
- serviceMonitor optional

Rules:
- Do not claim HA if using embedded/local DB.
- Clearly distinguish dev vs production deployment.
- No real cloud credentials.

Implementation:
- Update Helm chart values and templates where practical.
- Add examples:
  - examples/control-plane-ha-values.yaml

Run:
- helm lint if available
```
