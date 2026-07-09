variable "name" {
  description = "Base name for Azure network resources."
  type        = string
}

variable "environment" {
  description = "Environment identifier."
  type        = string
}

variable "location" {
  description = "Azure region where resources are created."
  type        = string
}

variable "resource_group_name" {
  description = "Resource group name to create or reuse."
  type        = string
}

variable "create_resource_group" {
  description = "Whether to create the resource group."
  type        = bool
  default     = true
}

variable "address_space" {
  description = "Virtual network address space."
  type        = list(string)
  default     = ["10.40.0.0/16"]
}

variable "subnet_prefixes" {
  description = "Subnet CIDR prefixes."
  type        = list(string)
  default     = ["10.40.1.0/24"]
}

variable "tags" {
  description = "Tags to apply to resources."
  type        = map(string)
  default     = {}
}
