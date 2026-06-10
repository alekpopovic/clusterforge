variable "aws_region" {
  description = "AWS region for the development EKS cluster."
  type        = string
  default     = "us-east-1"
}

variable "project" {
  description = "Project identifier."
  type        = string
  default     = "clusterforge"
}

variable "environment" {
  description = "Environment name."
  type        = string
  default     = "dev"
}

variable "vpc_cidr" {
  description = "Development VPC CIDR block."
  type        = string
  default     = "10.40.0.0/16"
}
