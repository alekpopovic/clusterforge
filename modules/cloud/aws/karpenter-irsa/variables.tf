variable "cluster_name" {
  description = "EKS cluster name Karpenter will manage."
  type        = string

  validation {
    condition     = length(trimspace(var.cluster_name)) > 0
    error_message = "Cluster name must not be empty."
  }
}

variable "oidc_provider_arn" {
  description = "ARN of the EKS IAM OIDC provider trusted by Karpenter."
  type        = string

  validation {
    condition     = length(trimspace(var.oidc_provider_arn)) > 0
    error_message = "OIDC provider ARN must not be empty."
  }
}

variable "oidc_provider_url" {
  description = "EKS OIDC provider issuer URL. May include or omit the https:// prefix."
  type        = string

  validation {
    condition     = length(trimspace(var.oidc_provider_url)) > 0
    error_message = "OIDC provider URL must not be empty."
  }
}

variable "namespace" {
  description = "Kubernetes namespace for the Karpenter service account."
  type        = string
  default     = "karpenter"

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "Namespace must not be empty."
  }
}

variable "service_account_name" {
  description = "Kubernetes service account name used by Karpenter."
  type        = string
  default     = "karpenter"

  validation {
    condition     = length(trimspace(var.service_account_name)) > 0
    error_message = "Service account name must not be empty."
  }
}

variable "tags" {
  description = "Tags applied to Karpenter IAM resources."
  type        = map(string)
  default     = {}
}
