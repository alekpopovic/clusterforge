variable "nomad_address" {
  description = "Nomad HTTP API endpoint."
  type        = string
  default     = "http://127.0.0.1:4646"
}

variable "nomad_region" {
  description = "Nomad region."
  type        = string
  default     = "global"
}

variable "datacenters" {
  description = "Nomad datacenters for the job."
  type        = list(string)
  default     = ["dc1"]
}

variable "namespace" {
  description = "Nomad namespace for the job."
  type        = string
  default     = "default"
}
