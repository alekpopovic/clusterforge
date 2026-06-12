locals {
  metadata = merge(
    { name = var.name },
    var.kind == "SecretStore" ? { namespace = var.service_account_ref_namespace } : {}
  )
}

resource "kubernetes_manifest" "this" {
  manifest = {
    apiVersion = "external-secrets.io/v1beta1"
    kind       = var.kind

    metadata = local.metadata

    spec = {
      provider = {
        aws = {
          service = var.service
          region  = var.region
          auth = {
            (var.auth_type) = {
              serviceAccountRef = {
                name      = var.service_account_ref_name
                namespace = var.service_account_ref_namespace
              }
            }
          }
        }
      }
    }
  }
}
