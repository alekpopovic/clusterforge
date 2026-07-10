module "payments_tenant" {
  source = "../../modules/platform/kubernetes/tenant"

  name       = "payments"
  namespaces = ["payments-dev"]

  labels = {
    "example.com/owner" = "payments-team"
  }

  pod_security = {
    enforce = "baseline"
    audit   = "restricted"
    warn    = "restricted"
  }

  resource_quota = {
    enabled = true
    hard = {
      "requests.cpu"    = "2"
      "requests.memory" = "4Gi"
      pods              = "20"
    }
  }

  limit_range = {
    enabled = true
    limits = [{
      type = "Container"
      default = {
        cpu    = "500m"
        memory = "512Mi"
      }
      default_request = {
        cpu    = "100m"
        memory = "128Mi"
      }
    }]
  }

  network_policy = {
    default_deny_ingress = false
    default_deny_egress  = false
  }

  rbac = {
    create = true
    subjects = [{
      kind = "Group"
      name = "payments-developers"
    }]
  }
}
