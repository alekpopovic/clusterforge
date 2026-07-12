# Prompt 196 — Docker images and container release

```text
Add Docker image build support for ClusterForge components.

Components:
- CLI optional
- Control Plane API server
- Runner
- Dashboard

Create:
- Dockerfile.control-plane
- Dockerfile.runner
- Dockerfile.dashboard
- .dockerignore
- docker-compose.yaml for local dev
- docs/docker-images.md

Image requirements:
- small runtime image
- non-root user
- no credentials baked in
- version labels
- source labels
- healthcheck where practical

docker-compose local dev:
- control-plane
- postgres
- dashboard
- optional runner

CI:
- build images on pull request
- do not push by default
- release workflow may push later if configured

Rules:
- Do not publish images automatically unless release workflow explicitly configured.
- Do not include secrets.
- Keep local compose for development only.

Tests:
- docker build if Docker available
- otherwise scripts should skip with clear message
```
