locals {
  namespace    = trimspace(var.namespace)
  release_name = "gatekeeper"
  labels = merge(var.labels, {
    "app.kubernetes.io/managed-by" = "terraform"
    "clusterforge.io/component"    = "gatekeeper"
  })
}

resource "kubernetes_namespace_v1" "this" {
  count = var.create_namespace ? 1 : 0
  metadata {
    name   = local.namespace
    labels = local.labels
  }
}

resource "helm_release" "this" {
  name       = local.release_name
  namespace  = local.namespace
  repository = "https://open-policy-agent.github.io/gatekeeper/charts"
  chart      = "gatekeeper"
  version    = var.chart_version == "" ? null : var.chart_version
  values     = var.values
  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_manifest" "constraint_template" {
  for_each   = var.constraint_templates
  manifest   = yamldecode(each.value)
  depends_on = [helm_release.this]
}

resource "kubernetes_manifest" "constraint" {
  for_each = var.constraints
  manifest = merge(yamldecode(each.value), {
    spec = merge(try(yamldecode(each.value).spec, {}), {
      enforcementAction = var.enforcement_action
    })
  })
  depends_on = [kubernetes_manifest.constraint_template]
}
