locals {
  name      = trimspace(var.name)
  namespace = trimspace(var.namespace)
  selector_labels = {
    "app.kubernetes.io/name"     = local.name
    "app.kubernetes.io/instance" = local.name
  }
  labels = merge(local.selector_labels, var.labels, {
    "app.kubernetes.io/managed-by" = "terraform"
    "clusterforge.io/workload"     = "rollout"
  })
  stable_service_name  = local.name
  preview_service_name = var.strategy.type == "canary" ? "${local.name}-canary" : "${local.name}-preview"
  canary_steps = [for step in var.strategy.canary_steps :
    step.set_weight != null ? { setWeight = step.set_weight } : {
      pause                               = step.pause_duration != null ? { duration = step.pause_duration } : {}
    }
  ]
  strategy = var.strategy.type == "canary" ? {
    canary = {
      stableService = local.stable_service_name
      canaryService = local.preview_service_name
      steps         = local.canary_steps
    }
    } : {
    blueGreen = merge({
      activeService         = local.stable_service_name
      previewService        = local.preview_service_name
      autoPromotionEnabled  = var.strategy.blue_green_config.auto_promotion_enabled
      scaleDownDelaySeconds = var.strategy.blue_green_config.scale_down_delay_seconds
      }, var.strategy.blue_green_config.preview_replica_count == null ? {} : {
      previewReplicaCount = var.strategy.blue_green_config.preview_replica_count
    })
  }
  container = merge({
    name  = local.name
    image = var.image
    ports = [{ name = "http", containerPort = var.port, protocol = "TCP" }]
    }, length(var.env) == 0 ? {} : {
    env = [for name, value in var.env : { name = name, value = value }]
    }, length(var.resources.requests) == 0 && length(var.resources.limits) == 0 ? {} : {
    resources = {
      requests = var.resources.requests
      limits   = var.resources.limits
    }
  })
}

resource "kubernetes_namespace_v1" "this" {
  count = var.create_namespace ? 1 : 0
  metadata {
    name   = local.namespace
    labels = local.labels
  }
}

resource "kubernetes_manifest" "rollout" {
  manifest = {
    apiVersion = "argoproj.io/v1alpha1"
    kind       = "Rollout"
    metadata = {
      name        = local.name
      namespace   = local.namespace
      labels      = local.labels
      annotations = var.annotations
    }
    spec = {
      replicas = var.replicas
      selector = { matchLabels = local.selector_labels }
      template = {
        metadata = {
          labels      = local.labels
          annotations = var.annotations
        }
        spec = { containers = [local.container] }
      }
      strategy = local.strategy
    }
  }
  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_service_v1" "stable" {
  metadata {
    name      = local.stable_service_name
    namespace = local.namespace
    labels    = local.labels
  }
  spec {
    selector = local.selector_labels
    port {
      name        = "http"
      port        = 80
      target_port = "http"
      protocol    = "TCP"
    }
  }
  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_service_v1" "preview" {
  metadata {
    name      = local.preview_service_name
    namespace = local.namespace
    labels    = local.labels
  }
  spec {
    selector = local.selector_labels
    port {
      name        = "http"
      port        = 80
      target_port = "http"
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
    annotations = var.ingress.annotations
  }
  spec {
    ingress_class_name = try(var.ingress.class_name, null)
    dynamic "tls" {
      for_each = var.ingress.tls ? [1] : []
      content {
        hosts       = [var.ingress.host]
        secret_name = "${local.name}-tls"
      }
    }
    rule {
      host = var.ingress.host
      http {
        path {
          path      = "/"
          path_type = "Prefix"
          backend {
            service {
              name = kubernetes_service_v1.stable.metadata[0].name
              port { name = "http" }
            }
          }
        }
      }
    }
  }
}
