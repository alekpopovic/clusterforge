# kubernetes-basic-app

Example root module for `modules/workloads/kubernetes/app`.

This example deploys `nginx` as a Kubernetes Deployment and ClusterIP Service.
It assumes the Kubernetes provider can connect to an existing cluster.

## Provider

The provider uses `var.kubeconfig_path`, defaulting to `~/.kube/config`.

```bash
terraform init
terraform validate
terraform plan
```

Use a test namespace or disposable cluster when experimenting. The module can
create the namespace and workload resources, so review the plan before apply.

No Kubernetes secrets are created by this example.
