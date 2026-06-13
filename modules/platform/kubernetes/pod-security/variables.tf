variable "namespaces" {
  description = "Namespace names and Pod Security Admission levels to apply."
  type = map(object({
    enforce = optional(string, "baseline")
    audit   = optional(string, "restricted")
    warn    = optional(string, "restricted")
  }))

  validation {
    condition = alltrue(flatten([
      for _, namespace in var.namespaces : [
        contains(["privileged", "baseline", "restricted"], namespace.enforce),
        contains(["privileged", "baseline", "restricted"], namespace.audit),
        contains(["privileged", "baseline", "restricted"], namespace.warn)
      ]
    ]))
    error_message = "Pod Security levels must be privileged, baseline, or restricted."
  }
}

variable "labels" {
  description = "Additional labels to apply to each namespace."
  type        = map(string)
  default     = {}
}
