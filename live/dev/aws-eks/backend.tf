# Example remote backend for a real development environment.
# Keep backend values environment-specific and do not commit secrets.
#
# terraform {
#   backend "s3" {
#     bucket         = "example-clusterforge-dev-terraform-state"
#     key            = "live/dev/aws-eks/terraform.tfstate"
#     region         = "us-east-1"
#     dynamodb_table = "example-clusterforge-dev-terraform-locks"
#     encrypt        = true
#   }
# }
