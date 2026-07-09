variable "namespace" {
  description = "Kubernetes namespace for the Velero Helm release."
  type        = string
  default     = "velero"

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "namespace must not be empty."
  }
}

variable "create_namespace" {
  description = "Whether to create the namespace before installing the Helm release."
  type        = bool
  default     = true
}

variable "chart_version" {
  description = "Optional Velero chart version. Leave empty to use the latest provider-resolved version."
  type        = string
  default     = ""
}

variable "velero_provider" {
  description = "Velero object storage provider. AWS is the first supported configuration."
  type        = string
  default     = "aws"

  validation {
    condition     = contains(["aws"], var.velero_provider)
    error_message = "Only aws is supported by this module today."
  }
}

variable "bucket" {
  description = "Object storage bucket name used by Velero backups."
  type        = string

  validation {
    condition     = length(trimspace(var.bucket)) > 0
    error_message = "bucket must not be empty."
  }
}

variable "backup_storage_location_name" {
  description = "Velero BackupStorageLocation name."
  type        = string
  default     = "default"
}

variable "volume_snapshot_location_name" {
  description = "Velero VolumeSnapshotLocation name."
  type        = string
  default     = "default"
}

variable "service_account_role_arn" {
  description = "Optional IAM role ARN annotated on the Velero service account for IRSA."
  type        = string
  default     = ""
}

variable "aws_plugin_image" {
  description = "Velero AWS plugin image used by the init container."
  type        = string
  default     = "velero/velero-plugin-for-aws:v1.10.1"
}

variable "values" {
  description = "Additional YAML values passed to the Velero Helm release."
  type        = list(string)
  default     = []
}

variable "labels" {
  description = "Labels applied to the namespace when create_namespace is true."
  type        = map(string)
  default     = {}
}
