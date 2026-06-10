# orchestrators/kubernetes/eks

Creates an AWS EKS cluster and one default managed node group.

Provider configuration stays in the root module. This module intentionally
does not configure Kubernetes, Helm, or kubectl providers; platform bootstrap
belongs in separate platform-layer modules.

## Example

```hcl
module "eks" {
  source = "../../../../modules/orchestrators/kubernetes/eks"

  name       = "clusterforge-dev-eks"
  subnet_ids = module.network.private_subnet_ids
  tags       = module.tags.tags
}
```
