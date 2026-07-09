## Prompt 14 — Nomad job modul

```text
Implement modules/workloads/nomad/service.

Purpose:
Deploy a service-style Nomad job using Terraform Nomad provider.

Provider:
- hashicorp/nomad.
- Provider configured in root.

Inputs:
- name: string
- datacenters: list(string)
- namespace: string, default "default"
- type: string, default "service"
- image: string
- count: number, default 1
- command: string, default ""
- args: list(string), default []
- env: map(string), default {}
- ports: list(object({
    label = string
    to = number
    static = optional(number)
  })), default []
- cpu: number, default 500
- memory: number, default 256
- service: object({
    enabled = bool
    name = optional(string)
    port_label = optional(string, "http")
    tags = optional(list(string), [])
  }), default enabled false

Resource:
- nomad_job

Implementation:
- Generate the Nomad job specification using templatefile or heredoc.
- Keep the template readable.
- Use Docker driver.
- Include network ports.
- Include service block only when enabled.
- Include env vars.
- Include resources.

Outputs:
- job_id
- job_name

README:
- Explain Nomad use case.
- Show example service with Docker image.
- Show ports and service registration example.
- Mention that Consul integration may be needed for service discovery.

Run terraform fmt -recursive.
```

---
