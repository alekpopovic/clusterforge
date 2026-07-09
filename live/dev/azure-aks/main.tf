module "network" {
  source = "../../../modules/cloud/azure/network"

  name                = var.name
  environment         = var.environment
  location            = var.location
  resource_group_name = var.resource_group_name
  address_space       = var.address_space
  subnet_prefixes     = var.subnet_prefixes
  tags                = var.tags
}

module "aks" {
  source = "../../../modules/orchestrators/kubernetes/aks"

  name                = var.name
  environment         = var.environment
  location            = var.location
  resource_group_name = module.network.resource_group_name
  subnet_id           = module.network.subnet_ids[0]
  kubernetes_version  = var.kubernetes_version
  dns_prefix          = var.dns_prefix
  default_node_pool   = var.default_node_pool
  tags                = var.tags
}
