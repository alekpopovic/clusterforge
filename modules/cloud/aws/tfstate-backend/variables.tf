variable "name" {
  description = "Logical name for the Terraform state backend."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "name must not be empty."
  }
}

variable "environment" {
  description = "Environment name for tagging and metadata."
  type        = string

  validation {
    condition     = length(trimspace(var.environment)) > 0
    error_message = "environment must not be empty."
  }
}

variable "bucket_name" {
  description = "Globally unique S3 bucket name for Terraform state."
  type        = string

  validation {
    condition     = length(trimspace(var.bucket_name)) > 0
    error_message = "bucket_name must not be empty."
  }
}

variable "dynamodb_table_name" {
  description = "DynamoDB table name for Terraform state locking."
  type        = string

  validation {
    condition     = length(trimspace(var.dynamodb_table_name)) > 0
    error_message = "dynamodb_table_name must not be empty."
  }
}

variable "force_destroy" {
  description = "Whether to allow Terraform to destroy the state bucket even when it contains objects. Keep false for production."
  type        = bool
  default     = false
}

variable "enable_versioning" {
  description = "Whether to enable S3 bucket versioning for state history protection."
  type        = bool
  default     = true
}

variable "enable_encryption" {
  description = "Whether to enable default server-side encryption on the state bucket."
  type        = bool
  default     = true
}

variable "kms_key_arn" {
  description = "Optional KMS key ARN for S3 state bucket encryption. When empty, AES256 SSE is used."
  type        = string
  default     = ""
}

variable "tags" {
  description = "Additional tags to apply to backend resources."
  type        = map(string)
  default     = {}
}
