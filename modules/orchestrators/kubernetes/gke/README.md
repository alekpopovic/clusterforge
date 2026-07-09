# orchestrators/kubernetes/gke

MVP GKE module with Workload Identity enabled.

Status: experimental.

```hcl
module "gke" {
  source = "../../../../modules/orchestrators/kubernetes/gke"

  project_id = "example-project"
  name       = "clusterforge-dev-gke"
  environment = "dev"
  region     = "europe-west1"
  network    = module.network.network_name
  subnetwork = module.network.subnetwork_name
}
```

Do not commit service account keys or generated kubeconfigs.
