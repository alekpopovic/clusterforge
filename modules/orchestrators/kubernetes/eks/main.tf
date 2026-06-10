locals {
  name        = trimspace(var.name)
  environment = trimspace(var.environment)

  common_tags = merge(var.tags, {
    Name        = local.name
    Environment = local.environment
  })

  node_groups = {
    for name, node_group in var.node_groups : name => merge(node_group, {
      subnet_ids = node_group.subnet_ids == null ? var.subnet_ids : node_group.subnet_ids
    })
  }

  enabled_addons = merge(
    var.enable_vpc_cni_addon ? {
      vpc-cni = {
        name                     = "vpc-cni"
        before_compute           = true
        service_account_role_arn = null
      }
    } : {},
    var.enable_coredns_addon ? {
      coredns = {
        name                     = "coredns"
        before_compute           = false
        service_account_role_arn = null
      }
    } : {},
    var.enable_kube_proxy_addon ? {
      kube-proxy = {
        name                     = "kube-proxy"
        before_compute           = false
        service_account_role_arn = null
      }
    } : {},
    var.enable_ebs_csi_driver_addon ? {
      aws-ebs-csi-driver = {
        name                     = "aws-ebs-csi-driver"
        before_compute           = false
        service_account_role_arn = null
      }
    } : {}
  )
}

data "aws_iam_policy_document" "cluster_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["eks.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "node_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "cluster" {
  name               = "${local.name}-cluster"
  assume_role_policy = data.aws_iam_policy_document.cluster_assume_role.json
  tags               = local.common_tags
}

resource "aws_iam_role_policy_attachment" "cluster" {
  role       = aws_iam_role.cluster.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
}

resource "aws_eks_cluster" "this" {
  name                          = local.name
  role_arn                      = aws_iam_role.cluster.arn
  version                       = var.kubernetes_version
  bootstrap_self_managed_addons = false
  enabled_cluster_log_types     = var.enabled_cluster_log_types
  tags                          = local.common_tags

  vpc_config {
    subnet_ids              = var.subnet_ids
    endpoint_public_access  = var.endpoint_public_access
    endpoint_private_access = var.endpoint_private_access
    public_access_cidrs     = var.public_access_cidrs
  }

  lifecycle {
    precondition {
      condition     = length(trimspace(var.vpc_id)) > 0
      error_message = "VPC ID must not be empty."
    }
  }

  depends_on = [
    aws_iam_role_policy_attachment.cluster
  ]
}

resource "aws_iam_role" "node" {
  name               = "${local.name}-node"
  assume_role_policy = data.aws_iam_policy_document.node_assume_role.json
  tags               = local.common_tags
}

resource "aws_iam_role_policy_attachment" "node_worker" {
  role       = aws_iam_role.node.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
}

resource "aws_iam_role_policy_attachment" "node_cni" {
  role       = aws_iam_role.node.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
}

resource "aws_iam_role_policy_attachment" "node_registry" {
  role       = aws_iam_role.node.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
}

resource "aws_eks_addon" "before_compute" {
  for_each = {
    for key, addon in local.enabled_addons : key => addon
    if addon.before_compute
  }

  cluster_name                = aws_eks_cluster.this.name
  addon_name                  = each.value.name
  resolve_conflicts_on_create = "OVERWRITE"
  resolve_conflicts_on_update = "PRESERVE"
  service_account_role_arn    = each.value.service_account_role_arn
  tags                        = local.common_tags
}

resource "aws_eks_node_group" "this" {
  for_each = local.node_groups

  cluster_name    = aws_eks_cluster.this.name
  node_group_name = each.key
  node_role_arn   = aws_iam_role.node.arn
  subnet_ids      = each.value.subnet_ids
  instance_types  = each.value.instance_types
  capacity_type   = each.value.capacity_type
  disk_size       = each.value.disk_size
  labels          = each.value.labels
  tags            = local.common_tags

  scaling_config {
    min_size     = each.value.min_size
    desired_size = each.value.desired_size
    max_size     = each.value.max_size
  }

  dynamic "taint" {
    for_each = each.value.taints

    content {
      key    = taint.value.key
      value  = try(taint.value.value, null)
      effect = taint.value.effect
    }
  }

  depends_on = [
    aws_eks_addon.before_compute,
    aws_iam_role_policy_attachment.node_worker,
    aws_iam_role_policy_attachment.node_cni,
    aws_iam_role_policy_attachment.node_registry
  ]
}

resource "aws_eks_addon" "after_compute" {
  for_each = {
    for key, addon in local.enabled_addons : key => addon
    if !addon.before_compute
  }

  cluster_name                = aws_eks_cluster.this.name
  addon_name                  = each.value.name
  resolve_conflicts_on_create = "OVERWRITE"
  resolve_conflicts_on_update = "PRESERVE"
  service_account_role_arn    = each.value.service_account_role_arn
  tags                        = local.common_tags

  depends_on = [
    aws_eks_node_group.this
  ]
}

# TODO: Add an IAM Roles for Service Accounts role for aws-ebs-csi-driver when
# the EBS CSI add-on is enabled. That requires OIDC provider and trust policy
# wiring and should stay explicit rather than hidden in a shortcut.
