locals {
  name         = trimspace(var.name)
  environment  = trimspace(var.environment)
  cluster_name = element(split("/", var.cluster_arn), length(split("/", var.cluster_arn)) - 1)
  log_group    = "/ecs/${local.name}"

  common_tags = merge(var.tags, {
    Name        = local.name
    Environment = local.environment
  })

  execution_role_arn = var.execution_role_arn != "" ? var.execution_role_arn : aws_iam_role.execution[0].arn
  task_role_arn      = var.task_role_arn != "" ? var.task_role_arn : aws_iam_role.task[0].arn

  container_name = try(var.load_balancer.container_name, null) == null ? local.name : var.load_balancer.container_name
  container_port = try(var.load_balancer.container_port, null) == null ? var.container_port : var.load_balancer.container_port
}

data "aws_region" "current" {}

data "aws_iam_policy_document" "ecs_tasks_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_cloudwatch_log_group" "this" {
  #checkov:skip=CKV_AWS_158:Encryption key selection is configurable and may use an approved external or provider-managed key.
  #checkov:skip=CKV_AWS_338:Encryption key selection is configurable and may use an approved external or provider-managed key.
  count = var.create_log_group ? 1 : 0

  name              = local.log_group
  retention_in_days = var.log_retention_in_days
  tags              = local.common_tags
}

resource "aws_iam_role" "execution" {
  count = var.execution_role_arn == "" ? 1 : 0

  name               = "${local.name}-execution"
  assume_role_policy = data.aws_iam_policy_document.ecs_tasks_assume_role.json
  tags               = local.common_tags
}

resource "aws_iam_role_policy_attachment" "execution" {
  count = var.execution_role_arn == "" ? 1 : 0

  role       = aws_iam_role.execution[0].name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role" "task" {
  count = var.task_role_arn == "" ? 1 : 0

  name               = "${local.name}-task"
  assume_role_policy = data.aws_iam_policy_document.ecs_tasks_assume_role.json
  tags               = local.common_tags
}

resource "aws_iam_role_policy_attachment" "task" {
  for_each = var.task_role_arn == "" ? toset(var.task_role_policy_arns) : toset([])

  role       = aws_iam_role.task[0].name
  policy_arn = each.value
}

resource "aws_iam_role_policy" "task" {
  for_each = var.task_role_arn == "" ? var.task_role_inline_policies : {}

  name   = each.key
  role   = aws_iam_role.task[0].id
  policy = each.value
}

resource "aws_ecs_task_definition" "this" {
  family                   = local.name
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = tostring(var.cpu)
  memory                   = tostring(var.memory)
  execution_role_arn       = local.execution_role_arn
  task_role_arn            = local.task_role_arn
  tags                     = local.common_tags

  runtime_platform {
    operating_system_family = var.runtime_platform.operating_system_family
    cpu_architecture        = var.runtime_platform.cpu_architecture
  }

  container_definitions = jsonencode([
    {
      name      = local.name
      image     = var.image
      essential = true
      portMappings = [
        {
          containerPort = var.container_port
          protocol      = lower(var.protocol)
        }
      ]
      environment = [
        for key, value in var.environment_variables : {
          name  = key
          value = value
        }
      ]
      secrets = [
        for secret in var.secrets : {
          name      = secret.name
          valueFrom = secret.value_from
        }
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-group         = local.log_group
          awslogs-region        = data.aws_region.current.region
          awslogs-stream-prefix = local.name
        }
      }
    }
  ])

  depends_on = [
    aws_cloudwatch_log_group.this,
    aws_iam_role_policy_attachment.execution
  ]
}

resource "aws_ecs_service" "this" {
  name            = local.name
  cluster         = var.cluster_arn
  task_definition = aws_ecs_task_definition.this.arn
  desired_count   = var.desired_count
  launch_type     = "FARGATE"
  tags            = local.common_tags

  deployment_controller {
    type = var.deployment_controller
  }

  network_configuration {
    subnets          = var.subnet_ids
    security_groups  = var.security_group_ids
    assign_public_ip = var.assign_public_ip
  }

  dynamic "load_balancer" {
    for_each = var.load_balancer.enabled ? [var.load_balancer] : []

    content {
      target_group_arn = load_balancer.value.target_group_arn
      container_name   = local.container_name
      container_port   = local.container_port
    }
  }

  lifecycle {
    ignore_changes = [
      desired_count
    ]
  }
}

resource "aws_appautoscaling_target" "this" {
  count = var.autoscaling.enabled ? 1 : 0

  max_capacity       = var.autoscaling.max_capacity
  min_capacity       = var.autoscaling.min_capacity
  resource_id        = "service/${local.cluster_name}/${aws_ecs_service.this.name}"
  scalable_dimension = "ecs:service:DesiredCount"
  service_namespace  = "ecs"
}

resource "aws_appautoscaling_policy" "cpu" {
  count = var.autoscaling.enabled ? 1 : 0

  name               = "${local.name}-cpu"
  policy_type        = "TargetTrackingScaling"
  resource_id        = aws_appautoscaling_target.this[0].resource_id
  scalable_dimension = aws_appautoscaling_target.this[0].scalable_dimension
  service_namespace  = aws_appautoscaling_target.this[0].service_namespace

  target_tracking_scaling_policy_configuration {
    target_value = var.autoscaling.cpu_target_value

    predefined_metric_specification {
      predefined_metric_type = "ECSServiceAverageCPUUtilization"
    }
  }
}
