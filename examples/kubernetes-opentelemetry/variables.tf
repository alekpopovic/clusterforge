variable "kubeconfig_path" {
  description = "Path to kubeconfig."
  type        = string
  default     = "~/.kube/config"
}
variable "kubeconfig_context" {
  description = "Kubeconfig context."
  type        = string
  default     = ""
}
variable "chart_version" {
  description = "Reviewed collector chart version."
  type        = string
  default     = ""
}
