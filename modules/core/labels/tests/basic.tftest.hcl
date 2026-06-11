run "basic_labels" {
  command = plan

  variables {
    project     = "clusterforge"
    environment = "dev"
    app         = "api"
    component   = "web"
    part_of     = "platform"
    managed_by  = "terraform"
    extra_labels = {
      "example.com/team" = "platform"
    }
  }

  assert {
    condition     = output.labels["app.kubernetes.io/name"] == "api"
    error_message = "Expected Kubernetes app name label."
  }

  assert {
    condition     = output.labels["clusterforge.io/environment"] == "dev"
    error_message = "Expected ClusterForge environment label."
  }
}
