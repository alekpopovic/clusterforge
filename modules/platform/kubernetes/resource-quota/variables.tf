variable "namespace" {
  description = "Existing Kubernetes namespace where the ResourceQuota is created."
  type        = string

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "namespace must not be empty."
  }
}

variable "hard" {
  description = "Hard namespace resource limits using Kubernetes ResourceQuota keys and quantity strings."
  type        = map(string)

  validation {
    condition     = length(var.hard) > 0 && alltrue([for key, value in var.hard : length(trimspace(key)) > 0 && length(trimspace(value)) > 0])
    error_message = "hard must contain at least one non-empty resource name and quantity."
  }
}

variable "labels" {
  description = "Additional labels applied to the ResourceQuota."
  type        = map(string)
  default     = {}
}
