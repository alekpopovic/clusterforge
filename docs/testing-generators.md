# Testing CLI Generators

ClusterForge keeps generator output stable with golden tests under
`cli/testdata/golden/`. These snapshots cover environment generation, backend
rendering, application rendering, and template pack overrides.

Run the generator tests with:

```bash
cd cli
go test ./...
```

Update golden files only when the generated Terraform change is intentional:

```bash
cd cli
go test ./... -update-golden
```

The same update flow is available for CI-like wrappers:

```bash
cd cli
CLUSTERFORGE_UPDATE_GOLDEN=true go test ./...
```

Golden files must not contain timestamps, absolute paths, credentials,
account IDs, kubeconfig content, or machine-specific values. Review snapshot
diffs as Terraform code: generated files should remain readable, stable, and
easy to edit after generation.
