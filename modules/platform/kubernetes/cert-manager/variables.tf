variable "namespace" {
  description = "Kubernetes namespace for the cert-manager Helm release."
  type        = string
  default     = "cert-manager"
}

variable "chart_version" {
  description = "Optional cert-manager chart version. Leave empty to use the latest provider-resolved version."
  type        = string
  default     = ""
}

variable "values" {
  description = "YAML values passed to the Helm release."
  type        = list(string)
  default     = []
}

variable "service_account_annotations" {
  description = "Annotations applied to the cert-manager service account, such as an EKS IRSA role ARN."
  type        = map(string)
  default     = {}
}

variable "labels" {
  description = "Labels applied to the namespace when create_namespace is true."
  type        = map(string)
  default     = {}
}

variable "create_namespace" {
  description = "Whether to create the namespace before installing the Helm release."
  type        = bool
  default     = true
}
