provider "docker" {
  host = var.docker_host
}

module "service" {
  source = "../../modules/workloads/docker/swarm-service"

  name     = "hello"
  image    = "nginx:1.27"
  replicas = 1

  ports = [
    {
      target_port    = 80
      published_port = 8080
    }
  ]

  labels = {
    "clusterforge.io/example" = "docker-swarm-service"
  }
}
