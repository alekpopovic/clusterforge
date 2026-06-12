variable "name" {
  description = "Name of the SecretStore or ClusterSecretStore."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "Name must not be empty."
  }
}

variable "region" {
  description = "AWS region for Secrets Manager or SSM Parameter Store."
  type        = string

  validation {
    condition     = length(trimspace(var.region)) > 0
    error_message = "Region must not be empty."
  }
}

variable "service" {
  description = "AWS external-secrets provider service."
  type        = string
  default     = "SecretsManager"

  validation {
    condition     = contains(["SecretsManager", "ParameterStore"], var.service)
    error_message = "Service must be SecretsManager or ParameterStore."
  }
}

variable "auth_type" {
  description = "External Secrets AWS auth type. The first implementation supports jwt for IRSA."
  type        = string
  default     = "jwt"

  validation {
    condition     = var.auth_type == "jwt"
    error_message = "Only jwt auth_type is currently supported."
  }
}

variable "service_account_ref_name" {
  description = "Service account name used by External Secrets Operator to authenticate with AWS."
  type        = string

  validation {
    condition     = length(trimspace(var.service_account_ref_name)) > 0
    error_message = "Service account reference name must not be empty."
  }
}

variable "service_account_ref_namespace" {
  description = "Namespace of the service account used by External Secrets Operator."
  type        = string

  validation {
    condition     = length(trimspace(var.service_account_ref_namespace)) > 0
    error_message = "Service account reference namespace must not be empty."
  }
}

variable "kind" {
  description = "External Secrets store kind to create."
  type        = string
  default     = "ClusterSecretStore"

  validation {
    condition     = contains(["ClusterSecretStore", "SecretStore"], var.kind)
    error_message = "Kind must be ClusterSecretStore or SecretStore."
  }
}
