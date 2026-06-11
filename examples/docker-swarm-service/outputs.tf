output "service_id" {
  description = "Docker Swarm service ID."
  value       = module.service.service_id
}

output "service_name" {
  description = "Docker Swarm service name."
  value       = module.service.service_name
}
