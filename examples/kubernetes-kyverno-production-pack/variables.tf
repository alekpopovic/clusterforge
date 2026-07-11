variable "kyverno_chart_version" {
  description = "Reviewed Kyverno Helm chart version."
  type        = string
}
variable "enforce" {
  description = "Opt in to blocking admission after completing the audit rollout."
  type        = bool
  default     = false
}
variable "allowed_registries" {
  description = "Approved image registry prefixes."
  type        = list(string)
  default     = ["registry.example.com"]
}
variable "require_image_digest" {
  description = "Require sha256 image digests."
  type        = bool
  default     = false
}
