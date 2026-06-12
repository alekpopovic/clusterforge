variable "namespace" {
  description = "Kubernetes namespace for the External Secrets Operator Helm release."
  type        = string
  default     = "external-secrets"

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "Namespace must not be empty."
  }
}

variable "create_namespace" {
  description = "Whether to create the namespace before installing the Helm release."
  type        = bool
  default     = true
}

variable "chart_version" {
  description = "Optional external-secrets chart version. Leave empty to use the latest provider-resolved version."
  type        = string
  default     = ""
}

variable "values" {
  description = "YAML values passed to the Helm release."
  type        = list(string)
  default     = []
}

variable "labels" {
  description = "Labels applied to the namespace when create_namespace is true."
  type        = map(string)
  default     = {}
}
