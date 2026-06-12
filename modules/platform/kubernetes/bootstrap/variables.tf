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

variable "argocd_enable_app_of_apps" {
  description = "Whether to create an Argo CD app-of-apps Application when Argo CD is enabled."
  type        = bool
  default     = false
}

variable "argocd_app_of_apps_name" {
  description = "Name of the Argo CD app-of-apps Application."
  type        = string
  default     = "cluster-apps"
}

variable "argocd_app_of_apps_repo_url" {
  description = "Git repository URL containing app-of-apps definitions. Do not include credentials."
  type        = string
  default     = ""

  validation {
    condition     = !var.enable_argocd || !var.argocd_enable_app_of_apps || length(trimspace(var.argocd_app_of_apps_repo_url)) > 0
    error_message = "argocd_app_of_apps_repo_url must not be empty when argocd_enable_app_of_apps is true."
  }
}

variable "argocd_app_of_apps_path" {
  description = "Path within the Git repository for app-of-apps definitions."
  type        = string
  default     = "apps"
}

variable "argocd_app_of_apps_revision" {
  description = "Git revision for the app-of-apps Application."
  type        = string
  default     = "HEAD"
}

variable "argocd_app_of_apps_destination_namespace" {
  description = "Destination namespace used by the app-of-apps Application."
  type        = string
  default     = "argocd"
}

variable "argocd_app_of_apps_project" {
  description = "Argo CD project for the app-of-apps Application."
  type        = string
  default     = "default"
}

variable "common_labels" {
  description = "Labels applied to namespaces created by enabled platform add-ons."
  type        = map(string)
  default     = {}
}
