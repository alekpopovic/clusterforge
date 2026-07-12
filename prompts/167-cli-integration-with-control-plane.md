# Prompt 167 — CLI integration with Control Plane

```text
Add ClusterForge CLI integration with the Control Plane API.

CLI commands:
- cf login
- cf logout
- cf context list
- cf context use <name>
- cf context current
- cf api status
- cf api sync
- cf api push-inventory
- cf api push-policy-results
- cf api push-drift-results
- cf api push-cost-report

Config:
- Store local CLI context in:
  ~/.clusterforge/config.yaml
or project-local:
  .cf/context.yaml

Do not store plaintext tokens if OS keychain support exists.
If not implemented, document token file risk and restrict file permissions.

cf login:
- asks for API URL
- asks for token via hidden input
- validates /api/v1/me
- stores context

cf api sync:
- reads clusterforge.yaml
- app manifests
- service catalog
- module catalog
- pushes sanitized inventory to control plane

Rules:
- No secrets.
- Redact sensitive metadata.
- Make API integration optional.
- Existing local CLI workflows must continue working without API.
- Add --offline or --no-api behavior where useful.

Tests:
- API client tests with httptest server
- login stores context
- sync pushes expected resources
- token redaction
- API unavailable returns useful error

Docs:
- docs/control-plane-cli.md

Run:
- cd cli && gofmt -w .
- cd cli && go test ./...
```
