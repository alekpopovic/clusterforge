variable "name" {
  description = "Nomad batch job name."
  type        = string
}
variable "datacenters" {
  description = "Target Nomad datacenters."
  type        = list(string)
  default     = ["dc1"]
}
variable "image" {
  description = "Docker image reference."
  type        = string
}
variable "args" {
  description = "Docker task arguments."
  type        = list(string)
  default     = []
}
variable "env" {
  description = "Non-secret environment variables."
  type        = map(string)
  default     = {}
}
variable "task_count" {
  description = "Number of batch allocations."
  type        = number
  default     = 1
  validation {
    condition     = var.task_count >= 1
    error_message = "task_count must be at least one."
  }
}
variable "cpu" {
  description = "CPU shares."
  type        = number
  default     = 500
}
variable "memory" {
  description = "Memory in MiB."
  type        = number
  default     = 256
}
