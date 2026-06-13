variable "namespace" {
  description = "Kubernetes namespace for the kube-prometheus-stack Helm release."
  type        = string
  default     = "monitoring"
}

variable "chart_version" {
  description = "Optional kube-prometheus-stack chart version. Leave empty to use the latest provider-resolved version."
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

variable "enable_grafana_ingress" {
  description = "Whether to enable Grafana ingress. Disabled by default so Grafana is not exposed publicly."
  type        = bool
  default     = false
}

variable "grafana_host" {
  description = "Hostname for Grafana ingress when enable_grafana_ingress is true."
  type        = string
  default     = ""

  validation {
    condition     = !var.enable_grafana_ingress || length(trimspace(var.grafana_host)) > 0
    error_message = "grafana_host must not be empty when enable_grafana_ingress is true."
  }
}

variable "storage_enabled" {
  description = "Whether to enable persistent storage for Prometheus."
  type        = bool
  default     = false
}

variable "storage_class_name" {
  description = "StorageClass name for Prometheus persistent storage. Leave empty to use the cluster default."
  type        = string
  default     = ""
}

variable "storage_size" {
  description = "Prometheus persistent volume request size when storage_enabled is true."
  type        = string
  default     = "20Gi"
}
