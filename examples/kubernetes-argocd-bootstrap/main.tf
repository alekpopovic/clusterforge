module "argocd" {
  source = "../../modules/platform/kubernetes/argocd"

  namespace        = "argocd"
  create_namespace = true

  enable_app_of_apps   = true
  app_of_apps_repo_url = var.gitops_repo_url
  app_of_apps_path     = var.gitops_path
  app_of_apps_revision = var.gitops_revision
}
