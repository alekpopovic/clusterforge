## Prompt 128 — Kubernetes upgrade planner

```text
Implement Kubernetes upgrade planner.

Goal:
Help operators plan cluster upgrades safely.

CLI:
- cf k8s upgrade plan <cluster-or-env>
- cf k8s upgrade check <cluster-or-env>
- cf k8s versions <cluster-or-env>

Inputs:
- current Kubernetes version from config, Terraform outputs, or kubectl if available.
- target version from flag:
  --target-version 1.31

Checks:
- version jump is one minor at a time unless provider supports otherwise
- EKS/AKS/GKE support status from local VERSION_MATRIX.md
- node group version alignment
- platform chart compatibility notes
- deprecated Kubernetes APIs in manifests if detectable
- Helm release versions if available
- workload API versions in generated Terraform files

Output:
- upgrade readiness
- blocking issues
- warnings
- recommended steps
- rollback notes

Docs:
- docs/kubernetes-upgrades.md

Rules:
- Do not perform upgrade.
- Planning only.
- Do not query internet.
- Use local support matrix and config.
- Be explicit when live cluster access is unavailable.

Tests:
- current 1.29 to target 1.30 allowed.
- current 1.29 to target 1.31 warns/block based on policy.
- unsupported target warns.
- deprecated API fixture detected.

Run:
- gofmt
- go test ./...
```


---
