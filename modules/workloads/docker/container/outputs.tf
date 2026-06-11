output "container_id" {
  description = "Docker container ID."
  value       = docker_container.this.id
}

output "container_name" {
  description = "Docker container name."
  value       = docker_container.this.name
}

output "image_id" {
  description = "Docker image ID used by the container."
  value       = docker_image.this.image_id
}
