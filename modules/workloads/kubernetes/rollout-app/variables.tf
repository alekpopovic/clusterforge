variable "name" {
  description = "Rollout and application resource name."
  type        = string
  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "name must not be empty."
  }
}

variable "namespace" {
  description = "Namespace for the Rollout application."
  type        = string
  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "namespace must not be empty."
  }
}

variable "create_namespace" {
  description = "Whether to create the application namespace."
  type        = bool
  default     = false
}

variable "image" {
  description = "Container image reference. Use an immutable tag or digest."
  type        = string
  validation {
    condition     = length(trimspace(var.image)) > 0
    error_message = "image must not be empty."
  }
}

variable "replicas" {
  description = "Desired Rollout replica count."
  type        = number
  default     = 2
  validation {
    condition     = var.replicas > 0
    error_message = "replicas must be greater than zero."
  }
}

variable "labels" {
  description = "Additional labels applied to workload resources."
  type        = map(string)
  default     = {}
}

variable "annotations" {
  description = "Annotations applied to the Rollout and pod template."
  type        = map(string)
  default     = {}
}

variable "port" {
  description = "Container port exposed by the Services."
  type        = number
  default     = 8080
  validation {
    condition     = var.port > 0 && var.port < 65536
    error_message = "port must be between 1 and 65535."
  }
}

variable "env" {
  description = "Plain environment variables. Do not put secrets here."
  type        = map(string)
  default     = {}
}

variable "resources" {
  description = "Container requests and limits."
  type = object({
    requests = optional(map(string), {})
    limits   = optional(map(string), {})
  })
  default = {}
}

variable "strategy" {
  description = "Opt-in canary or blueGreen rollout strategy."
  type = object({
    type = string
    canary_steps = optional(list(object({
      set_weight     = optional(number)
      pause          = optional(bool, false)
      pause_duration = optional(string)
    })), [])
    blue_green_config = optional(object({
      auto_promotion_enabled   = optional(bool, false)
      scale_down_delay_seconds = optional(number, 30)
      preview_replica_count    = optional(number)
    }), {})
  })
  validation {
    condition     = contains(["canary", "blueGreen"], var.strategy.type)
    error_message = "strategy.type must be canary or blueGreen."
  }
  validation {
    condition     = var.strategy.type != "canary" || alltrue([for step in var.strategy.canary_steps : (step.set_weight != null ? 1 : 0) + (step.pause || step.pause_duration != null ? 1 : 0) == 1])
    error_message = "Each canary step must define exactly one of set_weight or pause/pause_duration."
  }
}

variable "ingress" {
  description = "Optional ingress routing to the active/stable Service."
  type = object({
    enabled     = optional(bool, false)
    host        = optional(string)
    class_name  = optional(string)
    annotations = optional(map(string), {})
    tls         = optional(bool, true)
  })
  default = {}
  validation {
    condition     = !var.ingress.enabled || try(length(trimspace(var.ingress.host)) > 0, false)
    error_message = "ingress.host is required when ingress is enabled."
  }
}
