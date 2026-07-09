## Prompt 72 — RBAC and service account workload support

```text
Enhance Kubernetes workload modules with service account and RBAC support.

Target modules:
- modules/workloads/kubernetes/app
- modules/workloads/kubernetes/worker
- modules/workloads/kubernetes/cronjob

Add inputs:
- service_account:
    create: bool
    name: optional(string)
    annotations: optional(map(string))
    labels: optional(map(string))
    automount_token: optional(bool)

- rbac:
    create: bool
    rules: list(object({
      api_groups = list(string)
      resources = list(string)
      verbs = list(string)
      resource_names = optional(list(string))
    }))

Resources:
- kubernetes_service_account_v1 optional
- kubernetes_role_v1 optional
- kubernetes_role_binding_v1 optional

Behavior:
- If service_account.create=true and name empty, use workload name.
- If rbac.create=true, service account must exist or name must be provided.
- Support service account annotations for IRSA/Workload Identity.
- Do not create ClusterRole by default.

README:
- IRSA annotation example.
- Least privilege RBAC example.
- Warning about automounting tokens.

Tests/validation:
- terraform fmt
- examples updated

Rules:
- Backward compatible defaults.
- No broad permissions by default.
```

---
