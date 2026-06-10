variable "kubeconfig_path" {
  description = "Path to kubeconfig for the target cluster."
  type        = string
  default     = "~/.kube/config"
}

variable "namespace" {
  description = "Namespace for the example app."
  type        = string
  default     = "apps"
}
