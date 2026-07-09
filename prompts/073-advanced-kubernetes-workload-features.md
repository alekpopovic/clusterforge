## Prompt 73 — Advanced Kubernetes workload features

```text
Add advanced Kubernetes workload options.

Target modules:
- workloads/kubernetes/app
- workloads/kubernetes/worker

Features:
1. init_containers
2. sidecars
3. volumes
4. volume_mounts
5. node_selector
6. tolerations
7. affinity
8. topology_spread_constraints
9. pod_disruption_budget
10. priority_class_name

Inputs should be typed but not overly complex.

Rules:
- Keep defaults simple.
- Avoid huge unreadable dynamic blocks where possible.
- If a feature becomes too complex, document a helm-app or GitOps alternative.
- Preserve backward compatibility.

Add examples:
- examples/kubernetes-app-advanced
- examples/kubernetes-worker-with-sidecar

README:
- Document each feature briefly.
- Explain when to stop using generic module and switch to Helm/Kustomize/GitOps.

Run:
- terraform fmt -recursive
```

---
