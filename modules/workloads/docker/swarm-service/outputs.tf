output "service_id" {
  description = "Docker Swarm service ID."
  value       = docker_service.this.id
}

output "service_name" {
  description = "Docker Swarm service name."
  value       = docker_service.this.name
}
