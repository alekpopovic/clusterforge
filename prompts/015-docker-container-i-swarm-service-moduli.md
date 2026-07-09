## Prompt 15 — Docker container i Swarm service moduli

```text
Implement Docker workload modules:

1. modules/workloads/docker/container
2. modules/workloads/docker/swarm-service

Provider:
- kreuzwerker/docker.
- Provider configured in root.

modules/workloads/docker/container:
Inputs:
- name
- image
- command list(string)
- env map(string)
- ports list(object({
    internal = number
    external = optional(number)
    protocol = optional(string, "tcp")
  }))
- volumes list(object({
    host_path = string
    container_path = string
    read_only = optional(bool, false)
  }))
- restart_policy string default "unless-stopped"
- networks list(string) default []
- labels map(string) default {}

Resources:
- docker_image
- docker_container

Outputs:
- container_id
- container_name
- image_id

modules/workloads/docker/swarm-service:
Inputs:
- name
- image
- replicas number default 1
- env map(string)
- ports list(object({
    target_port = number
    published_port = number
    protocol = optional(string, "tcp")
  }))
- networks list(string) default []
- labels map(string) default {}

Resources:
- docker_image
- docker_service

Outputs:
- service_id
- service_name

README for each:
- Explain local/self-hosted use case.
- Explain this is not recommended as the main production target compared to Kubernetes/ECS/Nomad.
- Add simple examples.

Run terraform fmt -recursive.
```

---
