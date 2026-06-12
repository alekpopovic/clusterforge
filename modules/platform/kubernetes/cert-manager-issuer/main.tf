locals {
  solver_manifests = [
    for solver in var.solvers : merge(
      try(solver.http01_ingress_class, null) != null ? {
        http01 = {
          ingress = {
            class = solver.http01_ingress_class
          }
        }
      } : {},
      try(solver.dns01_route53, null) != null ? {
        dns01 = {
          route53 = merge(
            {
              region = solver.dns01_route53.region
            },
            try(solver.dns01_route53.hosted_zone_id, null) != null ? {
              hostedZoneID = solver.dns01_route53.hosted_zone_id
            } : {}
          )
        }
      } : {}
    )
  ]

  metadata = merge(
    {
      name = var.name
    },
    var.kind == "Issuer" ? {
      namespace = var.namespace
    } : {}
  )
}

resource "kubernetes_manifest" "this" {
  manifest = {
    apiVersion = "cert-manager.io/v1"
    kind       = var.kind

    metadata = local.metadata

    spec = {
      acme = {
        email  = var.email
        server = var.server

        privateKeySecretRef = {
          name = var.private_key_secret_name
        }

        solvers = local.solver_manifests
      }
    }
  }
}
