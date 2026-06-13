locals {
  pod_security_labels = {
    for name, levels in var.namespaces : name => merge(
      var.labels,
      {
        "pod-security.kubernetes.io/enforce" = levels.enforce
        "pod-security.kubernetes.io/audit"   = levels.audit
        "pod-security.kubernetes.io/warn"    = levels.warn
      }
    )
  }
}

resource "kubernetes_labels" "pod_security" {
  for_each = local.pod_security_labels

  api_version = "v1"
  kind        = "Namespace"
  metadata {
    name = each.key
  }
  labels = each.value
}
