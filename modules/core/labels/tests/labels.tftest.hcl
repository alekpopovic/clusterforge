run "app_kubernetes_labels_exist" {
  command = plan

  variables {
    project     = "clusterforge"
    environment = "dev"
    app         = "api"
    component   = "web"
    part_of     = "platform"
    managed_by  = "terraform"
  }

  assert {
    condition     = output.labels["app.kubernetes.io/name"] == "api"
    error_message = "Expected app.kubernetes.io/name label."
  }

  assert {
    condition     = output.labels["app.kubernetes.io/component"] == "web"
    error_message = "Expected app.kubernetes.io/component label."
  }

  assert {
    condition     = output.labels["app.kubernetes.io/part-of"] == "platform"
    error_message = "Expected app.kubernetes.io/part-of label."
  }

  assert {
    condition     = output.labels["app.kubernetes.io/managed-by"] == "terraform"
    error_message = "Expected app.kubernetes.io/managed-by label."
  }
}

run "clusterforge_labels_exist" {
  command = plan

  variables {
    project     = "ClusterForge"
    environment = "Dev"
  }

  assert {
    condition     = output.labels["clusterforge.io/project"] == "clusterforge"
    error_message = "Expected normalized ClusterForge project label."
  }

  assert {
    condition     = output.labels["clusterforge.io/environment"] == "dev"
    error_message = "Expected normalized ClusterForge environment label."
  }
}

run "extra_labels_merge_correctly" {
  command = plan

  variables {
    project     = "clusterforge"
    environment = "dev"
    app         = "api"
    extra_labels = {
      "example.com/team"          = "Platform Team"
      "app.kubernetes.io/name"    = "api-override"
      "clusterforge.io/ownership" = "Shared Services"
    }
  }

  assert {
    condition     = output.labels["example.com/team"] == "platform-team"
    error_message = "Expected extra label values to be normalized."
  }

  assert {
    condition     = output.labels["app.kubernetes.io/name"] == "api-override"
    error_message = "Expected extra labels to merge last and override app label values."
  }

  assert {
    condition     = output.labels["clusterforge.io/ownership"] == "shared-services"
    error_message = "Expected custom ClusterForge label to be included."
  }
}

run "generated_labels_are_kubernetes_compatible" {
  command = plan

  variables {
    project     = "Cluster Forge"
    environment = "Dev_Env"
    app         = "API Service"
    component   = "Web/API"
    part_of     = "Main Platform"
    managed_by  = "Terraform"
    extra_labels = {
      "example.com/release" = "Release Candidate 001"
    }
  }

  assert {
    condition = alltrue([
      for value in values(output.labels) :
      length(value) <= 63 && can(regex("^[A-Za-z0-9]([-A-Za-z0-9_.]*[A-Za-z0-9])?$", value))
    ])
    error_message = "Expected every generated label value to be Kubernetes-compatible."
  }

  assert {
    condition     = output.labels["clusterforge.io/project"] == "cluster-forge"
    error_message = "Expected project label value to be sanitized."
  }

  assert {
    condition     = output.labels["app.kubernetes.io/component"] == "web-api"
    error_message = "Expected component label value to be sanitized."
  }
}
