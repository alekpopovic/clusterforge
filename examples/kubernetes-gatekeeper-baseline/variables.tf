variable "kubeconfig_path" {
  description = "Path to an existing kubeconfig. Do not commit kubeconfig contents."
  type        = string
  default     = "~/.kube/config"
}

variable "kubeconfig_context" {
  description = "Optional kubeconfig context."
  type        = string
  default     = null
}

variable "gatekeeper_chart_version" {
  description = "Reviewed Gatekeeper Helm chart version."
  type        = string
  default     = ""
}

variable "enable_audit_constraint" {
  description = "Enable the dry-run label constraint after Gatekeeper CRDs exist."
  type        = bool
  default     = false
}
