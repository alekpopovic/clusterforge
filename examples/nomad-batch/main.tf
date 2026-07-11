provider "nomad" {
  address = var.nomad_address
}
module "backup" {
  source = "../../modules/workloads/nomad/batch"
  name   = "backup"
  image  = "alpine:3.20"
  args   = ["sh", "-c", "echo replace-with-a-real-reviewed-job"]
}
