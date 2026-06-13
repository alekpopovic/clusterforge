variable "namespace" {
  description = "Kubernetes namespace for the loki Helm release."
  type        = string
  default     = "logging"
}

variable "chart_version" {
  description = "Optional loki chart version. Leave empty to use the latest provider-resolved version."
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

variable "storage_enabled" {
  description = "Whether to enable persistent storage for Loki single binary mode."
  type        = bool
  default     = false
}

variable "storage_class_name" {
  description = "StorageClass name for Loki persistent storage. Leave empty to use the cluster default."
  type        = string
  default     = ""
}

variable "storage_size" {
  description = "Loki persistent volume request size when storage_enabled is true."
  type        = string
  default     = "20Gi"
}
