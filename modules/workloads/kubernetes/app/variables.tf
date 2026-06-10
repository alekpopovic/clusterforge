variable "name" {
  description = "Application name."
  type        = string
}

variable "namespace" {
  description = "Kubernetes namespace for the application."
  type        = string
  default     = "default"
}

variable "image" {
  description = "Container image reference."
  type        = string

  validation {
    condition     = length(var.image) > 0
    error_message = "Image must not be empty."
  }
}

variable "replicas" {
  description = "Number of desired application replicas."
  type        = number
  default     = 2

  validation {
    condition     = var.replicas >= 1
    error_message = "Replicas must be at least 1."
  }
}

variable "container_port" {
  description = "Container port exposed by the application."
  type        = number
  default     = 8080

  validation {
    condition     = var.container_port > 0 && var.container_port < 65536
    error_message = "Container port must be between 1 and 65535."
  }
}

variable "service_port" {
  description = "Service port exposed inside the cluster."
  type        = number
  default     = 80

  validation {
    condition     = var.service_port > 0 && var.service_port < 65536
    error_message = "Service port must be between 1 and 65535."
  }
}

variable "env" {
  description = "Plain environment variables. Do not pass secrets here."
  type        = map(string)
  default     = {}
}

variable "labels" {
  description = "Labels applied to Kubernetes resources."
  type        = map(string)
  default     = {}
}

variable "annotations" {
  description = "Annotations applied to Kubernetes resources."
  type        = map(string)
  default     = {}
}
