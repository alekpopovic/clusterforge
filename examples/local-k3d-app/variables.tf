variable "kubeconfig_path" {
  description = "Path to kubeconfig."
  type        = string
  default     = "~/.kube/config"
}

variable "kubeconfig_context" {
  description = "k3d kubeconfig context."
  type        = string
  default     = "k3d-clusterforge-local"
}
