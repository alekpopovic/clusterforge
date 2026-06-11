variable "name" {
  description = "Nomad job, group, and task name."
  type        = string

  validation {
    condition     = can(regex("^[a-z][a-z0-9-]{1,62}$", var.name))
    error_message = "Name must start with a lowercase letter and contain 2-63 lowercase letters, numbers, or hyphens."
  }
}

variable "datacenters" {
  description = "Nomad datacenters where the job can run."
  type        = list(string)

  validation {
    condition     = length(var.datacenters) > 0
    error_message = "At least one datacenter is required."
  }
}

variable "namespace" {
  description = "Nomad namespace for the job."
  type        = string
  default     = "default"

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "Namespace must not be empty."
  }
}

variable "type" {
  description = "Nomad job type."
  type        = string
  default     = "service"

  validation {
    condition     = contains(["service", "batch", "system", "sysbatch"], var.type)
    error_message = "Type must be one of service, batch, system, or sysbatch."
  }
}

variable "image" {
  description = "Docker image reference for the task."
  type        = string

  validation {
    condition     = length(trimspace(var.image)) > 0
    error_message = "Image must not be empty."
  }
}

variable "task_count" {
  description = "Number of task group instances."
  type        = number
  default     = 1

  validation {
    condition     = var.task_count >= 0
    error_message = "Task count must be greater than or equal to 0."
  }
}

variable "command" {
  description = "Optional command override for the Docker task."
  type        = string
  default     = ""
}

variable "args" {
  description = "Optional command arguments."
  type        = list(string)
  default     = []
}

variable "env" {
  description = "Environment variables for the task. Do not put secrets here."
  type        = map(string)
  default     = {}

  validation {
    condition     = alltrue([for key in keys(var.env) : can(regex("^[A-Za-z_][A-Za-z0-9_]*$", key))])
    error_message = "Environment variable names must be valid shell-style identifiers."
  }
}

variable "ports" {
  description = "Network ports exposed by the task group."
  type = list(object({
    label  = string
    to     = number
    static = optional(number)
  }))
  default = []

  validation {
    condition     = alltrue([for port in var.ports : can(regex("^[a-z][a-z0-9-]{0,62}$", port.label))])
    error_message = "Port labels must start with a lowercase letter and contain lowercase letters, numbers, or hyphens."
  }

  validation {
    condition     = alltrue([for port in var.ports : port.to > 0 && port.to < 65536])
    error_message = "Port target values must be between 1 and 65535."
  }

  validation {
    condition     = alltrue([for port in var.ports : port.static == null || (port.static > 0 && port.static < 65536)])
    error_message = "Static port values must be between 1 and 65535 when set."
  }
}

variable "cpu" {
  description = "CPU shares allocated to the task."
  type        = number
  default     = 500

  validation {
    condition     = var.cpu > 0
    error_message = "CPU must be greater than 0."
  }
}

variable "memory" {
  description = "Memory allocated to the task in MiB."
  type        = number
  default     = 256

  validation {
    condition     = var.memory > 0
    error_message = "Memory must be greater than 0."
  }
}

variable "service" {
  description = "Optional Nomad service registration."
  type = object({
    enabled    = bool
    name       = optional(string)
    port_label = optional(string, "http")
    tags       = optional(list(string), [])
  })
  default = {
    enabled = false
  }

  validation {
    condition     = !var.service.enabled || contains([for port in var.ports : port.label], var.service.port_label)
    error_message = "service.port_label must match one of the configured port labels when service.enabled is true."
  }
}
