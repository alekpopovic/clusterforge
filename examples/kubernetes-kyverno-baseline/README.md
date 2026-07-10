# Kubernetes Kyverno audit baseline

This example installs Kyverno and optionally enables the ClusterForge baseline
in `Audit` mode. It uses one replica per controller for a small non-production
cluster; production requires reviewed availability and resource settings.

For a fresh cluster, apply in two phases because policy CRDs must exist before
Terraform can plan `kubernetes_manifest` resources:

```bash
terraform init
terraform plan -var='kyverno_chart_version=<reviewed-version>'
terraform apply -var='kyverno_chart_version=<reviewed-version>'
terraform plan -var='kyverno_chart_version=<reviewed-version>' -var='enable_baseline_policies=true'
terraform apply -var='kyverno_chart_version=<reviewed-version>' -var='enable_baseline_policies=true'
```

Review PolicyReports and workload warnings before considering `Enforce`. Remove
policies before uninstalling Kyverno, then clean up with `terraform destroy`.
