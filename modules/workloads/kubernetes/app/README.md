# workloads/kubernetes/app

Deploys a generic web or service workload to Kubernetes using the Terraform
Kubernetes provider.

Provider configuration belongs in the root module. This module does not create
or configure a cluster.

## Secret Handling

Do not put secret values in Terraform variables. Use `secret_env` to reference
keys from existing Kubernetes Secrets.

## Terraform Or GitOps

Use this module when Terraform owns the workload lifecycle. If your cluster is
managed through Argo CD or another GitOps controller, prefer committing the
application manifests or Helm values to the GitOps repository instead of
managing the same workload from Terraform.

## Basic Deployment

```hcl
module "app" {
  source = "../../../modules/workloads/kubernetes/app"

  name      = "hello"
  namespace = "apps"
  image     = "nginx:1.27"

  ports = [{
    name           = "http"
    container_port = 80
  }]
}
```

## Ingress

```hcl
module "app" {
  source = "../../../modules/workloads/kubernetes/app"

  name      = "hello"
  namespace = "apps"
  image     = "nginx:1.27"

  ports = [{
    name           = "http"
    container_port = 80
  }]

  ingress = {
    enabled    = true
    class_name = "nginx"
    host       = "hello.example.com"
  }
}
```

## Secret Environment Variables

```hcl
module "app" {
  source = "../../../modules/workloads/kubernetes/app"

  name      = "api"
  namespace = "apps"
  image     = "example/api:1.0.0"

  secret_env = {
    DATABASE_URL = {
      secret_name = "api-database"
      secret_key  = "url"
    }
  }
}
```

## Autoscaling

```hcl
module "app" {
  source = "../../../modules/workloads/kubernetes/app"

  name      = "api"
  namespace = "apps"
  image     = "example/api:1.0.0"

  autoscaling = {
    enabled      = true
    min_replicas = 2
    max_replicas = 6
    cpu_percent  = 70
  }
}
```

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
