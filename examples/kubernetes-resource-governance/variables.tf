variable "namespace" {
  description = "Existing namespace that explicitly opts into the example governance settings."
  type        = string
  default     = "apps"
}

variable "kubeconfig_path" {
  description = "Path to an existing kubeconfig. Do not commit kubeconfig contents."
  type        = string
  default     = "~/.kube/config"
}

variable "kubeconfig_context" {
  description = "Optional kubeconfig context used by the Kubernetes provider."
  type        = string
  default     = null
}
