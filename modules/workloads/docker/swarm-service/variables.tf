variable "name" {
  description = "Docker Swarm service name."
  type        = string

  validation {
    condition     = can(regex("^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,127}$", var.name))
    error_message = "Name must start with a letter or number and contain only letters, numbers, underscores, dots, or hyphens."
  }
}

variable "image" {
  description = "Docker image reference to pull and run."
  type        = string

  validation {
    condition     = length(trimspace(var.image)) > 0
    error_message = "Image must not be empty."
  }
}

variable "replicas" {
  description = "Number of Swarm service replicas."
  type        = number
  default     = 1

  validation {
    condition     = var.replicas >= 0
    error_message = "Replicas must be greater than or equal to 0."
  }
}

variable "env" {
  description = "Environment variables for the service task. Do not put secrets here."
  type        = map(string)
  default     = {}
}

variable "ports" {
  description = "Swarm service port publishing rules."
  type = list(object({
    target_port    = number
    published_port = number
    protocol       = optional(string, "tcp")
  }))
  default = []

  validation {
    condition     = alltrue([for port in var.ports : port.target_port > 0 && port.target_port < 65536])
    error_message = "Target ports must be between 1 and 65535."
  }

  validation {
    condition     = alltrue([for port in var.ports : port.published_port > 0 && port.published_port < 65536])
    error_message = "Published ports must be between 1 and 65535."
  }

  validation {
    condition     = alltrue([for port in var.ports : contains(["tcp", "udp", "sctp"], lower(port.protocol))])
    error_message = "Port protocol must be tcp, udp, or sctp."
  }
}

variable "networks" {
  description = "Docker network names or IDs to attach to the service."
  type        = list(string)
  default     = []
}

variable "labels" {
  description = "Docker labels applied to the Swarm service."
  type        = map(string)
  default     = {}
}
