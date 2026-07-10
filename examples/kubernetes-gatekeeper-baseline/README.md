# Gatekeeper audit baseline example

The first apply installs Gatekeeper with no constraints. After its CRDs exist,
enable the example ConstraintTemplate and `dryrun` Namespace label constraint:

```bash
terraform init
terraform apply -var='gatekeeper_chart_version=<reviewed-version>'
terraform apply -var='gatekeeper_chart_version=<reviewed-version>' -var='enable_audit_constraint=true'
```

The constraint reports missing `example.com/owner` labels but does not deny
admission. Inspect audit results before any enforcement change. Remove constraints
before destroying Gatekeeper; Helm does not necessarily remove its CRDs.
