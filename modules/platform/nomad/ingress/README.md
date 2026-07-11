# Nomad ingress metadata

Produces reviewed service-registration tags for an operator-installed ingress
controller. It does not install, expose, or secure the ingress controller.

```hcl
module "ingress_tags" {
  source       = "../../modules/platform/nomad/ingress"
  service_name = "api"
}
```
