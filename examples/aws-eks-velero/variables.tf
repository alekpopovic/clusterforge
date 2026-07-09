variable "aws_region" {
  description = "AWS region used by the example."
  type        = string
  default     = "us-east-1"
}

variable "use_fake_credentials_for_plan" {
  description = "Use fake AWS credentials and skip AWS credential validation for local no-refresh plans. Do not use this for apply."
  type        = bool
  default     = true
}

variable "kubeconfig_path" {
  description = "Path to kubeconfig for the target EKS cluster."
  type        = string
  default     = "~/.kube/config"
}

variable "kubeconfig_context" {
  description = "Kubeconfig context for the target EKS cluster."
  type        = string
  default     = null
}

variable "velero_service_account_role_arn" {
  description = "IAM role ARN for Velero IRSA. Leave empty only for syntax validation."
  type        = string
  default     = ""
}
