variable "aws_region" {
  description = "AWS region used by the example."
  type        = string
  default     = "us-east-1"
}

variable "availability_zones" {
  description = "Availability zones used by the example network."
  type        = list(string)
  default     = ["us-east-1a", "us-east-1b"]
}
