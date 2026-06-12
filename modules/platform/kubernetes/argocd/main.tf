locals {
  release_name = "argo-cd"
}

resource "kubernetes_namespace_v1" "this" {
  count = var.create_namespace ? 1 : 0

  metadata {
    name   = var.namespace
    labels = var.labels
  }
}

resource "helm_release" "this" {
  name       = local.release_name
  namespace  = var.namespace
  repository = "https://argoproj.github.io/argo-helm"
  chart      = "argo-cd"
  version    = var.chart_version == "" ? null : var.chart_version
  values     = var.values

  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_manifest" "app_of_apps" {
  count = var.enable_app_of_apps ? 1 : 0

  manifest = {
    apiVersion = "argoproj.io/v1alpha1"
    kind       = "Application"

    metadata = {
      name      = var.app_of_apps_name
      namespace = var.namespace
      labels    = var.labels
    }

    spec = {
      project = var.app_of_apps_project

      source = {
        repoURL        = var.app_of_apps_repo_url
        path           = var.app_of_apps_path
        targetRevision = var.app_of_apps_revision
      }

      destination = {
        server    = "https://kubernetes.default.svc"
        namespace = var.app_of_apps_destination_namespace
      }

      syncPolicy = {
        automated = {
          prune    = false
          selfHeal = true
        }
      }
    }
  }

  depends_on = [helm_release.this]
}
