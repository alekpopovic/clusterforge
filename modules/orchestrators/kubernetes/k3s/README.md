# orchestrators/kubernetes/k3s

Experimental K3s cloud-init/user-data generator.

This module does not create servers, SSH to hosts, or store kubeconfig
credentials. Use the generated user data with an existing VM workflow.

```hcl
module "k3s" {
  source = "../../../../modules/orchestrators/kubernetes/k3s"

  cluster_name = "clusterforge-dev-k3s"
  environment  = "dev"
  server_count = 1
}
```
