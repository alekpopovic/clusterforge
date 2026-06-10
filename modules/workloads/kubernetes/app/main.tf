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
      "app.kubernetes.io/component"  = "app"
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
  ingress_tls_secret_name = try(var.ingress.tls_secret_name, null) == null ? (
    "${local.name}-tls"
  ) : var.ingress.tls_secret_name
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

          dynamic "port" {
            for_each = var.ports

            content {
              name           = port.value.name
              container_port = port.value.container_port
              protocol       = port.value.protocol
            }
          }

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

          dynamic "liveness_probe" {
            for_each = var.liveness_probe == null ? [] : [var.liveness_probe]

            content {
              initial_delay_seconds = liveness_probe.value.initial_delay_seconds
              period_seconds        = liveness_probe.value.period_seconds
              timeout_seconds       = liveness_probe.value.timeout_seconds
              failure_threshold     = liveness_probe.value.failure_threshold

              http_get {
                path = liveness_probe.value.path
                port = liveness_probe.value.port
              }
            }
          }

          dynamic "readiness_probe" {
            for_each = var.readiness_probe == null ? [] : [var.readiness_probe]

            content {
              initial_delay_seconds = readiness_probe.value.initial_delay_seconds
              period_seconds        = readiness_probe.value.period_seconds
              timeout_seconds       = readiness_probe.value.timeout_seconds
              failure_threshold     = readiness_probe.value.failure_threshold

              http_get {
                path = readiness_probe.value.path
                port = readiness_probe.value.port
              }
            }
          }
        }
      }
    }
  }

  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_service_v1" "this" {
  count = var.service.enabled ? 1 : 0

  metadata {
    name        = local.name
    namespace   = local.namespace
    labels      = local.labels
    annotations = var.annotations
  }

  spec {
    type     = var.service.type
    selector = local.selector_labels

    port {
      name        = var.service.target_port_name
      port        = var.service.port
      target_port = var.service.target_port_name
      protocol    = "TCP"
    }
  }

  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_ingress_v1" "this" {
  count = var.ingress.enabled ? 1 : 0

  metadata {
    name        = local.name
    namespace   = local.namespace
    labels      = local.labels
    annotations = merge(var.annotations, var.ingress.annotations)
  }

  spec {
    ingress_class_name = try(var.ingress.class_name, null)

    dynamic "tls" {
      for_each = var.ingress.tls ? [1] : []

      content {
        hosts       = [var.ingress.host]
        secret_name = local.ingress_tls_secret_name
      }
    }

    rule {
      host = var.ingress.host

      http {
        path {
          path      = var.ingress.path
          path_type = var.ingress.path_type

          backend {
            service {
              name = kubernetes_service_v1.this[0].metadata[0].name

              port {
                number = var.service.port
              }
            }
          }
        }
      }
    }
  }

  depends_on = [kubernetes_service_v1.this]
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
