variable "name" {
  description = "Docker container name."
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

variable "command" {
  description = "Command and arguments passed to the container."
  type        = list(string)
  default     = []
}

variable "env" {
  description = "Environment variables for the container. Do not put secrets here."
  type        = map(string)
  default     = {}
}

variable "ports" {
  description = "Container port mappings."
  type = list(object({
    internal = number
    external = optional(number)
    protocol = optional(string, "tcp")
  }))
  default = []

  validation {
    condition     = alltrue([for port in var.ports : port.internal > 0 && port.internal < 65536])
    error_message = "Internal ports must be between 1 and 65535."
  }

  validation {
    condition     = alltrue([for port in var.ports : port.external == null || (port.external > 0 && port.external < 65536)])
    error_message = "External ports must be between 1 and 65535 when set."
  }

  validation {
    condition     = alltrue([for port in var.ports : contains(["tcp", "udp", "sctp"], lower(port.protocol))])
    error_message = "Port protocol must be tcp, udp, or sctp."
  }
}

variable "volumes" {
  description = "Host path bind mounts."
  type = list(object({
    host_path      = string
    container_path = string
    read_only      = optional(bool, false)
  }))
  default = []

  validation {
    condition     = alltrue([for volume in var.volumes : length(trimspace(volume.host_path)) > 0 && length(trimspace(volume.container_path)) > 0])
    error_message = "Volume host_path and container_path must not be empty."
  }
}

variable "restart_policy" {
  description = "Docker restart policy."
  type        = string
  default     = "unless-stopped"

  validation {
    condition     = contains(["no", "on-failure", "always", "unless-stopped"], var.restart_policy)
    error_message = "Restart policy must be no, on-failure, always, or unless-stopped."
  }
}

variable "networks" {
  description = "Docker network names or IDs to attach to the container."
  type        = list(string)
  default     = []
}

variable "labels" {
  description = "Docker labels applied to the container."
  type        = map(string)
  default     = {}
}
