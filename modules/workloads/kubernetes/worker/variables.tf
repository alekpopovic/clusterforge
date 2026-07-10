variable "name" {
  description = "Kubernetes worker workload name."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "Name must not be empty."
  }
}

variable "namespace" {
  description = "Kubernetes namespace for the worker."
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

variable "image" {
  description = "Container image reference."
  type        = string

  validation {
    condition     = length(trimspace(var.image)) > 0
    error_message = "Image must not be empty."
  }
}

variable "replicas" {
  description = "Desired number of worker replicas when autoscaling is disabled."
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

variable "labels" {
  description = "Additional labels applied to worker resources."
  type        = map(string)
  default     = {}
}

variable "annotations" {
  description = "Annotations applied to the Deployment and pod template."
  type        = map(string)
  default     = {}
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

variable "termination_grace_period_seconds" {
  description = "Optional pod termination grace period in seconds."
  type        = number
  default     = null

  validation {
    condition     = var.termination_grace_period_seconds == null || var.termination_grace_period_seconds >= 0
    error_message = "termination_grace_period_seconds must be greater than or equal to 0."
  }
}

variable "service_account_name" {
  description = "Optional existing Kubernetes service account name for the worker pods."
  type        = string
  default     = ""
}

variable "service_account_annotations" {
  description = "Annotations for an optional workload service account, including EKS IRSA role annotations."
  type        = map(string)
  default     = {}
}

variable "autoscaling" {
  description = "Optional Horizontal Pod Autoscaler configuration."
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

  validation {
    condition     = !var.autoscaling.enabled || var.autoscaling.max_replicas > 0
    error_message = "Autoscaling max_replicas must be greater than 0."
  }
}
