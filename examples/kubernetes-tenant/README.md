# Kubernetes tenant example

This example creates one tenant namespace with Pod Security labels, a quota,
container defaults, and read-only group RBAC. Network default-deny remains off
to avoid disrupting workloads before allow policies are designed.

```bash
terraform init
terraform plan
terraform apply
terraform destroy
```

Use a disposable cluster or reviewed context. The example reads an existing
kubeconfig path but does not contain or generate credentials. Review quotas and
identity-provider group names before apply.
