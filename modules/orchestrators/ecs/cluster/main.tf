locals {
  name        = trimspace(var.name)
  environment = trimspace(var.environment)

  common_tags = merge(var.tags, {
    Name        = local.name
    Environment = local.environment
  })

  default_capacity_provider_strategy = length(var.default_capacity_provider_strategy) > 0 ? var.default_capacity_provider_strategy : [
    {
      capacity_provider = "FARGATE"
      weight            = 1
      base              = 0
    }
  ]
}

resource "aws_ecs_cluster" "this" {
  name = local.name
  tags = local.common_tags

  setting {
    name  = "containerInsights"
    value = var.enable_container_insights ? "enabled" : "disabled"
  }
}

resource "aws_ecs_cluster_capacity_providers" "this" {
  cluster_name       = aws_ecs_cluster.this.name
  capacity_providers = var.capacity_providers

  dynamic "default_capacity_provider_strategy" {
    for_each = local.default_capacity_provider_strategy

    content {
      capacity_provider = default_capacity_provider_strategy.value.capacity_provider
      weight            = try(default_capacity_provider_strategy.value.weight, null)
      base              = try(default_capacity_provider_strategy.value.base, null)
    }
  }
}
