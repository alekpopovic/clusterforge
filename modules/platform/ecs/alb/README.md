# platform/ecs/alb

## Purpose

Creates an AWS Application Load Balancer for ECS/Fargate services, including
the ALB security group, target groups, and HTTP/HTTPS listeners.

Provider configuration belongs in the root module. This module declares the
AWS provider requirement but does not configure the provider.

## Status

Implemented.

## Usage

```hcl
module "alb" {
  source = "../../../modules/platform/ecs/alb"

  name              = "clusterforge-dev"
  environment       = "dev"
  vpc_id            = module.network.vpc_id
  public_subnet_ids = module.network.public_subnet_ids

  target_groups = {
    web = {
      port              = 80
      health_check_path = "/"
    }
  }

  tags = module.tags.tags
}
```

## HTTPS Listener

```hcl
module "alb" {
  source = "../../../modules/platform/ecs/alb"

  name              = "clusterforge-dev"
  environment       = "dev"
  vpc_id            = module.network.vpc_id
  public_subnet_ids = module.network.public_subnet_ids

  enable_http     = true
  enable_https    = true
  certificate_arn = var.certificate_arn

  target_groups = {
    api = {
      port              = 8080
      health_check_path = "/health"
    }
  }
}
```

When both HTTP and HTTPS are enabled, the HTTP listener redirects to HTTPS.
HTTPS requires an ACM certificate ARN from the root environment.

## ECS Service Wiring

Pass a target group ARN into `modules/workloads/ecs/service`:

```hcl
module "service" {
  source = "../../../modules/workloads/ecs/service"

  name               = "api"
  environment        = "dev"
  cluster_arn        = module.ecs_cluster.cluster_arn
  subnet_ids         = module.network.private_subnet_ids
  security_group_ids = [aws_security_group.service.id]
  image              = "public.ecr.aws/nginx/nginx:latest"
  container_port     = 80

  load_balancer = {
    enabled          = true
    target_group_arn = module.alb.target_group_arns["api"]
    container_name   = "api"
    container_port   = 80
  }
}
```

## Notes

- This module does not create Route53 DNS records.
- This module does not request ACM certificates.
- Target groups use `target_type = "ip"` for Fargate `awsvpc` services.
- Keep service security groups in the root or service composition so each
  workload explicitly decides which ALB can reach it.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
