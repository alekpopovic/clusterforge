variable "namespace" {
  description = "Kubernetes namespace for Karpenter."
  type        = string
  default     = "karpenter"

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "Namespace must not be empty."
  }
}

variable "create_namespace" {
  description = "Whether to create the namespace before installing the Helm release."
  type        = bool
  default     = true
}

variable "chart_version" {
  description = "Karpenter chart version. Pin this before production use."
  type        = string
  default     = ""
}

variable "cluster_name" {
  description = "EKS cluster name."
  type        = string

  validation {
    condition     = length(trimspace(var.cluster_name)) > 0
    error_message = "Cluster name must not be empty."
  }
}

variable "cluster_endpoint" {
  description = "EKS cluster endpoint."
  type        = string

  validation {
    condition     = length(trimspace(var.cluster_endpoint)) > 0
    error_message = "Cluster endpoint must not be empty."
  }
}

variable "service_account_role_arn" {
  description = "IAM role ARN annotated onto the Karpenter service account."
  type        = string

  validation {
    condition     = length(trimspace(var.service_account_role_arn)) > 0
    error_message = "Service account role ARN must not be empty."
  }
}

variable "values" {
  description = "Additional YAML values passed to the Helm release."
  type        = list(string)
  default     = []
}

variable "labels" {
  description = "Labels applied to the namespace when create_namespace is true."
  type        = map(string)
  default     = {}
}
