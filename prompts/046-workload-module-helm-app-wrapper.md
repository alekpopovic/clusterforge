## Prompt 46 — Workload module: Helm app wrapper

```text
Implement modules/workloads/kubernetes/helm-app.

Purpose:
Deploy an application using a Helm chart, with ClusterForge naming and metadata conventions.

Provider:
- hashicorp/helm
- Provider configured in root.

Inputs:
- name: string
- namespace: string
- create_namespace: bool, default true
- repository: string
- chart: string
- chart_version: string, default ""
- values: list(string), default []
- set: map(string), default {}
- set_sensitive: map(string), default {}
- labels: map(string), default {}
- timeout: number, default 300
- atomic: bool, default false
- cleanup_on_fail: bool, default true
- wait: bool, default true

Resources:
- helm_release

Behavior:
- Use set blocks for set.
- Use set_sensitive for sensitive values, but document Terraform state implications.
- Do not force chart versions but strongly recommend pinning.

Outputs:
- release_name
- namespace
- chart
- version
- status if available

README:
- Basic Helm app example.
- Values file example.
- Warning about set_sensitive and Terraform state.
- Explain when to prefer Argo CD instead.

Create example:
- examples/kubernetes-helm-app

Run:
- terraform fmt -recursive
```

---
