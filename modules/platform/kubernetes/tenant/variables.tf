variable "name" {
  description = "Logical tenant name used in labels and namespace-scoped resource names."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "name must not be empty."
  }
}

variable "namespaces" {
  description = "Namespace names owned by the tenant."
  type        = list(string)

  validation {
    condition     = length(var.namespaces) > 0 && length(distinct(var.namespaces)) == length(var.namespaces) && alltrue([for namespace in var.namespaces : length(trimspace(namespace)) > 0])
    error_message = "namespaces must contain at least one unique, non-empty namespace name."
  }
}

variable "labels" {
  description = "Additional labels applied to tenant resources."
  type        = map(string)
  default     = {}
}

variable "annotations" {
  description = "Annotations applied to tenant namespaces."
  type        = map(string)
  default     = {}
}

variable "pod_security" {
  description = "Pod Security Admission levels applied to tenant namespaces."
  type = object({
    enforce = optional(string, "baseline")
    audit   = optional(string, "restricted")
    warn    = optional(string, "restricted")
  })
  default = {}

  validation {
    condition     = alltrue([for level in [var.pod_security.enforce, var.pod_security.audit, var.pod_security.warn] : contains(["privileged", "baseline", "restricted"], level)])
    error_message = "pod_security levels must be privileged, baseline, or restricted."
  }
}

variable "resource_quota" {
  description = "Optional ResourceQuota configuration applied independently to each tenant namespace."
  type = object({
    enabled = optional(bool, false)
    hard    = optional(map(string), {})
  })
  default = {}

  validation {
    condition     = !var.resource_quota.enabled || length(var.resource_quota.hard) > 0
    error_message = "resource_quota.hard must not be empty when resource quota is enabled."
  }
}

variable "limit_range" {
  description = "Optional LimitRange configuration applied independently to each tenant namespace."
  type = object({
    enabled = optional(bool, false)
    limits = optional(list(object({
      type            = string
      default         = optional(map(string), {})
      default_request = optional(map(string), {})
      max             = optional(map(string), {})
      min             = optional(map(string), {})
    })), [])
  })
  default = {}

  validation {
    condition     = !var.limit_range.enabled || length(var.limit_range.limits) > 0
    error_message = "limit_range.limits must not be empty when the limit range is enabled."
  }
}

variable "network_policy" {
  description = "Opt-in namespace network isolation settings."
  type = object({
    default_deny_ingress = optional(bool, false)
    default_deny_egress  = optional(bool, false)
    allow_dns_egress     = optional(bool, true)
  })
  default = {}
}

variable "rbac" {
  description = "Optional minimal namespace Role and RoleBinding configuration."
  type = object({
    create = optional(bool, false)
    subjects = optional(list(object({
      kind      = string
      name      = string
      namespace = optional(string)
      api_group = optional(string, "rbac.authorization.k8s.io")
    })), [])
    rules = optional(list(object({
      api_groups     = list(string)
      resources      = list(string)
      verbs          = list(string)
      resource_names = optional(list(string))
      })), [
      {
        api_groups = [""]
        resources  = ["configmaps", "pods", "services"]
        verbs      = ["get", "list", "watch"]
      }
    ])
  })
  default = {}

  validation {
    condition     = !var.rbac.create || length(var.rbac.subjects) > 0
    error_message = "rbac.subjects must not be empty when RBAC creation is enabled."
  }

  validation {
    condition     = alltrue([for subject in var.rbac.subjects : contains(["User", "Group", "ServiceAccount"], subject.kind) && length(trimspace(subject.name)) > 0])
    error_message = "RBAC subjects must have kind User, Group, or ServiceAccount and a non-empty name."
  }
}
