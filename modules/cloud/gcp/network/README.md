# cloud/gcp/network

MVP GCP VPC and subnet module for GKE.

Status: experimental.

```hcl
module "network" {
  source = "../../../modules/cloud/gcp/network"

  project_id  = "example-project"
  name        = "clusterforge-dev-gke"
  environment = "dev"
  region      = "europe-west1"
}
```

Provider configuration belongs in the root module.
