variable "name" {
  description = "Name for the Application Load Balancer and related resources."
  type        = string

  validation {
    condition     = can(regex("^[a-z][a-z0-9-]{1,30}$", var.name))
    error_message = "Name must start with a lowercase letter and contain 2-31 lowercase letters, numbers, or hyphens."
  }
}

variable "environment" {
  description = "Environment identifier such as dev, staging, or prod."
  type        = string

  validation {
    condition     = can(regex("^[a-z][a-z0-9-]{1,30}$", var.environment))
    error_message = "Environment must start with a lowercase letter and contain 2-31 lowercase letters, numbers, or hyphens."
  }
}

variable "vpc_id" {
  description = "VPC ID where the ALB and target groups are created."
  type        = string

  validation {
    condition     = length(trimspace(var.vpc_id)) > 0
    error_message = "vpc_id must not be empty."
  }
}

variable "public_subnet_ids" {
  description = "Public subnet IDs for the internet-facing ALB."
  type        = list(string)

  validation {
    condition     = length(var.public_subnet_ids) > 0
    error_message = "public_subnet_ids must contain at least one subnet ID."
  }
}

variable "allowed_cidr_blocks" {
  description = "CIDR blocks allowed to reach enabled ALB listeners."
  type        = list(string)
  default     = ["0.0.0.0/0"]

  validation {
    condition     = length(var.allowed_cidr_blocks) > 0
    error_message = "allowed_cidr_blocks must contain at least one CIDR block."
  }
}

variable "certificate_arn" {
  description = "ACM certificate ARN for HTTPS listeners. Required when enable_https is true."
  type        = string
  default     = ""
}

variable "enable_http" {
  description = "Whether to create an HTTP listener on port 80."
  type        = bool
  default     = true

  validation {
    condition     = var.enable_http || var.enable_https
    error_message = "At least one listener must be enabled with enable_http or enable_https."
  }
}

variable "enable_https" {
  description = "Whether to create an HTTPS listener on port 443. Requires certificate_arn."
  type        = bool
  default     = false

  validation {
    condition     = !var.enable_https || length(trimspace(var.certificate_arn)) > 0
    error_message = "certificate_arn must be set when enable_https is true."
  }
}

variable "target_groups" {
  description = "Target groups keyed by logical service name."
  type = map(object({
    port                 = number
    protocol             = optional(string, "HTTP")
    health_check_path    = optional(string, "/")
    health_check_matcher = optional(string, "200-399")
  }))

  validation {
    condition     = length(var.target_groups) > 0
    error_message = "target_groups must contain at least one target group."
  }

  validation {
    condition = alltrue([
      for key in keys(var.target_groups) : can(regex("^[a-z][a-z0-9-]{0,15}$", key))
    ])
    error_message = "Each target group key must start with a lowercase letter and contain up to 16 lowercase letters, numbers, or hyphens."
  }

  validation {
    condition = alltrue([
      for key in keys(var.target_groups) : length("${var.name}-${key}") <= 32
    ])
    error_message = "Each target group name formed as name-key must be 32 characters or fewer."
  }

  validation {
    condition = alltrue([
      for target_group in values(var.target_groups) : target_group.port > 0 && target_group.port <= 65535
    ])
    error_message = "Each target group port must be between 1 and 65535."
  }

  validation {
    condition = alltrue([
      for target_group in values(var.target_groups) : contains(["HTTP", "HTTPS"], upper(target_group.protocol))
    ])
    error_message = "Each target group protocol must be HTTP or HTTPS."
  }
}

variable "tags" {
  description = "Tags applied to ALB resources."
  type        = map(string)
  default     = {}
}
