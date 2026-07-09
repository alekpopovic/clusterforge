## Prompt 10 — Kubernetes workload app modul

```text
Implement modules/workloads/kubernetes/app.

Purpose:
Deploy a generic web/service workload to Kubernetes using Terraform Kubernetes provider.

Provider:
- Use hashicorp/kubernetes.
- Provider must be configured in root.

Inputs:
- name: string
- namespace: string
- create_namespace: bool, default true
- labels: map(string), default {}
- annotations: map(string), default {}

Image:
- image: string
- image_pull_policy: string, default "IfNotPresent"
- image_pull_secrets: list(string), default []

Runtime:
- replicas: number, default 1
- command: list(string), default []
- args: list(string), default []
- env: map(string), default {}
- secret_env: map(object({
    secret_name = string
    secret_key  = string
  })), default {}

Ports:
- ports: list(object({
    name = string
    container_port = number
    protocol = optional(string, "TCP")
  })), default []

Resources:
- resources: object({
    cpu_request = optional(string)
    memory_request = optional(string)
    cpu_limit = optional(string)
    memory_limit = optional(string)
  }), default {}

Probes:
- liveness_probe: optional object
- readiness_probe: optional object
Keep probe schema simple:
  path
  port
  initial_delay_seconds
  period_seconds
  timeout_seconds
  failure_threshold

Service:
- service: object({
    enabled = optional(bool, true)
    type = optional(string, "ClusterIP")
    port = optional(number, 80)
    target_port_name = optional(string, "http")
  }), default enabled

Ingress:
- ingress: object({
    enabled = bool
    class_name = optional(string)
    host = optional(string)
    path = optional(string, "/")
    path_type = optional(string, "Prefix")
    tls = optional(bool, true)
    tls_secret_name = optional(string)
    annotations = optional(map(string), {})
  }), default disabled

Autoscaling:
- autoscaling: object({
    enabled = bool
    min_replicas = optional(number, 1)
    max_replicas = optional(number, 3)
    cpu_percent = optional(number, 70)
  }), default disabled

Resources to create:
- kubernetes_namespace_v1 when create_namespace=true
- kubernetes_deployment_v1
- kubernetes_service_v1 when service.enabled=true
- kubernetes_ingress_v1 when ingress.enabled=true
- kubernetes_horizontal_pod_autoscaler_v2 when autoscaling.enabled=true

Behavior:
- Add app.kubernetes.io labels.
- Merge custom labels.
- Avoid putting secret values in Terraform. Only reference existing Kubernetes secrets.
- Support multiple ports.
- Ensure service selectors match pod labels.
- If ingress TLS enabled and tls_secret_name not set, generate secret name from app name.

Validation:
- name and namespace non-empty.
- replicas >= 0.
- image non-empty.
- service type in ClusterIP, NodePort, LoadBalancer.
- ingress host required when ingress.enabled=true.
- autoscaling min <= max.

Outputs:
- name
- namespace
- deployment_name
- service_name
- ingress_name
- labels

README:
- Include basic deployment example.
- Include ingress example.
- Include secret_env example.
- Include autoscaling example.
- Explain when to use GitOps instead.

Create examples/kubernetes-basic-app using this module.

Run terraform fmt -recursive.
```

---
