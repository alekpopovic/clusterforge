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

variable "zone_id" {
  description = "Example existing Route53 hosted zone ID. Replace before applying."
  type        = string
  default     = "Z000000000000EXAMPLE"
}

variable "alb_dns_name" {
  description = "Example ALB DNS name used for an alias record. Replace with module.alb.alb_dns_name in real compositions."
  type        = string
  default     = "example-alb-123456789.us-east-1.elb.amazonaws.com"
}

variable "alb_zone_id" {
  description = "Example ALB hosted zone ID used for an alias record. Replace with module.alb.alb_zone_id in real compositions."
  type        = string
  default     = "Z35SXDOTRQ7X7K"
}
