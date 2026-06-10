variable "name" {
  description = "Logical name for this module instance."
  type        = string
  default     = null

  validation {
    condition     = var.name == null || can(regex("^[a-z][a-z0-9-]{1,62}$", var.name))
    error_message = "Name must start with a lowercase letter and contain 2-63 lowercase letters, numbers, or hyphens."
  }
}

variable "environment" {
  description = "Environment identifier such as dev, staging, or prod."
  type        = string
  default     = null

  validation {
    condition     = var.environment == null || can(regex("^[a-z][a-z0-9-]{1,30}$", var.environment))
    error_message = "Environment must start with a lowercase letter and contain 2-31 lowercase letters, numbers, or hyphens."
  }
}

variable "labels" {
  description = "Labels to apply to resources created by this module once implemented."
  type        = map(string)
  default     = {}
}
