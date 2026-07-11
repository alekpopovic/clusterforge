variable "gatekeeper_chart_version" {
  description = "Reviewed Gatekeeper Helm chart version."
  type        = string
}
variable "enforce" {
  description = "Opt in to deny after dry-run findings have been remediated."
  type        = bool
  default     = false
}
