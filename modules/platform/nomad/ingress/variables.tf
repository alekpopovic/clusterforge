variable "service_name" {
  description = "Nomad service registration name."
  type        = string
}
variable "router" {
  description = "Ingress/service tag prefix."
  type        = string
  default     = "traefik"
}
variable "entrypoints" {
  description = "Optional ingress entrypoint list."
  type        = string
  default     = "websecure"
}
variable "extra_tags" {
  description = "Additional reviewed service discovery tags."
  type        = list(string)
  default     = []
}
