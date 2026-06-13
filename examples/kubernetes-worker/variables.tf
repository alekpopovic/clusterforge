variable "kubeconfig_path" {
  description = "Path to kubeconfig for the target cluster."
  type        = string
  default     = "~/.kube/config"
}

variable "namespace" {
  description = "Namespace for the worker example."
  type        = string
  default     = "apps"
}
