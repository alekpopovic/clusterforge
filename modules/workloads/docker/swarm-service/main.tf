locals {
  name = trimspace(var.name)

  env = [
    for key, value in var.env : "${key}=${value}"
  ]
}

resource "docker_image" "this" {
  name = var.image
}

resource "docker_service" "this" {
  name = local.name

  task_spec {
    container_spec {
      image = docker_image.this.repo_digest
      env   = local.env

      dynamic "labels" {
        for_each = var.labels

        content {
          label = labels.key
          value = labels.value
        }
      }
    }

    dynamic "networks_advanced" {
      for_each = var.networks

      content {
        name = networks_advanced.value
      }
    }
  }

  mode {
    replicated {
      replicas = var.replicas
    }
  }

  endpoint_spec {
    dynamic "ports" {
      for_each = var.ports

      content {
        target_port    = ports.value.target_port
        published_port = ports.value.published_port
        protocol       = lower(ports.value.protocol)
      }
    }
  }
}
