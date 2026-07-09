variable "name" {
  description = "Logical name for Velero backup bucket resources."
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

variable "bucket_name" {
  description = "Globally unique S3 bucket name for Velero backups."
  type        = string

  validation {
    condition     = length(trimspace(var.bucket_name)) > 0
    error_message = "bucket_name must not be empty."
  }
}

variable "force_destroy" {
  description = "Whether to allow Terraform to delete a non-empty backup bucket. Keep false for production."
  type        = bool
  default     = false
}

variable "kms_key_arn" {
  description = "Optional KMS key ARN for bucket encryption. When empty, AES256 SSE is used."
  type        = string
  default     = ""
}

variable "lifecycle_expiration_days" {
  description = "Optional number of days after which old backup objects expire. Null disables lifecycle expiration."
  type        = number
  default     = null
}

variable "tags" {
  description = "Tags applied to S3 backup resources."
  type        = map(string)
  default     = {}
}
