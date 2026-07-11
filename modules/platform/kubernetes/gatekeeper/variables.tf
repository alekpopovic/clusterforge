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

variable "enforcement_action" {
  description = "Enforcement action applied to all supplied constraints. dryrun is the safe default; deny is explicitly opt-in."
  type        = string
  default     = "dryrun"

  validation {
    condition     = contains(["dryrun", "warn", "deny"], var.enforcement_action)
    error_message = "enforcement_action must be dryrun, warn, or deny."
  }
}

variable "labels" {
  description = "Labels applied to the namespace."
  type        = map(string)
  default     = {}
}
