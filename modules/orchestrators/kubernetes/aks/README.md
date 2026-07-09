# orchestrators/kubernetes/aks

MVP Azure AKS module.

Status: experimental.

```hcl
module "aks" {
  source = "../../../../modules/orchestrators/kubernetes/aks"

  name                = "clusterforge-dev-aks"
  environment         = "dev"
  location            = "westeurope"
  resource_group_name = module.network.resource_group_name
  subnet_id           = module.network.subnet_ids[0]
  dns_prefix          = "clusterforge-dev"
}
```

Provider configuration belongs in the root module. Kubeconfig outputs are
sensitive and must not be committed.
