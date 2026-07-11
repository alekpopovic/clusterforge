variable "namespace" {
  description = "Dedicated Kubernetes namespace for Kyverno."
  type        = string
  default     = "kyverno"

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "namespace must not be empty."
  }
}

variable "create_namespace" {
  description = "Whether to create the dedicated Kyverno namespace."
  type        = bool
  default     = true
}

variable "chart_version" {
  description = "Optional Kyverno chart version. Pin this in production; empty uses the provider-resolved chart version."
  type        = string
  default     = ""
}

variable "values" {
  description = "Additional YAML values passed to the Kyverno Helm release."
  type        = list(string)
  default     = []
}

variable "enable_baseline_policies" {
  description = "Whether to create the optional baseline ClusterPolicies after Kyverno CRDs exist."
  type        = bool
  default     = false
}

variable "baseline_failure_action" {
  description = "Failure action for baseline validation rules. Audit is the safe default; Enforce blocks new violations."
  type        = string
  default     = "Audit"

  validation {
    condition     = contains(["Audit", "Enforce"], var.baseline_failure_action)
    error_message = "baseline_failure_action must be Audit or Enforce."
  }
}

variable "enable_require_resources_policy" {
  description = "Whether the baseline includes a policy requiring CPU and memory requests and limits."
  type        = bool
  default     = true
}

variable "enable_disallow_latest_tag_policy" {
  description = "Whether the baseline includes a policy disallowing the latest image tag."
  type        = bool
  default     = true
}

variable "enable_pod_security_extended_policy" {
  description = "Whether the baseline disallows host networking/hostPath and requires non-root containers with dropped capabilities."
  type        = bool
  default     = false
}

variable "allowed_registries" {
  description = "Optional registry prefixes allowed by the image policy. Empty disables registry allowlisting."
  type        = list(string)
  default     = []
}

variable "require_image_digest" {
  description = "Whether container images must use an immutable sha256 digest. Intended for reviewed production rollouts."
  type        = bool
  default     = false
}

variable "policies" {
  description = "Additional policy manifests keyed by a stable logical name. Values must be YAML and require installed CRDs."
  type        = map(string)
  default     = {}
}

variable "labels" {
  description = "Labels applied to the namespace and baseline policies."
  type        = map(string)
  default     = {}
}
