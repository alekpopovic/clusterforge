locals {
  name      = trimspace(var.name)
  namespace = trimspace(var.namespace)

  selector_labels = {
    "app.kubernetes.io/name"     = local.name
    "app.kubernetes.io/instance" = local.name
  }

  labels = merge(
    {
      "app.kubernetes.io/name"       = local.name
      "app.kubernetes.io/instance"   = local.name
      "app.kubernetes.io/component"  = "worker"
      "app.kubernetes.io/managed-by" = "terraform"
    },
    var.labels
  )

  resource_requests = {
    for key, value in {
      cpu    = try(var.resources.cpu_request, null)
      memory = try(var.resources.memory_request, null)
    } : key => value if value != null
  }

  resource_limits = {
    for key, value in {
      cpu    = try(var.resources.cpu_limit, null)
      memory = try(var.resources.memory_limit, null)
    } : key => value if value != null
  }

  has_resources = length(local.resource_requests) > 0 || length(local.resource_limits) > 0
}

resource "kubernetes_namespace_v1" "this" {
  count = var.create_namespace ? 1 : 0

  metadata {
    name   = local.namespace
    labels = local.labels
  }
}

resource "kubernetes_deployment_v1" "this" {
  metadata {
    name        = local.name
    namespace   = local.namespace
    labels      = local.labels
    annotations = var.annotations
  }

  spec {
    replicas = var.autoscaling.enabled ? null : var.replicas

    selector {
      match_labels = local.selector_labels
    }

    template {
      metadata {
        labels      = local.labels
        annotations = var.annotations
      }

      spec {
        service_account_name             = var.service_account_name == "" ? null : var.service_account_name
        termination_grace_period_seconds = var.termination_grace_period_seconds

        dynamic "image_pull_secrets" {
          for_each = toset(var.image_pull_secrets)

          content {
            name = image_pull_secrets.value
          }
        }

        container {
          name              = local.name
          image             = var.image
          image_pull_policy = var.image_pull_policy
          command           = var.command
          args              = var.args

          dynamic "env" {
            for_each = var.env

            content {
              name  = env.key
              value = env.value
            }
          }

          dynamic "env" {
            for_each = var.secret_env

            content {
              name = env.key

              value_from {
                secret_key_ref {
                  name = env.value.secret_name
                  key  = env.value.secret_key
                }
              }
            }
          }

          dynamic "resources" {
            for_each = local.has_resources ? [1] : []

            content {
              limits   = local.resource_limits
              requests = local.resource_requests
            }
          }
        }
      }
    }
  }

  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_horizontal_pod_autoscaler_v2" "this" {
  count = var.autoscaling.enabled ? 1 : 0

  metadata {
    name      = local.name
    namespace = local.namespace
    labels    = local.labels
  }

  spec {
    min_replicas = var.autoscaling.min_replicas
    max_replicas = var.autoscaling.max_replicas

    scale_target_ref {
      api_version = "apps/v1"
      kind        = "Deployment"
      name        = kubernetes_deployment_v1.this.metadata[0].name
    }

    metric {
      type = "Resource"

      resource {
        name = "cpu"

        target {
          type                = "Utilization"
          average_utilization = var.autoscaling.cpu_percent
        }
      }
    }
  }

  depends_on = [kubernetes_namespace_v1.this]
}
