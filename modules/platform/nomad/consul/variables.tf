variable "address" {
  description = "Consul HTTP address reachable by Nomad."
  type        = string
  default     = "127.0.0.1:8500"
}
variable "auto_advertise" {
  description = "Enable Nomad service advertisement through Consul."
  type        = bool
  default     = true
}
variable "server_service_name" {
  description = "Consul service name for Nomad servers."
  type        = string
  default     = "nomad"
}
variable "client_service_name" {
  description = "Consul service name for Nomad clients."
  type        = string
  default     = "nomad-client"
}
variable "token_reference" {
  description = "Reference placeholder for a Consul ACL token; never pass a real token into Terraform."
  type        = string
  default     = "CONSUL_HTTP_TOKEN_FROM_ENV"
}
