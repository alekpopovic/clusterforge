locals {
  release_name = "karpenter"

  default_values = yamlencode({
    settings = {
      clusterName     = var.cluster_name
      clusterEndpoint = var.cluster_endpoint
    }

    serviceAccount = {
      annotations = {
        "eks.amazonaws.com/role-arn" = var.service_account_role_arn
      }
    }
  })
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
  repository = "oci://public.ecr.aws/karpenter"
  chart      = "karpenter"
  version    = var.chart_version == "" ? null : var.chart_version
  values     = concat([local.default_values], var.values)

  depends_on = [kubernetes_namespace_v1.this]
}
