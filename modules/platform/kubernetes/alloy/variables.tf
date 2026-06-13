variable "namespace" {
  description = "Kubernetes namespace for the Grafana Alloy Helm release."
  type        = string
  default     = "logging"
}

variable "chart_version" {
  description = "Optional Grafana Alloy chart version. Pin this before production use."
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

variable "create_namespace" {
  description = "Whether to create the namespace before installing the Helm release."
  type        = bool
  default     = true
}
