## Prompt 155 — Air-gapped and offline support

```text
Design and implement initial air-gapped/offline support.

Goal:
Help users prepare ClusterForge deployments in restricted environments.

Create docs:
- docs/air-gapped.md
- docs/offline-bundles.md

CLI:
- cf bundle create
- cf bundle create --env prod
- cf bundle inspect <bundle>
- cf bundle verify <bundle>

Bundle contents:
- generated Terraform files
- module source snapshot
- provider lock file guidance
- Helm chart list
- container image list
- app manifests
- policy packs
- template packs
- docs/runbooks summary

Do not include:
- secrets
- tfstate
- cloud credentials
- kubeconfig files

Optional:
- Generate image mirror list:
  images.txt
- Generate Helm chart list:
  helm-charts.txt
- Generate provider list:
  providers.txt

Rules:
- Do not download all artifacts in MVP unless explicitly implemented.
- Start by generating manifest bundle.
- Redact sensitive values.
- Add checksum file for bundle contents.

Tests:
- Bundle create from sample project.
- Bundle excludes secrets and tfstate.
- Bundle verify checks checksum.

Run:
- gofmt
- go test ./...
```


---
