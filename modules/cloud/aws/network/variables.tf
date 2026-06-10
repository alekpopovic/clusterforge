variable "name" {
  description = "Name prefix for VPC networking resources."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "Name must not be empty."
  }
}

variable "environment" {
  description = "Environment name for tagging and resource names."
  type        = string

  validation {
    condition     = length(trimspace(var.environment)) > 0
    error_message = "Environment must not be empty."
  }
}

variable "cidr_block" {
  description = "CIDR block for the VPC."
  type        = string

  validation {
    condition     = length(trimspace(var.cidr_block)) > 0 && can(cidrnetmask(var.cidr_block))
    error_message = "CIDR block must be a non-empty valid IPv4 CIDR block."
  }
}

variable "availability_zones" {
  description = "Availability zones used for public and private subnets."
  type        = list(string)

  validation {
    condition     = length(var.availability_zones) > 0 && alltrue([for zone in var.availability_zones : length(trimspace(zone)) > 0])
    error_message = "At least one non-empty availability zone is required."
  }
}

variable "public_subnet_cidrs" {
  description = "CIDR blocks for public subnets. Must align one-to-one with availability_zones."
  type        = list(string)

  validation {
    condition     = alltrue([for cidr in var.public_subnet_cidrs : can(cidrnetmask(cidr))])
    error_message = "Each public subnet CIDR must be a valid IPv4 CIDR block."
  }
}

variable "private_subnet_cidrs" {
  description = "CIDR blocks for private subnets. Must align one-to-one with availability_zones."
  type        = list(string)

  validation {
    condition     = alltrue([for cidr in var.private_subnet_cidrs : can(cidrnetmask(cidr))])
    error_message = "Each private subnet CIDR must be a valid IPv4 CIDR block."
  }
}

variable "enable_nat_gateway" {
  description = "Whether to create NAT gateway resources and private default routes."
  type        = bool
  default     = true
}

variable "single_nat_gateway" {
  description = "Whether to create one shared NAT gateway instead of one NAT gateway per availability zone."
  type        = bool
  default     = true
}

variable "enable_dns_hostnames" {
  description = "Whether DNS hostnames are enabled in the VPC."
  type        = bool
  default     = true
}

variable "enable_dns_support" {
  description = "Whether DNS resolution is enabled in the VPC."
  type        = bool
  default     = true
}

variable "tags" {
  description = "Tags applied to all supported AWS resources."
  type        = map(string)
  default     = {}
}

variable "public_subnet_tags" {
  description = "Additional tags applied to public subnets, such as EKS cluster discovery tags."
  type        = map(string)
  default     = {}
}

variable "private_subnet_tags" {
  description = "Additional tags applied to private subnets, such as EKS cluster discovery tags."
  type        = map(string)
  default     = {}
}
