variable "kubeconfig_path" {
  description = "Path to an existing kubeconfig. Do not commit kubeconfig contents."
  type        = string
  default     = "~/.kube/config"
}

variable "kubeconfig_context" {
  description = "Optional kubeconfig context used by Kubernetes and Helm providers."
  type        = string
  default     = null
}

variable "kyverno_chart_version" {
  description = "Reviewed Kyverno Helm chart version. Set this explicitly before production use."
  type        = string
  default     = ""
}

variable "enable_baseline_policies" {
  description = "Enable audit-mode baseline policies after the Kyverno CRDs have been installed."
  type        = bool
  default     = false
}
