variable "name" {
  description = "Name prefix for Redis resources."
  type        = string
}

variable "environment" {
  description = "Environment name for tagging."
  type        = string
}

variable "vpc_id" {
  description = "VPC ID where the Redis security group is created."
  type        = string
}

variable "subnet_ids" {
  description = "Private subnet IDs for the ElastiCache subnet group."
  type        = list(string)
}

variable "allowed_security_group_ids" {
  description = "Security group IDs allowed to connect to Redis."
  type        = list(string)
}

variable "node_type" {
  description = "ElastiCache node type."
  type        = string
  default     = "cache.t4g.micro"
}

variable "engine_version" {
  description = "Redis engine version."
  type        = string
  default     = "7.1"
}

variable "num_cache_nodes" {
  description = "Number of cache nodes in the replication group."
  type        = number
  default     = 1
}

variable "automatic_failover_enabled" {
  description = "Whether automatic failover is enabled."
  type        = bool
  default     = false
}

variable "multi_az_enabled" {
  description = "Whether Multi-AZ is enabled."
  type        = bool
  default     = false
}

variable "at_rest_encryption_enabled" {
  description = "Whether at-rest encryption is enabled."
  type        = bool
  default     = true
}

variable "transit_encryption_enabled" {
  description = "Whether in-transit encryption is enabled."
  type        = bool
  default     = true
}

variable "auth_token_secret_arn" {
  description = "Optional Secrets Manager secret ARN where an application can read the Redis auth token. The token value is not read or output by this module."
  type        = string
  default     = ""
}

variable "parameter_group_name" {
  description = "Optional ElastiCache parameter group name."
  type        = string
  default     = ""
}

variable "tags" {
  description = "Tags applied to Redis resources."
  type        = map(string)
  default     = {}
}
