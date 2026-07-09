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

variable "oidc_provider_arn" {
  description = "EKS IAM OIDC provider ARN."
  type        = string
  default     = "arn:aws:iam::000000000000:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/EXAMPLE"
}

variable "oidc_provider_url" {
  description = "EKS OIDC provider URL."
  type        = string
  default     = "https://oidc.eks.us-east-1.amazonaws.com/id/EXAMPLE"
}

variable "hosted_zone_ids" {
  description = "Route53 hosted zone IDs ExternalDNS may manage."
  type        = list(string)
  default     = ["ZREPLACEEXAMPLE"]
}
