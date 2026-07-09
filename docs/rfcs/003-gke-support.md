# RFC 003: GKE Support

## Goals

- GCP network module
- GKE cluster module
- node pools
- Workload Identity
- optional Cloud DNS later

## Proposed Modules

- `modules/cloud/gcp/network`
- `modules/cloud/gcp/iam`
- `modules/cloud/gcp/dns`
- `modules/orchestrators/kubernetes/gke`

Provider configuration belongs in root modules. Kubernetes and Helm providers
should be configured from GKE outputs after cluster creation.

## Security

Use Workload Identity and least-privilege service accounts. Do not commit key
files or kubeconfigs.

## CLI Impact

```bash
cf env create dev --cloud gcp --orchestrator gke --region europe-west1
cf generate dev
```

## Risks

Regional versus zonal clusters, private cluster networking, IAM permissions,
and firewall assumptions need dedicated smoke tests.
