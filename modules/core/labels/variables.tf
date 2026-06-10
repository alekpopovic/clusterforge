variable "project" {
  description = "Project name used for ClusterForge Kubernetes labels."
  type        = string

  validation {
    condition     = length(trimspace(var.project)) > 0
    error_message = "Project must not be empty."
  }
}

variable "environment" {
  description = "Environment name used for ClusterForge Kubernetes labels."
  type        = string

  validation {
    condition     = length(trimspace(var.environment)) > 0
    error_message = "Environment must not be empty."
  }
}

variable "app" {
  description = "Optional app.kubernetes.io/name label value."
  type        = string
  default     = ""
}

variable "component" {
  description = "Optional app.kubernetes.io/component label value."
  type        = string
  default     = ""
}

variable "part_of" {
  description = "Optional app.kubernetes.io/part-of label value."
  type        = string
  default     = ""
}

variable "managed_by" {
  description = "app.kubernetes.io/managed-by label value."
  type        = string
  default     = "terraform"

  validation {
    condition     = length(trimspace(var.managed_by)) > 0
    error_message = "Managed_by must not be empty."
  }
}

variable "extra_labels" {
  description = "Additional Kubernetes labels. These are sanitized and merged last."
  type        = map(string)
  default     = {}

  validation {
    condition = alltrue([
      for key in keys(var.extra_labels) :
      length(key) <= 317 &&
      length(element(reverse(split("/", key)), 0)) <= 63 &&
      can(regex("^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?[A-Za-z0-9]([-A-Za-z0-9_.]*[A-Za-z0-9])?$", key))
    ])
    error_message = "Each extra_labels key must be a valid Kubernetes label key."
  }
}
