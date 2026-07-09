# cloud/azure/network

MVP Azure network module for ClusterForge AKS environments.

Status: experimental.

```hcl
module "network" {
  source = "../../../modules/cloud/azure/network"

  name                = "clusterforge-dev-aks"
  environment         = "dev"
  location            = "westeurope"
  resource_group_name = "rg-clusterforge-dev"
}
```

Provider configuration belongs in the root module.
