locals {
  jobspec = <<-EOT
    job "${var.name}" {
      datacenters = ${jsonencode(var.datacenters)}
      type = "batch"
      group "${var.name}" {
        count = ${var.task_count}
        task "${var.name}" {
          driver = "docker"
          config {
            image = ${jsonencode(var.image)}
            args  = ${jsonencode(var.args)}
          }
          env = ${jsonencode(var.env)}
          resources {
            cpu    = ${var.cpu}
            memory = ${var.memory}
          }
        }
      }
    }
  EOT
}
resource "nomad_job" "this" {
  jobspec = local.jobspec
}
