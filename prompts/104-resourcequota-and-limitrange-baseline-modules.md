## Prompt 104 — ResourceQuota and LimitRange baseline modules

```text
Create standalone Kubernetes quota baseline modules.

Modules:
- modules/platform/kubernetes/resource-quota
- modules/platform/kubernetes/limit-range

resource-quota inputs:
- namespace
- hard map(string)
- labels map(string)

limit-range inputs:
- namespace
- limits list(object({
    type = string
    default = optional(map(string))
    default_request = optional(map(string))
    max = optional(map(string))
    min = optional(map(string))
  }))
- labels map(string)

Outputs:
- namespace
- resource_quota_name
- limit_range_name

Examples:
- examples/kubernetes-resource-governance

Docs:
- docs/kubernetes-resource-governance.md
- Explain requests/limits.
- Explain namespace quotas.
- Explain HPA interaction.
- Explain common pitfalls.

Rules:
- Do not apply globally by default.
- Keep modules explicit.
- Avoid breaking workloads without user opt-in.

Run:
- terraform fmt -recursive
```

---
