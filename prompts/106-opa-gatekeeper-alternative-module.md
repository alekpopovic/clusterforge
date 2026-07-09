## Prompt 106 — OPA Gatekeeper alternative module

```text
Add OPA Gatekeeper support as an alternative to Kyverno.

Create module:
- modules/platform/kubernetes/gatekeeper

Purpose:
Install Gatekeeper via Helm and provide optional constraints/templates support.

Inputs:
- namespace default "gatekeeper-system"
- create_namespace default true
- chart_version default ""
- values list(string)
- constraint_templates map(string) default {}
- constraints map(string) default {}

Outputs:
- namespace
- release_name

Docs:
- docs/kubernetes-policy-gatekeeper.md
- Compare Kyverno and Gatekeeper.
- Explain when to choose each.
- Explain audit vs enforcement.
- Explain rollout strategy.

Example:
- examples/kubernetes-gatekeeper-baseline

Rules:
- Do not install both Kyverno and Gatekeeper by default.
- Do not enforce strict policies by default.
- Keep module optional.

Run:
- terraform fmt -recursive
```

---
