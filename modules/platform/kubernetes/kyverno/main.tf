locals {
  release_name = "kyverno"
  namespace    = trimspace(var.namespace)
  labels = merge(var.labels, {
    "app.kubernetes.io/managed-by" = "terraform"
    "clusterforge.io/component"    = "kyverno"
  })

  privileged_policy = {
    apiVersion = "kyverno.io/v1"
    kind       = "ClusterPolicy"
    metadata = {
      name   = "clusterforge-disallow-privileged-containers"
      labels = local.labels
    }
    spec = {
      background  = true
      emitWarning = var.baseline_failure_action == "Audit"
      rules = [{
        name = "disallow-privileged-containers"
        match = {
          any = [{ resources = { kinds = ["Pod"] } }]
        }
        validate = {
          failureAction = var.baseline_failure_action
          message       = "Privileged containers are not allowed."
          pattern = {
            spec = {
              containers = [{
                name = "*"
                "=(securityContext)" = {
                  "=(privileged)" = "false"
                }
              }]
            }
          }
        }
      }]
    }
  }

  resources_policy = {
    apiVersion = "kyverno.io/v1"
    kind       = "ClusterPolicy"
    metadata = {
      name   = "clusterforge-require-resources"
      labels = local.labels
    }
    spec = {
      background  = true
      emitWarning = var.baseline_failure_action == "Audit"
      rules = [{
        name = "require-container-requests-and-limits"
        match = {
          any = [{ resources = { kinds = ["Pod"] } }]
        }
        validate = {
          failureAction = var.baseline_failure_action
          message       = "Containers must define CPU and memory requests and limits."
          pattern = {
            spec = {
              containers = [{
                name = "*"
                resources = {
                  requests = { cpu = "?*", memory = "?*" }
                  limits   = { cpu = "?*", memory = "?*" }
                }
              }]
            }
          }
        }
      }]
    }
  }

  latest_tag_policy = {
    apiVersion = "kyverno.io/v1"
    kind       = "ClusterPolicy"
    metadata = {
      name   = "clusterforge-disallow-latest-tag"
      labels = local.labels
    }
    spec = {
      background  = true
      emitWarning = var.baseline_failure_action == "Audit"
      rules = [{
        name = "disallow-latest-tag"
        match = {
          any = [{ resources = { kinds = ["Pod"] } }]
        }
        validate = {
          failureAction = var.baseline_failure_action
          message       = "Container images must not use the latest tag."
          pattern = {
            spec = {
              containers = [{ name = "*", image = "!*:latest" }]
            }
          }
        }
      }]
    }
  }

  pod_security_extended_policy = {
    apiVersion = "kyverno.io/v1"
    kind       = "ClusterPolicy"
    metadata = {
      name   = "clusterforge-pod-security-extended"
      labels = local.labels
    }
    spec = {
      background  = true
      emitWarning = var.baseline_failure_action == "Audit"
      rules = [
        {
          name  = "restrict-pod-and-container-security"
          match = { any = [{ resources = { kinds = ["Pod"] } }] }
          validate = {
            failureAction = var.baseline_failure_action
            message       = "Pods must avoid host networking/hostPath and run as non-root with all capabilities dropped."
            pattern = {
              spec = {
                "=(hostNetwork)" = false
                containers = [{
                  name = "*"
                  securityContext = {
                    runAsNonRoot = true
                    capabilities = { drop = ["ALL"] }
                  }
                }]
              }
            }
          }
        },
        {
          name  = "disallow-hostpath-volumes"
          match = { any = [{ resources = { kinds = ["Pod"] } }] }
          validate = {
            failureAction = var.baseline_failure_action
            message       = "hostPath volumes are not allowed."
            foreach = [{
              list = "request.object.spec.volumes || `[]`"
              deny = { conditions = { any = [{ key = "{{ element.hostPath || '' }}", operator = "NotEquals", value = "" }] } }
            }]
          }
        }
      ]
    }
  }

  registry_policy = {
    apiVersion = "kyverno.io/v1"
    kind       = "ClusterPolicy"
    metadata = {
      name   = "clusterforge-require-approved-registry"
      labels = local.labels
    }
    spec = {
      background  = true
      emitWarning = var.baseline_failure_action == "Audit"
      rules = [{
        name  = "require-approved-registry"
        match = { any = [{ resources = { kinds = ["Pod"] } }] }
        validate = {
          failureAction = var.baseline_failure_action
          message       = "Container images must come from an approved registry."
          foreach = [{
            list = "request.object.spec.containers"
            deny = { conditions = { all = [{ key = "{{ element.image }}", operator = "AnyNotIn", value = [for registry in var.allowed_registries : "${registry}/*"] }] } }
          }]
        }
      }]
    }
  }

  digest_policy = {
    apiVersion = "kyverno.io/v1"
    kind       = "ClusterPolicy"
    metadata = {
      name   = "clusterforge-require-image-digest"
      labels = local.labels
    }
    spec = {
      background  = true
      emitWarning = var.baseline_failure_action == "Audit"
      rules = [{
        name  = "require-sha256-digest"
        match = { any = [{ resources = { kinds = ["Pod"] } }] }
        validate = {
          failureAction = var.baseline_failure_action
          message       = "Container images must be pinned by sha256 digest."
          pattern       = { spec = { containers = [{ name = "*", image = "*@sha256:*" }] } }
        }
      }]
    }
  }

  baseline_policies = var.enable_baseline_policies ? merge(
    { privileged = local.privileged_policy },
    var.enable_require_resources_policy ? { resources = local.resources_policy } : {},
    var.enable_disallow_latest_tag_policy ? { latest_tag = local.latest_tag_policy } : {},
    var.enable_pod_security_extended_policy ? { pod_security_extended = local.pod_security_extended_policy } : {},
    length(var.allowed_registries) > 0 ? { registry = local.registry_policy } : {},
    var.require_image_digest ? { digest = local.digest_policy } : {}
  ) : {}
}

resource "kubernetes_namespace_v1" "this" {
  count = var.create_namespace ? 1 : 0

  metadata {
    name   = local.namespace
    labels = local.labels
  }
}

resource "helm_release" "this" {
  name       = local.release_name
  namespace  = local.namespace
  repository = "https://kyverno.github.io/kyverno/"
  chart      = "kyverno"
  version    = var.chart_version == "" ? null : var.chart_version
  values     = var.values

  depends_on = [kubernetes_namespace_v1.this]
}

resource "kubernetes_manifest" "baseline" {
  for_each = local.baseline_policies

  manifest = each.value

  depends_on = [helm_release.this]
}

resource "kubernetes_manifest" "custom" {
  for_each = var.policies

  manifest = yamldecode(each.value)

  depends_on = [helm_release.this]
}
