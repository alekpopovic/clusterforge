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

  create_cluster_kms_key         = var.enable_cluster_encryption && var.create_kms_key && var.kms_key_arn == ""
  cluster_encryption_key_arn     = var.kms_key_arn != "" ? var.kms_key_arn : (local.create_cluster_kms_key ? aws_kms_key.cluster[0].arn : null)
  create_control_plane_log_group = length(var.enabled_cluster_log_types) > 0

  create_ebs_csi_irsa_role = var.enable_ebs_csi_driver_addon && var.create_ebs_csi_irsa_role
  oidc_provider_url        = try(aws_eks_cluster.this.identity[0].oidc[0].issuer, null)
  oidc_issuer_hostpath     = local.oidc_provider_url == null ? null : trimsuffix(replace(local.oidc_provider_url, "https://", ""), "/")
  oidc_provider_arn        = var.enable_irsa ? aws_iam_openid_connect_provider.this[0].arn : null

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
        service_account_role_arn = local.create_ebs_csi_irsa_role ? module.ebs_csi_irsa[0].role_arn : null
      }
    } : {}
  )
}

resource "aws_kms_key" "cluster" {
  count = local.create_cluster_kms_key ? 1 : 0

  description             = "EKS secrets encryption key for ${local.name}"
  deletion_window_in_days = 30
  enable_key_rotation     = true
  tags                    = local.common_tags
}

resource "aws_kms_alias" "cluster" {
  count = local.create_cluster_kms_key ? 1 : 0

  name          = "alias/${local.name}-eks-secrets"
  target_key_id = aws_kms_key.cluster[0].key_id
}

resource "aws_cloudwatch_log_group" "control_plane" {
  count = local.create_control_plane_log_group ? 1 : 0

  name              = "/aws/eks/${local.name}/cluster"
  retention_in_days = var.cluster_log_retention_days
  tags              = local.common_tags
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

  dynamic "encryption_config" {
    for_each = var.enable_cluster_encryption && local.cluster_encryption_key_arn != null ? [local.cluster_encryption_key_arn] : []

    content {
      resources = ["secrets"]

      provider {
        key_arn = encryption_config.value
      }
    }
  }

  lifecycle {
    precondition {
      condition     = length(trimspace(var.vpc_id)) > 0
      error_message = "VPC ID must not be empty."
    }

    precondition {
      condition     = !var.enable_cluster_encryption || local.cluster_encryption_key_arn != null
      error_message = "enable_cluster_encryption requires kms_key_arn or create_kms_key=true."
    }
  }

  depends_on = [
    aws_cloudwatch_log_group.control_plane,
    aws_iam_role_policy_attachment.cluster
  ]
}

data "tls_certificate" "oidc" {
  count = var.enable_irsa ? 1 : 0

  url = aws_eks_cluster.this.identity[0].oidc[0].issuer
}

resource "aws_iam_openid_connect_provider" "this" {
  count = var.enable_irsa ? 1 : 0

  url             = aws_eks_cluster.this.identity[0].oidc[0].issuer
  client_id_list  = ["sts.amazonaws.com"]
  thumbprint_list = [data.tls_certificate.oidc[0].certificates[0].sha1_fingerprint]
  tags            = local.common_tags
}

module "ebs_csi_irsa" {
  count = local.create_ebs_csi_irsa_role ? 1 : 0

  source = "../../../cloud/aws/irsa-role"

  name                 = "${local.name}-ebs-csi"
  environment          = local.environment
  oidc_provider_arn    = local.oidc_provider_arn
  oidc_provider_url    = local.oidc_provider_url
  namespace            = "kube-system"
  service_account_name = "ebs-csi-controller-sa"
  tags                 = local.common_tags

  policy_arns = [
    "arn:aws:iam::aws:policy/service-role/AmazonEBSCSIDriverPolicy"
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

  cluster_name         = aws_eks_cluster.this.name
  node_group_name      = each.key
  node_role_arn        = aws_iam_role.node.arn
  subnet_ids           = each.value.subnet_ids
  instance_types       = each.value.instance_types
  ami_type             = var.node_group_ami_type
  capacity_type        = each.value.capacity_type
  disk_size            = each.value.disk_size
  release_version      = var.node_group_release_version
  force_update_version = var.node_group_force_update_version
  labels               = each.value.labels
  tags                 = local.common_tags

  scaling_config {
    min_size     = each.value.min_size
    desired_size = each.value.desired_size
    max_size     = each.value.max_size
  }

  update_config {
    max_unavailable            = try(var.node_group_update_config.max_unavailable, null)
    max_unavailable_percentage = try(var.node_group_update_config.max_unavailable_percentage, null)
  }

  dynamic "remote_access" {
    for_each = var.node_group_remote_access == null ? [] : [var.node_group_remote_access]

    content {
      ec2_ssh_key               = remote_access.value.ec2_ssh_key
      source_security_group_ids = remote_access.value.source_security_group_ids
    }
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
