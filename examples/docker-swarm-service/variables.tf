variable "docker_host" {
  description = "Docker daemon endpoint. Use a Swarm manager for docker_service resources."
  type        = string
  default     = "unix:///var/run/docker.sock"
}
