locals {
  release_name = "kyverno"
  namespace    = trimspace(var.namespace)
  labels = merge(var.labels, {
    "app.kubernetes.io/managed-by" = "terraform"
    "clusterforge.io/component"    = "kyverno"
  })

  privileged_policy = {
    apiVersion = "kyverno.io/v1"
    kind       = "ClusterPolicy"
    metadata = {
      name   = "clusterforge-disallow-privileged-containers"
      labels = local.labels
    }
    spec = {
      background  = true
      emitWarning = var.baseline_failure_action == "Audit"
      rules = [{
        name = "disallow-privileged-containers"
        match = {
          any = [{ resources = { kinds = ["Pod"] } }]
        }
        validate = {
          failureAction = var.baseline_failure_action
          message       = "Privileged containers are not allowed."
          pattern = {
            spec = {
              containers = [{
                name = "*"
                "=(securityContext)" = {
                  "=(privileged)" = "false"
                }
              }]
            }
          }
        }
      }]
    }
  }

  resources_policy = {
    apiVersion = "kyverno.io/v1"
    kind       = "ClusterPolicy"
    metadata = {
      name   = "clusterforge-require-resources"
      labels = local.labels
    }
    spec = {
      background  = true
      emitWarning = var.baseline_failure_action == "Audit"
      rules = [{
        name = "require-container-requests-and-limits"
        match = {
          any = [{ resources = { kinds = ["Pod"] } }]
        }
        validate = {
          failureAction = var.baseline_failure_action
          message       = "Containers must define CPU and memory requests and limits."
          pattern = {
            spec = {
              containers = [{
                name = "*"
                resources = {
                  requests = { cpu = "?*", memory = "?*" }
                  limits   = { cpu = "?*", memory = "?*" }
                }
              }]
            }
          }
        }
      }]
    }
  }

  latest_tag_policy = {
    apiVersion = "kyverno.io/v1"
    kind       = "ClusterPolicy"
    metadata = {
      name   = "clusterforge-disallow-latest-tag"
      labels = local.labels
    }
    spec = {
      background  = true
      emitWarning = var.baseline_failure_action == "Audit"
      rules = [{
        name = "disallow-latest-tag"
        match = {
          any = [{ resources = { kinds = ["Pod"] } }]
        }
        validate = {
          failureAction = var.baseline_failure_action
          message       = "Container images must not use the latest tag."
          pattern = {
            spec = {
              containers = [{ name = "*", image = "!*:latest" }]
            }
          }
        }
      }]
    }
  }

  baseline_policies = var.enable_baseline_policies ? merge(
    { privileged = local.privileged_policy },
    var.enable_require_resources_policy ? { resources = local.resources_policy } : {},
    var.enable_disallow_latest_tag_policy ? { latest_tag = local.latest_tag_policy } : {}
  ) : {}
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
  repository = "https://kyverno.github.io/kyverno/"
  chart      = "kyverno"
  version    = var.chart_version == "" ? null : var.chart_version
  values     = var.values

  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_manifest" "baseline" {
  for_each = local.baseline_policies

  manifest = each.value

  depends_on = [helm_release.this]
}

resource "kubernetes_manifest" "custom" {
  for_each = var.policies

  manifest = yamldecode(each.value)

  depends_on = [helm_release.this]
}
