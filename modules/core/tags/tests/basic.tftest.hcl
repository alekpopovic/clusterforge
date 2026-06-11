run "basic_tags" {
  command = plan

  variables {
    project     = "clusterforge"
    environment = "dev"
    owner       = "platform"
    cost_center = "engineering"
    component   = "network"
    managed_by  = "terraform"
    extra_tags = {
      Team = "platform"
    }
  }

  assert {
    condition     = output.tags.Project == "clusterforge"
    error_message = "Expected Project tag."
  }

  assert {
    condition     = output.tags.Team == "platform"
    error_message = "Expected extra tag to be included."
  }
}
