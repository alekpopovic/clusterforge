# RFC 002: Azure AKS Support

## Goals

- Azure network module
- AKS cluster module
- node pool configuration
- optional Azure DNS later
- managed identity
- workload identity
- Kubernetes provider output config

## Non-Goals

- service mesh
- advanced private clusters
- Azure Policy deep integration
- multi-region AKS

## Proposed Modules

- `modules/cloud/azure/network`: resource group, virtual network, subnets
- `modules/cloud/azure/dns`: future Azure DNS zones and records
- `modules/cloud/azure/identity`: future managed identity helpers
- `modules/orchestrators/kubernetes/aks`: AKS cluster and node pools

Provider configuration belongs in root modules. `azurerm` is required;
`azuread` should be added only when identity workflows require it.

## CLI Impact

```bash
cf env create dev --cloud azure --orchestrator aks --region westeurope
cf generate dev
```

## Risks

Azure provider complexity, identity permissions, private cluster networking,
and DNS/certificate differences need real account validation.
