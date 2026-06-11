# workloads/ecs/service

## Purpose

Deploys a generic AWS ECS Fargate service. The module creates a task
definition, service, optional CloudWatch log group, optional IAM roles, and
optional service autoscaling.

Provider configuration belongs in the root module. This module declares the
AWS provider requirement but does not configure the provider.

## ECS vs Kubernetes Workloads

Use this module for workloads that run directly on ECS/Fargate. Use
`modules/workloads/kubernetes/app` when the workload runs inside a Kubernetes
cluster such as EKS.

## Fargate CPU And Memory

The module validates common Fargate CPU and memory combinations. If AWS adds
new combinations, update the validation before using them.

## Basic Service

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
}
```

## Service With ALB Target Group

```hcl
module "service" {
  source = "../../../modules/workloads/ecs/service"

  name               = "api"
  environment        = "dev"
  cluster_arn        = module.ecs_cluster.cluster_arn
  subnet_ids         = module.network.private_subnet_ids
  security_group_ids = [aws_security_group.service.id]
  image              = "123456789012.dkr.ecr.us-east-1.amazonaws.com/api:1.0.0"
  container_port     = 8080

  load_balancer = {
    enabled          = true
    target_group_arn = aws_lb_target_group.api.arn
    container_port   = 8080
  }
}
```

## Secrets

Do not pass secret values in Terraform variables. Use `secrets` to reference
existing SSM Parameter Store or Secrets Manager ARNs.

```hcl
module "service" {
  source = "../../../modules/workloads/ecs/service"

  name               = "api"
  environment        = "dev"
  cluster_arn        = module.ecs_cluster.cluster_arn
  subnet_ids         = module.network.private_subnet_ids
  security_group_ids = [aws_security_group.service.id]
  image              = "123456789012.dkr.ecr.us-east-1.amazonaws.com/api:1.0.0"
  container_port     = 8080

  secrets = [
    {
      name       = "DATABASE_URL"
      value_from = "arn:aws:ssm:us-east-1:123456789012:parameter/dev/api/database-url"
    }
  ]
}
```

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
