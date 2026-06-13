# platform/kubernetes/bootstrap

## Purpose

Convenience composition module for common Kubernetes platform add-ons. It
installs selected add-ons by calling focused Helm-wrapper child modules.

This module does not create a Kubernetes cluster. It assumes the root
environment has already configured Kubernetes and Helm providers from an
existing cluster, such as EKS.

## Terraform And GitOps Boundary

Use this module for initial platform bootstrap when Terraform owns the add-on
installation lifecycle. Once Argo CD or another GitOps controller is installed,
teams may choose to move ongoing add-on configuration into GitOps. Avoid
managing the same Helm release from Terraform and GitOps at the same time.

This module installs platform add-ons, not business applications. Application
deployments belong in workload modules or GitOps application definitions.

## CRD And Helm Lifecycle Warning

Several platform add-ons install or depend on CRDs, especially cert-manager,
kube-prometheus-stack, and Argo CD. Review chart upgrade notes before changing
chart versions or values. Terraform tracks Helm releases, but CRD lifecycle can
still require careful manual review.

## Example

```hcl
module "platform_bootstrap" {
  source = "../../../modules/platform/kubernetes/bootstrap"

  enable_ingress_nginx    = true
  enable_cert_manager     = true
  enable_external_secrets = true
  enable_pod_security     = true
  enable_karpenter        = false
  enable_prometheus_stack = true
  enable_loki             = true
  enable_log_agent        = true
  enable_argocd           = true

  pod_security_namespaces = {
    apps = {
      enforce = "baseline"
      audit   = "restricted"
      warn    = "restricted"
    }
  }

  argocd_enable_app_of_apps   = true
  argocd_app_of_apps_repo_url = "https://github.com/example/platform-gitops.git"
  argocd_app_of_apps_path     = "gitops/apps"

  common_labels = {
    "clusterforge.io/environment" = "dev"
    "clusterforge.io/managed-by"  = "terraform"
  }
}
```

Enable External Secrets Operator when Kubernetes workloads should consume
secrets synced from external secret stores such as AWS Secrets Manager or SSM
Parameter Store. Terraform should manage references and operator configuration,
not raw secret values.

Enable Karpenter only for EKS clusters where the AWS IRSA role and node class
strategy have been reviewed. Keep at least one managed node group or equivalent
bootstrap capacity for system add-ons and the Karpenter controller.

Enable `enable_log_agent` to install Grafana Alloy alongside Loki. Provide
explicit Alloy values for production log collection rules and destinations.

Pod Security and NetworkPolicy baselines are opt-in. Do not enable default-deny
network policies globally without testing namespace-specific allow policies and
confirming the cluster CNI enforces NetworkPolicy.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
