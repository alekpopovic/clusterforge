## Prompt 67 — Implement Azure network and AKS MVP

```text
Implement Azure AKS MVP based on docs/rfcs/002-aks-support.md.

Create modules:
- modules/cloud/azure/network
- modules/orchestrators/kubernetes/aks

Azure network module:
Inputs:
- name
- environment
- location
- resource_group_name
- create_resource_group default true
- address_space
- subnet_prefixes
- tags

Resources:
- azurerm_resource_group optional
- azurerm_virtual_network
- azurerm_subnet

Outputs:
- resource_group_name
- vnet_id
- subnet_ids

AKS module:
Inputs:
- name
- environment
- location
- resource_group_name
- subnet_id
- kubernetes_version default ""
- dns_prefix
- default_node_pool object
- tags

Resources:
- azurerm_kubernetes_cluster
- system-assigned or user-assigned identity; choose simple default
- default node pool

Outputs:
- cluster_name
- resource_group_name
- kube_config_raw sensitive
- host
- client_certificate sensitive if available
- cluster_ca_certificate sensitive if available

Create:
- examples/azure-aks-minimal
- live/dev/azure-aks template
- CLI generator support:
  - cloud azure + orchestrator aks

Rules:
- Provider configured in root.
- Do not store real kubeconfig in examples.
- Mark sensitive outputs.
- Keep MVP simple.

Run:
- terraform fmt -recursive
- gofmt
- go test ./...
```

---
