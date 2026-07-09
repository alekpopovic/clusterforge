# Modules

ClusterForge modules live under `modules/` and are intended to stay readable
when consumed by generated roots or direct Terraform composition.

Use local paths while developing:

```hcl
source = "../../../modules/cloud/aws/network"
```

Use Git refs for early consumers:

```hcl
source = "git::https://github.com/<org>/clusterforge.git//modules/cloud/aws/network?ref=v0.1.0"
```

Pin versions in production. Do not use `main` as a production module source.
