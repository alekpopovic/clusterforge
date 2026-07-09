variable "name" {
  description = "Name prefix for VPC endpoint resources."
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

variable "vpc_id" {
  description = "VPC ID where endpoints are created."
  type        = string

  validation {
    condition     = length(trimspace(var.vpc_id)) > 0
    error_message = "vpc_id must not be empty."
  }
}

variable "subnet_ids" {
  description = "Subnet IDs for interface endpoints."
  type        = list(string)
  default     = []
}

variable "route_table_ids" {
  description = "Route table IDs for gateway endpoints."
  type        = list(string)
  default     = []
}

variable "security_group_ids" {
  description = "Existing security group IDs attached to interface endpoints when create_security_group is false."
  type        = list(string)
  default     = []
}

variable "create_security_group" {
  description = "Whether to create a security group for interface endpoints."
  type        = bool
  default     = true
}

variable "allowed_security_group_ids" {
  description = "Security group IDs allowed to connect to created interface endpoints on TCP 443."
  type        = list(string)
  default     = []
}

variable "gateway_endpoints" {
  description = "AWS service suffixes for gateway endpoints, such as s3."
  type        = list(string)
  default     = ["s3"]
}

variable "interface_endpoints" {
  description = "AWS service suffixes for interface endpoints, such as ecr.api, ecr.dkr, logs, sts, ec2, eks, secretsmanager, or ssm."
  type        = list(string)
  default     = []
}

variable "private_dns_enabled" {
  description = "Whether private DNS is enabled for interface endpoints."
  type        = bool
  default     = true
}

variable "tags" {
  description = "Tags applied to supported endpoint resources."
  type        = map(string)
  default     = {}
}
