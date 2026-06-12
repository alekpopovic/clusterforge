# platform/kubernetes/argocd

## Purpose

Installs Argo CD with Helm and can optionally create a bootstrap Argo CD
Application for an app-of-apps GitOps pattern.

This module assumes Kubernetes and Helm providers are configured in the root
module. Terraform should install the controller and minimal bootstrap objects;
frequent application deployment changes should usually move to GitOps after
bootstrap.

## Status

Implemented.

## Usage

```hcl
module "argocd" {
  source = "../../../modules/platform/kubernetes/argocd"

  namespace = "argocd"
}
```

## App-Of-Apps Example

```hcl
module "argocd" {
  source = "../../../modules/platform/kubernetes/argocd"

  namespace = "argocd"

  enable_app_of_apps    = true
  app_of_apps_repo_url  = "https://github.com/example/platform-gitops.git"
  app_of_apps_path      = "gitops/apps"
  app_of_apps_revision  = "main"
  app_of_apps_project   = "default"
}
```

Do not put Git credentials in Terraform. Configure private repository access
through Argo CD repository credentials, a secret integration, or another
approved secret management flow.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
