variable "kubeconfig_path" {
  description = "Path to the kubeconfig file used by the Kubernetes provider."
  type        = string
  default     = "~/.kube/config"
}
