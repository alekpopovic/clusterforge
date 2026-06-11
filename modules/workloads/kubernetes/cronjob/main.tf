locals {
  name      = trimspace(var.name)
  namespace = trimspace(var.namespace)

  labels = merge(
    {
      "app.kubernetes.io/name"       = local.name
      "app.kubernetes.io/instance"   = local.name
      "app.kubernetes.io/component"  = "cronjob"
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

resource "kubernetes_cron_job_v1" "this" {
  metadata {
    name        = local.name
    namespace   = local.namespace
    labels      = local.labels
    annotations = var.annotations
  }

  spec {
    schedule                      = var.schedule
    concurrency_policy            = var.concurrency_policy
    successful_jobs_history_limit = var.successful_jobs_history_limit
    failed_jobs_history_limit     = var.failed_jobs_history_limit

    job_template {
      metadata {
        labels      = local.labels
        annotations = var.annotations
      }

      spec {
        template {
          metadata {
            labels      = local.labels
            annotations = var.annotations
          }

          spec {
            restart_policy = var.restart_policy

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
    }
  }

  depends_on = [kubernetes_namespace_v1.this]
}
