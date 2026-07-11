## Prompt 154 — Edge deployment support RFC

```text
Create edge deployment support RFC.

Create:
- docs/rfcs/014-edge-deployments.md

Goal:
Evaluate ClusterForge support for small edge clusters and constrained environments.

Targets:
- K3s
- RKE2
- MicroK8s if considered
- single-node Kubernetes
- remote Docker host experimental

Cover:
1. Use cases:
   - retail locations
   - IoT gateways
   - on-prem mini clusters
   - disconnected environments
   - lab environments

2. Constraints:
   - limited CPU/memory
   - unreliable connectivity
   - offline upgrades
   - local storage
   - simplified observability
   - local registry

3. Proposed modules:
   - local registry mirror
   - lightweight observability
   - edge app workload profile
   - backup to remote or local target
   - GitOps pull model

4. CLI:
   - cf edge init
   - cf edge bundle
   - cf edge status

5. Security:
   - secrets distribution
   - local credentials
   - device identity
   - update signing

Do not implement code yet.
Update roadmap.
```


---
