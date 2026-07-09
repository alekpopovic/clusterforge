variable "kubeconfig_path" {
  description = "Path to kubeconfig."
  type        = string
  default     = "~/.kube/config"
}

variable "kubeconfig_context" {
  description = "Existing cluster kubeconfig context."
  type        = string
  default     = ""
}
