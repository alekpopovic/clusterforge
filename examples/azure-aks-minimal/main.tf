module "network" {
  source = "../../modules/cloud/azure/network"

  name                = var.name
  environment         = "dev"
  location            = var.location
  resource_group_name = var.resource_group_name
}

module "aks" {
  source = "../../modules/orchestrators/kubernetes/aks"

  name                = var.name
  environment         = "dev"
  location            = var.location
  resource_group_name = module.network.resource_group_name
  subnet_id           = module.network.subnet_ids[0]
  dns_prefix          = var.name
}
