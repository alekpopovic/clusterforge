variable "name" {
  description = "CodeDeploy application and deployment group base name."
  type        = string
}

variable "environment" {
  description = "Environment identifier."
  type        = string
}

variable "cluster_name" {
  description = "ECS cluster name."
  type        = string
}

variable "service_name" {
  description = "ECS service name."
  type        = string
}

variable "alb_listener_arn" {
  description = "Production ALB listener ARN."
  type        = string
}

variable "blue_target_group_name" {
  description = "Blue target group name."
  type        = string
}

variable "blue_target_group_arn" {
  description = "Blue target group ARN."
  type        = string
}

variable "green_target_group_name" {
  description = "Green target group name."
  type        = string
}

variable "green_target_group_arn" {
  description = "Green target group ARN."
  type        = string
}

variable "termination_wait_time" {
  description = "Minutes to wait before terminating the old task set."
  type        = number
  default     = 5
}

variable "auto_rollback_enabled" {
  description = "Whether CodeDeploy auto rollback is enabled."
  type        = bool
  default     = true
}

variable "tags" {
  description = "Tags applied to supported resources."
  type        = map(string)
  default     = {}
}
