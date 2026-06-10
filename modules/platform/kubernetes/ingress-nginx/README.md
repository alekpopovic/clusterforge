# platform/kubernetes/ingress-nginx

Installs ingress-nginx with Helm.

This module assumes Kubernetes and Helm providers are configured in the root
module.

```hcl
module "ingress_nginx" {
  source = "../../../modules/platform/kubernetes/ingress-nginx"

  namespace        = "ingress-nginx"
  create_namespace = true
  values           = []
}
```
