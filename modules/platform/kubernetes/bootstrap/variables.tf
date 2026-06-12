variable "namespace_prefix" {
  description = "Optional prefix added to each platform add-on namespace."
  type        = string
  default     = ""
}

variable "enable_ingress_nginx" {
  description = "Whether to install ingress-nginx."
  type        = bool
  default     = true
}

variable "enable_cert_manager" {
  description = "Whether to install cert-manager."
  type        = bool
  default     = true
}

variable "enable_external_dns" {
  description = "Whether to install external-dns."
  type        = bool
  default     = false
}

variable "enable_external_secrets" {
  description = "Whether to install External Secrets Operator."
  type        = bool
  default     = false
}

variable "enable_metrics_server" {
  description = "Whether to install metrics-server."
  type        = bool
  default     = true
}

variable "enable_prometheus_stack" {
  description = "Whether to install kube-prometheus-stack."
  type        = bool
  default     = false
}

variable "enable_loki" {
  description = "Whether to install Loki."
  type        = bool
  default     = false
}

variable "enable_argocd" {
  description = "Whether to install Argo CD."
  type        = bool
  default     = false
}

variable "common_labels" {
  description = "Labels applied to namespaces created by enabled platform add-ons."
  type        = map(string)
  default     = {}
}
