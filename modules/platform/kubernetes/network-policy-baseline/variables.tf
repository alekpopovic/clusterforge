variable "namespace" {
  description = "Namespace where baseline NetworkPolicy resources are created."
  type        = string

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "namespace must not be empty."
  }
}

variable "default_deny_ingress" {
  description = "Whether to create a default deny ingress NetworkPolicy."
  type        = bool
  default     = true
}

variable "default_deny_egress" {
  description = "Whether to create a default deny egress NetworkPolicy."
  type        = bool
  default     = false
}

variable "allow_dns_egress" {
  description = "Whether to allow DNS egress to kube-system pods labeled k8s-app=kube-dns."
  type        = bool
  default     = true
}

variable "labels" {
  description = "Labels applied to NetworkPolicy resources."
  type        = map(string)
  default     = {}
}
