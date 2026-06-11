run "basic_name_generation" {
  command = plan

  variables {
    project     = "ClusterForge"
    environment = "Dev"
    component   = "API"
    name        = "Web"
    extra_parts = ["Blue"]
    suffix      = "001"
  }

  assert {
    condition     = output.name == "clusterforge-dev-api-web-blue-001"
    error_message = "Expected normalized full name output."
  }

  assert {
    condition     = output.dns_safe_name == "clusterforge-dev-api-web-blue-001"
    error_message = "Expected DNS-safe output."
  }
}
