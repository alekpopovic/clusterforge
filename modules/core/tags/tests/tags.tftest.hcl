run "required_tags_exist" {
  command = plan

  variables {
    project     = "clusterforge"
    environment = "dev"
  }

  assert {
    condition     = output.tags.Project == "clusterforge"
    error_message = "Expected Project tag."
  }

  assert {
    condition     = output.tags.Environment == "dev"
    error_message = "Expected Environment tag."
  }

  assert {
    condition     = output.tags.ManagedBy == "terraform"
    error_message = "Expected default ManagedBy tag."
  }
}

run "optional_tags_omitted_when_empty" {
  command = plan

  variables {
    project     = "clusterforge"
    environment = "dev"
    owner       = ""
    cost_center = ""
    component   = ""
  }

  assert {
    condition     = !contains(keys(output.tags), "Owner")
    error_message = "Expected empty Owner tag to be omitted."
  }

  assert {
    condition     = !contains(keys(output.tags), "CostCenter")
    error_message = "Expected empty CostCenter tag to be omitted."
  }

  assert {
    condition     = !contains(keys(output.tags), "Component")
    error_message = "Expected empty Component tag to be omitted."
  }
}

run "extra_tags_merge_correctly" {
  command = plan

  variables {
    project     = "clusterforge"
    environment = "dev"
    owner       = "platform"
    cost_center = "engineering"
    component   = "network"
    extra_tags = {
      Team       = "platform"
      Compliance = "internal"
    }
  }

  assert {
    condition     = output.tags.Owner == "platform"
    error_message = "Expected optional Owner tag."
  }

  assert {
    condition     = output.tags.CostCenter == "engineering"
    error_message = "Expected optional CostCenter tag."
  }

  assert {
    condition     = output.tags.Component == "network"
    error_message = "Expected optional Component tag."
  }

  assert {
    condition     = output.tags.Team == "platform" && output.tags.Compliance == "internal"
    error_message = "Expected extra tags to be included."
  }
}

run "extra_tags_override_standard_tags" {
  command = plan

  variables {
    project     = "clusterforge"
    environment = "dev"
    managed_by  = "terraform"
    extra_tags = {
      Project   = "override-project"
      ManagedBy = "platform-pipeline"
    }
  }

  assert {
    condition     = output.tags.Project == "override-project"
    error_message = "Expected extra_tags to override the Project tag."
  }

  assert {
    condition     = output.tags.ManagedBy == "platform-pipeline"
    error_message = "Expected extra_tags to override the ManagedBy tag."
  }
}
