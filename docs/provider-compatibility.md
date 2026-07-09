# Provider Compatibility

ClusterForge distinguishes supported versions from tested combinations.

- Supported versions are documented in `VERSION_MATRIX.md`.
- Tested combinations are the focused CI matrix and local validation commands.
- Experimental providers may validate syntax without real cloud smoke evidence.

## CI Matrix

`.github/workflows/provider-compatibility.yml` validates representative stacks
without cloud credentials:

| Stack | Root | Notes |
| --- | --- | --- |
| AWS | `examples/aws-eks-minimal` | Uses `init -backend=false` and `validate`; no apply. |
| Kubernetes/Helm | `examples/existing-kubernetes-platform-bootstrap` | Provider schema validation only; no cluster required. |
| Azure | `examples/azure-aks-minimal` | Experimental; no Azure credentials required for validate. |
| GCP | `examples/gcp-gke-minimal` | Experimental; no GCP credentials required for validate. |
| Nomad | `examples/nomad-service` | Provider schema validation only. |
| Docker | `examples/docker-swarm-service` | Provider schema validation only. |

The matrix uses supported Terraform and OpenTofu versions from
`VERSION_MATRIX.md` where practical. It intentionally avoids every possible
provider combination to keep CI useful and stable.

## Local Run

Run the same style of checks locally:

```bash
terraform fmt -check -recursive
terraform -chdir=examples/aws-eks-minimal init -backend=false
terraform -chdir=examples/aws-eks-minimal validate
```

Use OpenTofu by replacing `terraform` with `tofu`.

## Skip Rules

- Real cloud data sources that require credentials should be skipped or moved
  to opt-in smoke tests.
- Real apply operations never run in compatibility CI.
- Provider plugin download failures should be reported clearly rather than
  hidden as successful validation.
