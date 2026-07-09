variable "name" {
  description = "Kubernetes workload name."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "Name must not be empty."
  }
}

variable "namespace" {
  description = "Kubernetes namespace for the workload."
  type        = string

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "Namespace must not be empty."
  }
}

variable "create_namespace" {
  description = "Whether to create the namespace."
  type        = bool
  default     = true
}

variable "labels" {
  description = "Additional labels applied to workload resources."
  type        = map(string)
  default     = {}
}

variable "annotations" {
  description = "Annotations applied to workload resources."
  type        = map(string)
  default     = {}
}

variable "image" {
  description = "Container image reference."
  type        = string

  validation {
    condition     = length(trimspace(var.image)) > 0
    error_message = "Image must not be empty."
  }
}

variable "image_pull_policy" {
  description = "Image pull policy for the container."
  type        = string
  default     = "IfNotPresent"

  validation {
    condition     = contains(["Always", "IfNotPresent", "Never"], var.image_pull_policy)
    error_message = "Image pull policy must be Always, IfNotPresent, or Never."
  }
}

variable "image_pull_secrets" {
  description = "Names of existing image pull secrets."
  type        = list(string)
  default     = []
}

variable "replicas" {
  description = "Desired number of replicas when autoscaling is disabled."
  type        = number
  default     = 1

  validation {
    condition     = var.replicas >= 0
    error_message = "Replicas must be greater than or equal to 0."
  }
}

variable "command" {
  description = "Optional container command."
  type        = list(string)
  default     = []
}

variable "args" {
  description = "Optional container args."
  type        = list(string)
  default     = []
}

variable "env" {
  description = "Plain environment variables. Do not put secrets here."
  type        = map(string)
  default     = {}
}

variable "secret_env" {
  description = "Environment variables sourced from existing Kubernetes secrets."
  type = map(object({
    secret_name = string
    secret_key  = string
  }))
  default = {}
}

variable "ports" {
  description = "Container ports exposed by the workload."
  type = list(object({
    name           = string
    container_port = number
    protocol       = optional(string, "TCP")
  }))
  default = []
}

variable "resources" {
  description = "Container resource requests and limits."
  type = object({
    cpu_request    = optional(string)
    memory_request = optional(string)
    cpu_limit      = optional(string)
    memory_limit   = optional(string)
  })
  default = {}
}

variable "liveness_probe" {
  description = "Optional HTTP liveness probe."
  type = object({
    path                  = string
    port                  = string
    initial_delay_seconds = number
    period_seconds        = number
    timeout_seconds       = number
    failure_threshold     = number
  })
  default = null
}

variable "readiness_probe" {
  description = "Optional HTTP readiness probe."
  type = object({
    path                  = string
    port                  = string
    initial_delay_seconds = number
    period_seconds        = number
    timeout_seconds       = number
    failure_threshold     = number
  })
  default = null
}

variable "service" {
  description = "Service configuration."
  type = object({
    enabled          = optional(bool, true)
    type             = optional(string, "ClusterIP")
    port             = optional(number, 80)
    target_port_name = optional(string, "http")
  })
  default = {}

  validation {
    condition     = contains(["ClusterIP", "NodePort", "LoadBalancer"], var.service.type)
    error_message = "Service type must be ClusterIP, NodePort, or LoadBalancer."
  }
}

variable "ingress" {
  description = "Ingress configuration."
  type = object({
    enabled         = bool
    class_name      = optional(string)
    host            = optional(string)
    path            = optional(string, "/")
    path_type       = optional(string, "Prefix")
    tls             = optional(bool, true)
    tls_secret_name = optional(string)
    annotations     = optional(map(string), {})
  })
  default = {
    enabled = false
  }

  validation {
    condition     = !var.ingress.enabled || try(length(trimspace(var.ingress.host)) > 0, false)
    error_message = "Ingress host is required when ingress is enabled."
  }

  validation {
    condition     = !var.ingress.enabled || var.service.enabled
    error_message = "Service must be enabled when ingress is enabled."
  }
}

variable "autoscaling" {
  description = "Horizontal pod autoscaling configuration."
  type = object({
    enabled      = bool
    min_replicas = optional(number, 1)
    max_replicas = optional(number, 3)
    cpu_percent  = optional(number, 70)
  })
  default = {
    enabled = false
  }

  validation {
    condition     = !var.autoscaling.enabled || var.autoscaling.min_replicas <= var.autoscaling.max_replicas
    error_message = "Autoscaling min_replicas must be less than or equal to max_replicas."
  }
}

variable "service_account" {
  description = "Optional service account configuration for the workload."
  type = object({
    create          = optional(bool, false)
    name            = optional(string)
    annotations     = optional(map(string), {})
    labels          = optional(map(string), {})
    automount_token = optional(bool)
  })
  default = {}
}

variable "rbac" {
  description = "Optional namespace-scoped RBAC Role and RoleBinding rules for the workload service account."
  type = object({
    create = optional(bool, false)
    rules = optional(list(object({
      api_groups     = list(string)
      resources      = list(string)
      verbs          = list(string)
      resource_names = optional(list(string))
    })), [])
  })
  default = {}

  validation {
    condition     = !var.rbac.create || var.service_account.create || try(length(trimspace(var.service_account.name)) > 0, false)
    error_message = "RBAC requires a created service account or an existing service_account.name."
  }
}
