---
title: GitOps
permalink: /gitops/
---

# GitOps

ClusterForge supports a practical Terraform plus GitOps boundary.

Terraform should create durable infrastructure and initial platform bootstrap:

- Cloud foundation resources
- Container orchestrators such as EKS or ECS
- Cluster platform add-ons such as ingress, cert-manager, External Secrets,
  and Argo CD
- Minimal bootstrap resources needed to hand control to GitOps

Argo CD should manage frequent application and platform configuration changes
after bootstrap:

- Application deployments
- Helm values that change often
- Environment overlays
- Progressive delivery resources
- App-specific policies and manifests

Plain Terraform workload modules remain available for smaller deployments,
simple clusters, and teams that do not want GitOps yet.

## App-Of-Apps

The app-of-apps pattern uses one Argo CD `Application` to point at a Git path
that contains other `Application` definitions. Terraform can create that first
Application, and Argo CD reconciles the rest from Git.

ClusterForge supports this in `modules/platform/kubernetes/argocd`:

```hcl
module "argocd" {
  source = "../modules/platform/kubernetes/argocd"

  namespace = "argocd"

  enable_app_of_apps   = true
  app_of_apps_repo_url = "https://github.com/example/platform-gitops.git"
  app_of_apps_path     = "gitops/apps"
  app_of_apps_revision = "main"
}
```

Do not include Git credentials in Terraform. Configure repository credentials
through Argo CD, External Secrets Operator, cloud secret managers, or another
approved secret path.

## Recommended Repository Layout

```text
gitops/
  apps/
    cluster-apps.yaml
    ingress.yaml
    observability.yaml
    services.yaml
  environments/
    dev/
    staging/
    prod/
  projects/
    platform-project.yaml
    workloads-project.yaml
```

`apps/` contains Argo CD `Application` definitions. `environments/` contains
overlays or values per environment. `projects/` contains Argo CD `AppProject`
boundaries for teams or platform areas.

## Security Notes

- Keep Git credentials out of Terraform variables, state, and generated files.
- Use read-only deploy keys or fine-scoped tokens when private repositories are
  required.
- Store credentials in a secret manager and sync them through a controlled
  secret workflow.
- Review production app-of-apps changes like code changes; a small Git diff can
  update many applications.
