locals {
  tags = merge(var.tags, {
    Environment = var.environment
    ManagedBy   = "clusterforge"
  })
}

resource "azurerm_resource_group" "this" {
  count    = var.create_resource_group ? 1 : 0
  name     = var.resource_group_name
  location = var.location
  tags     = local.tags
}

resource "azurerm_virtual_network" "this" {
  name                = "${var.name}-vnet"
  location            = var.location
  resource_group_name = var.resource_group_name
  address_space       = var.address_space
  tags                = local.tags

  depends_on = [azurerm_resource_group.this]
}

resource "azurerm_subnet" "this" {
  count                = length(var.subnet_prefixes)
  name                 = "${var.name}-subnet-${count.index + 1}"
  resource_group_name  = var.resource_group_name
  virtual_network_name = azurerm_virtual_network.this.name
  address_prefixes     = [var.subnet_prefixes[count.index]]
}
