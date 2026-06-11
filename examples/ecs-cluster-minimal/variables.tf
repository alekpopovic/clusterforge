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
