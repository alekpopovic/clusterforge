variable "project" {
  description = "Project tag value."
  type        = string
}

variable "environment" {
  description = "Environment tag value."
  type        = string
}

variable "managed_by" {
  description = "Tool responsible for managing the resources."
  type        = string
  default     = "terraform"
}

variable "extra_tags" {
  description = "Additional tags to merge with the standard tags."
  type        = map(string)
  default     = {}
}
