run "basic_name_generation" {
  command = plan

  variables {
    project     = "clusterforge"
    environment = "dev"
    component   = "api"
    name        = "web"
  }

  assert {
    condition     = output.name == "clusterforge-dev-api-web"
    error_message = "Expected basic generated name."
  }

  assert {
    condition     = output.full_name == "clusterforge-dev-api-web"
    error_message = "Expected full_name to match untruncated name."
  }

  assert {
    condition     = output.parts == ["clusterforge", "dev", "api", "web"]
    error_message = "Expected non-empty input parts in order."
  }
}

run "lowercase_conversion" {
  command = plan

  variables {
    project     = "ClusterForge"
    environment = "Dev"
    component   = "API"
    name        = "Web"
  }

  assert {
    condition     = output.name == "clusterforge-dev-api-web"
    error_message = "Expected lowercase output by default."
  }
}

run "max_length_truncation" {
  command = plan

  variables {
    project     = "clusterforge"
    environment = "dev"
    component   = "api"
    name        = "verylongworkload"
    max_length  = 20
  }

  assert {
    condition     = output.name == "clusterforge-dev-api"
    error_message = "Expected name to be truncated to max_length and trimmed."
  }

  assert {
    condition     = length(output.name) <= 20
    error_message = "Expected name length to respect max_length."
  }
}

run "extra_parts_and_suffix" {
  command = plan

  variables {
    project     = "clusterforge"
    environment = "staging"
    component   = "worker"
    name        = "queue"
    extra_parts = ["blue", "v2"]
    suffix      = "001"
  }

  assert {
    condition     = output.name == "clusterforge-staging-worker-queue-blue-v2-001"
    error_message = "Expected extra parts and suffix to be appended."
  }
}

run "dns_and_label_safe_names" {
  command = plan

  variables {
    project     = "Cluster_Forge"
    environment = "Dev"
    component   = "API_Service"
    name        = "Web_App"
    separator   = "_"
  }

  assert {
    condition     = output.name == "cluster_forge_dev_api_service_web_app"
    error_message = "Expected primary name to preserve configured underscores."
  }

  assert {
    condition     = output.dns_safe_name == "cluster-forge-dev-api-service-web-app"
    error_message = "Expected DNS-safe name to use lowercase dashes and no underscores."
  }

  assert {
    condition     = output.labels_safe_name == "cluster-forge-dev-api-service-web-app"
    error_message = "Expected labels-safe name to use lowercase dashes and no underscores."
  }

  assert {
    condition     = can(regex("^[a-z0-9]([-a-z0-9]*[a-z0-9])?$", output.dns_safe_name))
    error_message = "Expected DNS-safe name to be DNS-label compatible."
  }
}

run "invalid_separator_fails" {
  command = plan

  variables {
    project     = "clusterforge"
    environment = "dev"
    component   = "api"
    name        = "web"
    separator   = "."
  }

  expect_failures = [var.separator]
}

run "empty_required_values_fail" {
  command = plan

  variables {
    project     = ""
    environment = ""
    component   = ""
    name        = ""
  }

  expect_failures = [
    var.project,
    var.environment,
    var.component,
    var.name,
  ]
}
