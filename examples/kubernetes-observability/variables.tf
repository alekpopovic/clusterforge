variable "kubeconfig_path" {
  description = "Path to a local kubeconfig for an existing Kubernetes cluster."
  type        = string
  default     = "~/.kube/config"
}

variable "grafana_host" {
  description = "Optional Grafana ingress hostname. Leave empty to keep ingress disabled."
  type        = string
  default     = ""
}

variable "storage_class_name" {
  description = "Optional StorageClass for Prometheus and Loki persistence."
  type        = string
  default     = ""
}

variable "enable_persistent_storage" {
  description = "Whether to enable persistent storage for Prometheus and Loki."
  type        = bool
  default     = false
}

variable "alloy_values" {
  description = "YAML values passed to Grafana Alloy. Provide explicit production log collection config here."
  type        = list(string)
  default     = []
}
