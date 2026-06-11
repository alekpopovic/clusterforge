variable "name" {
  description = "Kubernetes CronJob name."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "Name must not be empty."
  }
}

variable "namespace" {
  description = "Kubernetes namespace for the CronJob."
  type        = string

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "Namespace must not be empty."
  }
}

variable "create_namespace" {
  description = "Whether to create the namespace."
  type        = bool
  default     = true
}

variable "image" {
  description = "Container image reference."
  type        = string

  validation {
    condition     = length(trimspace(var.image)) > 0
    error_message = "Image must not be empty."
  }
}

variable "schedule" {
  description = "Cron schedule expression."
  type        = string

  validation {
    condition     = length(trimspace(var.schedule)) > 0
    error_message = "Schedule must not be empty."
  }
}

variable "command" {
  description = "Optional container command."
  type        = list(string)
  default     = []
}

variable "args" {
  description = "Optional container args."
  type        = list(string)
  default     = []
}

variable "env" {
  description = "Plain environment variables. Do not put secrets here."
  type        = map(string)
  default     = {}
}

variable "secret_env" {
  description = "Environment variables sourced from existing Kubernetes secrets."
  type = map(object({
    secret_name = string
    secret_key  = string
  }))
  default = {}
}

variable "labels" {
  description = "Additional labels applied to CronJob resources."
  type        = map(string)
  default     = {}
}

variable "annotations" {
  description = "Annotations applied to the CronJob and pod template."
  type        = map(string)
  default     = {}
}

variable "concurrency_policy" {
  description = "How Kubernetes treats concurrent executions."
  type        = string
  default     = "Forbid"

  validation {
    condition     = contains(["Allow", "Forbid", "Replace"], var.concurrency_policy)
    error_message = "Concurrency policy must be Allow, Forbid, or Replace."
  }
}

variable "successful_jobs_history_limit" {
  description = "Number of successful jobs to retain."
  type        = number
  default     = 3
}

variable "failed_jobs_history_limit" {
  description = "Number of failed jobs to retain."
  type        = number
  default     = 3
}

variable "restart_policy" {
  description = "Pod restart policy for CronJob pods."
  type        = string
  default     = "OnFailure"

  validation {
    condition     = contains(["OnFailure", "Never"], var.restart_policy)
    error_message = "Restart policy must be OnFailure or Never."
  }
}

variable "resources" {
  description = "Container resource requests and limits."
  type = object({
    cpu_request    = optional(string)
    memory_request = optional(string)
    cpu_limit      = optional(string)
    memory_limit   = optional(string)
  })
  default = {}
}

variable "image_pull_policy" {
  description = "Image pull policy for the container."
  type        = string
  default     = "IfNotPresent"

  validation {
    condition     = contains(["Always", "IfNotPresent", "Never"], var.image_pull_policy)
    error_message = "Image pull policy must be Always, IfNotPresent, or Never."
  }
}

variable "image_pull_secrets" {
  description = "Names of existing image pull secrets."
  type        = list(string)
  default     = []
}
