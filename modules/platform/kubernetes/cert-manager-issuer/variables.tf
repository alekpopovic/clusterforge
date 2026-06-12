variable "name" {
  description = "Name of the cert-manager Issuer or ClusterIssuer."
  type        = string
  default     = "letsencrypt-prod"

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "name must not be empty."
  }
}

variable "kind" {
  description = "cert-manager resource kind to create. Must be Issuer or ClusterIssuer."
  type        = string
  default     = "ClusterIssuer"

  validation {
    condition     = contains(["Issuer", "ClusterIssuer"], var.kind)
    error_message = "kind must be Issuer or ClusterIssuer."
  }
}

variable "namespace" {
  description = "Namespace for Issuer resources. Leave empty for ClusterIssuer."
  type        = string
  default     = ""

  validation {
    condition     = var.kind != "Issuer" || length(trimspace(var.namespace)) > 0
    error_message = "namespace is required when kind is Issuer."
  }
}

variable "email" {
  description = "ACME account email address used by cert-manager."
  type        = string

  validation {
    condition     = length(trimspace(var.email)) > 0
    error_message = "email must not be empty."
  }
}

variable "server" {
  description = "ACME server URL."
  type        = string
  default     = "https://acme-v02.api.letsencrypt.org/directory"

  validation {
    condition     = length(trimspace(var.server)) > 0
    error_message = "server must not be empty."
  }
}

variable "private_key_secret_name" {
  description = "Kubernetes Secret name where cert-manager stores the ACME account private key."
  type        = string
  default     = "letsencrypt-prod"

  validation {
    condition     = length(trimspace(var.private_key_secret_name)) > 0
    error_message = "private_key_secret_name must not be empty."
  }
}

variable "solvers" {
  description = "ACME challenge solvers for the issuer."
  type = list(object({
    http01_ingress_class = optional(string)
    dns01_route53 = optional(object({
      region         = string
      hosted_zone_id = optional(string)
    }))
  }))
  default = []

  validation {
    condition     = length(var.solvers) > 0
    error_message = "At least one solver is required."
  }

  validation {
    condition = alltrue([
      for solver in var.solvers : (
        (try(solver.http01_ingress_class, null) != null && try(solver.dns01_route53, null) == null) ||
        (try(solver.http01_ingress_class, null) == null && try(solver.dns01_route53, null) != null)
      )
    ])
    error_message = "Each solver must set exactly one of http01_ingress_class or dns01_route53."
  }

  validation {
    condition = alltrue([
      for solver in var.solvers : try(solver.dns01_route53, null) == null || length(trimspace(solver.dns01_route53.region)) > 0
    ])
    error_message = "dns01_route53.region must not be empty when dns01_route53 is set."
  }
}
