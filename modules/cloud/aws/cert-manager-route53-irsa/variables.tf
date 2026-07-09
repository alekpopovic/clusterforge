variable "name" {
  description = "IAM role name for cert-manager Route53 DNS01."
  type        = string
}

variable "environment" {
  description = "Environment name for tagging."
  type        = string
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
  description = "Kubernetes namespace containing the cert-manager service account."
  type        = string
  default     = "cert-manager"
}

variable "service_account_name" {
  description = "cert-manager Kubernetes service account name."
  type        = string
  default     = "cert-manager"
}

variable "hosted_zone_ids" {
  description = "Route53 hosted zone IDs cert-manager may modify for DNS01 challenges."
  type        = list(string)
}

variable "allow_all_hosted_zones" {
  description = "Explicitly allow cert-manager to change all hosted zones. Keep false for production."
  type        = bool
  default     = false
}

variable "tags" {
  description = "Tags applied to IAM resources."
  type        = map(string)
  default     = {}
}
