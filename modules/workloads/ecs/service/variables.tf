variable "name" {
  description = "ECS service and task family name."
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

variable "cluster_arn" {
  description = "ARN of the ECS cluster where the service runs."
  type        = string

  validation {
    condition     = length(trimspace(var.cluster_arn)) > 0
    error_message = "Cluster ARN must not be empty."
  }
}

variable "subnet_ids" {
  description = "Subnet IDs for the ECS service awsvpc network configuration."
  type        = list(string)

  validation {
    condition     = length(var.subnet_ids) > 0
    error_message = "At least one subnet ID is required."
  }
}

variable "security_group_ids" {
  description = "Security group IDs for the ECS service awsvpc network configuration."
  type        = list(string)
}

variable "assign_public_ip" {
  description = "Whether tasks receive public IP addresses."
  type        = bool
  default     = false
}

variable "desired_count" {
  description = "Desired task count when autoscaling is disabled."
  type        = number
  default     = 1

  validation {
    condition     = var.desired_count >= 0
    error_message = "Desired count must be greater than or equal to 0."
  }
}

variable "image" {
  description = "Container image reference."
  type        = string

  validation {
    condition     = length(trimspace(var.image)) > 0
    error_message = "Image must not be empty."
  }
}

variable "cpu" {
  description = "Fargate task CPU units."
  type        = number
  default     = 256
}

variable "memory" {
  description = "Fargate task memory in MiB."
  type        = number
  default     = 512

  validation {
    condition = (
      (var.cpu == 256 && contains([512, 1024, 2048], var.memory)) ||
      (var.cpu == 512 && var.memory >= 1024 && var.memory <= 4096 && var.memory % 1024 == 0) ||
      (var.cpu == 1024 && var.memory >= 2048 && var.memory <= 8192 && var.memory % 1024 == 0) ||
      (var.cpu == 2048 && var.memory >= 4096 && var.memory <= 16384 && var.memory % 1024 == 0) ||
      (var.cpu == 4096 && var.memory >= 8192 && var.memory <= 30720 && var.memory % 1024 == 0) ||
      (var.cpu == 8192 && var.memory >= 16384 && var.memory <= 61440 && var.memory % 4096 == 0) ||
      (var.cpu == 16384 && var.memory >= 32768 && var.memory <= 122880 && var.memory % 8192 == 0)
    )
    error_message = "CPU and memory must use a valid Fargate combination."
  }
}

variable "container_port" {
  description = "Container port exposed by the task."
  type        = number

  validation {
    condition     = var.container_port > 0 && var.container_port < 65536
    error_message = "Container port must be between 1 and 65535."
  }
}

variable "protocol" {
  description = "Container port protocol."
  type        = string
  default     = "tcp"

  validation {
    condition     = contains(["tcp", "udp"], lower(var.protocol))
    error_message = "Protocol must be tcp or udp."
  }
}

variable "environment_variables" {
  description = "Plain environment variables. Do not put secrets here."
  type        = map(string)
  default     = {}
}

variable "secrets" {
  description = "Secrets exposed to the container by ARN references, such as SSM Parameter or Secrets Manager ARNs."
  type = list(object({
    name       = string
    value_from = string
  }))
  default = []
}

variable "execution_role_arn" {
  description = "Existing ECS task execution role ARN. Leave empty to create one."
  type        = string
  default     = ""
}

variable "task_role_arn" {
  description = "Existing ECS task role ARN. Leave empty to create one."
  type        = string
  default     = ""
}

variable "create_log_group" {
  description = "Whether to create the CloudWatch log group."
  type        = bool
  default     = true
}

variable "log_retention_in_days" {
  description = "CloudWatch log retention in days when create_log_group is true."
  type        = number
  default     = 30
}

variable "tags" {
  description = "Tags applied to supported AWS resources."
  type        = map(string)
  default     = {}
}

variable "load_balancer" {
  description = "Optional ECS service load balancer configuration."
  type = object({
    enabled          = bool
    target_group_arn = optional(string)
    container_name   = optional(string)
    container_port   = optional(number)
  })
  default = {
    enabled = false
  }

  validation {
    condition     = !var.load_balancer.enabled || try(length(trimspace(var.load_balancer.target_group_arn)) > 0, false)
    error_message = "target_group_arn is required when load_balancer.enabled is true."
  }
}

variable "autoscaling" {
  description = "Optional ECS service autoscaling configuration."
  type = object({
    enabled          = bool
    min_capacity     = optional(number, 1)
    max_capacity     = optional(number, 3)
    cpu_target_value = optional(number, 70)
  })
  default = {
    enabled = false
  }

  validation {
    condition     = !var.autoscaling.enabled || var.autoscaling.min_capacity <= var.autoscaling.max_capacity
    error_message = "Autoscaling min_capacity must be less than or equal to max_capacity."
  }
}
