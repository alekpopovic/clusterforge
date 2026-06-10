# kubernetes-basic-app

Example root module for `modules/workloads/kubernetes/app`.

The Kubernetes provider is configured in this root module using
`var.kubeconfig_path`.

```bash
terraform init
terraform plan
```

This example deploys `nginx` as a Deployment and ClusterIP Service.
