locals {
  name = trimspace(var.name)

  env = [
    for key, value in var.env : "${key}=${value}"
  ]
}

resource "docker_image" "this" {
  name = var.image
}

resource "docker_container" "this" {
  name  = local.name
  image = docker_image.this.image_id

  command = var.command
  env     = local.env

  restart = var.restart_policy

  dynamic "ports" {
    for_each = var.ports

    content {
      internal = ports.value.internal
      external = ports.value.external
      protocol = lower(ports.value.protocol)
    }
  }

  dynamic "volumes" {
    for_each = var.volumes

    content {
      host_path      = volumes.value.host_path
      container_path = volumes.value.container_path
      read_only      = volumes.value.read_only
    }
  }

  dynamic "networks_advanced" {
    for_each = var.networks

    content {
      name = networks_advanced.value
    }
  }

  dynamic "labels" {
    for_each = var.labels

    content {
      label = labels.key
      value = labels.value
    }
  }
}
