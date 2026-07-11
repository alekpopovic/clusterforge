## Prompt 135 — Docker target policy and lifecycle decision

```text
Review Docker and Docker Swarm target strategy.

Goal:
Decide whether Docker targets remain supported, experimental, or deprecated.

Create:
- docs/rfcs/010-docker-target-strategy.md
- docs/docker-target.md

Cover:
1. Supported use cases:
   - local experiments
   - simple self-hosted
   - migration/legacy
   - not recommended for large production setups

2. Docker Engine module status
3. Docker Swarm module status
4. Security limitations
5. Networking limitations
6. Secret handling limitations
7. Upgrade/rollback limitations
8. Recommendation vs Kubernetes/ECS/Nomad

CLI:
- Ensure Docker target is marked experimental.
- cf env create with docker should print warning.
- cf doctor should warn for Docker production environments.

Tests:
- Docker target warning.
- Docker prod warning.
- Docs updated.

Rules:
- Do not remove existing Docker modules unless intentionally deprecated.
- Be clear and honest about limitations.
```


---
