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

variable "argo_rollouts_chart_version" {
  description = "Reviewed Argo Rollouts Helm chart version."
  type        = string
  default     = ""
}

variable "enable_rollout_app" {
  description = "Create the example app after Argo Rollouts CRDs exist."
  type        = bool
  default     = false
}
