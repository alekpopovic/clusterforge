variable "name" {
  description = "Logical name for the KMS key."
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

variable "description" {
  description = "KMS key description."
  type        = string
  default     = "ClusterForge managed KMS key."
}

variable "alias_name" {
  description = "KMS alias name. May be provided with or without the alias/ prefix."
  type        = string

  validation {
    condition     = length(trimspace(var.alias_name)) > 0
    error_message = "alias_name must not be empty."
  }
}

variable "deletion_window_in_days" {
  description = "KMS key deletion recovery window in days."
  type        = number
  default     = 30

  validation {
    condition     = var.deletion_window_in_days >= 7 && var.deletion_window_in_days <= 30
    error_message = "deletion_window_in_days must be between 7 and 30."
  }
}

variable "enable_key_rotation" {
  description = "Whether annual KMS key rotation is enabled."
  type        = bool
  default     = true
}

variable "policy_json" {
  description = "Optional KMS key policy JSON. Leave empty to use the AWS provider default policy behavior."
  type        = string
  default     = ""
}

variable "tags" {
  description = "Additional tags applied to KMS resources."
  type        = map(string)
  default     = {}
}
