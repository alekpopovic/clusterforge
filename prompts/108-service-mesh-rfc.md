## Prompt 108 — Service mesh RFC

```text
Create a service mesh RFC for ClusterForge.

Create:
- docs/rfcs/008-service-mesh.md

Evaluate:
- Istio
- Linkerd
- Consul service mesh

Cover:
1. Goals:
   - mTLS
   - traffic shifting
   - observability
   - policy
   - service-to-service security

2. Non-goals for first implementation:
   - multi-cluster mesh
   - advanced zero-trust automation
   - enterprise control planes

3. Proposed modules:
   - modules/platform/kubernetes/istio
   - modules/platform/kubernetes/linkerd
   - modules/platform/kubernetes/consul-service-mesh

4. Workload impact:
   - sidecar injection labels
   - namespace configuration
   - ingress gateway
   - telemetry

5. Operational risks:
   - complexity
   - upgrades
   - debugging
   - resource overhead
   - app compatibility

6. Recommendation:
   - choose one first implementation
   - keep mesh optional
   - do not enable by default

Do not implement code yet.
Update docs/roadmap.md.
```

---
