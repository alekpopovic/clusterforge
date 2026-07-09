variable "kubeconfig_path" {
  description = "Path to kubeconfig."
  type        = string
  default     = "~/.kube/config"
}

variable "kubeconfig_context" {
  description = "Kind kubeconfig context."
  type        = string
  default     = "kind-clusterforge-local"
}
