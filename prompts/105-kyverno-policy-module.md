## Prompt 105 — Kyverno policy module

```text
Add Kyverno platform policy support.

Create module:
- modules/platform/kubernetes/kyverno

Purpose:
Install Kyverno via Helm and optionally apply baseline ClusterPolicies.

Inputs:
- namespace default "kyverno"
- create_namespace default true
- chart_version default ""
- values list(string)
- enable_baseline_policies default false
- policies map(string) default {}

Behavior:
- Install Kyverno Helm chart.
- If enable_baseline_policies=true, create basic policies:
  - disallow privileged containers
  - require resource requests/limits advisory or enforce configurable
  - disallow latest tag configurable
- Allow user-provided policy YAML via kubernetes_manifest if practical.

Outputs:
- namespace
- release_name

Docs:
- docs/kubernetes-policy-kyverno.md
- Explain enforce vs audit.
- Explain production rollout strategy.
- Explain how to avoid breaking existing workloads.

Example:
- examples/kubernetes-kyverno-baseline

Rules:
- Default must not break workloads.
- Policies should start in audit mode unless explicitly enforced.
- Document CRD dependency.

Run:
- terraform fmt -recursive
```

---
