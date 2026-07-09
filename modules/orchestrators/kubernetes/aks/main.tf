locals {
  kubernetes_version = var.kubernetes_version == "" ? null : var.kubernetes_version
  tags = merge(var.tags, {
    Environment = var.environment
    ManagedBy   = "clusterforge"
  })
}

resource "azurerm_kubernetes_cluster" "this" {
  name                = var.name
  location            = var.location
  resource_group_name = var.resource_group_name
  dns_prefix          = var.dns_prefix
  kubernetes_version  = local.kubernetes_version
  tags                = local.tags

  default_node_pool {
    name           = var.default_node_pool.name
    node_count     = var.default_node_pool.node_count
    vm_size        = var.default_node_pool.vm_size
    vnet_subnet_id = var.subnet_id
  }

  identity {
    type = "SystemAssigned"
  }
}
