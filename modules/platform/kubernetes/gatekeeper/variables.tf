variable "namespace" {
  description = "Dedicated namespace for the Gatekeeper Helm release."
  type        = string
  default     = "gatekeeper-system"
  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "namespace must not be empty."
  }
}

variable "create_namespace" {
  description = "Whether to create the Gatekeeper namespace."
  type        = bool
  default     = true
}

variable "chart_version" {
  description = "Optional Gatekeeper chart version. Pin a reviewed version in production."
  type        = string
  default     = ""
}

variable "values" {
  description = "Additional YAML values passed to the Gatekeeper Helm release."
  type        = list(string)
  default     = []
}

variable "constraint_templates" {
  description = "ConstraintTemplate YAML manifests keyed by stable logical name. Empty by default."
  type        = map(string)
  default     = {}
}

variable "constraints" {
  description = "Constraint YAML manifests keyed by stable logical name. Prefer enforcementAction dryrun during rollout."
  type        = map(string)
  default     = {}
}

variable "labels" {
  description = "Labels applied to the namespace."
  type        = map(string)
  default     = {}
}
