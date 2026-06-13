locals {
  release_name = "kube-prometheus-stack"

  generated_values = merge(
    var.enable_grafana_ingress ? {
      grafana = {
        ingress = {
          enabled = true
          hosts   = [var.grafana_host]
        }
      }
    } : {},
    var.storage_enabled ? {
      prometheus = {
        prometheusSpec = {
          storageSpec = {
            volumeClaimTemplate = {
              spec = {
                storageClassName = var.storage_class_name == "" ? null : var.storage_class_name
                accessModes      = ["ReadWriteOnce"]
                resources = {
                  requests = {
                    storage = var.storage_size
                  }
                }
              }
            }
          }
        }
      }
    } : {}
  )

  generated_values_yaml = length(keys(local.generated_values)) > 0 ? [yamlencode(local.generated_values)] : []
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
  repository = "https://prometheus-community.github.io/helm-charts"
  chart      = "kube-prometheus-stack"
  version    = var.chart_version == "" ? null : var.chart_version
  values     = concat(local.generated_values_yaml, var.values)

  depends_on = [kubernetes_namespace_v1.this]
}
