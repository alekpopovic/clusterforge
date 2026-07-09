locals {
  release_name = "velero"

  service_account_annotations = trimspace(var.service_account_role_arn) == "" ? {} : {
    "eks.amazonaws.com/role-arn" = var.service_account_role_arn
  }

  default_values = yamlencode({
    credentials = {
      useSecret = false
    }
    serviceAccount = {
      server = {
        annotations = local.service_account_annotations
      }
    }
    configuration = {
      backupStorageLocation = [
        {
          name     = var.backup_storage_location_name
          provider = var.velero_provider
          bucket   = var.bucket
        }
      ]
      volumeSnapshotLocation = [
        {
          name     = var.volume_snapshot_location_name
          provider = var.velero_provider
        }
      ]
    }
    initContainers = [
      {
        name  = "velero-plugin-for-aws"
        image = var.aws_plugin_image
        volumeMounts = [
          {
            mountPath = "/target"
            name      = "plugins"
          }
        ]
      }
    ]
  })
}

resource "kubernetes_namespace_v1" "this" {
  count = var.create_namespace ? 1 : 0

  metadata {
    name   = var.namespace
    labels = var.labels
  }
}

resource "helm_release" "this" {
  name       = local.release_name
  namespace  = var.namespace
  repository = "https://vmware-tanzu.github.io/helm-charts"
  chart      = "velero"
  version    = var.chart_version == "" ? null : var.chart_version
  values     = concat([local.default_values], var.values)

  depends_on = [kubernetes_namespace_v1.this]
}
