variable "name" {
  description = "Name prefix for PostgreSQL resources and DB identifier."
  type        = string
}

variable "environment" {
  description = "Environment name for tagging."
  type        = string
}

variable "vpc_id" {
  description = "VPC ID where the database security group is created."
  type        = string
}

variable "subnet_ids" {
  description = "Private subnet IDs for the DB subnet group."
  type        = list(string)
}

variable "allowed_security_group_ids" {
  description = "Security group IDs allowed to connect to PostgreSQL."
  type        = list(string)
}

variable "engine_version" {
  description = "PostgreSQL engine version."
  type        = string
  default     = "16.3"
}

variable "instance_class" {
  description = "RDS instance class."
  type        = string
  default     = "db.t4g.micro"
}

variable "allocated_storage" {
  description = "Initial allocated storage in GiB."
  type        = number
  default     = 20
}

variable "max_allocated_storage" {
  description = "Maximum autoscaled storage in GiB. Set 0 to disable storage autoscaling."
  type        = number
  default     = 100
}

variable "database_name" {
  description = "Initial database name."
  type        = string
}

variable "master_username" {
  description = "Master database username."
  type        = string
  default     = "postgres"
}

variable "manage_master_user_password" {
  description = "Whether RDS manages the master user password in AWS Secrets Manager."
  type        = bool
  default     = true
}

variable "master_password" {
  description = "Master password when manage_master_user_password is false. This value is sensitive and will be stored in state."
  type        = string
  default     = ""
  sensitive   = true
}

variable "multi_az" {
  description = "Whether to enable Multi-AZ deployment."
  type        = bool
  default     = false
}

variable "backup_retention_period" {
  description = "Automated backup retention in days."
  type        = number
  default     = 7
}

variable "deletion_protection" {
  description = "Whether deletion protection is enabled."
  type        = bool
  default     = true
}

variable "skip_final_snapshot" {
  description = "Whether to skip the final DB snapshot on destroy."
  type        = bool
  default     = false
}

variable "storage_encrypted" {
  description = "Whether storage encryption is enabled."
  type        = bool
  default     = true
}

variable "kms_key_arn" {
  description = "Optional KMS key ARN for storage encryption and AWS-managed master password."
  type        = string
  default     = ""
}

variable "tags" {
  description = "Tags applied to RDS resources."
  type        = map(string)
  default     = {}
}
