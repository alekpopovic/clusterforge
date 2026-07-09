variable "name" {
  description = "IAM role name for ExternalDNS."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "name must not be empty."
  }
}

variable "environment" {
  description = "Environment name for tagging."
  type        = string

  validation {
    condition     = length(trimspace(var.environment)) > 0
    error_message = "environment must not be empty."
  }
}

variable "oidc_provider_arn" {
  description = "ARN of the IAM OIDC provider trusted by the role."
  type        = string
}

variable "oidc_provider_url" {
  description = "OIDC provider issuer URL. May include or omit the https:// prefix."
  type        = string
}

variable "namespace" {
  description = "Kubernetes namespace containing the ExternalDNS service account."
  type        = string
  default     = "external-dns"
}

variable "service_account_name" {
  description = "ExternalDNS Kubernetes service account name."
  type        = string
  default     = "external-dns"
}

variable "hosted_zone_ids" {
  description = "Route53 hosted zone IDs ExternalDNS may change."
  type        = list(string)
  default     = []
}

variable "allow_all_hosted_zones" {
  description = "Explicitly allow ExternalDNS to change all hosted zones. Keep false for production."
  type        = bool
  default     = false
}

variable "policy_mode" {
  description = "ExternalDNS policy mode: sync allows deletes, upsert-only prevents deletes."
  type        = string
  default     = "sync"

  validation {
    condition     = contains(["sync", "upsert-only"], var.policy_mode)
    error_message = "policy_mode must be sync or upsert-only."
  }
}

variable "tags" {
  description = "Tags applied to IAM resources."
  type        = map(string)
  default     = {}
}
