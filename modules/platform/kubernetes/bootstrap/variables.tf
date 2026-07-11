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

variable "enable_pod_security" {
  description = "Whether to apply Pod Security Admission labels to configured namespaces."
  type        = bool
  default     = false
}

variable "pod_security_namespaces" {
  description = "Namespace names and Pod Security Admission levels to apply when enable_pod_security is true."
  type = map(object({
    enforce = optional(string, "baseline")
    audit   = optional(string, "restricted")
    warn    = optional(string, "restricted")
  }))
  default = {}

  validation {
    condition     = !var.enable_pod_security || length(var.pod_security_namespaces) > 0
    error_message = "pod_security_namespaces must not be empty when enable_pod_security is true."
  }
}

variable "enable_network_policy_baseline" {
  description = "Whether to create baseline NetworkPolicy resources in configured namespaces."
  type        = bool
  default     = false
}

variable "network_policy_namespaces" {
  description = "Namespaces where baseline NetworkPolicy resources should be created."
  type        = list(string)
  default     = []

  validation {
    condition     = !var.enable_network_policy_baseline || length(var.network_policy_namespaces) > 0
    error_message = "network_policy_namespaces must not be empty when enable_network_policy_baseline is true."
  }
}

variable "network_policy_default_deny_ingress" {
  description = "Whether baseline NetworkPolicy should deny ingress by default."
  type        = bool
  default     = true
}

variable "network_policy_default_deny_egress" {
  description = "Whether baseline NetworkPolicy should deny egress by default."
  type        = bool
  default     = false
}

variable "network_policy_allow_dns_egress" {
  description = "Whether baseline NetworkPolicy should allow DNS egress when default deny egress is enabled."
  type        = bool
  default     = true
}

variable "enable_karpenter" {
  description = "Whether to install Karpenter for EKS node autoscaling."
  type        = bool
  default     = false
}

variable "karpenter_chart_version" {
  description = "Optional Karpenter Helm chart version. Pin this before production use."
  type        = string
  default     = ""
}

variable "karpenter_cluster_name" {
  description = "EKS cluster name passed to Karpenter when enable_karpenter is true."
  type        = string
  default     = ""

  validation {
    condition     = !var.enable_karpenter || length(trimspace(var.karpenter_cluster_name)) > 0
    error_message = "karpenter_cluster_name must not be empty when enable_karpenter is true."
  }
}

variable "karpenter_cluster_endpoint" {
  description = "EKS cluster endpoint passed to Karpenter when enable_karpenter is true."
  type        = string
  default     = ""

  validation {
    condition     = !var.enable_karpenter || length(trimspace(var.karpenter_cluster_endpoint)) > 0
    error_message = "karpenter_cluster_endpoint must not be empty when enable_karpenter is true."
  }
}

variable "karpenter_service_account_role_arn" {
  description = "IAM role ARN annotated on the Karpenter service account when enable_karpenter is true."
  type        = string
  default     = ""

  validation {
    condition     = !var.enable_karpenter || length(trimspace(var.karpenter_service_account_role_arn)) > 0
    error_message = "karpenter_service_account_role_arn must not be empty when enable_karpenter is true."
  }
}

variable "karpenter_values" {
  description = "Additional YAML values passed to the Karpenter Helm release."
  type        = list(string)
  default     = []
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

variable "enable_log_agent" {
  description = "Whether to install Grafana Alloy as the Kubernetes log/telemetry agent."
  type        = bool
  default     = false
}

variable "log_agent_chart_version" {
  description = "Optional Grafana Alloy Helm chart version. Pin this before production use."
  type        = string
  default     = ""
}

variable "log_agent_values" {
  description = "Additional YAML values passed to the Grafana Alloy Helm release."
  type        = list(string)
  default     = []
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

variable "enable_opentelemetry_collector" {
  description = "Whether to install OpenTelemetry Collector."
  type        = bool
  default     = false
}
variable "opentelemetry_collector_chart_version" {
  description = "Optional collector chart version."
  type        = string
  default     = ""
}
variable "opentelemetry_collector_mode" {
  description = "Collector deployment mode."
  type        = string
  default     = "deployment"
}
variable "opentelemetry_collector_values" {
  description = "Additional collector Helm values."
  type        = list(string)
  default     = []
}
variable "opentelemetry_collector_service_account_annotations" {
  description = "Collector service account annotations."
  type        = map(string)
  default     = {}
}
