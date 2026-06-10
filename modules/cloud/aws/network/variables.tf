variable "name" {
  description = "Name prefix for network resources."
  type        = string
}

variable "vpc_cidr" {
  description = "CIDR block for the VPC."
  type        = string
  default     = "10.40.0.0/16"
}

variable "availability_zones" {
  description = "Availability zones to use. When empty, the first az_count available zones are selected."
  type        = list(string)
  default     = []
}

variable "az_count" {
  description = "Number of availability zones to use when availability_zones is empty."
  type        = number
  default     = 2

  validation {
    condition     = var.az_count >= 2 && var.az_count <= 6
    error_message = "az_count must be between 2 and 6."
  }
}

variable "public_subnet_cidrs" {
  description = "CIDR blocks for public subnets. Leave empty to derive them from vpc_cidr."
  type        = list(string)
  default     = []
}

variable "private_subnet_cidrs" {
  description = "CIDR blocks for private subnets. Leave empty to derive them from vpc_cidr."
  type        = list(string)
  default     = []
}

variable "enable_nat_gateway" {
  description = "Whether to create one NAT gateway for private subnet egress."
  type        = bool
  default     = true
}

variable "tags" {
  description = "Tags applied to all supported AWS resources."
  type        = map(string)
  default     = {}
}
