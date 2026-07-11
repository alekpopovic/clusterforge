variable "namespace" {
  description = "Namespace for the collector."
  type        = string
  default     = "observability"
}
variable "create_namespace" {
  description = "Whether to create the namespace."
  type        = bool
  default     = true
}
variable "chart_version" {
  description = "Optional chart version; pin for production."
  type        = string
  default     = ""
}
variable "mode" {
  description = "Collector deployment mode."
  type        = string
  default     = "deployment"
  validation {
    condition     = contains(["deployment", "daemonset", "statefulset"], var.mode)
    error_message = "mode must be deployment, daemonset, or statefulset."
  }
}
variable "values" {
  description = "Additional reviewed Helm values YAML."
  type        = list(string)
  default     = []
}
variable "presets" {
  description = "Optional upstream collector presets."
  type = object({
    logsCollection       = optional(object({ enabled = bool }), { enabled = false })
    kubernetesAttributes = optional(object({ enabled = bool }), { enabled = false })
    kubeletMetrics       = optional(object({ enabled = bool }), { enabled = false })
    hostMetrics          = optional(object({ enabled = bool }), { enabled = false })
  })
  default = null
}
variable "service_account_annotations" {
  description = "Service account annotations, for example workload identity references."
  type        = map(string)
  default     = {}
}
variable "labels" {
  description = "Namespace labels."
  type        = map(string)
  default     = {}
}
