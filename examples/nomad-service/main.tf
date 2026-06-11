provider "nomad" {
  address = var.nomad_address
  region  = var.nomad_region
}

module "service" {
  source = "../../modules/workloads/nomad/service"

  name        = "hello"
  datacenters = var.datacenters
  namespace   = var.namespace
  image       = "nginx:1.27"
  task_count  = 1

  ports = [
    {
      label = "http"
      to    = 80
    }
  ]

  service = {
    enabled    = true
    name       = "hello"
    port_label = "http"
    tags       = ["clusterforge-example"]
  }
}
