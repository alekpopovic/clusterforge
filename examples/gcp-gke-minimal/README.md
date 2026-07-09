# gcp-gke-minimal

Experimental GKE example. It can create billable GCP resources.

```bash
terraform init
terraform plan -var='project_id=example-project'
```

Do not commit service account keys or kubeconfigs.
