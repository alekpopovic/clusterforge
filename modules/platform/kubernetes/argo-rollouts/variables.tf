variable "namespace" {
  description = "Namespace for the Argo Rollouts controller."
  type        = string
  default     = "argo-rollouts"
  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "namespace must not be empty."
  }
}

variable "create_namespace" {
  description = "Whether to create the controller namespace."
  type        = bool
  default     = true
}

variable "chart_version" {
  description = "Optional argo-rollouts chart version. Pin a reviewed version in production."
  type        = string
  default     = ""
}

variable "values" {
  description = "Additional YAML values passed to the Helm release."
  type        = list(string)
  default     = []
}

variable "enable_dashboard_ingress" {
  description = "Whether to enable the optional Rollouts dashboard and expose it through chart-managed ingress."
  type        = bool
  default     = false
}

variable "dashboard_host" {
  description = "Dashboard ingress hostname when dashboard ingress is enabled."
  type        = string
  default     = ""
  validation {
    condition     = !var.enable_dashboard_ingress || length(trimspace(var.dashboard_host)) > 0
    error_message = "dashboard_host is required when dashboard ingress is enabled."
  }
}

variable "labels" {
  description = "Labels applied to the namespace."
  type        = map(string)
  default     = {}
}
