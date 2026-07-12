locals {
  name        = "clusterforge-ecs-alb-dev"
  environment = "dev"
}

module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = local.environment
  component   = "ecs"
}

module "network" {
  source = "../../modules/cloud/aws/network"

  name               = local.name
  environment        = local.environment
  cidr_block         = "10.60.0.0/16"
  availability_zones = var.availability_zones

  public_subnet_cidrs  = ["10.60.0.0/20", "10.60.16.0/20"]
  private_subnet_cidrs = ["10.60.128.0/20", "10.60.144.0/20"]

  enable_nat_gateway = true
  single_nat_gateway = true
  tags               = module.tags.tags
}

module "ecs_cluster" {
  source = "../../modules/orchestrators/ecs/cluster"

  name        = local.name
  environment = local.environment
  tags        = module.tags.tags
}

module "alb" {
  source = "../../modules/platform/ecs/alb"

  name              = local.name
  environment       = local.environment
  vpc_id            = module.network.vpc_id
  public_subnet_ids = module.network.public_subnet_ids

  enable_http     = true
  enable_https    = var.certificate_arn != ""
  certificate_arn = var.certificate_arn

  target_groups = {
    web = {
      port              = 80
      health_check_path = "/"
    }
  }

  tags = module.tags.tags
}

#trivy:ignore:AWS-0104
resource "aws_security_group" "service" {
  #checkov:skip=CKV2_AWS_5:The security group is exported or passed to a child module; static graph analysis cannot resolve that attachment.
  #checkov:skip=CKV_AWS_382:The security group is exported or passed to a child module; static graph analysis cannot resolve that attachment.
  name        = "${local.name}-service"
  description = "Example ECS service security group."
  vpc_id      = module.network.vpc_id

  ingress {
    description     = "Allow traffic from the ALB."
    from_port       = 80
    to_port         = 80
    protocol        = "tcp"
    security_groups = [module.alb.security_group_id]
  }

  egress {
    description = "Allow outbound traffic for image pulls and service calls."
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = module.tags.tags
}

module "service" {
  source = "../../modules/workloads/ecs/service"

  name               = "web"
  environment        = local.environment
  cluster_arn        = module.ecs_cluster.cluster_arn
  subnet_ids         = module.network.private_subnet_ids
  security_group_ids = [aws_security_group.service.id]
  image              = "public.ecr.aws/nginx/nginx:latest"
  container_port     = 80

  load_balancer = {
    enabled          = true
    target_group_arn = module.alb.target_group_arns["web"]
    container_name   = "web"
    container_port   = 80
  }

  environment_variables = {
    APP_ENV = local.environment
  }

  tags = module.tags.tags
}
