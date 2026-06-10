locals {
  labels = merge({
    "app.kubernetes.io/name"       = var.name
    "app.kubernetes.io/managed-by" = "terraform"
  }, var.labels)
}

resource "kubernetes_deployment_v1" "this" {
  metadata {
    name        = var.name
    namespace   = var.namespace
    labels      = local.labels
    annotations = var.annotations
  }

  spec {
    replicas = var.replicas

    selector {
      match_labels = {
        "app.kubernetes.io/name" = var.name
      }
    }

    template {
      metadata {
        labels = local.labels
      }

      spec {
        container {
          name  = var.name
          image = var.image

          port {
            container_port = var.container_port
          }

          dynamic "env" {
            for_each = var.env

            content {
              name  = env.key
              value = env.value
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service_v1" "this" {
  metadata {
    name        = var.name
    namespace   = var.namespace
    labels      = local.labels
    annotations = var.annotations
  }

  spec {
    selector = {
      "app.kubernetes.io/name" = var.name
    }

    port {
      port        = var.service_port
      target_port = var.container_port
    }
  }
}
