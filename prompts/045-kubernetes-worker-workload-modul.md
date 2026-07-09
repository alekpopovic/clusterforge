## Prompt 45 — Kubernetes worker workload modul

```text
Implement Kubernetes worker workload module.

Create module:
- modules/workloads/kubernetes/worker

Purpose:
Deploy a background worker process to Kubernetes.

Similar to app module, but:
- No Service by default.
- No Ingress.
- Supports replicas.
- Supports env and secret_env.
- Supports resources.
- Supports image pull secrets.
- Supports command and args.
- Supports pod annotations and labels.

Inputs:
- name
- namespace
- create_namespace
- image
- replicas default 1
- command
- args
- env
- secret_env
- resources
- labels
- annotations
- image_pull_policy
- image_pull_secrets
- termination_grace_period_seconds optional
- service_account_name optional

Resources:
- kubernetes_namespace_v1 optional
- kubernetes_deployment_v1

Optional:
- autoscaling using HPA if metrics available.

Outputs:
- name
- namespace
- deployment_name
- labels

README:
- Queue worker example.
- Secret reference example.
- Autoscaling note.
- Difference between app, worker and cronjob.

Create example:
- examples/kubernetes-worker

Run:
- terraform fmt -recursive
```

---
