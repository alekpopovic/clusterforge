variable "region" {
  description = "AWS region for backend bootstrap resources."
  type        = string
  default     = "eu-central-1"
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
  default     = "clusterforge-terraform-locks"
}

variable "tags" {
  description = "Tags to apply to backend bootstrap resources."
  type        = map(string)
  default = {
    Project   = "clusterforge"
    ManagedBy = "terraform"
  }
}
