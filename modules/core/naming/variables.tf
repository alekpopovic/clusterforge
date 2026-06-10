variable "project" {
  description = "Short project identifier used in generated names."
  type        = string

  validation {
    condition     = can(regex("^[a-z][a-z0-9-]{1,30}$", var.project))
    error_message = "Project must start with a lowercase letter and contain 2-31 lowercase letters, numbers, or hyphens."
  }
}

variable "environment" {
  description = "Deployment environment such as dev, staging, or prod."
  type        = string

  validation {
    condition     = can(regex("^[a-z][a-z0-9-]{1,20}$", var.environment))
    error_message = "Environment must start with a lowercase letter and contain 2-21 lowercase letters, numbers, or hyphens."
  }
}

variable "name" {
  description = "Resource or component name."
  type        = string

  validation {
    condition     = can(regex("^[a-z][a-z0-9-]{1,40}$", var.name))
    error_message = "Name must start with a lowercase letter and contain 2-41 lowercase letters, numbers, or hyphens."
  }
}

variable "delimiter" {
  description = "Delimiter used between name parts."
  type        = string
  default     = "-"

  validation {
    condition     = contains(["-", "_"], var.delimiter)
    error_message = "Delimiter must be either '-' or '_'."
  }
}

variable "max_length" {
  description = "Maximum length for the generated resource name."
  type        = number
  default     = 63

  validation {
    condition     = var.max_length >= 16 && var.max_length <= 128
    error_message = "Max length must be between 16 and 128."
  }
}
