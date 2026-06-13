# workloads/kubernetes/helm-app

Deploys an application Helm chart using the Terraform Helm provider with
ClusterForge naming and metadata conventions.

Provider configuration belongs in the root module. This module does not create
or configure a cluster.

## Basic Helm App

```hcl
module "helm_app" {
  source = "../../../modules/workloads/kubernetes/helm-app"

  name       = "podinfo"
  namespace  = "apps"
  repository = "https://stefanprodan.github.io/podinfo"
  chart      = "podinfo"

  chart_version = "6.7.1"
}
```

Pin `chart_version` before production use. Leaving it empty allows the provider
to resolve the latest available chart version, which is convenient for
experiments but risky for repeatable environments.

## Values Example

```hcl
module "helm_app" {
  source = "../../../modules/workloads/kubernetes/helm-app"

  name       = "podinfo"
  namespace  = "apps"
  repository = "https://stefanprodan.github.io/podinfo"
  chart      = "podinfo"

  values = [
    yamlencode({
      replicaCount = 2
      service = {
        type = "ClusterIP"
      }
    })
  ]

  set = {
    "ui.message" = "hello from ClusterForge"
  }
}
```

## Sensitive Values Warning

`set_sensitive` marks values as sensitive in Terraform output, but those values
can still be retained in Terraform state or passed through provider internals.
Prefer charts that reference existing Kubernetes Secrets, External Secrets
Operator, or cloud secret managers instead of passing raw secret values through
Terraform.

```hcl
set_sensitive = {
  "existingSecretToken" = var.token
}
```

Use this only when you understand the state implications.

## Terraform Or Argo CD

Use this module when Terraform owns a small Helm application lifecycle. Prefer
Argo CD or another GitOps controller when applications change frequently, need
progressive delivery workflows, or are managed by app teams outside the
infrastructure plan/apply process.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
