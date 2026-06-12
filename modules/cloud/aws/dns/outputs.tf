output "zone_id" {
  description = "Route53 hosted zone ID used by this module."
  value       = local.effective_zone_id
}

output "zone_name" {
  description = "Route53 hosted zone name used by this module."
  value       = local.effective_zone_name
}

output "name_servers" {
  description = "Name servers for created public hosted zones. Empty when using an existing zone."
  value       = var.create_zone ? aws_route53_zone.this[0].name_servers : []
}

output "record_fqdns" {
  description = "FQDNs of records created by this module, keyed by record key."
  value = {
    for key, record in aws_route53_record.this : key => record.fqdn
  }
}
