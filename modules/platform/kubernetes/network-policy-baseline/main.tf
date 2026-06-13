locals {
  labels = merge(
    {
      "app.kubernetes.io/managed-by" = "terraform"
      "clusterforge.io/component"    = "network-policy-baseline"
    },
    var.labels
  )
}

resource "kubernetes_network_policy_v1" "default_deny_ingress" {
  count = var.default_deny_ingress ? 1 : 0

  metadata {
    name      = "default-deny-ingress"
    namespace = var.namespace
    labels    = local.labels
  }

  spec {
    pod_selector {}
    policy_types = ["Ingress"]
  }
}

resource "kubernetes_network_policy_v1" "default_deny_egress" {
  count = var.default_deny_egress ? 1 : 0

  metadata {
    name      = "default-deny-egress"
    namespace = var.namespace
    labels    = local.labels
  }

  spec {
    pod_selector {}
    policy_types = ["Egress"]
  }
}

resource "kubernetes_network_policy_v1" "allow_dns_egress" {
  count = var.default_deny_egress && var.allow_dns_egress ? 1 : 0

  metadata {
    name      = "allow-dns-egress"
    namespace = var.namespace
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

        pod_selector {
          match_labels = {
            "k8s-app" = "kube-dns"
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
}
