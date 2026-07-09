## Prompt 107 — Progressive delivery with Argo Rollouts

```text
Add progressive delivery support using Argo Rollouts.

Create module:
- modules/platform/kubernetes/argo-rollouts

Purpose:
Install Argo Rollouts via Helm.

Inputs:
- namespace default "argo-rollouts"
- create_namespace default true
- chart_version default ""
- values list(string)
- enable_dashboard_ingress default false
- dashboard_host default ""

Outputs:
- namespace
- release_name

Create workload module:
- modules/workloads/kubernetes/rollout-app

Purpose:
Deploy an Argo Rollout instead of a standard Deployment.

Inputs:
- similar to workloads/kubernetes/app
- strategy:
    type: canary|blueGreen
    canary_steps optional
    blue_green_config optional

Resources:
- kubernetes_manifest for Rollout
- kubernetes_service_v1
- optional ingress

Docs:
- docs/progressive-delivery.md
- Explain standard Deployment vs Rollout.
- Explain canary.
- Explain blue/green.
- Explain metrics provider requirements.
- Explain rollback.

Example:
- examples/kubernetes-argo-rollouts-canary

Rules:
- Do not replace default app module.
- Rollout must be opt-in.
- Keep first strategy simple.

Run:
- terraform fmt -recursive
```

---
