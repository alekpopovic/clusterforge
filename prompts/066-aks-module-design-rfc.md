## Prompt 66 — AKS module design RFC

```text
Create an RFC for Azure AKS support.

Goal:
Design AKS support before implementation.

Create:
- docs/rfcs/002-aks-support.md

Include:
1. Goals
   - Azure network
   - AKS cluster
   - node pools
   - Azure DNS optional
   - managed identity
   - workload identity
   - Kubernetes provider output config

2. Non-goals for first implementation
   - service mesh
   - advanced private clusters
   - Azure Policy deep integration
   - multi-region

3. Proposed modules:
   - modules/cloud/azure/network
   - modules/cloud/azure/dns
   - modules/cloud/azure/identity
   - modules/orchestrators/kubernetes/aks

4. Inputs and outputs for each module.

5. Provider strategy:
   - azurerm provider in root
   - azuread provider only if needed
   - kubernetes/helm configured in root after cluster

6. Security model:
   - managed identity
   - workload identity
   - no secrets in tfvars

7. Examples:
   - examples/azure-aks-minimal
   - live/dev/azure-aks

8. CLI impact:
   - cf env create dev --cloud azure --orchestrator aks --region westeurope
   - cf generate dev

9. Risks:
   - Azure provider complexity
   - identity permissions
   - private cluster complexity
   - DNS/certificate differences

Do not implement code in this prompt.
Only create the RFC and update docs/roadmap.md.
```

---
