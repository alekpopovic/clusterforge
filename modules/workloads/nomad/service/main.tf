locals {
  name         = lower(trimspace(var.name))
  service_name = try(var.service.name, null) == null ? local.name : var.service.name

  job_spec = templatefile("${path.module}/job.nomad.hcl.tftpl", {
    name            = local.name
    datacenters     = jsonencode(var.datacenters)
    namespace       = jsonencode(var.namespace)
    type            = jsonencode(var.type)
    image           = jsonencode(var.image)
    count           = var.task_count
    command         = jsonencode(var.command)
    command_enabled = var.command != ""
    args            = jsonencode(var.args)
    env = [
      for key, value in var.env : {
        key   = key
        value = jsonencode(value)
      }
    ]
    ports       = var.ports
    port_labels = jsonencode([for port in var.ports : port.label])
    cpu         = var.cpu
    memory      = var.memory
    service = {
      enabled    = var.service.enabled
      name       = jsonencode(local.service_name)
      port_label = jsonencode(var.service.port_label)
      tags       = jsonencode(var.service.tags)
    }
  })
}

resource "nomad_job" "this" {
  jobspec = local.job_spec
}
