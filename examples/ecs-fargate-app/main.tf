module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "dev"
  component   = "ecs"
}

module "network" {
  source = "../../modules/cloud/aws/network"

  name               = "clusterforge-ecs-dev"
  environment        = "dev"
  cidr_block         = "10.50.0.0/16"
  availability_zones = var.availability_zones

  public_subnet_cidrs  = ["10.50.0.0/20", "10.50.16.0/20"]
  private_subnet_cidrs = ["10.50.128.0/20", "10.50.144.0/20"]

  enable_nat_gateway = true
  single_nat_gateway = true
  tags               = module.tags.tags
}

#trivy:ignore:AWS-0104
resource "aws_security_group" "service" {
  #checkov:skip=CKV_AWS_382:Unrestricted egress is an explicit compatibility default; production policy must narrow destination rules.
  name        = "clusterforge-ecs-dev-service"
  description = "Example ECS service security group."
  vpc_id      = module.network.vpc_id

  egress {
    description = "Allow outbound traffic for image pulls and service calls."
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = module.tags.tags
}

module "ecs_cluster" {
  source = "../../modules/orchestrators/ecs/cluster"

  name        = "clusterforge-ecs-dev"
  environment = "dev"
  tags        = module.tags.tags
}

module "service" {
  source = "../../modules/workloads/ecs/service"

  name               = "hello"
  environment        = "dev"
  cluster_arn        = module.ecs_cluster.cluster_arn
  subnet_ids         = module.network.private_subnet_ids
  security_group_ids = [aws_security_group.service.id]
  image              = "public.ecr.aws/nginx/nginx:latest"
  container_port     = 80

  environment_variables = {
    APP_ENV = "dev"
  }

  tags = module.tags.tags
}
