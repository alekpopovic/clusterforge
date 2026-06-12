locals {
  cluster_name = trimspace(var.cluster_name)
  role_name    = "${local.cluster_name}-karpenter"

  common_tags = merge(var.tags, {
    Name = local.role_name
  })
}

data "aws_partition" "current" {}

data "aws_iam_policy_document" "controller" {
  # Read-only discovery permissions are intentionally broad because Karpenter
  # evaluates instance offerings, launch templates, subnets, security groups,
  # AMIs, and pricing data while making scheduling decisions.
  statement {
    sid = "KarpenterReadOnlyDiscovery"
    actions = [
      "ec2:DescribeAvailabilityZones",
      "ec2:DescribeImages",
      "ec2:DescribeInstanceTypeOfferings",
      "ec2:DescribeInstanceTypes",
      "ec2:DescribeInstances",
      "ec2:DescribeLaunchTemplates",
      "ec2:DescribeSecurityGroups",
      "ec2:DescribeSpotPriceHistory",
      "ec2:DescribeSubnets",
      "eks:DescribeCluster",
      "pricing:GetProducts",
      "ssm:GetParameter"
    ]
    resources = ["*"]
  }

  # Instance launch permissions cannot be resource-scoped cleanly before EC2
  # creates resources. Conditions constrain launches to resources tagged for
  # this cluster.
  statement {
    sid = "KarpenterScopedEC2Launch"
    actions = [
      "ec2:CreateFleet",
      "ec2:RunInstances"
    ]
    resources = ["*"]

    condition {
      test     = "StringEquals"
      variable = "aws:RequestTag/kubernetes.io/cluster/${local.cluster_name}"
      values   = ["owned"]
    }
  }

  statement {
    sid = "KarpenterScopedEC2Tagging"
    actions = [
      "ec2:CreateTags"
    ]
    resources = ["*"]

    condition {
      test     = "StringEquals"
      variable = "aws:RequestTag/kubernetes.io/cluster/${local.cluster_name}"
      values   = ["owned"]
    }
  }

  statement {
    sid = "KarpenterScopedEC2Termination"
    actions = [
      "ec2:DeleteLaunchTemplate",
      "ec2:TerminateInstances"
    ]
    resources = ["*"]

    condition {
      test     = "StringEquals"
      variable = "aws:ResourceTag/kubernetes.io/cluster/${local.cluster_name}"
      values   = ["owned"]
    }
  }

  # Karpenter must pass the node role selected in EC2NodeClass. This is broad
  # until the project exposes a specific node role ARN input.
  statement {
    sid       = "KarpenterPassNodeRole"
    actions   = ["iam:PassRole"]
    resources = ["*"]

    condition {
      test     = "StringEquals"
      variable = "iam:PassedToService"
      values   = ["ec2.${data.aws_partition.current.dns_suffix}"]
    }
  }
}

module "role" {
  source = "../irsa-role"

  name                 = local.role_name
  environment          = lookup(var.tags, "Environment", "unknown")
  oidc_provider_arn    = var.oidc_provider_arn
  oidc_provider_url    = var.oidc_provider_url
  namespace            = var.namespace
  service_account_name = var.service_account_name
  tags                 = local.common_tags

  inline_policies = {
    KarpenterControllerPolicy = data.aws_iam_policy_document.controller.json
  }
}
