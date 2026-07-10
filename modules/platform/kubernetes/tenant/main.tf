locals {
  name       = trimspace(var.name)
  namespaces = toset(var.namespaces)
  labels = merge(var.labels, {
    "app.kubernetes.io/managed-by" = "terraform"
    "clusterforge.io/tenant"       = local.name
  })
  namespace_labels = merge(local.labels, {
    "pod-security.kubernetes.io/enforce" = var.pod_security.enforce
    "pod-security.kubernetes.io/audit"   = var.pod_security.audit
    "pod-security.kubernetes.io/warn"    = var.pod_security.warn
  })
}

resource "kubernetes_namespace_v1" "this" {
  for_each = local.namespaces

  metadata {
    name        = each.value
    labels      = local.namespace_labels
    annotations = var.annotations
  }
}

resource "kubernetes_resource_quota_v1" "this" {
  for_each = var.resource_quota.enabled ? local.namespaces : toset([])

  metadata {
    name      = "${local.name}-quota"
    namespace = each.value
    labels    = local.labels
  }

  spec {
    hard = var.resource_quota.hard
  }

  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_limit_range_v1" "this" {
  for_each = var.limit_range.enabled ? local.namespaces : toset([])

  metadata {
    name      = "${local.name}-limits"
    namespace = each.value
    labels    = local.labels
  }

  spec {
    dynamic "limit" {
      for_each = var.limit_range.limits

      content {
        type            = limit.value.type
        default         = limit.value.default
        default_request = limit.value.default_request
        max             = limit.value.max
        min             = limit.value.min
      }
    }
  }

  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_network_policy_v1" "default_deny_ingress" {
  for_each = var.network_policy.default_deny_ingress ? local.namespaces : toset([])

  metadata {
    name      = "default-deny-ingress"
    namespace = each.value
    labels    = local.labels
  }

  spec {
    pod_selector {}
    policy_types = ["Ingress"]
  }

  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_network_policy_v1" "default_deny_egress" {
  for_each = var.network_policy.default_deny_egress ? local.namespaces : toset([])

  metadata {
    name      = "default-deny-egress"
    namespace = each.value
    labels    = local.labels
  }

  spec {
    pod_selector {}
    policy_types = ["Egress"]
  }

  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_network_policy_v1" "allow_dns_egress" {
  for_each = var.network_policy.default_deny_egress && var.network_policy.allow_dns_egress ? local.namespaces : toset([])

  metadata {
    name      = "allow-dns-egress"
    namespace = each.value
    labels    = local.labels
  }

  spec {
    pod_selector {}
    policy_types = ["Egress"]

    egress {
      to {
        namespace_selector {
          match_labels = {
            "kubernetes.io/metadata.name" = "kube-system"
          }
        }
      }

      ports {
        port     = "53"
        protocol = "UDP"
      }

      ports {
        port     = "53"
        protocol = "TCP"
      }
    }
  }

  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_role_v1" "this" {
  for_each = var.rbac.create ? local.namespaces : toset([])

  metadata {
    name      = "${local.name}-role"
    namespace = each.value
    labels    = local.labels
  }

  dynamic "rule" {
    for_each = var.rbac.rules

    content {
      api_groups     = rule.value.api_groups
      resources      = rule.value.resources
      verbs          = rule.value.verbs
      resource_names = try(rule.value.resource_names, null)
    }
  }

  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_role_binding_v1" "this" {
  for_each = var.rbac.create ? local.namespaces : toset([])

  metadata {
    name      = "${local.name}-binding"
    namespace = each.value
    labels    = local.labels
  }

  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "Role"
    name      = kubernetes_role_v1.this[each.key].metadata[0].name
  }

  dynamic "subject" {
    for_each = var.rbac.subjects

    content {
      kind      = subject.value.kind
      name      = subject.value.name
      namespace = subject.value.kind == "ServiceAccount" ? coalesce(try(subject.value.namespace, null), each.value) : null
      api_group = subject.value.kind == "ServiceAccount" ? "" : subject.value.api_group
    }
  }
}
