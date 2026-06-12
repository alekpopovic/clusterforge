variable "name" {
  description = "IAM role name."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "Name must not be empty."
  }
}

variable "environment" {
  description = "Environment name for tagging."
  type        = string

  validation {
    condition     = length(trimspace(var.environment)) > 0
    error_message = "Environment must not be empty."
  }
}

variable "oidc_provider_arn" {
  description = "ARN of the IAM OIDC provider trusted by the role."
  type        = string

  validation {
    condition     = length(trimspace(var.oidc_provider_arn)) > 0
    error_message = "OIDC provider ARN must not be empty."
  }
}

variable "oidc_provider_url" {
  description = "OIDC provider issuer URL. May include or omit the https:// prefix."
  type        = string

  validation {
    condition     = length(trimspace(var.oidc_provider_url)) > 0
    error_message = "OIDC provider URL must not be empty."
  }
}

variable "namespace" {
  description = "Kubernetes namespace containing the trusted service account."
  type        = string

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "Namespace must not be empty."
  }
}

variable "service_account_name" {
  description = "Kubernetes service account name trusted by the role."
  type        = string

  validation {
    condition     = length(trimspace(var.service_account_name)) > 0
    error_message = "Service account name must not be empty."
  }
}

variable "policy_arns" {
  description = "Managed IAM policy ARNs to attach to the role."
  type        = list(string)
  default     = []
}

variable "inline_policies" {
  description = "Map of inline IAM policy names to JSON policy documents."
  type        = map(string)
  default     = {}
}

variable "tags" {
  description = "Tags applied to the IAM role."
  type        = map(string)
  default     = {}
}
