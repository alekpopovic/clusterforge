terraform {
  required_version = ">= 1.6.0"
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = ">= 2.0, < 3.0"
    }
  }
}
