variable "project" {
  description = "Project identifier used as the first name segment."
  type        = string

  validation {
    condition     = length(trimspace(var.project)) > 0
    error_message = "Project must not be empty."
  }
}

variable "environment" {
  description = "Environment identifier used as the second name segment."
  type        = string

  validation {
    condition     = length(trimspace(var.environment)) > 0
    error_message = "Environment must not be empty."
  }
}

variable "component" {
  description = "Component or layer identifier such as network, eks, ingress, or app."
  type        = string

  validation {
    condition     = length(trimspace(var.component)) > 0
    error_message = "Component must not be empty."
  }
}

variable "name" {
  description = "Specific resource, platform component, or workload name."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "Name must not be empty."
  }
}

variable "separator" {
  description = "Separator used when joining name parts."
  type        = string
  default     = "-"

  validation {
    condition     = contains(["-", "_", ""], var.separator)
    error_message = "Separator must be '-', '_' or ''."
  }
}

variable "max_length" {
  description = "Maximum length for the generated name."
  type        = number
  default     = 63

  validation {
    condition     = var.max_length >= 16 && var.max_length <= 128
    error_message = "Max length must be between 16 and 128."
  }
}

variable "lowercase" {
  description = "Whether to lowercase the generated name."
  type        = bool
  default     = true
}

variable "extra_parts" {
  description = "Additional name parts appended after name and before suffix."
  type        = list(string)
  default     = []
}

variable "suffix" {
  description = "Optional suffix appended after extra_parts."
  type        = string
  default     = ""
}
