provider "aws" {
  region = var.aws_region

  default_tags {
    tags = module.tags.tags
  }
}
