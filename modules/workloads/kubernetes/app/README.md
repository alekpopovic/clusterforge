# workloads/kubernetes/app

Deploys a basic Kubernetes application as a Deployment and ClusterIP Service.

Provider configuration must be declared in the root module. Do not pass secrets
through `env`; use a dedicated secrets module or an external secret controller.

## Example

```hcl
module "app" {
  source = "../../../modules/workloads/kubernetes/app"

  name  = "hello"
  image = "nginx:1.27"
}
```
