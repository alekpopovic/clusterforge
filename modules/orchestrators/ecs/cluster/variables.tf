variable "name" {
  description = "ECS cluster name."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "Name must not be empty."
  }
}

variable "environment" {
  description = "Environment name for tagging."
  type        = string

  validation {
    condition     = length(trimspace(var.environment)) > 0
    error_message = "Environment must not be empty."
  }
}

variable "tags" {
  description = "Tags applied to supported AWS resources."
  type        = map(string)
  default     = {}
}

variable "enable_container_insights" {
  description = "Whether to enable ECS Container Insights for the cluster."
  type        = bool
  default     = true
}

variable "capacity_providers" {
  description = "Capacity providers attached to the ECS cluster."
  type        = list(string)
  default     = ["FARGATE", "FARGATE_SPOT"]

  validation {
    condition     = length(var.capacity_providers) > 0
    error_message = "At least one capacity provider is required."
  }
}

variable "default_capacity_provider_strategy" {
  description = "Default capacity provider strategy for ECS services that do not set one explicitly."
  type = list(object({
    capacity_provider = string
    weight            = optional(number)
    base              = optional(number)
  }))
  default = []
}
