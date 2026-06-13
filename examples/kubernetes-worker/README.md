# kubernetes-worker

Example root module for `modules/workloads/kubernetes/worker`.

This example deploys a simple background worker as a Kubernetes Deployment with
no Service and no Ingress. It assumes the Kubernetes provider can connect to an
existing cluster.

## Provider

The provider uses `var.kubeconfig_path`, defaulting to `~/.kube/config`.

```bash
terraform init
terraform validate
terraform plan
```

The `secret_env` block references an existing Kubernetes Secret named
`worker-secrets`. The example does not create or store secret values.

Use a test namespace or disposable cluster when experimenting. Review the plan
before apply.
