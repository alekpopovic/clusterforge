# workloads/kubernetes/worker

Deploys a background worker process to Kubernetes using the Terraform
Kubernetes provider.

For EKS IRSA, set `service_account_name` and the
`eks.amazonaws.com/role-arn` key in `service_account_annotations`. The module
creates that annotated service account; create the IAM role separately.

Provider configuration belongs in the root module. This module does not create
or configure a cluster.

## Difference From App And CronJob

- `workloads/kubernetes/app` is for long-running web/service workloads and can
  create a Service, Ingress, and HPA.
- `workloads/kubernetes/worker` is for long-running background consumers and
  creates a Deployment with no Service or Ingress.
- `workloads/kubernetes/cronjob` is for scheduled jobs that run on a cron
  expression.

## Queue Worker Example

```hcl
module "worker" {
  source = "../../../modules/workloads/kubernetes/worker"

  name      = "email-worker"
  namespace = "apps"
  image     = "example/email-worker:1.0.0"
  replicas  = 2

  command = ["./worker"]
  args    = ["--queue", "emails"]

  env = {
    LOG_LEVEL = "info"
    QUEUE     = "emails"
  }

  resources = {
    cpu_request    = "100m"
    memory_request = "128Mi"
    cpu_limit      = "500m"
    memory_limit   = "512Mi"
  }
}
```

## Secret References

Do not put secret values in Terraform variables. Use `secret_env` to reference
keys from existing Kubernetes Secrets.

```hcl
module "worker" {
  source = "../../../modules/workloads/kubernetes/worker"

  name      = "billing-worker"
  namespace = "apps"
  image     = "example/billing-worker:1.0.0"

  secret_env = {
    DATABASE_URL = {
      secret_name = "billing-secrets"
      secret_key  = "database-url"
    }
  }
}
```

## Autoscaling

This module supports a basic CPU-based Horizontal Pod Autoscaler:

```hcl
autoscaling = {
  enabled      = true
  min_replicas = 2
  max_replicas = 8
  cpu_percent  = 70
}
```

Queue-depth or event-driven autoscaling usually needs a dedicated scaler such
as KEDA and should be added through a platform module or explicit workload
configuration.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
