# Asset inventory export

`cf inventory export <env>` emits redacted JSON from ClusterForge config,
Terraform configuration/module sources, app manifests and Terraform state when
`terraform show -json` is locally accessible. Use `--stack`, `--fleet`,
`--format json|csv|markdown`, and `--output <file>` for CMDB/reporting workflows.

The exporter makes no cloud API calls. It exports resource identity, provider,
module, environment/stack/cloud/region, non-sensitive tags and a sensitive flag;
it never exports arbitrary resource attributes. Secret-looking tag keys are
redacted and output files use restrictive permissions. State can still expose
metadata, so review and protect exports before distribution.
