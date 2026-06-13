variable "name" {
  description = "Helm release name."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "name must not be empty."
  }
}

variable "namespace" {
  description = "Kubernetes namespace for the Helm release."
  type        = string

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "namespace must not be empty."
  }
}

variable "create_namespace" {
  description = "Whether to create the namespace before installing the Helm release."
  type        = bool
  default     = true
}

variable "repository" {
  description = "Helm chart repository URL."
  type        = string

  validation {
    condition     = length(trimspace(var.repository)) > 0
    error_message = "repository must not be empty."
  }
}

variable "chart" {
  description = "Helm chart name."
  type        = string

  validation {
    condition     = length(trimspace(var.chart)) > 0
    error_message = "chart must not be empty."
  }
}

variable "chart_version" {
  description = "Optional Helm chart version. Pin this before production use."
  type        = string
  default     = ""
}

variable "values" {
  description = "YAML values passed to the Helm release."
  type        = list(string)
  default     = []
}

variable "set" {
  description = "Non-sensitive Helm set values."
  type        = map(string)
  default     = {}
}

variable "set_sensitive" {
  description = "Sensitive Helm set values. Values may still be retained in Terraform state; prefer existing secrets where possible."
  type        = map(string)
  default     = {}
  sensitive   = true
}

variable "labels" {
  description = "Labels applied to the namespace when create_namespace is true."
  type        = map(string)
  default     = {}
}

variable "timeout" {
  description = "Time in seconds to wait for Helm operations."
  type        = number
  default     = 300

  validation {
    condition     = var.timeout > 0
    error_message = "timeout must be greater than 0."
  }
}

variable "atomic" {
  description = "Whether Helm should roll back changes on failed install or upgrade."
  type        = bool
  default     = false
}

variable "cleanup_on_fail" {
  description = "Whether Helm should delete newly-created resources when an install or upgrade fails."
  type        = bool
  default     = true
}

variable "wait" {
  description = "Whether Terraform should wait for Helm resources to become ready."
  type        = bool
  default     = true
}
