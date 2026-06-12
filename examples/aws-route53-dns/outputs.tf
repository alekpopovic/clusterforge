output "zone_id" {
  description = "Route53 hosted zone ID used by the example."
  value       = module.dns.zone_id
}

output "record_fqdns" {
  description = "FQDNs of records managed by the example."
  value       = module.dns.record_fqdns
}
