variable "namespace" {
  description = "Existing Kubernetes namespace where the LimitRange is created."
  type        = string

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "namespace must not be empty."
  }
}

variable "limits" {
  description = "LimitRange entries for containers, pods, or persistent volume claims."
  type = list(object({
    type            = string
    default         = optional(map(string), {})
    default_request = optional(map(string), {})
    max             = optional(map(string), {})
    min             = optional(map(string), {})
  }))

  validation {
    condition     = length(var.limits) > 0 && alltrue([for limit in var.limits : contains(["Container", "Pod", "PersistentVolumeClaim"], limit.type)])
    error_message = "limits must contain at least one entry with type Container, Pod, or PersistentVolumeClaim."
  }
}

variable "labels" {
  description = "Additional labels applied to the LimitRange."
  type        = map(string)
  default     = {}
}
