# orchestrators/kubernetes/rke2

Experimental RKE2 cloud-init/user-data generator.

This module does not create servers, SSH to hosts, or store kubeconfig
credentials. Use the generated user data with an existing VM workflow.

```hcl
module "rke2" {
  source = "../../../../modules/orchestrators/kubernetes/rke2"

  cluster_name = "clusterforge-dev-rke2"
  environment  = "dev"
  server_count = 1
}
```
