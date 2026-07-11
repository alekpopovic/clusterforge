module "cluster_config" {
  source           = "../../modules/orchestrators/nomad/cluster"
  name             = "existing-nomad"
  environment      = "dev"
  server_addresses = ["10.0.0.10:4647"]
}
