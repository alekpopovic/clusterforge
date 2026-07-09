## Prompt 49 — Pod Security i NetworkPolicy module

```text
Add Kubernetes security baseline modules.

Create modules:
1. modules/platform/kubernetes/pod-security
2. modules/platform/kubernetes/network-policy-baseline

pod-security module:
Purpose:
Apply Kubernetes Pod Security Admission labels to namespaces.

Inputs:
- namespaces: map(object({
    enforce = optional(string, "baseline")
    audit = optional(string, "restricted")
    warn = optional(string, "restricted")
  }))
- labels: map(string), default {}

Resources:
- kubernetes_labels or kubernetes_namespace_v1 management approach.
Choose an approach that does not accidentally take over existing namespaces destructively.
Document limitations.

Validation:
- enforce/audit/warn must be privileged, baseline, or restricted.

network-policy-baseline:
Purpose:
Create default NetworkPolicy resources.

Inputs:
- namespace: string
- default_deny_ingress: bool, default true
- default_deny_egress: bool, default false
- allow_dns_egress: bool, default true
- labels: map(string), default {}

Resources:
- kubernetes_network_policy_v1

README:
- Explain Pod Security levels.
- Explain default deny tradeoffs.
- Explain that CNI must support NetworkPolicy.
- Include examples.

Update bootstrap module:
- enable_pod_security
- enable_network_policy_baseline optional

Rules:
- Do not break existing workloads by default.
- Default deny should not be globally enabled by bootstrap unless explicitly requested.

Run:
- terraform fmt -recursive
```

---
