# workloads/kubernetes/cronjob

Deploys a Kubernetes CronJob using the Terraform Kubernetes provider.

For EKS IRSA, set `service_account_name` and the
`eks.amazonaws.com/role-arn` key in `service_account_annotations`. The module
creates that annotated service account; create the IAM role separately.

Provider configuration belongs in the root module. This module does not create
or configure a cluster.

## Secret Handling

Do not put secret values in Terraform variables. Use `secret_env` to reference
keys from existing Kubernetes Secrets. This keeps secret values out of Terraform
configuration and avoids writing them directly into state through this module.

## Basic Example

```hcl
module "cleanup" {
  source = "../../../modules/workloads/kubernetes/cronjob"

  name      = "cleanup"
  namespace = "jobs"
  image     = "busybox:1.36"
  schedule  = "0 * * * *"
  command   = ["/bin/sh", "-c"]
  args      = ["echo cleanup"]
}
```

## Environment And Secret References

```hcl
module "report" {
  source = "../../../modules/workloads/kubernetes/cronjob"

  name      = "report"
  namespace = "jobs"
  image     = "example/report:1.0.0"
  schedule  = "*/15 * * * *"

  env = {
    LOG_LEVEL = "info"
  }

  secret_env = {
    API_TOKEN = {
      secret_name = "report-api"
      secret_key  = "token"
    }
  }
}
```

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
