output "endpoint_ids" {
  description = "VPC endpoint IDs keyed by service suffix."
  value = merge(
    { for service, endpoint in aws_vpc_endpoint.gateway : service => endpoint.id },
    { for service, endpoint in aws_vpc_endpoint.interface : service => endpoint.id }
  )
}

output "endpoint_dns_entries" {
  description = "Interface endpoint DNS entries keyed by service suffix. Gateway endpoints do not have DNS entries."
  value = {
    for service, endpoint in aws_vpc_endpoint.interface : service => endpoint.dns_entry
  }
}

output "security_group_id" {
  description = "Created interface endpoint security group ID, or null when not created."
  value       = local.create_interface_security_group ? aws_security_group.interface_endpoints[0].id : null
}
