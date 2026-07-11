variable "name" {
  description = "Logical Nomad cluster name."
  type        = string
}
variable "environment" {
  description = "Environment identifier."
  type        = string
}
variable "datacenter" {
  description = "Nomad datacenter name."
  type        = string
  default     = "dc1"
}
variable "data_dir" {
  description = "Nomad data directory."
  type        = string
  default     = "/opt/nomad"
}
variable "bind_addr" {
  description = "Bind address expression."
  type        = string
  default     = "0.0.0.0"
}
variable "server_count" {
  description = "Expected Nomad server quorum size."
  type        = number
  default     = 3
  validation {
    condition     = var.server_count >= 1 && var.server_count % 2 == 1
    error_message = "server_count must be a positive odd number."
  }
}
variable "server_addresses" {
  description = "Nomad server addresses used by clients."
  type        = list(string)
  default     = []
}
