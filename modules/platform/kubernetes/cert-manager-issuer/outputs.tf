output "name" {
  description = "Name of the cert-manager Issuer or ClusterIssuer."
  value       = var.name
}

output "kind" {
  description = "Created cert-manager resource kind."
  value       = var.kind
}

output "namespace" {
  description = "Namespace of the Issuer, or null for ClusterIssuer."
  value       = var.kind == "Issuer" ? var.namespace : null
}

output "manifest" {
  description = "Rendered cert-manager issuer manifest."
  value       = kubernetes_manifest.this.manifest
}
