# Kubernetes resource governance example

This example opts one existing namespace into aggregate quotas and per-container
defaults. Adjust all quantities to measured workload needs before apply.

```bash
terraform init
terraform plan -var='namespace=apps'
terraform apply -var='namespace=apps'
terraform destroy -var='namespace=apps'
```

Use a reviewed kubeconfig context. The example does not create a namespace, so
accidental global rollout is not possible from these modules.
