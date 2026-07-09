## Prompt 103 — Kubernetes tenant model

```text
Implement Kubernetes tenant model.

Goal:
Support teams/apps as tenants with namespace, quota, limit range, network policy, and optional RBAC.

Create module:
- modules/platform/kubernetes/tenant

Inputs:
- name
- namespaces: list(string)
- labels map(string)
- annotations map(string)
- pod_security:
    enforce
    audit
    warn
- resource_quota optional
- limit_range optional
- network_policy:
    default_deny_ingress
    default_deny_egress
    allow_dns_egress
- rbac:
    create
    subjects

Resources:
- kubernetes_namespace_v1
- kubernetes_resource_quota_v1 optional
- kubernetes_limit_range_v1 optional
- kubernetes_network_policy_v1 optional
- kubernetes_role_v1 optional
- kubernetes_role_binding_v1 optional

Docs:
- docs/kubernetes-tenancy.md
- Explain namespace-per-team vs namespace-per-app.
- Explain quotas.
- Explain pod security.
- Explain network policies.
- Explain RBAC limitations.

Example:
- examples/kubernetes-tenant

Rules:
- Do not create ClusterRole by default.
- Defaults should not break workloads.
- Default deny must be opt-in.
- Keep permissions minimal.

Run:
- terraform fmt -recursive
```

---
