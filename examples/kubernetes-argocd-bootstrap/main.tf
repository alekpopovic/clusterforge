module "argocd" {
  source = "../../modules/platform/kubernetes/argocd"

  namespace        = var.namespace
  create_namespace = true
  labels           = var.labels

  enable_app_of_apps   = var.enable_app_of_apps
  app_of_apps_repo_url = var.gitops_repo_url
  app_of_apps_path     = var.gitops_path
  app_of_apps_revision = var.gitops_revision
}
