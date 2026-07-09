output "resource_group_name" {
  description = "Resource group used by the network."
  value       = var.resource_group_name
}

output "vnet_id" {
  description = "Azure virtual network ID."
  value       = azurerm_virtual_network.this.id
}

output "subnet_ids" {
  description = "Azure subnet IDs."
  value       = azurerm_subnet.this[*].id
}
