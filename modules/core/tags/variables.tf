variable "project" {
  description = "Project name used for the Project cloud tag."
  type        = string

  validation {
    condition     = length(trimspace(var.project)) > 0
    error_message = "Project must not be empty."
  }
}

variable "environment" {
  description = "Environment name used for the Environment cloud tag."
  type        = string

  validation {
    condition     = length(trimspace(var.environment)) > 0
    error_message = "Environment must not be empty."
  }
}

variable "owner" {
  description = "Optional owner tag value."
  type        = string
  default     = ""
}

variable "cost_center" {
  description = "Optional cost center tag value."
  type        = string
  default     = ""
}

variable "managed_by" {
  description = "Tool or workflow responsible for managing these resources."
  type        = string
  default     = "terraform"

  validation {
    condition     = length(trimspace(var.managed_by)) > 0
    error_message = "Managed_by must not be empty."
  }
}

variable "component" {
  description = "Optional component tag value."
  type        = string
  default     = ""
}

variable "extra_tags" {
  description = "Additional cloud tags. These are merged last and may override standard tags."
  type        = map(string)
  default     = {}
}
