module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "dev"
  component   = "ecr"
}

module "ecr" {
  source = "../../modules/cloud/aws/ecr"

  repositories = {
    api = {
      lifecycle_policy_json = jsonencode({
        rules = [
          {
            rulePriority = 1
            description  = "Keep the last 30 images."
            selection = {
              tagStatus   = "any"
              countType   = "imageCountMoreThan"
              countNumber = 30
            }
            action = {
              type = "expire"
            }
          }
        ]
      })
    }
    worker = {}
  }

  tags = module.tags.tags
}
