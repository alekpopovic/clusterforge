variable "name" {
  description = "Logical name for this root placeholder."
  type        = string
  default     = "kubernetes-with-ingress"
}

variable "acme_email" {
  description = "Email address used for the example cert-manager ACME account."
  type        = string
  default     = "platform@example.com"
}
