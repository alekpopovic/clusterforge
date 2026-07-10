# Argo Rollouts canary example

Install the controller first, then enable the app after its CRD exists:

```bash
terraform init
terraform apply -var='argo_rollouts_chart_version=<reviewed-version>'
terraform apply -var='argo_rollouts_chart_version=<reviewed-version>' -var='enable_rollout_app=true'
```

The demo uses pod-count-based canary steps and a manual pause; it does not claim
precise traffic percentages or automated analysis. Inspect and promote it with
the optional `kubectl argo rollouts` plugin. Disable the app before destroying
the controller.
