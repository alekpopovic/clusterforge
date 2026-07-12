# Prompt 188 — Control Plane Helm chart

```text
Create Helm chart for ClusterForge Control Plane.

Path:
- charts/clusterforge-control-plane/

Chart components:
- API server Deployment
- Service
- ConfigMap
- Secret example
- Ingress optional
- ServiceAccount
- NetworkPolicy optional
- PodDisruptionBudget optional
- PostgreSQL external configuration
- optional embedded PostgreSQL disabled by default or documented carefully

Values:
- image.repository
- image.tag
- replicaCount
- config
- auth
- database
- ingress
- resources
- nodeSelector
- tolerations
- affinity
- serviceAccount
- networkPolicy

Security defaults:
- runAsNonRoot
- no privileged
- readOnlyRootFilesystem if practical
- secrets via existingSecret
- no default production token

Docs:
- charts/clusterforge-control-plane/README.md
- docs/control-plane-deployment.md

Rules:
- No real secrets.
- Embedded database not recommended for production.
- Do not expose publicly by default.

Run:
- helm lint if helm available
```
